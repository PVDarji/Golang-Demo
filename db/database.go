package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

//Mongo session
var (
	Mongo *mgo.Session
)

//CheckConnection of database
func CheckConnection() bool {
	if Mongo == nil {
		Connect()
	}

	if Mongo != nil {
		return true
	}

	return false
}

//Connect to database
func Connect() {

	var err error

	if Mongo, err = mgo.Dial("127.0.0.1:27017"); err != nil {
		log.Println("MongoDB Driver Error", err)
		return
	}
	// Prevents these errors: read tcp 127.0.0.1:27017: i/o timeout
	// Mongo.SetSocketTimeout(1 * time.Hour)

	// Check if is alive
	if err = Mongo.Ping(); err != nil {
		log.Println("Database Error", err)
	}
	Mongo.SetMode(mgo.Monotonic, true)
}

//MogUserSession user db
func MogUserSession() *mgo.Collection {
	if CheckConnection() {
		session := Mongo.Copy()
		// defer session.Close()
		c := session.DB("demo").C("users")

		return c
	}
	return nil
}
