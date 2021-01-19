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

// GetRemoteHost get the message host ex 629731239383
func GetRemoteHost(remoteJID string) string {
	if remoteJid.MatchString(remoteJID) {
		data := remoteJid.FindStringSubmatch(remoteJID)
		return data[1]
	}

	return ""
}

// GetRemoteIdentifier get the message identifier ex: group message, private message...
func GetRemoteIdentifier(remoteJID string) string {
	if remoteJid.MatchString(remoteJID) {
		data := remoteJid.FindStringSubmatch(remoteJID)
		return data[2]
	}

	return ""
}

// GetRemoteJID get it
func GetRemoteJID(msg whatsapp.TextMessage) string {
	return msg.Info.RemoteJid
}

// GetSenderNumber get sender info based on messsage
func GetSenderNumber(msg whatsapp.TextMessage) string {
	msgIdentifier, senderNum := GetRemoteIdentifier(msg.Info.RemoteJid), ""
	switch msgIdentifier {
	case PrivateMessage:
		senderNum = GetRemoteHost(msg.Info.RemoteJid)
	case GroupMessage:
		if msg.Info.FromMe {
			return config.Get.Whatsapp.SourceNumber
		}
		senderNum = GetRemoteHost(msg.Info.Source.GetParticipant())
	}

	return senderNum
}
