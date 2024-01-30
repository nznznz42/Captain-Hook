package main

func main() {
	var vals = readConfigFile("example.toml")

	body := vals.readBody()

	println(body)
}
