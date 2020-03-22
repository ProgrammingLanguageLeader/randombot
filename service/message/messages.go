package message

import "github.com/nicksnyder/go-i18n/v2/i18n"

const choiceMessageTemplate = `Choice between
{{.Choices}}
===============
{{.Result}}
`

var (
	chooseOptionsMessage = i18n.Message{
		ID:    "message_choose-options",
		Other: "Choose one of the following options",
	}

	unsupportedCommandMessage = i18n.Message{
		ID:    "message_unsupported-command",
		Other: "Unsupported command",
	}

	incorrectInputMessage = i18n.Message{
		ID:    "message_incorrect-input",
		Other: "Incorrect input",
	}

	dontUnderstandMessage = i18n.Message{
		ID:    "message_dont-understand",
		Other: "Sorry, I don't understand you",
	}

	somethingWentWrongMessage = i18n.Message{
		ID:    "message_something-went-wrong",
		Other: "Something went wrong... Please, try again later",
	}

	registrationMessage = i18n.Message{
		ID:    "message_register-success",
		Other: "Hello! You have successfully registered!",
	}

	flippingCoinHeadsMessage = i18n.Message{
		ID:    "message_flipping-coin-heads",
		Other: "it's heads!",
	}

	flippingCoinTailsMessage = i18n.Message{
		ID:    "message_flipping-coin-tails",
		Other: "it's tails!",
	}

	flippingCoinMessage = i18n.Message{
		ID:    "message_flipping-coin",
		Other: "You have flipped a coin: {{.Result}}",
	}

	rollingDiceMessage = i18n.Message{
		ID:    "message_rolling-dice",
		Other: "You have rolled the dice: {{.First}} and {{.Second}}",
	}

	randomizingMessage = i18n.Message{
		ID:    "message_randomizing",
		Other: "Random number from the range [{{.Min}}...{{.Max}}]: {{.Result}}",
	}

	makingChoiceMessage = i18n.Message{
		ID:    "message_making-choice",
		Other: choiceMessageTemplate,
	}

	settingsMessage = i18n.Message{
		ID:    "message_settings",
		Other: "Choose one of the following options",
	}

	helpMessage = i18n.Message{
		ID: "message_help",
		Other: "This is a Telegram bot that enables you to generate some kind of pseudorandom values. " +
			"For example, it has functionality for virtual \"rolling the dice\", \"flipping coin\", " +
			"choice between set of options, etc",
	}

	chooseLanguageMessage = i18n.Message{
		ID:    "message_choose-language",
		Other: "Choose a language",
	}

	enterMinAndMaxNumbersMessage = i18n.Message{
		ID:    "message_enter-min-and-max-numbers",
		Other: "Enter minimum and maximum numbers space separated",
	}

	enterChoiceVariantsMessage = i18n.Message{
		ID:    "message_enter-choice-variants",
		Other: "Enter the choice variants. One item - one line",
	}

	settingsWereSuccessfullyUpdated = i18n.Message{
		ID:    "message_settings-were-successfully-updated",
		Other: "Settings were successfully updated",
	}
)
