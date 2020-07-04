package controler

import (
	"logindemo/db"
	"logindemo/model"
	"logindemo/utils"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

//SignUpUser register new user
func SignUpUser(w http.ResponseWriter, r *http.Request) {

	userdb := db.MogUserSession()

	username := r.FormValue("user_name")
	emailAddress := r.FormValue("email_address")
	contactNumber := r.FormValue("contact_number")

	if username == "" {
		db.RespondJSON(w, "Please enter user name", utils.STATUFAILS, nil)
		return
	}

	if emailAddress == "" {
		db.RespondJSON(w, "Please enter email address", utils.STATUFAILS, nil)
		return
	}

	if contactNumber == "" {
		db.RespondJSON(w, "Please enter contact number", utils.STATUFAILS, nil)
		return
	}

	var userRs model.UserModel

	err := userdb.Find(bson.M{"contact_number": contactNumber}).One(&userRs)

	if err == nil {
		db.RespondJSON(w, "Contact number alrady exited", utils.STATUFAILS, nil)
		return
	}

	err = userdb.Find(bson.M{"email_address": emailAddress}).One(&userRs)

	if err == nil {
		db.RespondJSON(w, "Email address alrady exited", utils.STATUFAILS, nil)
		return
	}

	err = userdb.Insert(&model.UserModel{UserName: username,
		EmailAddress:  emailAddress,
		ContactNumber: contactNumber})

	if err != nil {
		db.RespondJSON(w, "Something went wrong please try again", utils.STATUFAILS, nil)
	} else {
		var userRS model.UserModelID
		err = userdb.Find(bson.M{"contact_number": contactNumber}).One(&userRS)

		token, _ := db.CreateToken(userRS.ID)

		userdb.Update(bson.M{"contact_number": contactNumber}, bson.M{"$set": bson.M{"token": token}})

		err = userdb.Find(bson.M{"contact_number": contactNumber}).One(&userRS)

		db.RespondJSON(w, "Create user successfuly", utils.STATUSSUCESS, userRS)
	}

}

//LoginUser login user
func LoginUser(w http.ResponseWriter, r *http.Request) {
	userdb := db.MogUserSession()
	contactNumber := r.FormValue("contact_number")

	var userRs model.UserModel

	err := userdb.Find(bson.M{"contact_number": contactNumber}).One(&userRs)

	if err != nil {
		db.RespondJSON(w, "Please check your contact number", utils.STATUFAILS, nil)
		return
	}

	db.RespondJSON(w, "Login User", utils.STATUSSUCESS, userRs)

}

//AllUsers All Users
func AllUsers(w http.ResponseWriter, r *http.Request) {
	var userlist []bson.M
	userdb := db.MogUserSession()
	userdb.Find(nil).All(&userlist)
	db.RespondJSON(w, "All User", utils.STATUSSUCESS, userlist)
}
