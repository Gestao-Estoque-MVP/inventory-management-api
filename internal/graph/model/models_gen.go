// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Address struct {
	ID      string `json:"id"`
	Street  string `json:"street"`
	Number  string `json:"number"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zipCode"`
	UserID  int    `json:"userId"`
}

type ContactInfo struct {
	ID    string `json:"id"`
	Name  string `json:"Name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type NewAddress struct {
	Street  string `json:"street"`
	Number  string `json:"number"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zipCode"`
	UserID  int    `json:"userId"`
}

type NewContactInfo struct {
	Name  string `json:"Name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type NewPreUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type NewUser struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Password       string `json:"password"`
	DocumentType   string `json:"document_type"`
	DocumentNumber string `json:"document_number"`
}

type PreUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	Phone          string     `json:"phone"`
	Password       string     `json:"password"`
	DocumentType   string     `json:"document_type"`
	DocumentNumber string     `json:"document_number"`
	Addresses      []*Address `json:"addresses"`
}
