package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.bug.st/serial"
)

func getConf() (SerialConf, string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.catctl")
	viper.SetDefault("Serial", map[string]any{"dev": "/dev/ttyUSB0",
		"baudRate": 38400, "parity": "none", "dataBits": 8,
		"stopBits": 1, "rts": false, "dtr": false})
	viper.SetDefault("Radio", "FT450D")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load config file.  Using defaults which are probably wrong for you.\n")
	}
	radio := viper.GetString("Radio")
	sc := SerialConf{
		dev:      viper.GetString("Serial.dev"),
		baudRate: viper.GetInt("Serial.baudRate"),
		rts:      viper.GetBool("Serial.rts"),
		dtr:      viper.GetBool("Serial.dts"),
		parity:   serial.NoParity,
		stopBits: serial.OneStopBit,
		dataBits: viper.GetInt("Serial.dataBits"),
	}
	return sc, radio

}
