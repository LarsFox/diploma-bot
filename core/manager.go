package core

import (
	"log"

	"github.com/LarsFox/diploma-bot/tg"
)

// Manager works with other packages
type Manager interface {
	Listen()
}

type manager struct {
	tgClient tg.Client
	delays   map[int]int64
}

// NewManager returns a new manager
func NewManager(tgClient tg.Client) Manager {
	delays := make(map[int]int64)
	return &manager{tgClient: tgClient, delays: delays}
}

// Listen listens the incoming messages
func (m *manager) Listen() {
	log.Println("Listening...")

	/*
		TESTING AREA STARTS
	*/

	/*
		TESTING AREA ENDS
	*/

	msgChan, _ := m.tgClient.GetMessagesChan()
	for msg := range msgChan {
		if msg != nil {
			go m.handleMsg(msg)
		}
	}
}

func (m *manager) SendError(chatID int, err error) {
	log.Println(err)
	m.tgClient.SendMessage(chatID, errorString, "Markdown", nil)
	// bugsnag.Notify(err)
}
