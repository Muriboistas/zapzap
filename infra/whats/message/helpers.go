package message

import (
	"fmt"
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
	// ContactMessage received contact private message identifier
	ContactMessage = "c.us"
)

// ValidateRemoteJID ...
func ValidateRemoteJID(remoteJID string) bool {
	return remoteJid.MatchString(remoteJID)
}

// GetRemoteJID get it
func GetRemoteJID(msg whatsapp.TextMessage) string {
	return msg.Info.RemoteJid
}

// GetSenderJID is used to get the JId
func GetSenderJID(msg whatsapp.TextMessage) string {
	msgIdentifier, senderJID := GetRemoteIdentifier(msg.Info.RemoteJid), ""

	//if the sender was the host
	if msg.Info.FromMe {
		return fmt.Sprintf("%s@%s", GetRemoteHost(config.Get.Whatsapp.RootNumber), PrivateMessage)
	}

	switch msgIdentifier {
	case PrivateMessage:
		senderJID = msg.Info.RemoteJid
	case GroupMessage:
		senderJID = msg.Info.Source.GetParticipant()
	}

	return senderJID
}

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

// GetSenderNumber get sender info based on messsage
func GetSenderNumber(msg whatsapp.TextMessage) string {
	return GetRemoteHost(GetSenderJID(msg))
}
