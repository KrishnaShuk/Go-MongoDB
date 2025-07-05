package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/krishnashuk/Go-MongoDB/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":8080", r)

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb:27017") // Notice: 'mongodb' is the service name
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to MongoDB at mongodb:27017\n")
	s.SetMode(mgo.Monotonic, true)
	return s
}
