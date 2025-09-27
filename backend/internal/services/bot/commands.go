package bot

import (
	"fmt"

	tele "gopkg.in/telebot.v4"
)

const greeting = `Хеллоу, %s! 🎧
Тут только свежие тусовки и ивенты.
Выбирай, подписывайся и туси без лишнего шума.

Жми кнопку 👇`

func (s *Service) startCommand(c tele.Context) error {
	markup := &tele.ReplyMarkup{}
	markup.Inline(
		markup.Row(tele.Btn{
			Text: "Тык в приложение",
			WebApp: &tele.WebApp{
				URL: s.webAppUrl,
			},
		}),
	)

	return c.Send(fmt.Sprintf(greeting, c.Sender().FirstName), markup)
}
