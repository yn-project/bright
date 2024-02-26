package user

import (
	"context"
	"fmt"
	"time"

	converter "yun.tea/block/bright/user/pkg/converter/user"
	crud "yun.tea/block/bright/user/pkg/crud/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"yun.tea/block/bright/common/logger"
	proto "yun.tea/block/bright/proto/bright/user"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"yun.tea/block/bright/common/ctredis"
	"yun.tea/block/bright/config"
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
		err := fmt.Errorf("invalid user")
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

	if Token != token {
		err = fmt.Errorf("invalid logined")
		logger.Sugar().Errorw("Logined", "input token", Token, "session token", token, "error", err)
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
	loginApis := config.GetConfig().AuthApis.LoginApis
	NoLoginApis := config.GetConfig().AuthApis.NoLoginApis
	publicApis := config.GetConfig().AuthApis.PublicApis
	fmt.Println("loginApis: ", loginApis)
	fmt.Println("NoLoginApis: ", NoLoginApis)
	fmt.Println("publicApis: ", publicApis)

	apisMap := map[string]string{}
	for _, item := range NoLoginApis {
		apisMap[item] = "NOLOGIN"
	}
	for _, item := range loginApis {
		apisMap[item] = "LOGIN"
	}
	for _, item := range publicApis {
		apisMap[item] = "PUBLICLOGIN"
	}

	if in.Resource == "" {
		err := fmt.Errorf("invalid Resource")
		logger.Sugar().Errorw("Authenticate", "error", err)
		return &proto.AuthenticateResponse{}, status.Error(codes.Internal, err.Error())
	}

	apiType, ok := apisMap[in.Resource]
	if !ok {
		err := fmt.Errorf("invalid api: %v", in.Resource)
		logger.Sugar().Errorw("Authenticate", "error", err)
		return &proto.AuthenticateResponse{}, status.Error(codes.Internal, err.Error())
	}
	fmt.Println("apiType: ", apiType)

	authResult := false
	switch apiType {
	case "NOLOGIN":
		authResult = true
	case "LOGIN":
		err := loginCheck(ctx, *in.UserID, in.Token)
		if err != nil {
			logger.Sugar().Errorw("Authenticate", "error", err)
			authResult = false
		} else {
			authResult = true
		}
	case "PUBLICLOGIN":
		authResult = true
	default:
		err := fmt.Errorf("invalid api: %v, %v", in.Resource, apiType)
		logger.Sugar().Errorw("Authenticate", "error", err)
		return &proto.AuthenticateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.AuthenticateResponse{
		Info: authResult,
	}, nil
}
