package remind_command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/madflojo/tasks"
	"gobot/task_scheduler"
	"strconv"
	"time"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "remind",
			Description: "Remind you of something",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "message",
					Description: "The message to remind you of",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "jours",
					Description: "The time to remind you",
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "heures",
					Description: "The time to remind you",
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "minutes",
					Description: "The time to remind you",
				},
			},
		},
	}
)

func RemindCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	message := i.ApplicationCommandData().Options[0].StringValue()

	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	fmt.Println(optionMap)

	var days int
	var hours int
	var minutes int
	if optionMap["jours"] != nil {
		days, _ = strconv.Atoi(optionMap["jours"].StringValue())
	} else {
		days = 0
	}

	if optionMap["heures"] != nil {
		hours, _ = strconv.Atoi(optionMap["heures"].StringValue())
	} else {
		hours = 0
	}

	if optionMap["minutes"] != nil {
		minutes, _ = strconv.Atoi(optionMap["minutes"].StringValue())
	} else {
		minutes = 0
	}

	fmt.Println(fmt.Sprintf("[%s] New Reminder created", time.Now().Format("2006-01-02 15:04:05")))

	_, err := task_scheduler.TaskScheduler.Add(
		&tasks.Task{
			RunOnce:  true,
			Interval: time.Duration(days)*24*time.Hour + time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute,
			TaskFunc: func() error {
				s.ChannelMessageSend(i.Interaction.ChannelID, "Reminder: "+message)
				return nil
			},
		},
	)

	if err != nil {
		fmt.Println(err)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "An error occurred",
			},
		})
		return
	}

	t := time.Duration(days)*24*time.Hour + time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Je te rappellerai dans %s", t.String()),
		},
	})
}

var (
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"remind": RemindCommandHandler,
	}
)
