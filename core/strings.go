package core

const (
	errorString          = "Ой! Возникла ошибка"
	markdown             = "Markdown"
	buttonChangeLanguage = "🇬🇧 / 🇷🇺"
)

const (
	buttonEn               = "🇬🇧 English"
	buttonEnCancel         = "🔙 Never mind"
	buttonEnNoCanFinish    = "🚫 I cannot finish a task"
	buttonEnNoCanDo        = "🏢 I cannot work today"
	buttonEnMoney          = "💰 I have a question about my payments"
	buttonEnPrbSecurity    = "👮 I have a problem with a security guard"
	buttonEnPrbUnavailable = "⛔ The shop is closed"
	buttonEnPrbLowCharge   = "🔋 My battery is low"
	buttonEnPrbAppCrash    = "💥 The app has crashed"
	buttonEnYesItWorked    = "✅ Yes, and that worked"
	buttonEnYesButNothing  = "❌ Yes, but it never helped"
	buttonEnPrbCannotSee   = "📃 I cannot see any tasks available"
	buttonEnMoneyFull      = "💰 Full list of payments"
)

const (
	buttonRu               = "🇷🇺 Русский" // emoji
	buttonRuCancel         = "🔙 Назад"
	buttonRuNoCanFinish    = "🚫 Не могу завершить задание"
	buttonRuNoCanDo        = "🏢 Не могу выйти на работу"
	buttonRuMoney          = "💰 У меня вопрос по выплатам"
	buttonRuPrbSecurity    = "👮 Проблема с персоналом магазина"
	buttonRuPrbUnavailable = "⛔ Магазин закрыт"
	buttonRuPrbLowCharge   = "🔋 Садится батарея"
	buttonRuPrbAppCrash    = "💥 Приложение не работает"
	buttonRuPrbCannotSee   = "📃 Не вижу заданий"
	buttonRuYesItWorked    = "✅ Да, заработало"
	buttonRuYesButNothing  = "❌ Нет, ничего не помогло"
	buttonRuMoneyFull      = "💰 Полный список выплат"
)

const (
	replyDefault       = "Привет, %s %s! Выбери язык, чтобы начать\n\nHello, %s %s! Select language to start"
	replyEnOkay        = "Okay! How can I help you?"
	replyRuOkay        = "Окей! Чем могу помочь?"
	replyEnNoTasks     = "It seems that your account has no reserved tasks so far! So there is no problem :)"
	replyRuNoTasks     = "Похоже, на Ваш аккаунт не зарезервировано никаких заданий! Все в порядке :)"
	replyRuContactSoon = "Хорошо, сейчас с Вами свяжется агент поддержки."
	replyEnContactSoon = "Ok, a support agent will call you soon."
	replyRuHelpComing  = "Ваше обращение записано, помощь уже в пути!"
	replyEnHelpComing  = "Your ticket has been sent, the help is on its way!"
)

// All replies to certain button pressed
var msgReplies = map[string]string{
	buttonEn:               replyEnOkay,
	buttonRu:               replyRuOkay,
	buttonEnCancel:         replyEnOkay,
	buttonRuCancel:         replyRuOkay,
	buttonEnNoCanFinish:    "What happened?",
	buttonRuNoCanFinish:    "Что случилось?",
	buttonEnNoCanDo:        replyEnNoTasks,
	buttonRuNoCanDo:        replyRuNoTasks,
	buttonRuMoney:          "Пожалуйста, ознакомьтесь [с графиком выплат по проектам](https://docs.google.com/spreadsheets/d/109fmy0O2hL77rxmbFKhmXiJVwiG71gfjhcGJuQz9MZc/edit?usp=sharing).\n\nДля уточнения информации по конкретной выплате воспользуйтесь общим списком.",
	buttonEnMoney:          "Please, have a look at [the payments schedule](https://docs.google.com/spreadsheets/d/109fmy0O2hL77rxmbFKhmXiJVwiG71gfjhcGJuQz9MZc/edit?usp=sharing).\n\nTo get information about a certain payment, please use the full list.",
	buttonEnMoneyFull:      "It seems that your account has no completed payments so far! Please use our app to create some :)",
	buttonRuMoneyFull:      "Похоже, Ваш аккаунт не производил никаких выплат! Попробуйте осуществить выплату через приложение :)",
	buttonRuPrbSecurity:    replyRuContactSoon,
	buttonEnPrbSecurity:    replyEnContactSoon,
	buttonRuPrbUnavailable: "Хорошо, сделайте пару фотоподтверждений закрытого магазина, и отправьте их в отчете. Ваши задания будут перенесены.",
	buttonEnPrbUnavailable: "Ok, take several photo proofs that the shop is closed and send them as a report. You tasks will be repinned.",
	buttonRuPrbLowCharge:   replyRuNoTasks,
	buttonEnPrbLowCharge:   replyEnNoTasks,
	buttonRuPrbAppCrash:    "Попробуйте перезагрузить приложение и перезайти в него.",
	buttonEnPrbAppCrash:    "Try rebooting the app and relogging in.",
	buttonRuPrbCannotSee:   "Мы проверим закрепление. Попробуйте обновить список заданий через 5 минут.",
	buttonEnPrbCannotSee:   "We will check your reservation. Try again in 5 minutes.",
	buttonRuYesItWorked:    replyRuOkay,
	buttonEnYesItWorked:    replyEnOkay,
	buttonRuYesButNothing:  replyRuContactSoon,
	buttonEnYesButNothing:  replyEnContactSoon,
}

var problems = []string{
	buttonEnYesButNothing,
	buttonEnPrbCannotSee,
	buttonEnPrbLowCharge,
	buttonEnPrbSecurity,
	buttonEnPrbUnavailable,

	buttonRuYesButNothing,
	buttonRuPrbCannotSee,
	buttonRuPrbLowCharge,
	buttonRuPrbSecurity,
	buttonRuPrbUnavailable,
}

var ruProblems = []string{
	buttonRuNoCanFinish,
	buttonRuYesButNothing,
	buttonRuPrbCannotSee,
	buttonRuPrbLowCharge,
	buttonRuPrbSecurity,
	buttonRuPrbUnavailable,
}

var problemsReqs = []string{
	buttonRuNoCanFinish,
	buttonEnNoCanFinish,
}
