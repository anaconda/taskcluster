//go:build (multiuser && darwin) || (multiuser && linux) || simple

package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mcuadros/go-defaults"
)

func TestInteractiveArtifact(t *testing.T) {
	setup(t)
	payload := GenericWorkerPayload{
		Command:    returnExitCode(0),
		MaxRunTime: 10,
		Features: FeatureFlags{
			Interactive: true,
		},
	}
	defaults.SetDefaults(&payload)
	td := testTask(t)

	taskID := submitAndAssert(t, td, payload, "completed", "completed")

	expectedArtifacts := ExpectedArtifacts{
		"public/logs/live_backing.log": {
			ContentType:     "text/plain; charset=utf-8",
			ContentEncoding: "gzip",
			Expires:         td.Expires,
		},
		"public/logs/live.log": {
			Extracts: []string{
				"exit 0",
				"=== Task Finished ===",
			},
			ContentType:     "text/plain; charset=utf-8",
			ContentEncoding: "gzip",
			Expires:         td.Expires,
		},
		"private/generic-worker/shell.html": {
			ContentType:      "text/html; charset=utf-8",
			SkipContentCheck: true,
		},
	}

	expectedArtifacts.Validate(t, taskID, 0)
}

func TestInteractiveCommand(t *testing.T) {
	setup(t)
	payload := GenericWorkerPayload{
		Command:    sleep(5),
		MaxRunTime: 10,
		Features: FeatureFlags{
			Interactive: true,
		},
	}
	defaults.SetDefaults(&payload)
	td := testTask(t)

	done := make(chan string, 1)
	go func() {
		done <- submitAndAssert(t, td, payload, "completed", "completed")
	}()

	// Wait for server to start
	timeout := time.After(10 * time.Second)
	tick := time.Tick(500 * time.Millisecond)

	var conn *websocket.Conn
	var err error

	for {
		select {
		case <-timeout:
			// Timeout reached
			t.Fatal("timeout waiting for server to start")
		case <-tick:
			// Try to connect to the server
			url := fmt.Sprintf("ws://localhost:%v/shell/%v", config.InteractivePort, os.Getenv("INTERACTIVE_ACCESS_TOKEN"))
			conn, _, err = websocket.DefaultDialer.Dial(url, nil)
			if err == nil {
				err = conn.WriteMessage(websocket.TextMessage, []byte("echo hello\n"))
				if err != nil {
					t.Fatalf("write error: %v", err)
				}

				var output []byte
				_, output, err = conn.ReadMessage()
				if err != nil {
					t.Fatalf("read error: %v", err)
				}
				expected := "hello\n"
				if string(output) != expected {
					t.Fatalf("unexpected output: %v\nexpected: %v", string(output), expected)
				}

				err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Closing connection"))
				if err != nil {
					t.Fatalf("Error sending WebSocket close message: %v", err)
				}

				_, _, err = conn.ReadMessage()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
						t.Fatalf("Unexpected close error: %v", err)
					}
				}

				err = conn.Close()
				if err != nil {
					t.Fatalf("Error closing WebSocket connection: %v", err)
				}

				<-done
				return
			} else {
				t.Logf("error connecting to server: %v", err)
			}
		}
	}
}

func TestInteractiveWrongSecret(t *testing.T) {
	setup(t)
	payload := GenericWorkerPayload{
		Command:    sleep(5),
		MaxRunTime: 10,
		Features: FeatureFlags{
			Interactive: true,
		},
	}
	defaults.SetDefaults(&payload)
	td := testTask(t)

	done := make(chan string, 1)
	go func() {
		done <- submitAndAssert(t, td, payload, "completed", "completed")
	}()

	// Wait for server to start
	timeout := time.After(10 * time.Second)
	tick := time.Tick(500 * time.Millisecond)

	for {
		select {
		case <-done:
			return
		case <-timeout:
			// Timeout reached, could not connect to server
			// which should be the case since we are using the wrong secret
			return
		case <-tick:
			// Try to connect to the server
			url := fmt.Sprintf("ws://localhost:%v/shell/%v", config.InteractivePort, "bad-secret")
			_, _, err := websocket.DefaultDialer.Dial(url, nil)
			if err == nil {
				t.Fatal("expected error connecting to server")
				return
			}
		}
	}
}
