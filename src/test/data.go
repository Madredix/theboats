package test

import (
	"io/ioutil"
	"os"
	"testing"
)

func GetTestData(path string, t *testing.T) (b []byte) {
	file, err := os.Open(path) // nolint:gosec
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if b, err = ioutil.ReadAll(file); err != nil {
		t.Fatal(err)
	}
	return
}
