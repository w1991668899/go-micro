package i18n

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var LanguageMap *i18n.Bundle

func Localizer(sourcePath string) {
	bundle := &i18n.Bundle{}

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile(fmt.Sprintf("%s/toolkit/i18n/en.json", sourcePath))
	bundle.MustLoadMessageFile(fmt.Sprintf("%s/toolkit/i18n/zh.json", sourcePath))

	LanguageMap = bundle
}

func LocalizerLanguage(errMsg, languageType string) string {
	if LanguageMap == nil {
		return errMsg
	}

	loc := i18n.NewLocalizer(LanguageMap, languageType)

	localized, err := loc.Localize(&i18n.LocalizeConfig{
		MessageID: errMsg,
	})
	if err != nil {
		return errMsg
	}
	return localized
	//return loc.MustLocalize(&tool_i18n.LocalizeConfig{
	//	MessageID: errMsg,
	//})
}
