package telegram

const msgHelp = `Я могу отслеживать цену и количество товара на странице.

In order to save the page, just send me al link to it.

In order to get a random page from your list, send me command /rnd.
Caution! After that, this page will be removed from your list!

Команды:
	/rnd - тест
	/rnd - тест
`

const msgHello = "Привет! 👾\n\n" + msgHelp

const (
	msgUnknownCommand = "Спокойной ночи 😘"
	msgNoSavedPages   = "You have no saved pages 🙊"
	msgSaved          = "Saved! 👌"
	msgAlreadyExists  = "You have already have this page in your list 🤗"
)
