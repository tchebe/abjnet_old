package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	pb "github.com/zjjt/abjnet/user_service/proto/user"
)

var (
	key = []byte("it's}a{fucking*Secret%lol")
)

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	repo repository
}

func newtokenService(repo repository) *TokenService {
	return &TokenService{repo}
}

//Decode -check the token and see if it is valid and retrieves the details encoded within
func (srv *TokenService) Decode(token string) (*CustomClaims, error) {
	//parse the token
	tokentype, err := jwt.ParseWithClaims(string(token), &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	//validate the token and return the customclaims
	if claims, ok := tokentype.Claims.(*CustomClaims); ok && tokentype.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

//Encode -encode the data and creates the jwt
func (srv *TokenService) Encode(user *pb.User) (string, error) {
	//get the custom time from environment
	timeenv, err := strconv.Atoi(os.Getenv("TOKENEXPIRE"))
	if err != nil {
		log.Fatal("Please check the TOKENEXPIRE environment variable")
	}
	exprireTime := time.Now().Add(time.Minute * time.Duration(timeenv)).Unix()
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: exprireTime,
			Issuer:    "abjnet.server.user",
		},
	}
	//create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
