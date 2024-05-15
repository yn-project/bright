package user

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	converter "yun.tea/block/bright/user/pkg/converter/user"
	crud "yun.tea/block/bright/user/pkg/crud/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"yun.tea/block/bright/common/logger"
	proto "yun.tea/block/bright/proto/bright/user"

	"bytes"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/metadata"
	"yun.tea/block/bright/common/ctredis"
	"yun.tea/block/bright/user/pkg/db/ent"
)

const (
	LoginExpireTime = time.Minute * 30
	PrefixLogin     = "PrefixLoginUser"
)

func checkReq(in *proto.UserReq) error {
	if in.Name == nil || *in.Name == "" {
		return fmt.Errorf("name is null")
	}
	if in.Password == nil || *in.Password == "" {
		return fmt.Errorf("passwd is null")
	}
	return nil
}

func setPasswd(in *proto.UserReq) (*crud.UserReq, error) {
	reqInfo := &crud.UserReq{
		Name:     in.Name,
		Password: in.Password,
		Remark:   in.Remark,
	}
	currentTime := time.Now()
	formattedTime := currentTime.Format("20060102150405")
	reqInfo.Salt = &formattedTime

	passwd := *reqInfo.Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwd+formattedTime), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newPasswd := string(hashedPassword)
	reqInfo.Password = &newPasswd
	return reqInfo, nil
}

func (s *Server) CreateUser(ctx context.Context, in *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	var err error
	req := in.GetInfo()
	if req == nil {
		err := fmt.Errorf("invalid req")
		logger.Sugar().Errorw("CreateUser", "error", err)
		return &proto.CreateUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	if err = checkReq(req); err != nil {
		logger.Sugar().Errorw("CreateUser", "error", err)
		return &proto.CreateUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	reqInfo, err := setPasswd(req)
	if err != nil {
		logger.Sugar().Errorw("CreateUser", "error", err)
		return &proto.CreateUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := crud.Create(ctx, reqInfo)
	if err != nil {
		logger.Sugar().Errorw("CreateUser", "error", err)
		return &proto.CreateUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, in *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	var err error
	req := in.GetInfo()
	if req == nil {
		err := fmt.Errorf("invalid req")
		logger.Sugar().Errorw("UpdateUser", "error", err)
		return &proto.UpdateUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	if req.ID == nil || *req.ID == "" {
		err := fmt.Errorf("invalid id")
		logger.Sugar().Errorw("UpdateUser", "error", err)
		return &proto.UpdateUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	id, err := uuid.Parse(req.GetID())
	if err != nil {
		logger.Sugar().Errorw("UpdateUser", "ID", req.GetID(), "error", err)
		return &proto.UpdateUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		err := fmt.Errorf("invalid user")
		logger.Sugar().Errorw("UpdateUser", "error", err)
		return &proto.UpdateUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	if info == nil {
		err := fmt.Errorf("invalid user")
		logger.Sugar().Errorw("UpdateUser", "error", err)
		return &proto.UpdateUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	reqInfo := &crud.UserReq{
		ID:     req.ID,
		Remark: req.Remark,
	}
	if req.Password != nil && *req.Password != "" {
		reqPasswdInfo, err := setPasswd(req)
		if err != nil {
			logger.Sugar().Errorw("UpdateUser", "error", err)
			return &proto.UpdateUserResponse{}, status.Error(codes.Internal, err.Error())
		}
		reqInfo.Password = reqPasswdInfo.Password
		reqInfo.Salt = reqPasswdInfo.Salt
	}

	info, err = crud.Update(ctx, reqInfo)
	if err != nil {
		logger.Sugar().Errorw("UpdateUser", "error", err)
		return &proto.UpdateUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	if req.Remark != nil || req.Password != nil {
		// 设置token:用name+logintime随机计算
		// 设置redis session
		var token string
		err = ctredis.EnableNilGet(fmt.Sprintf("%v/%v", PrefixLogin, info.ID.String()), &token)
		if err != nil {
			if err.Error() != "nil" {
				logger.Sugar().Errorw("UpdateUser", "ID", req.GetID(), "error", err)
				err = fmt.Errorf("invalid update")
				return &proto.UpdateUserResponse{}, status.Error(codes.Internal, err.Error())
			}
		}

		if token != "" {
			err = ctredis.Del(fmt.Sprintf("%v/%v", PrefixLogin, info.ID.String()))
			if err != nil {
				logger.Sugar().Errorw("UpdateUser", "ID", req.GetID(), "error", err)
				err = fmt.Errorf("invalid update")
				return &proto.UpdateUserResponse{}, status.Error(codes.Internal, err.Error())
			}

		}
	}

	return &proto.UpdateUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetUser(ctx context.Context, in *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetUser", "ID", in.GetID(), "error", err)
		return &proto.GetUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetUser", "ID", in.GetID(), "error", err)
		return &proto.GetUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetUsers(ctx context.Context, in *proto.GetUsersRequest) (*proto.GetUsersResponse, error) {
	var err error

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetUsers", "error", err)
		return &proto.GetUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetUsersResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteUser", "ID", in.GetID(), "error", err)
		return &proto.DeleteUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteUser", "ID", in.GetID(), "error", err)
		return &proto.DeleteUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	if info == nil {
		err := fmt.Errorf("invalid user")
		logger.Sugar().Errorw("DeleteUser", "ID", in.GetID(), "error", err)
		return &proto.DeleteUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	if info.Name == "root" {
		err := fmt.Errorf("can not delete root user")
		logger.Sugar().Errorw("DeleteUser", "ID", in.GetID(), "error", err)
		return &proto.DeleteUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err = crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteUser", "ID", in.GetID(), "error", err)
		return &proto.DeleteUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteUserResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) Login(ctx context.Context, in *proto.LoginRequest) (*proto.LoginResponse, error) {
	var err error

	if in.Name == "" {
		err := fmt.Errorf("invalid name")
		logger.Sugar().Errorw("Login", "error", err)
		return &proto.LoginResponse{}, status.Error(codes.Internal, err.Error())
	}
	if in.Password == "" {
		err := fmt.Errorf("invalid password")
		logger.Sugar().Errorw("Login", "error", err)
		return &proto.LoginResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := crud.RowByName(ctx, in.Name)
	if err != nil {
		err := fmt.Errorf("invalid user")
		logger.Sugar().Errorw("Login", "error", err)
		return &proto.LoginResponse{}, status.Error(codes.Internal, err.Error())
	}
	if info == nil {
		err := fmt.Errorf("invalid user")
		logger.Sugar().Errorw("Login", "error", err)
		return &proto.LoginResponse{}, status.Error(codes.Internal, err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(in.Password+info.Salt))
	if err != nil {
		err := fmt.Errorf("invalid password")
		logger.Sugar().Errorw("Login", "error", err)
		return &proto.LoginResponse{}, status.Error(codes.Internal, err.Error())
	}

	user := converter.Ent2Grpc(info)

	metadata, err := MetadataFromContext(ctx)
	if err != nil {
		logger.Sugar().Errorw("Login", "error", err)
		err := fmt.Errorf("invalid login")
		return &proto.LoginResponse{}, status.Error(codes.Internal, err.Error())
	}
	metadata.UserID = info.ID
	metadata.User = user

	token, err := createToken(metadata)
	if err != nil {
		return nil, err
	}

	err = ctredis.Set(fmt.Sprintf("%v/%v", PrefixLogin, info.ID.String()), token, LoginExpireTime)
	if err != nil {
		logger.Sugar().Errorw("Login", "error", err)
		err := fmt.Errorf("invalid login")
		return &proto.LoginResponse{}, status.Error(codes.Internal, err.Error())
	}
	user.Token = token

	return &proto.LoginResponse{
		Info: user,
	}, nil
}

func (s *Server) Logout(ctx context.Context, in *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	var err error

	id, err := uuid.Parse(in.GetUserID())
	if err != nil {
		logger.Sugar().Errorw("Logout", "ID", in.GetUserID(), "error", err)
		return &proto.LogoutResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("Logout", "ID", in.GetUserID(), "error", err)
		return &proto.LogoutResponse{}, status.Error(codes.Internal, err.Error())
	}
	if info == nil {
		err = fmt.Errorf("invalid user")
		logger.Sugar().Errorw("Logout", "ID", in.GetUserID(), "error", err)
		return &proto.LogoutResponse{}, status.Error(codes.Internal, err.Error())
	}

	var token string
	err = ctredis.Get(fmt.Sprintf("%v/%v", PrefixLogin, info.ID.String()), &token)
	if err != nil {
		logger.Sugar().Errorw("Logout", "ID", in.GetUserID(), "error", err)
		err = fmt.Errorf("invalid logout")
		return &proto.LogoutResponse{}, status.Error(codes.Internal, err.Error())
	}

	if token != "" {
		err = ctredis.Del(fmt.Sprintf("%v/%v", PrefixLogin, info.ID.String()))
		if err != nil {
			logger.Sugar().Errorw("Logout", "ID", in.GetUserID(), "error", err)
			err = fmt.Errorf("invalid logout")
			return &proto.LogoutResponse{}, status.Error(codes.Internal, err.Error())
		}
	}

	return &proto.LogoutResponse{}, nil
}

func (s *Server) Logined(ctx context.Context, in *proto.LoginedRequest) (*proto.LoginedResponse, error) {
	var err error

	if in.Token == "" {
		err := fmt.Errorf("invalid token")
		logger.Sugar().Errorw("Logined", "error", err)
		return &proto.LoginedResponse{}, status.Error(codes.Internal, err.Error())
	}

	if in.UserID == "" {
		err := fmt.Errorf("invalid userid")
		logger.Sugar().Errorw("Logined", "error", err)
		return &proto.LoginedResponse{}, status.Error(codes.Internal, err.Error())
	}

	id, err := uuid.Parse(in.GetUserID())
	if err != nil {
		logger.Sugar().Errorw("Logined", "ID", in.GetUserID(), "error", err)
		return &proto.LoginedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		err := fmt.Errorf("invalid user")
		logger.Sugar().Errorw("Logined", "error", err)
		return &proto.LoginedResponse{}, status.Error(codes.Internal, err.Error())
	}
	if info == nil {
		err := fmt.Errorf("invalid user")
		logger.Sugar().Errorw("Logined", "error", err)
		return &proto.LoginedResponse{}, status.Error(codes.Internal, err.Error())
	}

	var token string
	err = ctredis.Get(fmt.Sprintf("%v/%v", PrefixLogin, info.ID.String()), &token)
	if err != nil {
		logger.Sugar().Errorw("Logined", "ID", in.GetUserID(), "error", err)
		err = fmt.Errorf("invalid logined")
		return &proto.LoginedResponse{}, status.Error(codes.Internal, err.Error())
	}

	if token == "" {
		err = fmt.Errorf("invalid logined")
		logger.Sugar().Errorw("Logined", "ID", in.GetUserID(), "error", err)
		return &proto.LoginedResponse{}, status.Error(codes.Internal, err.Error())
	}

	if in.Token != token {
		err = fmt.Errorf("invalid logined")
		logger.Sugar().Errorw("Logined", "ID", in.GetUserID(), "error", err)
		return &proto.LoginedResponse{}, status.Error(codes.Internal, err.Error())
	}

	user := converter.Ent2Grpc(info)

	metadata, err := MetadataFromContext(ctx)
	if err != nil {
		logger.Sugar().Errorw("Logined", "error", err)
		err := fmt.Errorf("invalid logined")
		return &proto.LoginedResponse{}, status.Error(codes.Internal, err.Error())
	}
	metadata.UserID = info.ID
	metadata.User = user

	err = VerifyToken(metadata, token)
	if err != nil {
		logger.Sugar().Errorw("Logined", "error", err)
		err := fmt.Errorf("invalid logined")
		return &proto.LoginedResponse{}, status.Error(codes.Internal, err.Error())
	}

	err = ctredis.Set(fmt.Sprintf("%v/%v", PrefixLogin, info.ID.String()), token, LoginExpireTime)
	if err != nil {
		logger.Sugar().Errorw("Logined", "error", err)
		err := fmt.Errorf("invalid logined")
		return &proto.LoginedResponse{}, status.Error(codes.Internal, err.Error())
	}

	user.Token = token

	return &proto.LoginedResponse{
		Info: user,
	}, nil
}

func loginCheck(ctx context.Context, UserID, Token string) error {
	var err error

	if Token == "" {
		err := fmt.Errorf("invalid token")
		logger.Sugar().Errorw("Authenticate", "error", err)
		return err
	}

	id, err := uuid.Parse(UserID)
	if err != nil {
		logger.Sugar().Errorw("Authenticate", "ID", UserID, "error", err)
		return err
	}

	info, err := crud.Row(ctx, id)
	if err != nil {
		err := fmt.Errorf("invalid user: %v", id)
		logger.Sugar().Errorw("Authenticate", "error", err)
		return err
	}

	var token string
	err = ctredis.Get(fmt.Sprintf("%v/%v", PrefixLogin, info.ID.String()), &token)
	if err != nil {
		logger.Sugar().Errorw("Logined", "ID", UserID, "error", err)
		err = fmt.Errorf("invalid authenticate")
		return err
	}

	if token == "" {
		err = fmt.Errorf("invalid logined")
		logger.Sugar().Errorw("Logined", "token", token, "error", err)
		return err
	}

	metadata, err := MetadataFromContext(ctx)
	if err != nil {
		logger.Sugar().Errorw("Logined", "error", err)
		err := fmt.Errorf("invalid authenticate")
		return err
	}
	metadata.UserID = info.ID
	metadata.User = converter.Ent2Grpc(info)

	if Token != token {
		err = fmt.Errorf("invalid logined")
		logger.Sugar().Errorw("Logined", "input token", Token, "session token", token, "error", err)
		return err
	}

	err = VerifyToken(metadata, token)
	if err != nil {
		logger.Sugar().Errorw("Logined", "error", err)
		err := fmt.Errorf("invalid authenticate")
		return err
	}

	err = ctredis.Set(fmt.Sprintf("%v/%v", PrefixLogin, info.ID.String()), token, LoginExpireTime)
	if err != nil {
		logger.Sugar().Errorw("Logined", "error", err)
		err := fmt.Errorf("invalid authenticate")
		return err
	}

	return nil
}

func (s *Server) Authenticate(ctx context.Context, in *proto.AuthenticateRequest) (*proto.AuthenticateResponse, error) {
	fmt.Println("Authenticate==========")
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for key, item := range md {
			fmt.Println("md------: ", key, " ------value: ", item)
		}
	}
	if !ok {
		fmt.Println("MATEDATA== NULL")
	}
	authResult := true
	err := loginCheck(ctx, *in.UserID, in.Token)
	if err != nil {
		logger.Sugar().Errorw("Authenticate", "error", err)
		authResult = false
	}

	return &proto.AuthenticateResponse{
		Info: authResult,
	}, nil
}

type AuthCodeRequestData struct {
	GrantCode  string `json:"grantCode"`
	RememberMe bool   `json:"rememberMe"`
	TenantID   string `json:"tenantId"`
	AppID      string `json:"appId"`
}

type AuthCodeResponseData struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		LoginToken struct {
			MaxAge   uint32 `json:"maxAge"`
			Name     string `json:"name"`
			HttpOnly bool   `json:"httpOnly"`
			Secure   bool   `json:"secure"`
			Value    string `json:"value"`
		} `json:"ennUnifiedAuthorizationCookie"`
		ShortToken string `json:"ennUnifiedCsrfToken"`
	} `json:"data"`
}

const tokenApiURL = "/unify/auth/authentication"

func getAuthToken(authCode, authTenantID string) (string, string, error) {
	appID := os.Getenv("SSO_APP_ID")
	if appID == "" {
		err := fmt.Errorf("invalid appID")
		logger.Sugar().Errorw("AuthLogin", "error", err)
		return "", "", fmt.Errorf("invalid appID")
	}
	data := AuthCodeRequestData{
		GrantCode:  authCode,
		RememberMe: false,
		TenantID:   authTenantID,
		AppID:      appID,
	}

	// 将结构体编码为JSON
	jsonValue, err := json.Marshal(data)
	if err != nil {
		err := fmt.Errorf("error marshaling JSON: %v", err)
		logger.Sugar().Errorw("AuthLogin", "error", err)
		return "", "", err
	}

	tokenApiDomain := os.Getenv("SSO_TOKEN_API_DOMAIN")
	if tokenApiDomain == "" {
		err := fmt.Errorf("get null sso api")
		logger.Sugar().Errorw("AuthLogin", "error", err)
		return "", "", err
	}
	fmt.Println("tokenApiDomain: ", tokenApiDomain)
	req, err := http.NewRequest("POST", tokenApiDomain+tokenApiURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		logger.Sugar().Errorw("AuthLogin", "error", err)
		return "", "", err
	}

	// 设置请求头，告诉API我们发送的是JSON格式的数据
	req.Header.Set("Content-Type", "application/json")

	// 创建HTTP客户端
	client := &http.Client{}

	fmt.Println("req: ", req)
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		err := fmt.Errorf("error sending request: %v", err)
		return "", "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("error reading response body: %v", err)
		return "", "", err
	}

	// 打印响应状态码和响应体
	if strings.TrimSpace(resp.Status) != "200" {
		fmt.Println("resp.Status: ", resp.Status)
		fmt.Println("resp.Body: ", resp.Body)
		fmt.Println("resp: ", resp)
		return "", "", fmt.Errorf("authLogin failed: %v", resp.Status)
	}

	var responseData AuthCodeResponseData

	// 解析JSON响应
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		err := fmt.Errorf("error parsing JSON response: %v", err)
		return "", "", err
	}

	if !responseData.Success {
		err := fmt.Errorf("auth Failed: %v", responseData.Code)
		return "", "", err
	}

	return responseData.Data.LoginToken.Value, responseData.Data.ShortToken, nil
}

type AuthUserRequestData struct{}

type AuthUserResponseData struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		UserID       string `json:"userId"`
		Username     string `json:"username"`
		NickName     string `json:"nickName"`
		HasTenant    bool   `json:"hasTenant"`
		TenantID     string `json:"tenantId"`
		TerminalType string `json:"terminalType"`
	} `json:"data"`
}

const authUserApiURL = "/unify/auth/userInfoByCurrentToken"

func getAuthUser(loginToken, shortToken string) (string, error) {
	data := AuthUserRequestData{}

	// 将结构体编码为JSON
	jsonValue, err := json.Marshal(data)
	if err != nil {
		err := fmt.Errorf("error marshaling JSON: %v", err)
		logger.Sugar().Errorw("AuthLogin", "error", err)
		return "", err
	}

	authUserDomain := os.Getenv("SSO_AUTH_USER_API_DOMAIN")
	if authUserDomain == "" {
		err := fmt.Errorf("get null sso authuser api")
		logger.Sugar().Errorw("AuthLogin", "error", err)
		return "", err
	}

	// 创建HTTP的POST请求
	req, err := http.NewRequest("POST", authUserDomain+authUserApiURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		err := fmt.Errorf("error creating request: %v", err)
		logger.Sugar().Errorw("AuthLogin", "error", err)
		return "", err
	}

	// 设置请求头，告诉API我们发送的是JSON格式的数据
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ennunifiedcsrftoken", shortToken)
	req.Header.Set("ennunifiedauthorization", loginToken)

	// 创建HTTP客户端
	client := &http.Client{}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		err := fmt.Errorf("error sending request: %v", err)
		logger.Sugar().Errorw("AuthLogin", "error", err)
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("error reading response body: %v", err)
		logger.Sugar().Errorw("AuthLogin", "error", err)
		return "", err
	}

	// 打印响应状态码和响应体
	if strings.TrimSpace(resp.Status) != "200" {
		err := fmt.Errorf("auth Failed: %v", resp.Status)
		logger.Sugar().Errorw("AuthLogin", "error", err)
		return "", err
	}

	var responseData AuthUserResponseData

	// 解析JSON响应
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		err := fmt.Errorf("error parsing JSON response: %v", err)
		logger.Sugar().Errorw("AuthLogin", "error", err)
		return "", err
	}

	if !responseData.Success {
		err := fmt.Errorf("auth Failed")
		logger.Sugar().Errorw("AuthLogin", "error", err)
		return "", err
	}

	return responseData.Data.Username, nil
}

func (s *Server) AuthLogin(ctx context.Context, in *proto.AuthLoginRequest) (*proto.AuthLoginResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for key, item := range md {
			fmt.Println("md------: ", key, " ------value: ", item)
		}
	}
	if !ok {
		fmt.Println("MATEDATA== NULL")
	}
	if in.AuthCode == "" {
		err := fmt.Errorf("invalid token")
		logger.Sugar().Errorw("AuthLogin", "error", err)
		return &proto.AuthLoginResponse{}, status.Error(codes.Internal, err.Error())
	}

	if in.AuthTenantID == "" {
		err := fmt.Errorf("invalid userid")
		logger.Sugar().Errorw("AuthLogin", "error", err)
		return &proto.AuthLoginResponse{}, status.Error(codes.Internal, err.Error())
	}

	// 请求统一认证API授权接口获取长短Token
	longToken, shortToken, err := getAuthToken(in.AuthCode, in.AuthTenantID)
	if err != nil {
		logger.Sugar().Errorw("AuthLogin", "error", err)
		err := fmt.Errorf("invalid get token")
		return &proto.AuthLoginResponse{}, status.Error(codes.Internal, err.Error())
	}

	// 用长短Token获取用户信息
	authUserName, err := getAuthUser(longToken, shortToken)
	if err != nil {
		logger.Sugar().Errorw("AuthLogin", "error", err)
		err := fmt.Errorf("invalid get auth user info")
		return &proto.AuthLoginResponse{}, status.Error(codes.Internal, err.Error())
	}

	// 检查用户信息是否存在表中，存在则直接返回，不存在则创建一条记录并返回
	info, err := crud.RowByName(ctx, authUserName)
	if err != nil {
		if !ent.IsNotFound(err) {
			logger.Sugar().Errorw("AuthLogin", "error", err)
			return &proto.AuthLoginResponse{}, status.Error(codes.Internal, err.Error())
		}
	}

	if info == nil {
		// 新增user
		passwd := uuid.NewString()
		remark := "Auth Login User"
		info, err = crud.Create(ctx, &crud.UserReq{
			Name:     &authUserName,
			Password: &passwd,
			Salt:     &passwd,
			Remark:   &remark,
		})
		if err != nil {
			logger.Sugar().Errorw("AuthLogin", "error", err)
			return &proto.AuthLoginResponse{}, status.Error(codes.Internal, err.Error())
		}
	}

	user := converter.Ent2Grpc(info)

	metadata, err := MetadataFromContext(ctx)
	if err != nil {
		logger.Sugar().Errorw("Login", "error", err)
		err := fmt.Errorf("invalid login")
		return &proto.AuthLoginResponse{}, status.Error(codes.Internal, err.Error())
	}
	metadata.UserID = info.ID
	metadata.User = user

	token, err := createToken(metadata)
	if err != nil {
		return nil, err
	}

	err = ctredis.Set(fmt.Sprintf("%v/%v", PrefixLogin, info.ID.String()), token, LoginExpireTime)
	if err != nil {
		logger.Sugar().Errorw("Login", "error", err)
		err := fmt.Errorf("invalid login")
		return &proto.AuthLoginResponse{}, status.Error(codes.Internal, err.Error())
	}
	user.Token = token

	return &proto.AuthLoginResponse{
		Info: user,
	}, nil
}
