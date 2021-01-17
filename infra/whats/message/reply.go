package message

import "github.com/Rhymen/go-whatsapp"

// Reply send message to the sender
func Reply(text string, wac *whatsapp.Conn, msg whatsapp.TextMessage) error {
	content := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: GetRemoteJID(msg),
		},
		Text: text,
	}

	_, err := wac.Send(content)
	return err
}
