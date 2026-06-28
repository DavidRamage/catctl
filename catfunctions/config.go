package catfunctions

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"go.bug.st/serial"
)

func GetConf() (SerialConf, string, error) {
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
	if v.GetInt("stopbits") == 1 {
		sc.stopBits = serial.OneStopBit
	} else if v.GetInt("stopbits") == 2 {
		sc.stopBits = serial.TwoStopBits
	}
	sc.errorStr = v.GetString("errstr")
	sc.unit = v.GetString("unit")
	return sc, v.GetString("radio"), nil
}

func GetRadioData(radio string, key string, value string) (string, error) {
	v := viper.New()
	v.SetConfigName("radios")
	v.SetConfigType("yaml")
	v.AddConfigPath("$HOME/.catctl")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load config file.  Using defaults which are probably wrong for you.\n")
		os.Exit(-1)
	}
	lookup := fmt.Sprintf("%s.%s.%s", radio, key, value)
	result := v.GetString(lookup)
	if result == "" {
		fmt.Fprintf(os.Stderr, "Unable to find data for radio %s key %s and  value %s.  Check your radios.yaml file.\n", radio, key, value)
		os.Exit(-1)
	}
	return result, nil
}
