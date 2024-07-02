package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gobot/remind_command"
	"log"
	"os"
	"os/signal"
)

type Bot struct {
	botToken string
	guildId  string
	appId    string
}

func New(BotToken string, guildId string, appId string) *Bot {
	return &Bot{botToken: BotToken, guildId: guildId, appId: appId}
}

func (b *Bot) Run() {
	discord, err := discordgo.New("Bot " + b.botToken)
	checkNilErr(err)

	discord.AddHandler(newMessage)

	// Discord command registration

	_, err = discord.ApplicationCommandBulkOverwrite(b.appId, b.guildId, remind_command.Commands)
	checkNilErr(err)

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		handler, ok := remind_command.CommandHandlers[i.ApplicationCommandData().Name]
		if !ok {
			return
		}
		handler(s, i)
	})
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	discord.Open()
	defer discord.Close()

	fmt.Println("Bot running, press CTRL-C to exit.")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("Bot stopped.")
}

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message: ", e)
	}
}

func newMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

}
