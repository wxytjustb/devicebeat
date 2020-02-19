package beater

import (
	"fmt"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/wxytjustb/devicebeat/config"
)

// devicebeat configuration.
type devicebeat struct {
	done   chan struct{}
	config config.Config
	Client beat.Client
	controllerList	[]*PingerController
}

// New creates an instance of devicebeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &devicebeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts devicebeat.
func (bt *devicebeat) Run(b *beat.Beat) error {
	logp.Info("devicebeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.Client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}


	// 根据Devices配置生成controllers
	bt.controllerList = make([]*PingerController, len(bt.config.Devices), len(bt.config.Devices))

	for index, device := range bt.config.Devices{
		if (device.Method == "ping"){
			controller, _ := NewPingerController(b, bt.Client, device, bt.config)
			bt.controllerList[index] = controller
		}
	}


	for {
		select {
			case <-bt.done:
				return nil
		}
	}

	return nil
}

// Stop stops devicebeat.
func (bt *devicebeat) Stop() {
	bt.Client.Close()
	close(bt.done)
	for _, controller := range bt.controllerList{
		controller.Stop()
	}
}
