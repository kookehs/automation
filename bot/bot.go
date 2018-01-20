package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/kookehs/exp/api/complimentr"
	"github.com/kookehs/exp/api/gdax"
	"github.com/kookehs/exp/api/mattbas"
)

var (
	Auth map[string]string

	Magic8Ball = []string{
		"It is certain.", "It is decidely so.", "Without a doubt.", "Yes definitely.", "You may rely on it.",
		"As I see it; yes.", "Most likely", "Outlook good", "Yes.", "Signs point to yes.",
		"Reply hazy try again.", "Ask again later.", "Better not tell you now.", "Cannot predict now.", "Concentrate and ask again.",
		"Don't count on it.", "My reply is no.", "My sources say no.", "Outlook not so good.", "Very doubtful.",
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
	LoadJson("./auth.json", &Auth)
}

func main() {
	token := Auth["discord"]
	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Println("Error creating Discord session: ", err)
		return
	}

	discord.AddHandler(MessageCreate)

	if err = discord.Open(); err != nil {
		log.Println("Error opening connection: ", err)
		return
	}

	defer discord.Close()
	log.Println("Jarvis is online")
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func ExecuteCommand(s string) string {
	switch s {
	case "bitcoinPrice":
		return gdax.BTCUSD + "\n" + gdax.GetProductTicker(gdax.BTCUSD).String()
	case "magic8Ball":
		return Magic8Ball[rand.Intn(len(Magic8Ball))]
	default:
		return ""
	}
}

func GetChannel(c string, s *discordgo.Session) *discordgo.Channel {
	channel, err := s.Channel(c)

	if err != nil {
		log.Println("Error retreiving channel: ", err)
		return nil
	}

	return channel
}

func Mentioned(s string, u []*discordgo.User) bool {
	mentioned := false

	for _, user := range u {
		if user.ID == s {
			mentioned = true
			break
		}
	}

	return mentioned
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	log.Println(m.Author.ID + ": " + m.Content)

	if c := GetChannel(m.ChannelID, s); c != nil && !strings.Contains(c.Name, "crypto") {
		return
	}

	if !Mentioned(s.State.User.ID, m.Mentions) {
		return
	}

	sanitized := strings.ToLower(strings.Replace(m.ContentWithMentionsReplaced(), "@Jarvis", "", -1))
	message := ExecuteCommand(Classify(sanitized))

	if message == "" {
		sentiment := Sentiment(sanitized)

		if sentiment < -1 {
			message = mattbas.GetInsult()
		} else if sentiment > 1 {
			message = complimentr.GetCompliment()
		}
	}

	SendMessage(m.ChannelID, message, s)
}

func SendMessage(c, m string, s *discordgo.Session) {
	if len(m) == 0 {
		return
	}

	if _, err := s.ChannelMessageSend(c, m); err != nil {
		log.Println("Error sending message: ", err)
		return
	}
}
