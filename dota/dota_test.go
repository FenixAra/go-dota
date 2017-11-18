package dota

import (
	"os"
	"testing"
)

func TestNewDotaClient(t *testing.T) {
	client := NewDotaClient("key")
	if client.key != "key" {
		t.Error("The steam key is not set properly")
		t.FailNow()
	}
}

func TestNewDotaClientEnv(t *testing.T) {
	key := os.Getenv("STEAM_KEY")

	client := NewDotaClient("")
	if client.key != key {
		t.Error("The steam key is not set properly")
		t.FailNow()
	}
}
