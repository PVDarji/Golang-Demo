package config

//Configuration file
type Configuration struct {
	Dev struct {
		PORT     string `json:"PORT"`
		Imageurl string `json:"imageurl"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"dev"`
	Production struct {
		PORT     string `json:"PORT"`
		Imageurl string `json:"imageurl"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"production"`
}
