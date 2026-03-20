package catfunctions

import (
	"fmt"
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

func SendCommand(cnf SerialConf, cmd string) string {
	mode := &serial.Mode{
		BaudRate: cnf.baudRate,
		Parity:   cnf.parity,
		DataBits: cnf.dataBits,
		StopBits: cnf.stopBits,
	}
	fmt.Printf("%v\n", cnf)
	port, err := serial.Open(cnf.dev, mode)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	port.SetRTS(false)
	if !strings.HasSuffix(cmd, ";") {
		cmd += ";"
	}
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
		fmt.Println("received", n, "bytes")
		return string(buff[:n])
	}
	return ""
}
