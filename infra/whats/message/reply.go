package message

import (
	"github.com/Rhymen/go-whatsapp"
	"github.com/Rhymen/go-whatsapp/binary/proto"
)

// Reply send message to the sender
func Reply(text string, wac *whatsapp.Conn, msg whatsapp.TextMessage) error {
	content := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: GetRemoteJID(msg),
		},
		ContextInfo: whatsapp.ContextInfo{
			QuotedMessage: &proto.Message{
				Conversation: &msg.Text,
			},
			QuotedMessageID: msg.Info.Id,
			// Participant:     msg.Info.Source.GetParticipant(),
			Participant: GetSenderJID(msg),
		},
		Text: text,
	}
	_, err := wac.Send(content)
	return err
}
