package internationalization

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"os"
	"scm.tutorabc.com/tgo-framework/go-log"
	"scm.tutorabc.com/tgo-framework/toml"
	"strings"
)

var MsgBundle *i18n.Bundle

// go init 會直接執行
func init() {
	pwd, _ := os.Getwd()
	activePath := "internationalization/"
	if strings.Contains(pwd, "internationalization"){
		activePath = ""
	}

	MsgBundle = i18n.NewBundle(language.Chinese)
	MsgBundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// No need to load active.en.toml since we are providing default translations.
	MsgBundle.MustLoadMessageFile(activePath + "active.zh-TW.toml")
	MsgBundle.MustLoadMessageFile(activePath + "active.en-US.toml")
	//MsgBundle.MustLoadMessageFile("active.en-US.toml")
	log.Info("init MsgBundle")
}

func forTest(lang string) {
	accept := "en-US"
	localizer := i18n.NewLocalizer(MsgBundle, lang, accept)

	name := ""
	if name == "" {
		name = "Bob"
	}

	unreadEmailCount := 12

	helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "HelloPerson",
			Description: "你是誰",
			Other:       "你好 {{.Name}}",
		},
		TemplateData: map[string]string{
			"Name": name,
		},
	})

	myUnreadEmails := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "MyUnreadEmails",
			Description: "我的未讀郵件數量",
			Other:       "我有 {{.PluralCount}} 未讀的郵件",
		},
		PluralCount: unreadEmailCount,
	})

	personUnreadEmails := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "PersonUnreadEmails",
			Description: "一個人有未讀的數量",
			One:         "{{.Name}} 有 {{.UnreadEmailCount}} 未讀",
			Other:       "{{.Name}} has {{.UnreadEmailCount}} 未讀",
		},
		PluralCount: unreadEmailCount,
		TemplateData: map[string]interface{}{
			"Name":             name,
			"UnreadEmailCount": unreadEmailCount,
		},
	})

	fmt.Println(helloPerson)
	fmt.Println(myUnreadEmails)
	fmt.Println(personUnreadEmails)

	lobbyName, _ := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:          "lobbyName",
			Description: "大會堂名稱",
			One:         "",
			Other:       "Lobby Session55",
		},
	})
	fmt.Println(lobbyName)
}


func forCurrency() {
	fmt.Println("==")
	fmt.Println(currency.Symbol(currency.USD.Amount(10.1)))
	fmt.Println(currency.NarrowSymbol(currency.USD.Amount(10.1)))
	fmt.Println(currency.ISO.Kind(currency.Cash)(currency.USD.Amount(10.01)))

	fmt.Println("==")
	fmt.Println(currency.Symbol(currency.TWD.Amount(10.1)))
	fmt.Println(currency.NarrowSymbol(currency.TWD.Amount(10.1)))
	fmt.Println(currency.ISO.Kind(currency.Cash)(currency.TWD.Amount(10.01)))
}