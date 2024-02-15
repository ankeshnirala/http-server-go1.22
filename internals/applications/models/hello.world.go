package models

type HelloWorldHandler1Request struct {
	Email string `json:"email"`
}

type Data struct {
	Message string `josn:"message"`
}
