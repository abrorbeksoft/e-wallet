package main

import (
	"database/sql"
	"fmt"
	"github.com/abrorbeksoft/e-wallet.git/api"
	"github.com/abrorbeksoft/e-wallet.git/config"
	"github.com/abrorbeksoft/e-wallet.git/storage"
	_ "github.com/lib/pq"
)

func main()  {
	config:=config.Load();
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.PostgresUser,config.PostgresPassword,config.PostgresDBName)

	session, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	db:=storage.NewStorage(session)


	server:=api.New(&api.RouterOptions{
		Storage: db,
	})

	server.Run()

}


//package main
//
//import (
//	"crypto/hmac"
//	"crypto/sha1"
//	"crypto/subtle"
//	"encoding/hex"
//	"fmt"
//	"os"
//)
//
//func generateSignature(secretToken, payloadBody string) string {
//	mac := hmac.New(sha1.New, []byte(secretToken))
//	mac.Write([]byte(payloadBody))
//	expectedMAC := mac.Sum(nil)
//	return "sha1=" + hex.EncodeToString(expectedMAC)
//}
//
//func verifySignature(secretToken, payloadBody string, signatureToCompareWith string) bool {
//	signature := generateSignature(secretToken, payloadBody)
//	return subtle.ConstantTimeCompare([]byte(signature), []byte(signatureToCompareWith)) == 1
//}
//
//func main() {
//	testPayloadBody := `{"message":"test content"}`
//	testSignatureToCompareWith := `sha1=33a08e9b5e9c8d5e944d9288e9b499abb298344d`
//
//	fmt.Println("signature match? :", verifySignature(os.Getenv("SECRET_TOKEN"), testPayloadBody, testSignatureToCompareWith))
//}

