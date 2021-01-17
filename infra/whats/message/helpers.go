package message

import (
	"regexp"

	"github.com/muriboistas/zapzap/config"

	"github.com/Rhymen/go-whatsapp"
)

var remoteJid = regexp.MustCompile(`^([^@]+)@([\S]+)$`)

const (
	// PrivateMessage received private message identifier
	PrivateMessage = "s.whatsapp.net"
	// GroupMessage received group message identifier
	GroupMessage = "g.us"
	// BroadcastMessage received broadcasts message identifier
	BroadcastMessage = "broadcast"
	// NewContactMessage received new contact private message identifier
	NewContactMessage = "c.us"
)

// GetRemoteJID get it
func GetRemoteJID(msg whatsapp.TextMessage) string {
	return msg.Info.RemoteJid
}

// GetSenderNumber get sender info based on messsage
func GetSenderNumber(msg whatsapp.TextMessage) string {
	remoteJID := msg.Info.RemoteJid
	if remoteJid.MatchString(remoteJID) {
		var senderNum string
		data := remoteJid.FindStringSubmatch(remoteJID)
		msgType := data[2]
		switch msgType {
		case PrivateMessage:
			senderNum = data[1]
		case GroupMessage:
			if msg.Info.FromMe {
				return config.Get.Whatsapp.SourceNumber
			}
			data = remoteJid.FindStringSubmatch(msg.Info.Source.GetParticipant())
			senderNum = data[1]
		default:
			return ""
		}
		return senderNum
	}

	return ""
}
