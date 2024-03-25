package user

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	proto "yun.tea/block/bright/proto/bright/user"
)

type Metadata struct {
	UserID    uuid.UUID
	ClientIP  net.IP
	UserAgent string
	User      *proto.User
}

const tokenAccessSecret = "KdJotrSavIOArznhirWNfTgfEblWphLqLTVv"

func MetadataFromContext(ctx context.Context) (*Metadata, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("fail get metadata")
	}

	clientIP := ""
	if forwards, ok := meta["x-forwarded-for"]; ok {
		if len(forwards) > 0 {
			clientIP = strings.Split(forwards[0], ",")[0]
		}
	}

	userAgent := ""
	if agents, ok := meta["grpcgateway-user-agent"]; ok {
		if len(agents) > 0 {
			userAgent = agents[0]
		}
	}

	return &Metadata{
		ClientIP:  net.ParseIP(clientIP),
		UserAgent: userAgent,
	}, nil
}

func (meta *Metadata) ToJWTClaims() jwt.MapClaims {
	claims := jwt.MapClaims{}

	claims["user_id"] = meta.UserID.String()
	claims["client_ip"] = meta.ClientIP
	claims["user_agent"] = meta.UserAgent

	return claims
}

func createToken(meta *Metadata) (string, error) {
	claims := meta.ToJWTClaims()
	candidate := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := candidate.SignedString([]byte(tokenAccessSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (meta *Metadata) ValidateJWTClaims(claims jwt.MapClaims) error {
	userID, ok := claims["user_id"]
	if !ok || userID.(string) != meta.UserID.String() {
		return fmt.Errorf("invalid user id")
	}
	clientIP, ok := claims["client_ip"]
	if !ok || clientIP.(string) != meta.ClientIP.String() {
		return fmt.Errorf("invalid client ip, ok=%v, client_ip=%v, meta.client_ip=%v", ok, clientIP, meta.ClientIP)
	}
	userAgent, ok := claims["user_agent"]
	if !ok || userAgent.(string) != meta.UserAgent {
		return fmt.Errorf("invalid user agent")
	}
	return nil
}

func VerifyToken(meta *Metadata, token string) error {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		tokenAccessSecret := os.Getenv("LOGIN_TOKEN_ACCESS_SECRET")
		if tokenAccessSecret == "" {
			return "", fmt.Errorf("invalid login token access secret")
		}
		return []byte(tokenAccessSecret), nil
	})
	if err != nil {
		return err
	}

	if !jwtToken.Valid {
		return fmt.Errorf("invalid token")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("type assertion fail of jwt token")
	}

	err = meta.ValidateJWTClaims(claims)
	if err != nil {
		return err
	}

	return nil
}
