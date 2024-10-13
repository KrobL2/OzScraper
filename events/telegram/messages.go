package telegram

const msgHelp = `Я могу отслеживать цену и количество товара на странице.

In order to save the page, just send me al link to it.

In order to get a random page from your list, send me command /rnd.
Caution! After that, this page will be removed from your list!

Команды:
	/help - тест
	/start - тест
	/all - тест
`

const msgHello = "Привет! 👾\n\n" + msgHelp

const (
	msgUnknownCommand = "Неизестная команда 😘"
	msgNoSavedPages   = "Вы не отслеживаете товары 🙊"
	msgSaved          = "Сохранено! 👌"
	msgAlreadyExists  = "Вы уже отслеживаете этот товар 🤗"
)
