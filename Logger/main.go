package main

func main() {
	log := CreateChainOfResp()

	//log.Log(2, "hello world")
	log.Log(Info, "Application started")
	log.Log(Debug, "Debugging application flow")
	log.Log(Error, "Critical system failure!")
}
