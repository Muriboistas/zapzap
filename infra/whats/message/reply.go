package message

import (
	"io"

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
			Participant:     GetSenderJID(msg),
		},
		Text: text,
	}
	_, err := wac.Send(content)
	return err
}

// ReplyImg send image to the sender
func ReplyImg(img io.Reader, wac *whatsapp.Conn, msg whatsapp.TextMessage) error {
	content := whatsapp.ImageMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: GetRemoteJID(msg),
		},
		ContextInfo: whatsapp.ContextInfo{
			QuotedMessage: &proto.Message{
				Conversation: &msg.Text,
			},
			QuotedMessageID: msg.Info.Id,
			Participant:     GetSenderJID(msg),
		},
		Type:    "image/jpeg",
		Caption: "",
		Content: img,
	}
	_, err := wac.Send(content)
	return err
}
