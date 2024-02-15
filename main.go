package main

func main() {
	var vals = readConfigFile("example.toml")

	req, err := vals.constructRequest()
	if err != nil {
		panic("nooo")
	}

	s := NewServer(420, "/Configs")
	s.Start()
	s.sendRequest(req)
	s.Stop()

}
