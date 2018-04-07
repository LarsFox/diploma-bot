package core

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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

	// Setting language
	if lang, ok := languages[msg.Text]; ok {
		m.langs[msg.Chat.ID] = lang
		m.tgClient.SendMessage(msg.Chat.ID, replyMsgs[lang][keyOkay], markdown, keyboards[lang][keyOkay])
		return
	}

	// Getting language
	language, ok := m.langs[msg.Chat.ID]
	if !ok || msg.Text == buttonChangeLanguage || msg.Text == cmdStart {
		name := strings.Trim(msg.Chat.FirstName+" "+msg.Chat.LastName, " ")
		reply := fmt.Sprintf(replyDefault, name, name)
		m.tgClient.SendMessage(msg.Chat.ID, reply, markdown, keyboardDefault)
		return
	}

	// Problem request
	isProblem := stringInArray(msg.Text, problems) != -1
	isPrbReq := stringInArray(msg.Text, problemsReqs) != -1
	if isProblem || isPrbReq {
		if ticket, ok := m.tickets[msg.Chat.ID]; ok {
			// If the help is on the way
			if ticket != 0 {
				m.notifyAdmin(msg, language, keyHelpComing)
				return
			}
		}
		if isProblem {
			m.notifyAdmin(msg, language, replyKeys[msg.Text])
			return
		}
	}

	// Cannot do
	if stringInArray(msg.Text, noCanDo) != -1 {
		time.Sleep(time.Second * 2)
	}

	// Reply if there is a reply
	if key, ok := replyKeys[msg.Text]; ok {
		keyboard, ok := keyboards[language][key]
		if !ok {
			keyboard = keyboards[language][keyOkay]
		}
		m.tgClient.SendMessage(msg.Chat.ID, replyMsgs[language][key], markdown, keyboard)
		return
	}

	// Stars
	if stringInArray(msg.Text, stars) != -1 {
		m.tgClient.SendMessage(msg.Chat.ID, custom[language][keyStarsReceived], markdown, keyboards[language][keyOkay])
		text := fmt.Sprintf("%d %s %s gave %s", msg.Chat.ID, msg.Chat.FirstName, msg.Chat.LastName, msg.Text)
		m.tgClient.SendMessage(m.adminID, text, markdown, nil)
		return
	}

	// Custom input
	if msg.Chat.ID != m.adminID {
		m.notifyAdmin(msg, language, keyAcknowledged)
		return
	}

	// Admin controls
	switch {
	default:
		lines := strings.Split(msg.Text, "\n")
		replyID, err := strconv.Atoi(lines[0])
		if err != nil {
			return
		}
		m.tgClient.SendMessage(replyID, strings.Join(lines[1:len(lines)], "\n"), markdown, keyboards[language][keyOkay])

	case strings.HasPrefix(msg.Text, cmdClose):
		replyID, err := strconv.Atoi(strings.Split(msg.Text, "_")[1])
		if err != nil {
			return
		}

		m.tgClient.SendMessage(replyID, custom[m.langs[replyID]][keyTicketClosed], markdown, keyboardStars)
		m.tgClient.SendMessage(m.adminID, fmt.Sprintf("%d: ✅", m.tickets[replyID]), markdown, nil)
		m.tickets[replyID] = 0
	}
}

func (m *manager) notifyAdmin(msg *tg.Message, language string, key botKey) {
	if key == keyHelpComing {
		reply := fmt.Sprintf(replyMsgs[language][key], m.tickets[msg.Chat.ID])
		m.tgClient.SendMessage(msg.Chat.ID, reply, markdown, keyboards[language][keyOkay])
		return
	}

	text := fmt.Sprintf("%d %s %s %s\n%s\n\n/close_%d", msg.Chat.ID, msg.Chat.FirstName, msg.Chat.LastName, language, msg.Text, msg.Chat.ID)
	m.tgClient.SendMessage(m.adminID, text, "", nil)

	ticket, ok := m.tickets[msg.Chat.ID]
	if !ok {
		m.tickets[msg.Chat.ID] = 0
		ticket = 0
	}
	if ticket != 0 {
		return
	}

	m.ticket++
	ticket = m.ticket
	m.tickets[msg.Chat.ID] = ticket
	reply := replyMsgs[language][key] + fmt.Sprintf(custom[language][keyTicket], ticket)
	m.tgClient.SendMessage(msg.Chat.ID, reply, markdown, keyboards[language][keyOkay])
}
