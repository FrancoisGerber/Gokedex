package data

import (
	"os"
	"testing"
)

func TestDatabase(t *testing.T) {
	InitDB()

	if DB == nil {
		t.Fail()
	}

	CloseDB()
	CleanDB(t)
}

func CleanDB(t *testing.T) {
	err := os.Remove("api.db")
	if err != nil {
		t.Log("File not found")
	}
}
