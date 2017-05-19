package tuner

import (
	"fmt"
	"io"

	"github.com/mh-orange/tuner/api"
)

type Task interface {
	Progress() float32
	Cancel()
}

type task struct {
	total       float32
	cancelCh    chan bool
	completedCh chan float32
}

func (t *task) Progress() float32 {
	return <-t.completedCh / t.total
}

func (t *task) Cancel() {
	t.cancelCh <- true
}

type Tuner interface {
	Scan() (Task, error)
	Tune(channel api.Channel) error
	Stream() (io.Reader, error)
	Channels() []api.Channel
}

type DriverConfig interface{}
type DefaultDriverConfig interface{}

type TunerDriver interface {
	DefaultConfig() DefaultDriverConfig
	Connect(config DriverConfig) (Tuner, error)
}

var (
	tunerDrivers      map[string]TunerDriver
	ErrDriverNotFound = fmt.Errorf("Driver not found")
)

type DriverError struct {
	driverName string
	cause      error
}

func driverError(driverName string, err error) error {
	return DriverError{driverName, err}
}

func (de DriverError) Error() string {
	return fmt.Sprintf("%s: %v", de.driverName, de.cause)
}

func init() {
	tunerDrivers = make(map[string]TunerDriver)
}

func RegisterDriver(name string, driver TunerDriver) {
	tunerDrivers[name] = driver
}

func LookupDriver(name string) (TunerDriver, error) {
	if driver, found := tunerDrivers[name]; found {
		return driver, nil
	}
	return nil, driverError(name, ErrDriverNotFound)
}
