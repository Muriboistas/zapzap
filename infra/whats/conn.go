package whats

import (
	"encoding/gob"
	"errors"
	"log"
	"os"

	"time"

	cfg "github.com/muriboistas/zapzap/config"
	"github.com/muriboistas/zapzap/infra/qrcode"

	"github.com/Rhymen/go-whatsapp"
)

// Wac the whatsapp connection
var waConn *whatsapp.Conn

var config = &cfg.Get

// New create a new whatsapp connection
func New() error {
	//create new WhatsApp connection
	var err error
	waConn, err = whatsapp.NewConn(config.Whatsapp.TimeOutDuration * time.Second)
	if err != nil {
		return err
	}

	// Set client configs
	waConn.SetClientName(config.Whatsapp.LongClientName, config.Whatsapp.ShortClientName, config.Whatsapp.ClientVersion)
	waConn.SetClientVersion(2, 2021, 4)

	//Add handler
	waConn.AddHandler(&waHandler{waConn})

	// make the connection
	err = login(waConn)
	if err != nil {
		log.Println(err)
	}

	//verifies phone connectivity
	pong, err := waConn.AdminTest()
	if !pong || err != nil {
		return errors.New("error pinging in")
	}

	return nil
}

// Disconnect safely disconnect whatsapp connection
func Disconnect() error {
	_, err := waConn.Disconnect()

	return err
}

func login(wac *whatsapp.Conn) error {
	//load saved session
	session, err := readSession()
	if err == nil {
		//restore session
		session, err = wac.RestoreWithSession(session)
		if err == nil {
			return nil
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
			return err
		}
	}

	//save session
	err = writeSession(session)
	if err != nil {
		return err
	}

	return nil
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
