package main

import (
	"fmt"
	"github.com/abrorbeksoft/e-wallet.git/config"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

func main()  {

	tokenStr,_:=GenerateToken("username","password");
	claims, err:=ExtractClaims(tokenStr)
	if err != nil {
		fmt.Println("Error while extracting the claims")
	}
	login:=claims["login"]
	password:=claims["password"]
	fmt.Println(login)
	fmt.Println(password)

}

func ExtractClaims(tokenStr string) (jwt.MapClaims,error) {
	var (
		token *jwt.Token
		err error
	)
	token,err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return config.EncryptKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		err = fmt.Errorf("Invalid JWT Token")
		return nil, err
	}
	return claims, nil
}

func GenerateToken(login, password string) (string,error) {

	token:=jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["login"] = login
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Minute*60).Unix()

	tokenString, err := token.SignedString(config.EncryptKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}