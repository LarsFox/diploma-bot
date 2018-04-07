package core

import "github.com/LarsFox/diploma-bot/tg"

const keysRow = 3

var keyboardDefault = &tg.ReplyKeyboardMarkup{
	ResizeKeyboard: true,
	Keyboard: [][]tg.KeyboardButton{
		{{Text: buttonEn}, {Text: buttonRu}},
	},
}

var keyboardStars = &tg.ReplyKeyboardMarkup{
	ResizeKeyboard: true,
	Keyboard: [][]tg.KeyboardButton{
		{{Text: button5Stars}},
		{{Text: button4Stars}},
		{{Text: button3Stars}},
		{{Text: button2Stars}},
		{{Text: button1Star}},
	},
}

var keyboards = map[string]map[botKey]*tg.ReplyKeyboardMarkup{
	"en": {
		keyOkay: {
			ResizeKeyboard: true,
			Keyboard: [][]tg.KeyboardButton{
				{{Text: buttonEnNoCanFinish}},
				{{Text: buttonEnNoCanDo}},
				{{Text: buttonEnMoney}},
				{{Text: buttonChangeLanguage}},
			},
		},

		keyNoCanFinish: {
			ResizeKeyboard: true,
			Keyboard: [][]tg.KeyboardButton{
				{{Text: buttonEnPrbSecurity}},
				{{Text: buttonEnPrbUnavailable}},
				{{Text: buttonEnPrbLowCharge}},
				{{Text: buttonEnPrbAppCrash}},
				{{Text: buttonEnPrbCannotSee}},
				{{Text: buttonEnPrbOther}},
				{{Text: buttonEnCancel}},
			},
		},

		keyMoney: {
			ResizeKeyboard: true,
			Keyboard: [][]tg.KeyboardButton{
				{{Text: buttonEnMoneyFull}},
				{{Text: buttonEnCancel}},
			},
		},

		keyPrbAppCrash: {
			ResizeKeyboard: true,
			Keyboard: [][]tg.KeyboardButton{
				{{Text: buttonEnYesItWorked}},
				{{Text: buttonEnYesButNothing}},
			},
		},

		keyPrbOther: nil,
	},

	"ru": {
		keyOkay: {
			ResizeKeyboard: true,
			Keyboard: [][]tg.KeyboardButton{
				{{Text: buttonRuNoCanFinish}},
				{{Text: buttonRuNoCanDo}},
				{{Text: buttonRuMoney}},
				{{Text: buttonChangeLanguage}},
			},
		},

		keyNoCanFinish: {
			ResizeKeyboard: true,
			Keyboard: [][]tg.KeyboardButton{
				{{Text: buttonRuPrbSecurity}},
				{{Text: buttonRuPrbUnavailable}},
				{{Text: buttonRuPrbLowCharge}},
				{{Text: buttonRuPrbAppCrash}},
				{{Text: buttonRuPrbCannotSee}},
				{{Text: buttonRuPrbOther}},
				{{Text: buttonRuCancel}},
			},
		},

		keyMoney: {
			ResizeKeyboard: true,
			Keyboard: [][]tg.KeyboardButton{
				{{Text: buttonRuMoneyFull}},
				{{Text: buttonRuCancel}},
			},
		},

		keyPrbAppCrash: {
			ResizeKeyboard: true,
			Keyboard: [][]tg.KeyboardButton{
				{{Text: buttonRuYesItWorked}},
				{{Text: buttonRuYesButNothing}},
			},
		},

		keyPrbOther: nil,
	},
}
