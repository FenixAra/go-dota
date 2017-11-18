package dota

import (
	"testing"
)

func TestGetLeagueListing(t *testing.T) {
	client := NewDotaClient("")
	matchClient := NewMatchClient(client)

	leagues, err := matchClient.GetLeagueListing()
	if err != nil {
		t.Error("Unable to get league listing. Err:", err)
		t.FailNow()
	}

	if len(leagues) == 0 {
		t.Error("Unable to get league listing. Zero leagues are returned.")
		t.FailNow()
	}
}
