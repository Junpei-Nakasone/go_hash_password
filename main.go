package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	LoginId  string
	Password string
}

type Users []*User

var users Users

func register(loginId, pass string) {
	/*
	   bcrypt.MinCost = 4
	   bcrypt.MaxCost = 31
	   bcrypt.DefaultCost = 10
	*/
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	users = append(users, &User{LoginId: loginId, Password: string(hash)})
}

func login(loginId, password string) {
	var hashStr = ""
	start := time.Now()
	for _, user := range users {
		if loginId == user.LoginId {
			hashStr = user.Password
			break
		}
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashStr), []byte(password))
	end := time.Now()
	fmt.Printf("%fs\t", (end.Sub(start)).Seconds())
	if err == nil {
		// 成功
		fmt.Print("Success")
	} else {
		// 失敗
		fmt.Print("Failure")
	}
	fmt.Printf("\t%s/%s\n", loginId, password)
}

func main() {
	users = Users{}

	// 登録
	register("user1", "password1")
	register("user2", "password2")
	register("user3", "password3")
	register("user4", "password4")
	register("user5", "password5")

	// 認証
	login("user1", "password1")
	login("user2", "password2")
	login("user3", "password3")
	login("user4", "password4")
	login("user5", "password5")
	login("user6", "password1")
	login("user1", "")
	login("user3", "password1")
	login("user3", "password2")
	login("user3", "password3")
	login("user3", "password4")
}
