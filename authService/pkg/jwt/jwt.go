package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mohammaderm/authService/config"
	"github.com/mohammaderm/authService/pkg/logger"
)

type JwtClaims struct {
	Email string `json:"email"`
	Id    uint64 `json:"id"`
	jwt.StandardClaims
}

type jwtPkg struct {
	logger logger.Logger
	cfg    *config.Auth
}

type JwtInterface interface {
	GeneratePairToken(id uint64, email string) (map[string]string, error)
	// TokenValidate(claims JwtClaims) (bool, error)
	RenewTokens(refreshToken string) (map[string]string, error)
}

func Newjwt(cfg *config.Auth, logger logger.Logger) JwtInterface {
	return &jwtPkg{
		logger: logger,
		cfg:    cfg,
	}
}

func (j *jwtPkg) RenewTokens(refreshToken string) (map[string]string, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(j.cfg.Secretkey), nil
	})
	if err != nil {
		j.logger.Error("invalid token")
		return nil, err
	}
	var pairToken map[string]string
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		pairToken, err = j.GeneratePairToken(claims.Id, claims.Email)
		if err != nil {
			return nil, err
		}
	} else {
		j.logger.Error("invalid token")
		return nil, err
	}
	return pairToken, nil
}

func (j *jwtPkg) GeneratePairToken(id uint64, email string) (map[string]string, error) {
	// access token
	a_Cliams := JwtClaims{
		Email: email,
		Id:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * j.cfg.ExpireTime_a).Unix(),
			Issuer:    j.cfg.Issue,
		},
	}
	jwtToken_a := jwt.NewWithClaims(jwt.SigningMethodHS256, a_Cliams)
	accessToken, err := jwtToken_a.SignedString([]byte(j.cfg.Secretkey))
	if err != nil {
		j.logger.Error("can not create access Token", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}

	// refresh token
	r_Cliams := JwtClaims{
		Email: email,
		Id:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * j.cfg.ExpireTime_r).Unix(),
			Issuer:    j.cfg.Issue,
		},
	}
	jwtToken_r := jwt.NewWithClaims(jwt.SigningMethodHS256, r_Cliams)
	refreshToken, err := jwtToken_r.SignedString([]byte(j.cfg.Secretkey))
	if err != nil {
		j.logger.Error("can not create refresh Token", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	return map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}, nil

}
