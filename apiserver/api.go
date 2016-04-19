// Copyright 2016 The ratazana authors. All rights reserved.

package apiserver

// Application imports
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

// Structs
type Accounts struct {
	Id           string `json:"id"`
	Login        string `json:"login"`
	CreationDate string `json:"creation_date"`
	ModDate      string `json:"mod_date"`
	Status       string `json:"status"`
}

type AccessProfile struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type AccountCollection struct {
	Data []Accounts `json:"data"`
}

type Errors struct {
	Errors []*Error `json:"errors"`
}

type Error struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type DbDriver {
    Id bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
    Name string `json:"name"`
}

type DbDrivers {
    Data []DbDriver `json:"data"`
}

type DbConn struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type DbConnections struct {
	Data []DbConn `json:"data"`
}

// Global vars
var (
	ErrBadRequest     = &Error{"bad_request", 400, "Bad Request", "Invalid request body. Try using JSON"}
	ErrInternalServer = &Error{"internal_server_error", 500, "Internal server error", "Ooops! Something went wrong"}


)
