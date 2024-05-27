package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/AndreBalDm/go_final_project/env"
	"github.com/golang-jwt/jwt"
)

type AuthPass struct {
	Pass string `json:"password"`
}

type AuthPassError struct {
	MyTocken string `json:"token,omitempty"`
	Err      string `json:"error,omitempty"`
}

var AuthResult AuthPassError
var buf bytes.Buffer
var auth AuthPass

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jwt string // JWT-token from cookies
		// chek password
		pass := env.SetPass()
		if len(pass) > 0 {
			// get cookies
			cookie, err := r.Cookie("token")
			if err == nil {
				jwt = cookie.Value
			}
			if jwt != AuthResult.MyTocken {
				// back err 401
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}

func (srv Server) CheckPass(w http.ResponseWriter, r *http.Request) {
	//get data from the request
	_, err := buf.ReadFrom(r.Body)
	checkErr(err)
	//transfer data into a structure auth
	if err = json.Unmarshal(buf.Bytes(), &auth); err != nil {
		fmt.Println("recovery err")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//check TODO_PASSWORD anf the request
	if auth.Pass == os.Getenv("TODO_PASSWORD") {
		//do jwt token
		secret := []byte(auth.Pass)
		jwtToken := jwt.New(jwt.SigningMethodHS256)
		AuthResult.MyTocken, err = jwtToken.SignedString(secret)
		//fmt.Println("token=", AuthResult.MyTocken)
		checkErr(err)
	} else {
		AuthResult.Err = "invalid password"
	}
	//back tocken or err
	srv.Server.Response(AuthResult, w)
}
