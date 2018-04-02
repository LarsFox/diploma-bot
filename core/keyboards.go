package core

import "github.com/LarsFox/diploma-bot/tg"

const keysRow = 3

var keyboardDefault = &tg.ReplyKeyboardMarkup{
	ResizeKeyboard: true,
	Keyboard: [][]tg.KeyboardButton{
		{{Text: buttonEn}, {Text: buttonRu}},
	},
}

var (
	keyboardEnMenu = &tg.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tg.KeyboardButton{
			{{Text: buttonEnNoCanFinish}},
			{{Text: buttonEnNoCanDo}},
			{{Text: buttonEnMoney}},
			{{Text: buttonChangeLanguage}},
		},
	}

	keyboardRuMenu = &tg.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tg.KeyboardButton{
			{{Text: buttonRuNoCanFinish}},
			{{Text: buttonRuNoCanDo}},
			{{Text: buttonRuMoney}},
			{{Text: buttonChangeLanguage}},
		},
	}
)

var msgKeyboards = map[string]*tg.ReplyKeyboardMarkup{
	buttonEn:        keyboardEnMenu,
	buttonRu:        keyboardRuMenu,
	buttonEnCancel:  keyboardEnMenu,
	buttonRuCancel:  keyboardRuMenu,
	buttonEnNoCanDo: keyboardEnMenu,
	buttonRuNoCanDo: keyboardRuMenu,

	buttonRuPrbSecurity:    keyboardRuMenu,
	buttonRuPrbUnavailable: keyboardRuMenu,
	buttonRuPrbLowCharge:   keyboardRuMenu,
	buttonRuPrbCannotSee:   keyboardRuMenu,
	buttonRuYesItWorked:    keyboardRuMenu,
	buttonRuYesButNothing:  keyboardRuMenu,

	buttonEnPrbSecurity:    keyboardEnMenu,
	buttonEnPrbUnavailable: keyboardEnMenu,
	buttonEnPrbLowCharge:   keyboardEnMenu,
	buttonEnPrbCannotSee:   keyboardEnMenu,
	buttonEnYesItWorked:    keyboardEnMenu,
	buttonEnYesButNothing:  keyboardEnMenu,

	buttonEnNoCanFinish: &tg.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tg.KeyboardButton{
			{{Text: buttonEnPrbSecurity}},
			{{Text: buttonEnPrbUnavailable}},
			{{Text: buttonEnPrbLowCharge}},
			{{Text: buttonEnPrbAppCrash}},
			{{Text: buttonEnPrbCannotSee}},
			{{Text: buttonEnCancel}},
		},
	},
	buttonRuNoCanFinish: &tg.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tg.KeyboardButton{
			{{Text: buttonRuPrbSecurity}},
			{{Text: buttonRuPrbUnavailable}},
			{{Text: buttonRuPrbLowCharge}},
			{{Text: buttonRuPrbAppCrash}},
			{{Text: buttonRuPrbCannotSee}},
			{{Text: buttonRuCancel}},
		},
	},

	buttonRuMoney: &tg.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tg.KeyboardButton{
			{{Text: buttonRuMoneyFull}},
			{{Text: buttonRuCancel}},
		},
	},
	buttonEnMoney: &tg.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tg.KeyboardButton{
			{{Text: buttonEnMoneyFull}},
			{{Text: buttonEnCancel}},
		},
	},

	buttonRuPrbAppCrash: &tg.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tg.KeyboardButton{
			{{Text: buttonRuYesItWorked}},
			{{Text: buttonRuYesButNothing}},
		},
	},
	buttonEnPrbAppCrash: &tg.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tg.KeyboardButton{
			{{Text: buttonEnYesItWorked}},
			{{Text: buttonEnYesButNothing}},
		},
	},
}
