package services

import(
	"time"
    "fmt"
	"github.com/dgrijalva/jwt-go"
)


var jwtKey=[]byte("Secret_key")

type JWTClaims struct{
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string)(string,error){
	expirationTime:=time.Now().Add(24*time.Hour)
	claims:=&JWTClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Printf("%v",token)
	return token.SignedString(jwtKey)
}