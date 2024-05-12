package hookcore_test

import (
	"bytes"
	"fmt"
	hookcore "hooktest/hook-core"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

type localServerTestFake struct {
	req      http.Request
	received bool
	srv      *http.Server
}

func (l *localServerTestFake) Start() {
	fmt.Println(l.srv.ListenAndServe())
}

func (l *localServerTestFake) Close() {
	l.srv.Close()
}

func (l *localServerTestFake) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.req = *r
	l.received = true
}

func NewLocalServerTestFake(port string) *localServerTestFake {
	lsrv := &localServerTestFake{
		received: false,
	}
	srv := &http.Server{Addr: port, Handler: lsrv}
	lsrv.srv = srv
	return lsrv
}

type serverTestFake struct {
	ws  *websocket.Conn
	mux *http.ServeMux
}

func (s serverTestFake) WriteMessage(msg string) {
	s.ws.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (s serverTestFake) WriteEncodedRequest(body string) {
	b := bytes.NewBuffer([]byte(body))
	req, _ := http.NewRequest(http.MethodPost, "tempurl", b)
	msg, err := hookcore.SerializeRequest(req)
	if err != nil {
		log.Fatalf("unable to serialise request")
	}
	s.ws.WriteMessage(websocket.BinaryMessage, msg)
}

func (s *serverTestFake) Start() {
	http.ListenAndServe(":8080", s.mux)
}

func NewserverTestFake() *serverTestFake {
	s := &serverTestFake{}
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		ws, _ := upgrader.Upgrade(w, r, nil)
		s.ws = ws
		ws.WriteMessage(websocket.TextMessage, []byte("tempURL"))
	})
	s.mux = mux
	return s
}

func TestWhCLI(t *testing.T) {

	s := NewserverTestFake()
	go s.Start()
	t.Run("cli establishes websocket connection with the server", func(t *testing.T) {

		c := hookcore.Newclient("ws://localhost:8080/ws")
		defer c.Conn.Close()
		if c.Conn == nil {
			t.Fatal("cli didn't establish a connection with the server")
		}
	})

	t.Run("cli receives message from the server", func(t *testing.T) {

		buf := new(bytes.Buffer)
		c := hookcore.Newclient("ws://localhost:8080/ws")
		defer c.Conn.Close()
		want := "this is a temp message"
		s.WriteMessage(want)
		c.Read(buf, nil, 5555)
		if buf.String() == "" {
			t.Error("expected a message to be writtem")
		}
	})

	t.Run("cli prints the same message, in a new line it receives from the server", func(t *testing.T) {

		buf := new(bytes.Buffer)

		c := hookcore.Newclient("ws://localhost:8080/ws")
		defer c.Conn.Close()

		msg := "message sent"
		s.WriteMessage(msg)
		c.Read(buf, nil, 5555)
		want := "\n" + msg
		if buf.String() != want {
			t.Errorf("got %q, want %q", buf.String(), want)
		}
	})

	t.Run("cli prints only the specified fields of the request", func(t *testing.T) {

		lsrv := NewLocalServerTestFake(":5555")
		go lsrv.Start()
		defer lsrv.Close()

		buf := new(bytes.Buffer)

		c := hookcore.Newclient("ws://localhost:8080/ws")
		defer c.Conn.Close()

		s.WriteEncodedRequest("this is a test")
		fields := []string{"Body", "Method", "URL", "Header"}

		c.Read(buf, fields, 5555)
		got := buf.String()
		for _, field := range fields {
			if !strings.Contains(got, field) {
				t.Errorf("output does not contain field %q, got %q", field, got)
			}
		}
	})

	t.Run("client forwards the received request to locally running server", func(t *testing.T) {

		c := hookcore.Newclient("ws://localhost:8080/ws")
		defer c.Conn.Close()

		lsrv := NewLocalServerTestFake(":5555")
		go lsrv.Start()
		defer lsrv.Close()

		received := make(chan bool)

		s.WriteEncodedRequest("this is a test")

		go func() {
			buf := new(bytes.Buffer)
			c.Read(buf, []string{"Body"}, 5555)
			received <- true
		}()

		select {
		case <-received:

		case <-time.After(5 * time.Second):
			t.Error("Timeout waiting for local server to receive request")
		}

		if !lsrv.received {
			t.Error("Local server didn't receive any request")
		}
	})
}
