package tuner

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/mh-orange/tuner/api"
)

type DummyConfig struct {
	P1 string
	P2 string
}

type DummyTuner string

func (tuner DummyTuner) Stream() (io.Reader, error) {
	return nil, nil
}

func (tuner DummyTuner) Scan() (Task, error) {
	return nil, nil
}

func (tuner DummyTuner) Tune(api.Channel) error {
	return nil
}

func (tuner DummyTuner) Channels() []api.Channel {
	return nil
}

type DummyDriver string

func (d DummyDriver) DefaultConfig() DefaultDriverConfig {
	return &DummyConfig{
		P1: "default property 1",
		P2: "default property 2",
	}
}

func (d DummyDriver) Connect(config DriverConfig) (Tuner, error) {
	return DummyTuner(fmt.Sprintf("%s: Dummy Tuner", string(d))), nil
}

func registerDummyDrivers(names ...string) {
	for _, name := range names {
		RegisterDriver(name, DummyDriver(name))
	}
}

func TestRegisterLookupDriver(t *testing.T) {
	tests := []string{"driver1", "driver2"}
	registerDummyDrivers(tests...)

	for _, name := range tests {
		driver, err := LookupDriver(name)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		if dummyDriver, ok := driver.(DummyDriver); ok {
			if string(dummyDriver) != name {
				t.Errorf("Expected driver name %s but got %s", name, string(dummyDriver))
			}
		} else {
			t.Errorf("Expected instance of DummyDriver but got %v", reflect.TypeOf(driver))
		}
	}

	if _, err := LookupDriver("foo"); err != nil {
		expected := "foo: " + ErrDriverNotFound.Error()
		if err.Error() != expected {
			t.Errorf("Expected error %s but got %s", expected, err.Error())
		}
	} else {
		t.Errorf("Expected LookupDriver(\"foo\") to return an error, but it didn't")
	}
}
