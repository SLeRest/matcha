package main

func main() {
	config := config.GetConfig()
	Initialize(config)
	app.Run(":80")
}
