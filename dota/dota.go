package dota

import (
	"os"
)

const (
	baseURL = "http://api.steampowered.com"
	match   = "IDOTA2Match_570"
)

type DotaClient struct {
	key string
}

// NewDotaClient return a new initialised Dota client. You can specify the Steam API key by giving key or setting
// it in STEAM_KEY environment variable
//
// 	Example:
//		// Create Dota client with steam key
//		client := dota.NewDotaClient("key")
//
//		// Create Dota Client using STEAM_KEY environment variable
//		os.Setenv("STEAM_KEY", "key")
//		client := dota.NewDotaClient("")
func NewDotaClient(key string) *DotaClient {
	if key == "" {
		key = os.Getenv("STEAM_KEY")
	}

	return &DotaClient{
		key: key,
	}
}
