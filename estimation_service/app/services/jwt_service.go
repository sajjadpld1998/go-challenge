package services

import (
	"database/sql"
	"skeleton/config"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type JWTModel struct {
	Check bool
	JWT   struct {
		Id        uuid.UUID
		UserName  string
		FullName  string
		Email     sql.NullString
		Phone     sql.NullString
		Role      string
		Avatar    sql.NullString
		Wallet    sql.NullString
		BirthDate sql.NullTime
	}
}

type accessTokenClaim struct {
	Id        string `json:"id"`
	UserName  string `json:"user_name"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Avatar    string `json:"avatar"`
	Wallet    string `json:"wallet"`
	BirthDate string `json:"birth_date"`
	jwt.StandardClaims
}

func (obj Jwt) Init(context *gin.Context) (jwtData JWTModel) {
	Authorization := context.GetHeader("Authorization")

	Authorization = strings.Trim(Authorization, " ")

	if len(Authorization) > 0 {
		claim := &accessTokenClaim{}

		var jwtKey = []byte(config.GetConfig().Jwt.Key)

		token, err := jwt.ParseWithClaims(Authorization, claim, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err == nil && token.Valid {
			requestJwtData := token.Claims.(*accessTokenClaim)

			jwtData.Check = true

			jwtData.JWT.Id = uuid.MustParse(requestJwtData.Id)
			jwtData.JWT.UserName = requestJwtData.UserName
			jwtData.JWT.FullName = requestJwtData.FullName
			jwtData.JWT.Email = sql.NullString{Valid: len(requestJwtData.Email) > 0, String: requestJwtData.Email}
			jwtData.JWT.Phone = sql.NullString{Valid: len(requestJwtData.Phone) > 0, String: requestJwtData.Phone}
			jwtData.JWT.Role = requestJwtData.Role
			jwtData.JWT.Avatar = sql.NullString{Valid: len(requestJwtData.Avatar) > 0, String: requestJwtData.Avatar}
			jwtData.JWT.Wallet = sql.NullString{Valid: len(requestJwtData.Wallet) > 0, String: requestJwtData.Wallet}

			date, error := time.Parse("2006-01-02", requestJwtData.BirthDate)
			jwtData.JWT.BirthDate = sql.NullTime{Valid: error == nil, Time: date}
		} else {
			jwtData.Check = false
		}
	} else {
		jwtData.Check = false
	}

	return
}
