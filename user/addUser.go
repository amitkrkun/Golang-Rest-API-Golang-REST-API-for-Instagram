package user

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	auth "nimeshjohari02.com/restapi/auth"
	database "nimeshjohari02.com/restapi/database"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Users struct {
	id 			string 
	usrName		string
	Email 		string
	Password 	string
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if(r.Method == "POST") {
		conn := database.InitiateMongoClient()
		db := conn.Database("rest")
		collection := db.Collection("Users")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel() 

		r.ParseForm()
		user := Users{
			id: 		uuid.New().String(),
			usrName: 		r.Form["name"][0],
			Email: 		r.Form["email"][0],
			Password: 	r.Form["password"][0],
		}
		hash, err := auth.HashPassword(user.Password)
		if (err!=nil) {
w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(bson.M{"message": "Unable to hash password"})
			return 
		}
		result, err := collection.InsertOne(ctx, bson.M{
			"id": 		user.id,
			"name": 	user.usrName,
			"Email": 	user.Email,
			"Password": hash, 
		})
		if (err!=nil) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(bson.M{"message": "Unable to insert database"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bson.M{
			"message" : "User added successfully",
			"result" : result,
		})
} else {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(bson.M{"message": "Invalid request method"})
}
}