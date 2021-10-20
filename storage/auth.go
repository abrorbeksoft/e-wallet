package storage

import (
	"fmt"
	"github.com/abrorbeksoft/e-wallet.git/api/models"
)

func (d *dbRepo) CreateUser(user *models.CreateUser) (string,error) {
	var id string
	err:=d.session.QueryRow("INSERT INTO users " +
		"(id,name,surname,username,email,created_at,updated_at,password) " +
		"values (gen_random_uuid(),$1,$2,$3,$4,now(),now(),$5) returning id",
		user.Name,user.Surname,user.Username,user.Email,user.Password).Scan(&id)
	if err!=nil {
		fmt.Println(err)
		return "", err
	}
	return id, nil
}

func (d *dbRepo) UpdateUser()  {

}

func (d *dbRepo) GetUser(username string) (*models.User,error) {
	var user models.User
	err:=d.session.QueryRow("SELECT id,username,email,password from users where username=$1",username).Scan(&user.Id,&user.Username,&user.Email,&user.Password)
	if err!=nil {
		fmt.Println(err)
		return nil, err
	}
	return &user,nil
}