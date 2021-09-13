package gosend

import "log"

var (
	Agent GoSendAgent
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("GoSend initializing.")
	Agent = NewGoSendService()
	log.Println("GoSend initialized.")
}
