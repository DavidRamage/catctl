package catfunctions

import (
	"log"
	"os"

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

func SendCommand(cnf SerialConf, cmd string) string {
	mode := &serial.Mode{
		BaudRate: cnf.baudRate,
		Parity:   cnf.parity,
		DataBits: cnf.dataBits,
		StopBits: cnf.stopBits,
	}
	port, err := serial.Open(cnf.dev, mode)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	n, err := port.Write([]byte(cmd))
	if n == 0 {
		os.Exit(-1)
	}
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	buff := make([]byte, 32)
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		if n == 0 {
			break
		}
		return string(buff[:n])
	}
	return ""
}
