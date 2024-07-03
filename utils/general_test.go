package utils

import "testing"

func TestGeneral(t *testing.T) {
	setting, err := GetSetting("TokenSalt")

	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	if setting == "" {
		t.Log("Could not read setting.")
		t.Fail()
	}

	t.Log(setting)
}
