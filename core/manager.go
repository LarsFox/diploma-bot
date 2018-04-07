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
	adminID  int
	delays   map[int]int64
	langs    map[int]string
	ticket   int
	tickets  map[int]int
	tgClient tg.Client
}

// NewManager returns a new manager
func NewManager(tgClient tg.Client, adminID int) Manager {
	return &manager{
		adminID:  adminID,
		delays:   make(map[int]int64),
		langs:    map[int]string{adminID: "ru"},
		tickets:  make(map[int]int),
		tgClient: tgClient,
	}
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
