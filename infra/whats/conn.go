package whats

import (
	"encoding/gob"
	"os"
	"time"

	cfg "github.com/muriboistas/zapzap/config"
	"github.com/muriboistas/zapzap/infra/qrcode"

	"github.com/Rhymen/go-whatsapp"
)

var config = cfg.Get

// New create a new whatsapp connection
func New() (*whatsapp.Conn, error) {
	wac, err := whatsapp.NewConn(config.Whatsapp.TimeOutDuration * time.Second)
	if err != nil {
		return nil, err
	}

	//load saved session
	session, err := readSession()
	if err == nil {
		//restore session
		session, err = wac.RestoreWithSession(session)
		if err == nil {
			return wac, nil
		}
	}
	// if no saved session or failed restoration create it
	if err != nil {
		qr := make(chan string)
		go func() {
			qrc := <-qr
			if config.Qrcode.GeneratePNG {
				qrcode.Generate(qrc)
			}
			if config.Qrcode.PrintOnCLI {
				qrcode.Print(qrc)
			}
		}()
		session, err = wac.Login(qr)
		if err != nil {
			return nil, err
		}
	}

	//save session
	err = writeSession(session)
	if err != nil {
		return nil, err
	}

	return wac, nil
}

func readSession() (whatsapp.Session, error) {
	session := whatsapp.Session{}
	file, err := os.Open(config.Whatsapp.SessionPath + "/wac.gob")
	if err != nil {
		return session, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	if err := decoder.Decode(&session); err != nil {
		return session, err
	}

	return session, nil
}

func writeSession(session whatsapp.Session) error {
	file, err := os.Create(config.Whatsapp.SessionPath + "/wac.gob")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}

	return nil
}
