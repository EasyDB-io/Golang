package client

import (
	"fmt"
	"testing"
)

const db = "8622524c-ad0d-4a30-a4db-182037ff0c0c"
const token = "a07c7cf1-7643-48bb-be7e-c14c9d3158b8"

func TestPut(t *testing.T) {
	instance := DB{uuid: db, token: token}

	instance.Put("hello", "world")
}

func TestGet(t *testing.T) {
	instance := DB{uuid: db, token: token}

	fmt.Println(instance.Get("hello"))
}

func TestDelete(t *testing.T) {
	instance := DB{uuid: db, token: token}

	instance.Delete("hello")
}

func TestList(t *testing.T) {
	instance := DB{uuid: db, token: token}

	fmt.Println(instance.List())
}
