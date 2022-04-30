package config

import "github.com/checkr/goflagr"

var Flagr goflagr.Configuration

func initFlagrConfig() {
	Flagr = goflagr.Configuration{
		Host:     mustGetString("FLAGR_HOST"),
		BasePath: mustGetString("FLAGR_BASE_PATH"),
	}
}
