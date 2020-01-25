package keyboard

import "github.com/Syfaro/telegram-bot-api"

const (
	FlipCoin     string = "Flip a coin"
	RollDie      string = "Roll a die"
	RandomNumber string = "Random number"
	MakeChoice   string = "Make a choice"
	Settings     string = "Settings"
	About        string = "About"
)

var startKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(FlipCoin),
		tgbotapi.NewKeyboardButton(RollDie),
		tgbotapi.NewKeyboardButton(RandomNumber),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(MakeChoice),
		tgbotapi.NewKeyboardButton(Settings),
		tgbotapi.NewKeyboardButton(About),
	),
)

func GetStartKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	return &startKeyboard
}
