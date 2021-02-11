package message

import (
	"errors"

	"github.com/Rhymen/go-whatsapp"
)

// Send send message to the sender
func Send(text string, wac *whatsapp.Conn, msg whatsapp.TextMessage) error {
	remoteJId := GetRemoteJID(msg)
	if !ValidateRemoteJID(remoteJId) {
		return errors.New("Invalid RemoteJId")
	}
	content := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: remoteJId,
		},
		Text: text,
	}

	_, err := wac.Send(content)
	return err
}

// SendTo send message to the sender
func SendTo(remoteJId, text string, wac *whatsapp.Conn) error {
	if !ValidateRemoteJID(remoteJId) {
		return errors.New("Invalid RemoteJId")
	}
	content := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: remoteJId,
		},
		Text: text,
	}

	_, err := wac.Send(content)
	return err
}
