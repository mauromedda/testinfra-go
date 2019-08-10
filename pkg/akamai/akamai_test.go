package akamai

import (
	"log"
	"os"
	"testing"
	"time"

	mockserver "testinfra-go/pkg/akamai/tests"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()

	os.Exit(code)
}

func setup() {
	go func() {
		if err := mockserver.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	// wait for mock server to run
	time.Sleep(time.Millisecond * 10)
}

func teardown() {
	mockserver.Close()
}
