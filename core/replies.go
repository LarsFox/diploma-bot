package core

var replyMsgs = map[string]map[botKey]string{
	"en": map[botKey]string{
		keyOkay:         "Okay! How can I help you?",
		keyAcknowledged: "Your message has been sent to a support agent.",
		keyContactSoon:  "Ok, a support agent will contact you soon.",
		keyHelpComing:   "Your ticket `#%d` has been sent, the help is on its way!",
		keyNoCanFinish:  "What happened?",
		keyNoTasks:      "It seems that your account has no reserved tasks so far! So there is no problem :)",

		keyMoney:     "Please, have a look at [the payments schedule](https://docs.google.com/spreadsheets/d/109fmy0O2hL77rxmbFKhmXiJVwiG71gfjhcGJuQz9MZc/edit?usp=sharing).\n\nTo get information about a certain payment, please use the full list.",
		keyMoneyFull: "It seems that your account has no completed payments so far! Please use our app to create some :)",

		keyPrbAppCrash:    "Have you tried rebooting the app and relogging in?",
		keyPrbUnavailable: "Ok, take several photo proofs that the shop is closed and send them as a report. You tasks will be repinned.",
		keyPrbOther:       "Please describe your problem.",
	},

	"ru": map[botKey]string{
		keyOkay:         "Окей! Чем могу помочь?",
		keyAcknowledged: "Ваше сообщение передано агенту поддержки.",
		keyContactSoon:  "Хорошо, сейчас с Вами свяжется агент поддержки.",
		keyHelpComing:   "Ваше обращение `#%d` записано, помощь уже в пути!",
		keyNoCanFinish:  "Что случилось?",
		keyNoTasks:      "Похоже, на Ваш аккаунт не зарезервировано никаких заданий! Все в порядке :)",

		keyMoney:     "Пожалуйста, ознакомьтесь [с графиком выплат по проектам](https://docs.google.com/spreadsheets/d/109fmy0O2hL77rxmbFKhmXiJVwiG71gfjhcGJuQz9MZc/edit?usp=sharing).\n\nДля уточнения информации по конкретной выплате воспользуйтесь общим списком.",
		keyMoneyFull: "Похоже, Ваш аккаунт не производил никаких выплат! Попробуйте осуществить выплату через приложение :)",

		keyPrbAppCrash:    "Вы попробовали перезагрузить приложение и перезайти в него?",
		keyPrbUnavailable: "Хорошо, сделайте пару фотоподтверждений закрытого магазина, и отправьте их в отчете. Ваши задания будут перенесены.",
		keyPrbOther:       "Пожалуйста, опишите свою проблему.",
	},
}
