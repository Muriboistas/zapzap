package whats

import (
	"fmt"
	"strings"
	"time"

	"github.com/Rhymen/go-whatsapp"
)

var startedAt = time.Now()

type waHandler struct {
	c *whatsapp.Conn
}

//HandleError needs to be implemented to be a valid WhatsApp handler
func (h *waHandler) HandleError(err error) {}

//Optional to be implemented. Implement HandleXXXMessage for the types you need.
func (*waHandler) HandleTextMessage(message whatsapp.TextMessage) {
	msgTime := time.Unix(int64(message.Info.Timestamp), 0)
	if startedAt.Before(msgTime) {
		if strings.HasPrefix(message.Text, config.Command.Prefix) {
			fmt.Printf("RemoteJid: %v\nSource: %v\nText: %v\n", message.Info.RemoteJid, message.Info.Source.GetParticipant(), message.Text)
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
