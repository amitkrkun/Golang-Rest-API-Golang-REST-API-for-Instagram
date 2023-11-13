package main

import (
	"net/http"

	insta "nimeshjohari02.com/restapi/instagram"
	user "nimeshjohari02.com/restapi/user"
)
func main() {
	http.HandleFunc("/addInstaPost", insta.AddInstaPost)
	http.HandleFunc("/user/getUserById", user.GetUserById)
	http.HandleFunc("/user/addUser", user.AddUser)
	http.ListenAndServe(":8080", nil)
}
