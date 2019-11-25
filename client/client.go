package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const BASE_URL = "https://app.easydb.io"

type DB struct {
	database string
	token    string
}

func Connect(database string, token string) *DB {
	return &DB{database: database, token: token}
}

func (db *DB) Get(key string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/database/%s/%s", BASE_URL, db.database, key), nil)
	if err != nil {
		return nil, err
	}
	resp, err := db.doHTTP(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func (db *DB) doHTTP(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	req.Header.Set("token", db.token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("HTTP Error code %d", resp.StatusCode))
	}
	return resp, nil
}

func (db *DB) Put(key string, value interface{}) error {
	requestBody, err := json.Marshal(map[string]interface{}{
		"value": value,
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/database/%s/%s", BASE_URL, db.database, key),
		bytes.NewReader(requestBody),
	)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	_, err = db.doHTTP(req)
	return err
}

func (db *DB) List() (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/database/%s", BASE_URL, db.database), nil)
	if err != nil {
		return nil, err
	}
	resp, err := db.doHTTP(req)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *DB) Delete(key string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/database/%s/%s", BASE_URL, db.database, key), nil)
	if err != nil {
		return err
	}
	_, err = db.doHTTP(req)
	return err
}
