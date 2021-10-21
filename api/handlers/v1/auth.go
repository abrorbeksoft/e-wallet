package v1

import (
	"fmt"
	"github.com/abrorbeksoft/e-wallet.git/api/models"
	"github.com/abrorbeksoft/e-wallet.git/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	bcrypt "golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)
// @Tags Auth
// @Description use username and password to login
// @Summary login to the system
// @Produce json
// @Param login body models.LoginUser true "login info"
// @Success 200 {object} models.SuccessMessage
// @Success 400 {object} models.ErrorMessage
// @Router /v1/login [POST]
func (h *handlerv1) Login(c *gin.Context)  {
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{ "message" : "Username and Password is required" })
		return
	}

	var (
		user *models.User
		err error
	)
	user,err = h.Storage.GetUser(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{ "message" : "User not found" })
		return
	}

	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password)); err==nil {
		tokenString,_:=GenerateToken(user.Username,user.Id)
		c.JSON(http.StatusOK,gin.H{
			"message": "ok",
			"token": tokenString,
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message": "Username or email incorrect",
	})

}

func (h *handlerv1) Register(c *gin.Context)  {

	name:=c.PostForm("name")
	surname:=c.PostForm("surname")
	username:=c.PostForm("username")
	email:=c.PostForm("email")
	password:=c.PostForm("password")

	if name=="" || surname=="" || username=="" || email=="" || password=="" {
		c.JSON(http.StatusOK, gin.H{ "message" : "Maydonlarni toliq kiriting" })
		return
	}
	user,_:=h.Storage.GetUser(username)

	if user != nil {
		c.JSON(http.StatusOK, gin.H{ "message" : "Bunday username tarmoqda mavjud. Boshqa username tanlang" })
		return
	}

	hashedPassword,_:=bcrypt.GenerateFromPassword([]byte(password),14)
	id,err:=h.Storage.CreateUser(&models.CreateUser{
		Name: name,
		Surname: surname,
		Username: username,
		Email: email,
		Password: cast.ToString(hashedPassword),
	})
	if err!=nil {
		fmt.Println(err)
	}

	tokenString, _:=GenerateToken(username,id)
	fmt.Println(id)
	c.JSON(http.StatusOK,gin.H{
		"message":"ok",
		"token" : tokenString,
	})

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

func GenerateToken(login, id string) (string,error) {

	token:=jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["login"] = login
	claims["exp"] = time.Now().Add(time.Minute*60).Unix()

	tokenString, err := token.SignedString(config.EncryptKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}
