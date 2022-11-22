package main

import (
	"fmt"
	"lcorequest/go/pkg/mod/github.com/dgrijalva/jwt-go@v3.2.0+incompatible"
	"lcorequest/go/pkg/mod/golang.org/x/tools@v0.1.1/go/analysis/passes/nilfunc"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)


var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func GetJWT()(string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "shashank"
	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "jetgo.io"
	claims["exp"] = time.Now().Add(time.Minute*1).Unix()

	tokenstring, err := token.SignedString(MySigningKey)

	if err != nil {
		fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	
	
	}
	return tokenstring, nil


}

func Index(w http.ResponseWriter, r *http.Request){
	validToken, err :=  GetJWT()
	fmt.Println(validToken)
	if err != nil {
		fmt.Printf("failed to generate the token")
	}
	fmt.Fprintf(w, string(validToken))
}

func main(){
	http.HandleRequests(){
		http.HandleFunc("/", Index)

		log.Fatal(http.ListenAndServe(":8080", nil))
	} 

	func main(){
		handleRequests()

	}

}