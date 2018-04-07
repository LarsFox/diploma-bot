package core

const (
	errorString = "Ой! Возникла ошибка"
	markdown    = "Markdown"
)

const (
	cmdStart = "/start"
	cmdClose = "/close"
)

const (
	buttonChangeLanguage = "🇬🇧 / 🇷🇺"
	replyDefault         = "Привет, %s! Выбери язык, чтобы начать.\n\nHello, %s! Select language to start."
)

var custom = map[string]map[botKey]string{
	"en": map[botKey]string{
		keyTicket:        "\n\nYour ticket number is `%d`. To communicate with the support agent, send your messages to me :)",
		keyTicketClosed:  "A support agent has closed your ticket.\n\nPlease rate our support service.",
		keyStarsReceived: "Thanks for you feedback!",
	},
	"ru": map[botKey]string{
		keyTicket:        "\n\nНомер Вашего обращения — `%d`. Для связи с агентом поддержки пишите мне :)",
		keyTicketClosed:  "Агент поддержки закрыл Ваше обращение.\n\nПожалуйста, оцените работу нашего сервиса.",
		keyStarsReceived: "Спасибо за оценку!",
	},
}

var languages = map[string]string{
	buttonEn: "en",
	buttonRu: "ru",
}

var (
	problems = []string{
		buttonEnYesButNothing,
		buttonEnPrbCannotSee,
		buttonEnPrbSecurity,
		buttonEnPrbUnavailable,
		buttonRuYesButNothing,
		buttonRuPrbCannotSee,
		buttonRuPrbSecurity,
		buttonRuPrbUnavailable,
	}

	problemsReqs = []string{
		buttonRuNoCanFinish,
		buttonEnNoCanFinish,
	}

	noCanDo = []string{
		buttonEnNoCanDo,
		buttonEnPrbLowCharge,
		buttonEnPrbCannotSee,
		buttonRuNoCanDo,
		buttonRuPrbCannotSee,
		buttonRuPrbLowCharge,
	}

	stars = []string{
		button5Stars,
		button4Stars,
		button3Stars,
		button2Stars,
		button1Star,
	}
)
