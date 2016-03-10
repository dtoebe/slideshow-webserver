package socket

import (
	"log"
	"strconv"

	"github.com/igm/pubsub"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

//Services strusture of the services to be communicated
type Services struct {
	Service []service `json:"service"`
}

type service struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

var serv = make(chan Services)
var chat pubsub.Publisher
var message = make(chan string)
var hasMsg bool

//RecData receive data
func (s Services) RecData() string {
	hasMsg = true

	mess := "{\"status\": ["

	for i := 0; i < len(s.Service); i++ {
		mess += "{\"name\":\"" + s.Service[i].Name + "\","
		mess += "\"status\":\"" + strconv.FormatBool(s.Service[i].Status) + "\"}"
	}

	mess += "]}"

	return mess
}

//Server PubSub server to send slide show processes status
func Server(session sockjs.Session) {
	log.Println("[INF] New session established")

	closedSession := make(chan struct{})

	chat.Publish("{\"info\": \"User logged on\"}")
	defer chat.Publish("{\"info\":\"User Logged Out\"}")

	go func() {
		reader, _ := chat.SubChannel(nil)
		for {
			select {
			case <-closedSession:
				return
			case msg := <-reader:
				if err := session.Send(msg.(string)); err != nil {
					return
				}
			}
		}
	}()

	for {
		chat.Publish(<-message)
	}

}
