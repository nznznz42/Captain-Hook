package hookcore

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	URL        string
	Conn       *websocket.Conn
	httpClient *http.Client
}

func Newclient(serverURL string) *client {
	c := &client{}
	httpClient := &http.Client{}
	c.httpClient = httpClient
	c.Conn = NewConn(serverURL)
	c.URL = readURL(c.Conn)
	return c
}

func (c *client) Stream(w io.Writer, fields []string, port int) {
	for {
		c.Read(w, fields, port)
	}
}

func (c *client) Read(w io.Writer, fields []string, port int) {
	msgType, data, err := c.Conn.ReadMessage()
	if err != nil {
		log.Fatalf("\nerror reading message from server, %v\n", err)
		return
	}
	if msgType == websocket.TextMessage {
		fmt.Fprint(w, "\n"+string(data))
	} else if msgType == websocket.BinaryMessage {
		req, err := DeserializeRequest(data)
		if err != nil {
			log.Fatalf("Unable to Deserialise Request")
		}
		fmt.Fprint(w, ReadRequestFields(fields, *req))

		forwardRequest(c, req, port)
	}
}

func forwardRequest(c *client, req *http.Request, port int) {
	req.URL, _ = url.Parse(fmt.Sprintf("http://localhost:%d", port))
	req.RequestURI = ""
	_, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatalf("\ncli could not forwards message to local server, %v", err)
	}
}

func ReadRequestFields(fields []string, req http.Request) string {
	out := ""
	r := reflect.ValueOf(req)
	for _, f := range fields {
		if r.FieldByName(f) == reflect.ValueOf(nil) {
			fmt.Printf("does not have field, %s", f)
		}
		field := fmt.Sprintf("\n%s :%v", f, r.FieldByName(f).Interface())
		out += field
	}
	return out
}

func readURL(ws *websocket.Conn) string {
	result := make(chan string, 1)
	go func() {
		msgType, data, err := ws.ReadMessage()
		if err != nil {
			log.Fatalf("error reading URL from server, %v", err)
		}
		if msgType != websocket.TextMessage {
			log.Fatalf("expected to received URL from server, go message of type %d", msgType)
		}
		result <- string(data)
		close(result)
	}()

	select {
	case url := <-result:
		return url
	case <-time.After(5 * time.Second):
		log.Fatalf("took too long to read message from server")
	}
	return ""
}

func NewConn(wsLink string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(wsLink, nil)
	if err != nil {
		log.Fatalf("error establishing websocket connection: %v", err.Error())
	}
	return ws
}
