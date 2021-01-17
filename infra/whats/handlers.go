package whats

import (
	"fmt"
	"strings"
	"time"

	"github.com/Rhymen/go-whatsapp"
)

var startedAt = time.Now()
var setted bool

type waHandler struct {
	c *whatsapp.Conn
}

//HandleError needs to be implemented to be a valid WhatsApp handler
func (h *waHandler) HandleError(err error) {}

//Optional to be implemented. Implement HandleXXXMessage for the types you need.
func (*waHandler) HandleTextMessage(msg whatsapp.TextMessage) {
	msgTime := time.Unix(int64(msg.Info.Timestamp), 0)
	if startedAt.Before(msgTime) {
		if strings.HasPrefix(msg.Text, config.Command.Prefix) {
			// commands.ParseCommand(waConn, message)
			fmt.Println(msg)
			fmt.Println("SOURCE:", msg.Info.Source.GetParticipant())
			fmt.Println("SOURCE:", msg.Info.Source.String())
			fmt.Println("SOURCE:", msg.Info.Source.GetBroadcast())
			// fmt.Println(helpers.GetSenderNumber(message))
		}
	}
}

// func (*waHandler) HandleImageMessage(message whatsapp.ImageMessage) {}

// func (*waHandler) HandleDocumentMessage(message whatsapp.DocumentMessage) {}

// func (*waHandler) HandleVideoMessage(message whatsapp.VideoMessage) {}

// func (*waHandler) HandleAudioMessage(message whatsapp.AudioMessage) {}

// func (*waHandler) HandleJsonMessage(message string) {}

// func (*waHandler) HandleContactMessage(message whatsapp.ContactMessage) {}

// func (*waHandler) HandleBatteryMessage(msg whatsapp.BatteryMessage) {}

// func (*waHandler) HandleNewContact(contact whatsapp.Contact) {}
