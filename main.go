package main

import (
	"log"
	"os"
	"strings"
	"time"
	modules "walletgenbot/modules"

	tele "gopkg.in/telebot.v3"
)

func main() {
	BOT_TOKEN := os.Getenv("BOT_TOKEN")

	if BOT_TOKEN == "" {
		log.Fatal("Set BOT_TOKEN in env")
	}

	allowedUpdates := []string{"message", "callback_query"}
	pref := tele.Settings{
		Token:     BOT_TOKEN,
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second, AllowedUpdates: allowedUpdates},
		ParseMode: tele.ModeMarkdownV2,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	// Global-scoped middleware:
	//b.Use(middleware.Logger())

	commands := []tele.Command{
		{Text: "/start", Description: "Wake up the bot."},
		{Text: "/generate", Description: "Generate multi-coin HD wallet."},
		{Text: "/source", Description: "Source code of bot."},
	}

	b.SetCommands(commands)

	b.Me.CanJoinGroups = false

	// Universal markup builders.
	selector := &tele.ReplyMarkup{}
	selector1 := &tele.ReplyMarkup{}
	btnLen12 := selector.Data("12 words", "length_12")
	btnLen24 := selector.Data("24 words", "length_24")
	btnCancel := selector.Data("❌ Cancel", "cancel")

	selector.Inline(
		selector.Row(btnLen12, btnLen24),
		selector.Row(btnCancel),
	)

	savedKeys := selector1.Data("✅ I have Saved", "saved")

	selector1.Inline(selector.Row(savedKeys))

	b.Handle("/start", func(c tele.Context) error {
		msgL := []string{"Hi *", c.Chat().FirstName, "*, I can help you to generate multi\\-coin HD wallet at ease\\.\n\n⚠️ Please note, I don't store your generated keys\\. Be careful and save it somewhere\\.\n\n🗑️ Please delete the message after you saved keys\\."}
		msg := strBuilder(msgL)
		return c.Reply(msg)
	})

	b.Handle("/generate", func(c tele.Context) error {
		return c.Reply("🔐 Please choose the length of your Mnemonic key\\.", selector)
	})

	b.Handle("/source", func(c tele.Context) error {
		return c.Reply("👋 Here is my source code:\n\nRepository: [GitHub](https://github.com/sanjithacks/wallet-generator-bot)", tele.NoPreview)
	})

	b.Handle(&btnLen12, func(c tele.Context) error {
		wallet, err := modules.Wallets(128)

		if err != nil {
			log.Println(err)
			c.Respond(&tele.CallbackResponse{Text: "", ShowAlert: false})
			return c.Edit("😓 Unable to generate 12 words HD wallet\\.")
		}
		msgL := []string{"*🔥 Wallet Generated:*\n\nAddress:\n`", wallet.Address, "`\n\nPrivate Key:\n`", wallet.PrivateKey, "`\n\nMnemonic Phrase:\n`", wallet.Mnemonic, "`\n\n_⚠️ Exposing private keys and mnemonic can lead to loss of funds\\. We don't store or cannot recover it for you\\._"}
		msg := strBuilder(msgL)
		c.Respond(&tele.CallbackResponse{Text: "", ShowAlert: false})
		return c.Edit(msg, selector1)
	})

	b.Handle(&btnLen24, func(c tele.Context) error {
		wallet, err := modules.Wallets(256)

		if err != nil {
			log.Println(err)
			c.Respond(&tele.CallbackResponse{Text: "", ShowAlert: false})
			return c.Edit("😓 Unable to generate 24 words HD wallet\\.")
		}
		msgL := []string{"*🔥 Wallet Generated:*\n\nAddress:\n`", wallet.Address, "`\n\nPrivate Key:\n`", wallet.PrivateKey, "`\n\nMnemonic Phrase:\n`", wallet.Mnemonic, "`\n\n_⚠️ Exposing private keys and mnemonic can lead to loss of funds\\. We don't store or cannot recover it for you\\._"}
		msg := strBuilder(msgL)
		c.Respond(&tele.CallbackResponse{Text: "", ShowAlert: false})
		return c.Edit(msg, selector1)
	})

	b.Handle(&btnCancel, func(c tele.Context) error {
		c.Respond(&tele.CallbackResponse{Text: "", ShowAlert: false})
		return c.Delete()

	})

	b.Handle(&savedKeys, func(c tele.Context) error {
		c.Respond(&tele.CallbackResponse{Text: "", ShowAlert: false})
		return c.Edit("✅ You have saved keys\\.")

	})

	b.Handle(tele.OnCallback, func(c tele.Context) error {
		data := c.Callback().Data
		println(data)
		return c.Respond(&tele.CallbackResponse{})
	})

	b.Start()

}

func strBuilder(str []string) string {
	var sb strings.Builder
	for _, s := range str {
		sb.WriteString(s)
	}
	return sb.String()
}
