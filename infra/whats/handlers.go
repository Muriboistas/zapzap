package whats

import (
	"strings"
	"time"

	"github.com/muriboistas/zapzap/commands"
	"github.com/muriboistas/zapzap/infra/whats/broadcast"
	"github.com/muriboistas/zapzap/infra/whats/message"

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
			commands.ParseCommand(waConn, msg)
		}
	}

	remoteIdentifier := message.GetRemoteIdentifier(msg.Info.RemoteJid)
	switch remoteIdentifier {
	case message.BroadcastMessage:
		broadcast.Active[message.GetRemoteHost(msg.Info.RemoteJid)] = msg.Info.RemoteJid
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
