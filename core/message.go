package core

import (
	"errors"
	"fmt"
	"time"

	"github.com/LarsFox/diploma-bot/tg"
)

const (
	stateNothing = iota
	stateGettingPosts
	stateCheckingPost
	stateTranslators
	stateEditors
	stateCategories
)

// handleMsg is the main function for handling messages
func (m *manager) handleMsg(msg *tg.Message) {
	if msg == nil {
		m.SendError(msg.Chat.ID, errors.New("Nil pointer on msg"))
		return
	}

	isProblem := stringInArray(msg.Text, problems) != -1
	isPrbReq := stringInArray(msg.Text, problemsReqs) != -1

	if isProblem || isPrbReq {
		if delay, ok := m.delays[msg.Chat.ID]; ok {
			if delay > time.Now().Unix() {
				if stringInArray(msg.Text, ruProblems) != -1 {
					m.tgClient.SendMessage(msg.Chat.ID, replyRuHelpComing, markdown, keyboardRuMenu)
				} else {
					m.tgClient.SendMessage(msg.Chat.ID, replyEnHelpComing, markdown, keyboardEnMenu)
				}
				return
			}

		} else if isProblem {
			m.delays[msg.Chat.ID] = time.Now().Unix() + 300
		}
	}

	if text, ok := msgReplies[msg.Text]; ok {
		keyboard := msgKeyboards[msg.Text]
		m.tgClient.SendMessage(msg.Chat.ID, text, markdown, keyboard)
		return
	}

	reply := fmt.Sprintf(replyDefault, msg.Chat.FirstName, msg.Chat.LastName, msg.Chat.FirstName, msg.Chat.LastName)
	m.tgClient.SendMessage(msg.Chat.ID, reply, markdown, keyboardDefault)
}
