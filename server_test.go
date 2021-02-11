package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

// We're specifically trying to test the remote server capacity
// to handle a request larger than 32MB, which is why we're hitting
// an existing host (not just net/http/httputil)

const target = "http://localhost:8080"

// const target = ""

func TestSmallRequest(t *testing.T) {
	// 3 KB
	testRequest(t, 3*1024)
}

func TestLargeRequest(t *testing.T) {
	// 90MB
	testRequest(t, 90*1024*1024)
}

func testRequest(t *testing.T, nBytes int) {
	var bufReq bytes.Buffer
	totalWritten := 0
	sum := 0
	for totalWritten < nBytes {
		x := rand.Intn(100)
		sum += x
		n, _ := fmt.Fprintln(&bufReq, x)
		totalWritten += n
	}

	resp, err := http.Post(target, "text/plain", &bufReq)
	if err != nil {
		t.Fatal(err)
	}

	bufResp, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Error(err)
	}

	textResponse := strings.TrimSpace(string(bufResp))
	y, err := strconv.Atoi(textResponse)
	if err != nil {
		t.Errorf("Parsing response %q: %v", string(bufResp), err)
	}
	if y != sum {
		t.Errorf("Expected sum %d from server, got %d", sum, y)
	}
}
