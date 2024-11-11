package main

import (
	"os"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	app = application{}

	os.Exit(m.Run())
}
