package tuner

import (
	"fmt"
	"io"
	"reflect"

	"github.com/naoina/toml"
)

type TunerConfig struct {
	Driver              TunerDriver
	TunerSpecificConfig DriverConfig
}

func (tc *TunerConfig) UnmarshalTOML(decode func(interface{}) error) (err error) {
	var details map[string]interface{}

	if err = decode(&details); err == nil {
		if driverName, ok := details["driver"].(string); ok {
			var driver TunerDriver
			if driver, err = LookupDriver(driverName); err == nil {
				tc.Driver = driver
				tc.TunerSpecificConfig = driver.DefaultConfig()
				err = decode(tc.TunerSpecificConfig)
			}
		}
	}
	return err
}

type Config struct {
	Tuner []TunerConfig
}

func ReadConfig(reader io.Reader) (*Config, error) {
	toml.DefaultConfig.MissingField = func(typ reflect.Type, key string) (err error) {
		if key != "driver" {
			err = fmt.Errorf("Unknown key %s", key)
		}
		return
	}
	config := &Config{}
	decoder := toml.NewDecoder(reader)
	return config, decoder.Decode(config)
}
