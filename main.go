package main

func main() {
	s := NewServer("Logs/reqlog.log")

	c := readConfigFile("example.toml")

	req, err := c.constructRequest()
	if err != nil {
		panic(err)
	}

	s.sendRequest(req)
	s.Stop()
}
