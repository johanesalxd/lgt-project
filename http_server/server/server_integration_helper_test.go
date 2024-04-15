package server_test

import (
	"io"
	"os"
	"testing"
)

func createTempFile(t testing.TB, initData string) (io.ReadWriteSeeker, func()) {
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
