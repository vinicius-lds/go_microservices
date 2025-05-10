package main

import (
	"authentication/cmd/api/data"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	repo := data.NewPostgresTestRepository(nil)
	testApp.Repo = repo
	os.Exit(m.Run())
}
