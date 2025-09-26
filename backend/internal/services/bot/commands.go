package bot

import (
	"fmt"

	tele "gopkg.in/telebot.v4"
)

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
