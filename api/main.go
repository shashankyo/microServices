package main

import (
	"fmt"
	"lcorequest/go/pkg/mod/github.com/dgrijalva/jwt-go@v3.2.0+incompatible"
	"log"
	"net/http"
	"os"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "super secret information")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf(("invalid signing method"))
				}
				aud := "billing.jwtgo.io"
				checkAudience := token.claims.(jwt.MapClaims).VerifyAudience(aud, false)
				if !checkAudience {
					return nil, fmt.Errorf(("invalid audienvce"))
				}
				iss := "jwtgo.io"
				checkIss := token.Claims.(jwt.MapClaims).VerifyIsUser(iss, false)
				if !checkIss {
					return nil, fmt.Errorf(("invalid iss"))
				}

				return MySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "no authorization token provided")
		}

	})
}

func handleRequest() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Printf("server")
	handleRequest()
}
