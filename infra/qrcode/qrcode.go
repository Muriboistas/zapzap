package qrcode

import (
	"strings"

	cfg "github.com/muriboistas/zapzap/config"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/skip2/go-qrcode"
)

var config = cfg.Get.Qrcode

// Generate genetare a qrcode in a png image
func Generate(data string) error {
	var recoveryLevel qrcode.RecoveryLevel
	switch strings.ToLower(config.Quality) {
	case "low":
		recoveryLevel = qrcode.Low
	case "medium":
		recoveryLevel = qrcode.Medium
	case "high":
		recoveryLevel = qrcode.High
	case "highest":
		recoveryLevel = qrcode.Highest
	default:
		recoveryLevel = qrcode.Medium
	}

	err := qrcode.WriteFile(data, recoveryLevel, int(config.Size), config.FileName+"/qrcode.png")
	if err != nil {
		return err
	}

	return nil
}

// Print print our qr code on terminal
func Print(data string) {
	var recoveryLevel = qrcodeTerminal.QRCodeRecoveryLevels.Medium
	switch strings.ToLower(config.Quality) {
	case "low":
		recoveryLevel = qrcodeTerminal.QRCodeRecoveryLevels.Low
	case "medium":
		recoveryLevel = qrcodeTerminal.QRCodeRecoveryLevels.Medium
	case "high":
		recoveryLevel = qrcodeTerminal.QRCodeRecoveryLevels.High
	case "highest":
		recoveryLevel = qrcodeTerminal.QRCodeRecoveryLevels.Highest
	default:
		recoveryLevel = qrcodeTerminal.QRCodeRecoveryLevels.Medium
	}

	terminal := qrcodeTerminal.New2(
		qrcodeTerminal.ConsoleColors.BrightBlack,
		qrcodeTerminal.ConsoleColors.BrightWhite,
		recoveryLevel,
	)
	terminal.Get(data).Print()
}
