package drivers

import (
	"io"

	"github.com/mh-orange/tuner"
	"github.com/mh-orange/tuner/api"
)

const (
	STUB_DRIVER_NAME = "stub"
)

type StubDriverConfig struct {
}

type stubTuner struct {
	cfg *StubDriverConfig
}

func newStubTuner(cfg *StubDriverConfig) (tuner.Tuner, error) {
	return &stubTuner{cfg}, nil
}

func (s *stubTuner) Scan() (tuner.Task, error) {
	return nil, nil
}

func (s *stubTuner) Tune(channel api.Channel) error {
	return nil
}

func (s *stubTuner) Stream() (io.Reader, error) {
	return nil, nil
}

func (s *stubTuner) Channels() []api.Channel {
	return nil
}

type StubDriver string

func (sd StubDriver) DefaultConfig() tuner.DefaultDriverConfig {
	return &StubDriverConfig{}
}

func (sb StubDriver) Connect(config tuner.DriverConfig) (tuner.Tuner, error) {
	return newStubTuner(config.(*StubDriverConfig))
}

func init() {
	tuner.RegisterDriver(STUB_DRIVER_NAME, StubDriver(STUB_DRIVER_NAME))
}
