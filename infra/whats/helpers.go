package whats

import (
	"fmt"
	"regexp"
)

var remoteJid = regexp.MustCompile(`^((\w+)|(\w+)-(\w+))@([\w\W]+)$`)

const (
	// PrivateMessage received private message remotejid
	PrivateMessage = "s.whatsapp.net"
	// GroupMessage received group message remotejid
	GroupMessage = "g.us"
	// BroadcastMessage received broadcasts message remotejid
	BroadcastMessage = "broadcast"
	// NewContactMessage received new contact private message remotejid
	NewContactMessage = "c.us"
)

// GetRemoteJidInfo get infos
func GetRemoteJidInfo(str string) (map[string]string, error) {
	var err error
	info := map[string]string{"number": "", "uid": "", "msgType": ""}

	if remoteJid.MatchString(str) {
		res := remoteJid.FindStringSubmatch(str)

		info["msgType"] = res[5]
		switch info["msgType"] {
		case PrivateMessage:
			info["number"] = res[2]
		case GroupMessage:
			info["number"] = res[3]
			info["uid"] = res[4]
		case BroadcastMessage:
			info["uid"] = res[2]
		case NewContactMessage:
			info["number"] = res[2]
		default:
			err = fmt.Errorf("unknown message type: %s", info["msgType"])
		}
	} else {
		err = fmt.Errorf("string don't match: %s", str)
	}

	return info, err
}
