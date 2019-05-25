package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
	"testing"
)

func Test(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go runServer(t, &wg)
	wg.Wait()
	sendRequest(t)
}

func runServer(t *testing.T, wg *sync.WaitGroup) {
	l, err := net.Listen("tcp", "localhost:1313")
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	h := new(handler)
	mux.Handle("/test", h)
	go func() {
		err = http.Serve(l, mux)
		if err != nil {
			t.Fatal(err)
		}
	}()
	wg.Done()
	log.Println("Server is serving ...")
}

func sendRequest(t *testing.T) {
	url := "http://localhost:1313/test"
	resp, err := http.Post(url, "application/json", bytes.NewBufferString(`{"Name":"LocTD"}`))
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("Failed to read response body:", err)
	}

	var msg Message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		t.Error("Failed to read response body:", err)
	}

	log.Println("Received response:", msg)
}

type handler struct {
}

func (handler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	errHandler := func(writer http.ResponseWriter, err error) {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err = writer.Write([]byte("failed to read body request: " + err.Error()))
		if err != nil {
			log.Println("failed to send error response:", err)
			return
		}
	}

	err := req.ParseForm()
	if err != nil {
		errHandler(writer, err)
	}
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errHandler(writer, err)
	}

	var msg Message
	err = json.Unmarshal(b, &msg)
	if err != nil {
		errHandler(writer, err)
	}

	log.Println("Received req:", msg)

	_, err = writer.Write([]byte(`{"Name": "Server"}`))
	if err != nil {
		log.Println("failed to send response:", err)
		return
	}
}

type Message struct {
	Name string
}
