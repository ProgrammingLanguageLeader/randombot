package locale

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const baseDir = "locale"

var supportedLanguages = []string{"en", "ru"}

var bundleInstance *i18n.Bundle

func GetBundle() *i18n.Bundle {
	if bundleInstance == nil {
		return LoadTranslations()
	}
	return bundleInstance
}

func LoadTranslations() *i18n.Bundle {
	const format = "toml"
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc(format, toml.Unmarshal)
	for _, lang := range supportedLanguages {
		path := fmt.Sprintf("%s/translate.%s.%s", baseDir, lang, format)
		bundle.MustLoadMessageFile(path)
	}
	bundleInstance = bundle
	return bundle
}

func LocalizeSimpleMessage(message *i18n.Message, lang string) string {
	bundle := GetBundle()
	localizer := i18n.NewLocalizer(bundle, lang)
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: message.ID,
	})
}
