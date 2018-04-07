package core

type botKey int

const (
	keyOkay botKey = iota
	keyAcknowledged
	keyContactSoon
	keyHelpComing
	keyNoCanFinish
	keyNoTasks

	keyMoney
	keyMoneyFull

	keyPrbAppCrash
	keyPrbUnavailable
	keyPrbOther

	keyTicket
	keyTicketClosed
	keyStarsReceived
)

// All replies to certain button pressed
var replyKeys = map[string]botKey{
	buttonEn:               keyOkay,
	buttonEnCancel:         keyOkay,
	buttonEnNoCanFinish:    keyNoCanFinish,
	buttonEnNoCanDo:        keyNoTasks,
	buttonEnMoney:          keyMoney,
	buttonEnMoneyFull:      keyMoneyFull,
	buttonEnPrbSecurity:    keyContactSoon,
	buttonEnPrbUnavailable: keyPrbUnavailable,
	buttonEnPrbLowCharge:   keyNoTasks,
	buttonEnPrbAppCrash:    keyPrbAppCrash,
	buttonEnPrbCannotSee:   keyNoTasks,
	buttonEnYesItWorked:    keyOkay,
	buttonEnYesButNothing:  keyContactSoon,
	buttonEnPrbOther:       keyPrbOther,

	buttonRu:               keyOkay,
	buttonRuCancel:         keyOkay,
	buttonRuNoCanFinish:    keyNoCanFinish,
	buttonRuNoCanDo:        keyNoTasks,
	buttonRuMoney:          keyMoney,
	buttonRuMoneyFull:      keyMoneyFull,
	buttonRuPrbSecurity:    keyContactSoon,
	buttonRuPrbUnavailable: keyPrbUnavailable,
	buttonRuPrbLowCharge:   keyNoTasks,
	buttonRuPrbAppCrash:    keyPrbAppCrash,
	buttonRuPrbCannotSee:   keyNoTasks,
	buttonRuPrbOther:       keyPrbOther,
	buttonRuYesItWorked:    keyOkay,
	buttonRuYesButNothing:  keyContactSoon,
}
