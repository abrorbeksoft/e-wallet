package storage

import (
	"database/sql"
	"fmt"
	"github.com/abrorbeksoft/e-wallet.git/api/models"
	"github.com/abrorbeksoft/e-wallet.git/config"
	"time"
)

type dbRepo struct {
	session *sql.DB
}

func NewStorage(session *sql.DB) StorageI {
	return &dbRepo{
		session: session,
	}
}

func (d *dbRepo) CreateWallet(wallet *models.CreateWallet, userId string) (string, error) {
	var id string
	err:=d.session.QueryRow("INSERT INTO wallets " +
		"(id, type,amount,user_id,created_at,updated_at) " +
		"values (gen_random_uuid(),$1,0,$2,now(),now()) returning id",wallet.Type,userId).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	return id,nil
}



func (d *dbRepo) Hello(message string) string {
	var count int32
	d.session.QueryRow("SELECT count(id) FROM users").Scan(&count)
	return fmt.Sprintf("How are you %s. users in db=%d", message, count)
}

func (d *dbRepo) AllWallets(id string) []models.Wallet {
	var mywallets []models.Wallet
	rows,_:=d.session.Query("SELECT * FROM wallets where user_id=$1",id)
	for rows.Next() {
		wall:=models.Wallet{}
		err:=rows.Scan(&wall.Id,&wall.Type,&wall.Amount,&wall.UserId,&wall.CreatedAt,&wall.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		mywallets = append(mywallets,wall)
	}
	return mywallets
}

func (d *dbRepo) GetWallet(id string) *models.Wallet {
	var wallet models.Wallet
	err:=d.session.QueryRow("SELECT * FROM wallets where id=$1",id).Scan(&wallet.Id,&wallet.Type,&wallet.Amount,&wallet.UserId,&wallet.CreatedAt,&wallet.UpdatedAt)
	if err!= nil {
		fmt.Println(err)
		return nil
	}

	return &wallet

}

func (d *dbRepo) PaymentHistory(id string) *models.Payments {
	var payments []models.Payment
	location, _ := time.LoadLocation("Asia/Tashkent")
	now:=time.Now().In(location).Format("2006-01-02 15:04:05")
	monthAgo:= time.Now().AddDate(0,-1,0).In(location).Format("2006-01-02 15:04:05")

	rows,_:= d.session.Query("SELECT * FROM payments where created_at > $1 and created_at < $2 and wallet_id=$3",monthAgo, now, id)
	for rows.Next() {
		p:=models.Payment{}
		err := rows.Scan(&p.Id, &p.Amount,&p.Type,&p.WalletId, &p.CreatedAt, &p.UpdatedAt)
		if err != nil{
			fmt.Println(err)
			continue
		}
		payments = append(payments, p)
	}

	return &models.Payments{
		All: payments,
		Count: len(payments),
	}

}


func (d *dbRepo) AddMoney(id string,amount int64) (bool,string ) {
	con:=config.Load()
	//panic("implement me")
	var wallet models.Wallet;
	err:=d.session.QueryRow("SELECT * FROM wallets where id=$1",id).Scan(&wallet.Id, &wallet.Type, &wallet.Amount,&wallet.UserId,&wallet.CreatedAt,&wallet.UpdatedAt)
	if err!=nil {
		fmt.Println(err)
	}

	if wallet.Type == "unidentified" {
		if wallet.Amount+amount > con.Unidentified {
			return false, "Boshqa pul qosha olmaysiz"
		}
		d.session.Exec("UPDATE wallets SET amount=$1 where id=$2",wallet.Amount+amount, id)
		d.createPayment(&wallet,amount,"income")
		return true, "ok"
	}

	if wallet.Amount+amount > con.Identified {
		return false, "Boshqa pul qosha olmaysiz"
	}
	d.session.Exec("UPDATE wallets SET amount=$1 where id=$2",wallet.Amount+amount, id)
	d.createPayment(&wallet,amount,"income")
	return true, "ok"

}

func (d *dbRepo) RemoveMoney(id string,amount int64) (bool, string) {
	var wallet models.Wallet;
	err:=d.session.QueryRow("SELECT * FROM wallets where id=$1",id).Scan(&wallet.Id,&wallet.Type, &wallet.Amount, &wallet.UserId, &wallet.CreatedAt, &wallet.UpdatedAt)
	if err != nil {
		fmt.Println(err)
	}
	if wallet.Amount - amount < 0 {
		return false, "Hisobinggizda yetarli mablag' mavjud emas!"
	}
	_, err=d.session.Exec("UPDATE wallets SET amount=$1 where id=$2",wallet.Amount - amount, id)
	if err != nil {
		fmt.Println(err)
	}
	d.createPayment(&wallet, amount, "outcome")
	return true, "ok"
}

func (d *dbRepo) createPayment(wallet *models.Wallet, amount int64, paymentType string)  {
	location, _ := time.LoadLocation("Asia/Tashkent")
	now:=time.Now().In(location).Format("2006-01-02 15:04:05")
	_,err:=d.session.Exec("INSERT INTO payments (id,amount,type,wallet_id,created_at, updated_at) values (gen_random_uuid(),$1,$2,$3,$4,$5)",amount, paymentType, wallet.Id, now, now )
	if err != nil {
		fmt.Println(err)
	}
}