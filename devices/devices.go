package devices

import (
	keepkey "github.com/conejoninja/cerrojo/pb/keepkey/messages"
	keepkeytypes "github.com/conejoninja/cerrojo/pb/keepkey/types"
	trezor "github.com/conejoninja/cerrojo/pb/trezor/messages"
	trezortypes "github.com/conejoninja/cerrojo/pb/trezor/types"

	"github.com/conejoninja/cerrojo/pb/common"
	types "github.com/conejoninja/cerrojo/pb/types"
)

type Device struct {
	Messages common.Messager
	Types    types.Typer
	Info     Info
}
type Info struct {
	Name      string
	MasterKey string
	Vendor    uint16
	Product   uint16
	Interface uint8
}

var Devices map[string]Device

func init() {
	Devices = map[string]Device{
		"trezor": Device{
			Messages: &trezor.Getter{},
			Types:    &trezortypes.Getter{},
			Info: Info{
				Name:      "TREZOR",
				MasterKey: "2d650551248d792eabf628f451200d7f51cb63e46aadcbb1038aacb05e8c8aee2d650551248d792eabf628f451200d7f51cb63e46aadcbb1038aacb05e8c8aee",
				Vendor:    21324, //0x534c
				Product:   1,     // 0x0001
				Interface: 0,     // 0x00
			},
		},
		"keepkey": Device{
			Messages: &keepkey.Getter{},
			Types:    &keepkeytypes.Getter{},
			Info: Info{
				Name:      "KEEPKEY",
				MasterKey: "2d650551248d792eabf628f451200d7f51cb63e46aadcbb1038aacb05e8c8aee2d650551248d792eabf628f451200d7f51cb63e46aadcbb1038aacb05e8c8aee",
				Vendor:    11044, //0x2b24
				Product:   1,     // 0x0001
				Interface: 0,     // 0x00
			},
		},
	}

}

func GetDevices() map[string]Device {
	return Devices
}

func GetDevice(key string) Device {
	if _, ok := Devices[key]; ok {
		return Devices[key]
	}
	var d Device
	return d
}
