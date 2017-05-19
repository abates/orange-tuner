package drivers

import (
	"io"

	"github.com/mh-orange/tuner"
	"github.com/mh-orange/tuner/api"
)

const (
	PIPE_DRIVER_NAME = "pipe"
)

type PipeDriverConfig struct {
	Exec    string
	Loop    bool
	Channel string
}

type pipeTuner struct {
	cfg *PipeDriverConfig
}

func newPipeTuner(cfg *PipeDriverConfig) (tuner.Tuner, error) {
	tuner := &pipeTuner{cfg}
	return tuner, tuner.exec()
}

func (p *pipeTuner) exec() error {
	return nil
}

func (p *pipeTuner) Scan() (tuner.Task, error) {
	return nil, nil
}

func (p *pipeTuner) Tune(channel api.Channel) error {
	return nil
}

func (p *pipeTuner) Stream() (io.Reader, error) {
	return nil, nil
}

func (p *pipeTuner) Channels() []api.Channel {
	return nil
}

type PipeDriver string

func (p PipeDriver) DefaultConfig() tuner.DefaultDriverConfig {
	return &PipeDriverConfig{}
}

func (p PipeDriver) Connect(config tuner.DriverConfig) (tuner.Tuner, error) {
	return newPipeTuner(config.(*PipeDriverConfig))
}

func init() {
	tuner.RegisterDriver(PIPE_DRIVER_NAME, PipeDriver(PIPE_DRIVER_NAME))
}
