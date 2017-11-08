/*
Package bot contains implementation of the Calories bot.
The bot's Message handlers and functions are described in this package.
*/
package bot

import (
	"log"

	"github.com/bobheadxi/calories/facebook"
)

// Bot : The Calories bot of the app.
type Bot struct {
	API *facebook.API
}

// SetAPI : Assigns an instance of facebook.API to bot
func (b *Bot) SetAPI(api *facebook.API) {
	b.API = api
	b.API.MessageHandler = b.TestMessageReceivedAndReply
}

// TestMessageReceivedAndReply : Tests that bot receives messages and replies.
// DEPRECATE ASAP - replace with Bot handlers or something
func (b *Bot) TestMessageReceivedAndReply(event facebook.Event, sender facebook.Sender, msg facebook.ReceivedMessage) {
	b.API.SendTextMessage(sender.ID, "Hello!")
	log.Printf("Event: %+v", event)   // {ID:2066945410258565 Time:1510063491984}
	log.Printf("Sender: %+v", sender) // {ID:1657077300989984}
	log.Printf("Msg: %+v", msg)       // {ID:mid.$cAAcNE7mWyw1lyBGR51flsxJvj8_- Text:hello Seq:1028142}
}