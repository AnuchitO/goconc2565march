package main

// TODO: Implement relaying of message with Channel Direction

func sendMessage(ch1 chan<- string) {
	// send message on ch1
}

func relayMessage(ch1 <-chan string, ch2 chan<- string) {
	// recv message on ch1
	// send it on ch2
}

func main() {
	// create ch1 and ch2

	// spine goroutine sendMessage and relayMessage

	// recv message on ch2
}
