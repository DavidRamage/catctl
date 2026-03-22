package catfunctions

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"go.bug.st/serial"
)

func GetConf() (SerialConf, string) {
	v := viper.New()
	home, err := os.UserHomeDir()
	if err == nil {
		v.AddConfigPath(filepath.Join(home, ".catctl"))
	}
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetDefault("Serial", map[string]any{"dev": "/dev/ttyUSB0",
		"baudrate": 38400, "parity": "none", "databits": 8,
		"stopbits": 1, "rts": false, "dtr": false})
	v.SetDefault("radio", "ft450d")
	if err := v.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Config not found: %v. Using defaults.\n", err)
	}
	var sc SerialConf
	err = v.UnmarshalKey("Serial", &sc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to decode into struct, %v", err)
	}
	sc.dev = v.GetString("dev")
	sc.baudRate = v.GetInt("baudrate")
	sc.dataBits = v.GetInt("databits")
	if v.GetString("parity") == "even" {
		sc.parity = serial.EvenParity
	} else if v.GetString("parity") == "odd" {
		sc.parity = serial.OddParity
	} else {

		sc.parity = serial.NoParity
	}
	sc.dtr = v.GetBool("dtr")
	sc.rts = v.GetBool("rts")
	return sc, v.GetString("radio")
}

func GetCommand(radio string, cmd string) string {
	v := viper.New()
	v.SetConfigName("radios")
	v.SetConfigType("yaml")
	v.AddConfigPath("$HOME/.catctl")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load config file.  Using defaults which are probably wrong for you.\n")
		os.Exit(-1)
	}
	key := fmt.Sprintf("%s.%s", radio, cmd)
	result := v.GetString(key)
	if result == "" {
		fmt.Fprintf(os.Stderr, "Unable to find command for radio %s and command %s.  Check your radios.yaml file.\n", radio, cmd)
		os.Exit(-1)
	}
	return result
}
