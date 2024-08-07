package models

import (
	"Gokedex/data"
	"os"
	"testing"
)

func TestUser(t *testing.T) {
	data.InitDB()

	u := User{
		Id:       2,
		Username: "test",
		Email:    "test@test.com",
		Password: "test",
	}

	All(t, u)
	One(t, u)
	Create(t, u)

	u.Username = "changed"

	Update(t, u)
	Delete(t, u)

	data.CloseDB()
	CleanDB(t)
}

func CleanDB(t *testing.T) {
	err := os.Remove("api.db")
	if err != nil {
		t.Log("File not found")
	}
}

func All(t *testing.T, user User) {
	allUsers, err := user.GetAll()

	if err != nil || len(allUsers) == 0 {
		t.Log(err.Error())
		t.Fail()
	}
}

func One(t *testing.T, user User) {
	err := user.Get(1)

	if err != nil || user.Id != 1 {
		t.Log(err.Error())
		t.Fail()
	}
}

func Create(t *testing.T, user User) {
	err := user.Create()

	if err != nil || user.Id != 2 {
		t.Log(err.Error())
		t.Fail()
	}
}

func Update(t *testing.T, user User) {
	changedRows, err := user.Update(user.Id)

	if err != nil || changedRows == 0 {
		t.Log(err.Error())
		t.Fail()
	}
}

func Delete(t *testing.T, user User) {
	deleted, err := user.Delete(user.Id)

	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if deleted != true {
		t.Logf("User %v could not be deleted.", user.Id)
		t.Fail()
	}
}
