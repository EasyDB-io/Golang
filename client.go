package client

import (
	"github.com/imroc/req"
	"log"
	"net/http"
)

// DB - structure containing token and database UUID for each request
type DB struct {
	token string
	uuid  string
}

// Get - get value by key in a database
func (db *DB) Get(key string) string {

	url := "https://app.easydb.io/database/" + db.uuid + "/" + key

	header := make(http.Header)

	header.Set("token", db.token)

	r := req.New()

	res, err := r.Get(url, header)

	if err != nil {
		log.Fatalf("Error requesting database, %v", err)
	}

	if res.Response().StatusCode != 200 {
		log.Fatalf("GET Request is not ok. %v", res.Response().Status)
	}

	return res.String()
}

// Put - add new value to database
func (db *DB) Put(key string, value string) {

	url := "https://app.easydb.io/database/" + db.uuid

	headers := make(http.Header)

	headers.Set("token", db.token)

	headers.Set("content-type", "application/json")

	r := req.New()

	body := "{ key: " + key + ", value: " + value + " }"

	res, err := r.Post(url, body, headers)

	if err != nil {
		log.Fatalf("Couldn't put new value in database, %v", err)
	}

	if res.Response().StatusCode != 200 {
		log.Fatalf("POST Request is not ok. %v", res.Response().Status)
	}

}

// Delete - delete key from database
func (db *DB) Delete(key string) {
	url := "https://app.easydb.io/database/" + db.uuid

	headers := make(http.Header)

	headers.Set("token", db.token)

	headers.Set("content-type", "application/json")

	r := req.New()

	body := "{ key: " + key + " }"

	res, err := r.Delete(url, body, headers)

	if err != nil {
		log.Fatalf("Couldn't remove value from database, %v", err)
	}
	if res.Request().Response.StatusCode != 200 {
		log.Fatalf("DELETE request is not ok. %v", res.Response().Status)
	}
}

// List - list all pairs in a database
func (db *DB) List() interface{} {
	url := "https://app.easydb.io/database/" + db.uuid

	header := make(http.Header)

	header.Set("token", db.token)

	r := req.New()

	res, err := r.Get(url, header)

	if err != nil {
		log.Fatalf("Error requesting database, %v", err)
	}

	if res.Response().StatusCode != 200 {
		log.Fatalf("GET Request is not ok. %v", res.Response().Status)
	}

	return res.ToJSON()
}
