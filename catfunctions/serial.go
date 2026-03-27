package catfunctions

import (
	"log"
	"os"
	"strings"
	"time"

	"go.bug.st/serial"
)

type SerialConf struct {
	dev      string
	baudRate int
	parity   serial.Parity
	dataBits int
	stopBits serial.StopBits
	rts      bool
	dtr      bool
}

func SendCommand(cnf SerialConf, cmd string) (string, error) {
	//TODO:  Need to handle handle error messages from the radio
	mode := &serial.Mode{
		BaudRate:          cnf.baudRate,
		Parity:            cnf.parity,
		DataBits:          cnf.dataBits,
		StopBits:          cnf.stopBits,
		InitialStatusBits: &serial.ModemOutputBits{RTS: cnf.rts, DTR: cnf.dtr},
	}
	port, err := serial.Open(cnf.dev, mode)
	//sleep to let rts go low
	port.SetReadTimeout(500 * time.Millisecond)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	if !strings.HasSuffix(cmd, ";") {
		cmd += ";"
	}
	time.Sleep(100 * time.Millisecond)
	n, err := port.Write([]byte(cmd))
	if n == 0 {
		os.Exit(-1)
	}
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	time.Sleep(100 * time.Millisecond)
	buff := make([]byte, 64)
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		if n == 0 {
			break
		}
		return string(buff[:n]), nil
	}
	return "", nil
}
