// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Device struct {
	Name	string	`json:"name"`
	Address string  `json:"address"`
	Port 	string	`json:"port"`
	Method	string	`json:"method"`
}

type Config struct {
	Period 		time.Duration	`config:"period"`
	Devices 	[]Device		`config:"devices"`
	WorkZone	string			`config:"work_zone"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
	Devices: []Device{},
	WorkZone: "undefined",
}

