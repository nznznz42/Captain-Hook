package hookcore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Request struct {
	Method  string              `json:"method"`
	URL     string              `json:"url"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
}

func newRequest(req *http.Request) Request {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("unable to read request body")
	}
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	bodyString := string(bodyBytes)

	return Request{
		Method:  req.Method,
		URL:     req.URL.String(),
		Headers: req.Header,
		Body:    bodyString,
	}
}

func SerializeRequest(req *http.Request) ([]byte, error) {
	requestData := newRequest(req)

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func DeserializeRequest(data []byte) (*http.Request, error) {
	req := &Request{}
	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(req.Method, req.URL, strings.NewReader(req.Body))
	if err != nil {
		return nil, err
	}

	for key, values := range req.Headers {
		for _, value := range values {
			httpReq.Header.Add(key, value)
		}
	}

	return httpReq, nil
}
