package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const (
	prefix   = "!uncle"
	wisdom   = "wisdom"
	haha     = "haha"
	botToken = key // For the purpose of security, the key is hidden
)

var (
	funny = []string{
		"Sick of tea?!? That's like being sick of breathing!",
		"This tea is nothing more than hot leaf juice!",
	}
	proverbs = []string{
		"History Is Not Always Kind To Its Subjects.",
		"Life Happens Wherever You Are, Whether You Make It Or Not.",
		"Destiny Is A Funny Thing. You Never Know How Things Are Going To Work Out.",
		"There Is Nothing Wrong With A Life Of Peace And Prosperity. I Suggest You Think About What It Is That You Want From Your Life.",
	}
)

func main() {
	sess, err := setupDiscordSession()
	if err != nil {
		log.Fatalf("Error setting up Discord session: %v", err)
	}

	defer sess.Close()

	fmt.Println("Uncle Iroh's wisdon is ready!")

	waitForInterrupt()
}

func setupDiscordSession() (*discordgo.Session, error) {
	sess, err := discordgo.New("Bot " + botToken)
	if err != nil {
		return sess, fmt.Errorf("failed to create Discord session: %v", err)
	}

	sess.AddHandler(messageHandler)

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		return sess, fmt.Errorf("failed to open Discord session: %v", err)
	}

	return sess, err
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	args := strings.Fields(m.Content)
	if len(args) < 2 || args[0] != prefix {
		return
	}

	sendQuote(s, m.ChannelID, args[1])
}

func sendQuote(s *discordgo.Session, channelID string, postfix string) {
	switch postfix {
	case wisdom:
		selection := rand.Intn(len(proverbs))
		s.ChannelMessageSend(channelID, proverbs[selection])
		fmt.Printf("Message sent: %s\n", proverbs[selection])
	case haha:
		selection := rand.Intn(len(funny))
		s.ChannelMessageSend(channelID, funny[selection])
		fmt.Printf("Message sent: %s\n", funny[selection])
	}
}

func waitForInterrupt() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
