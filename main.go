package main

func main() {
	config := getConfig()
	bot := configureBot(config)
	launchMessageProcessing(bot)
}
