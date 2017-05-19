package tuner

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestUnmarshalTOML(t *testing.T) {
	registerDummyDrivers("test driver")

	tests := []struct {
		configText     []byte
		expectedConfig interface{}
		expectedErr    error
	}{
		{
			configText:     []byte("[[tuner]]\ndriver = \"test driver\"\nP1 = \"property 1\"\nP2 = \"property 2\""),
			expectedConfig: &DummyConfig{"property 1", "property 2"},
			expectedErr:    nil,
		},
		{
			configText:     []byte("[[tuner]]\ndriver = \"test driver\"\nP1 = \"property 1\"\nP3 = \"property 3\""),
			expectedConfig: &DummyConfig{"property 1", ""},
			expectedErr:    fmt.Errorf("Unknown key P3"),
		},
	}

	for i, test := range tests {
		config, err := ReadConfig(bytes.NewReader(test.configText))
		if err != nil && test.expectedErr != nil {
			if !strings.HasSuffix(err.Error(), test.expectedErr.Error()) {
				t.Errorf("Test %d: Expected error %v but got %v", i, test.expectedErr, err)
			}
		} else if err != nil || test.expectedErr != nil {
			t.Errorf("Test %d: Expected error %v but got %v", i, test.expectedErr, err)
		}

		if test.expectedErr == nil {
			if len(config.Tuner) != 1 {
				t.Errorf("Test %d: Expected one tuner, got %d", i, len(config.Tuner))
			} else {
				driverConfig := config.Tuner[0].TunerSpecificConfig
				if !reflect.DeepEqual(test.expectedConfig, driverConfig) {
					t.Errorf("Test %d: Expected %v got %v", i, test.expectedConfig, driverConfig)
				}
			}
		}
	}
}
