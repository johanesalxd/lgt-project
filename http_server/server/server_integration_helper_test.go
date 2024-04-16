package server_test

import (
	"os"
	"testing"
)

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("didn't expect an error but got one: %v", err)
	}
}

func createTempFile(t testing.TB, initData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("couldn't create temp file %v", err)
	}

	tmpFile.Write([]byte(initData))

	rmFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, rmFile
}
