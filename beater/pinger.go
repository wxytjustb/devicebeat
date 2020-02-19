package beater

import (
	"github.com/elastic/beats/libbeat/beat"
	"github.com/wxytjustb/devicebeat/config"
	"github.com/sparrc/go-ping"
	"time"
	"github.com/elastic/beats/libbeat/common"
)

type PingData struct {
	rtt				time.Duration
	create_time 	time.Time
}

type PingerController struct{
	bt 				*beat.Beat
	client			beat.Client
	device 			config.Device
	pinger			*ping.Pinger
	done			chan struct{}
	is_start_new	chan bool
	config 			config.Config
	data			chan PingData
}

func NewPingerController(bt *beat.Beat, client beat.Client, device config.Device, config config.Config) (*PingerController, error){
	controller := &PingerController{
		bt:				bt,
		client:			client,
		device:			device,
		done:			make(chan struct{}),
		is_start_new:	make(chan bool),
		config:			config,
		data:			make(chan PingData),
	}


	_, error := controller.newPinger()

	go controller.Run()

	return controller, error
}

func (controller *PingerController) Run() error{

	// 根据计时器采集数据，如没没有数据，发一次掉线的event

	ticker := time.NewTicker(controller.config.Period)

	for{
		select {
			case <- controller.done:
				return nil
			case <- controller.is_start_new:
				controller.newPinger()
			case <- ticker.C:

				go func() {

					// 设置0.1s(100毫秒)超时时间
					select {
						case data := <- controller.data:
							event := beat.Event{
								Timestamp: data.create_time,
								Fields: common.MapStr{
									"type":    	controller.bt.Info.Name,
									"rtt": 	   	data.rtt.Seconds(),
									"address":	controller.device.Address,
									"name": controller.device.Name,
									"work_zone": controller.config.WorkZone,
								},
							}
							controller.client.Publish(event)
						case <- time.After(100 * time.Millisecond):
							event := beat.Event{
								Timestamp: time.Now(),
								Fields: common.MapStr{
									"type":    	controller.bt.Info.Name,
									"rtt": 	   	0,
									"address":	controller.device.Address,
									"name": controller.device.Name,
									"work_zone": controller.config.WorkZone,
									//"counter": counter,
								},
							}
							controller.client.Publish(event)
					}


				}()

		}
	}
}

func (controller *PingerController) newPinger() (*ping.Pinger, error){



	pinger, error := ping.NewPinger(controller.device.Address)


	pinger.OnRecv = func(packet *ping.Packet) {
		controller.data <- PingData{
			packet.Rtt,
			time.Now(),
		}
	}

	pinger.OnFinish = func(statistics *ping.Statistics) {
		// 结束后需要重新开始一个
		controller.is_start_new <- true
	}

	pinger.Timeout = time.Hour * 24


	go pinger.Run()

	controller.pinger = pinger
	return pinger, error
}

// 停止执行
func (controller *PingerController) Stop(){

	// 关掉controller主循环
	close(controller.done)

	controller.pinger.Stop()
	close(controller.is_start_new)
}
