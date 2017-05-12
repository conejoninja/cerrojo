package transport

import (
	"encoding/binary"
	"log"
	"math"

	"github.com/karalabe/gousb/usb"
)

type GoUSBHID struct {
	device usb.Device
	epw    usb.Endpoint
	epr    usb.Endpoint
	set    bool
}

func (t *GoUSBHID) SetDevice(device usb.Device) {
	t.device = device
	epw, err := t.device.OpenEndpoint(1, 0, 0, 1) // change this
	t.epw = epw
	epr, err := t.device.OpenEndpoint(1, 0, 0, 129) // change this
	t.epr = epr
	if err != nil {
		log.Println("Open error: ", err)
	} else {
		t.set = true
	}
}

func (t *GoUSBHID) Close() {
	t.device.Close()
}

func (t *GoUSBHID) Write(msg []byte) {
	for len(msg) > 0 && t.set {
		blank := make([]byte, 64)
		l := int(math.Min(63, float64(len(msg))))
		tmp := append([]byte{63}, msg[:l]...)
		copy(blank, tmp)
		n, err := t.epw.Write(blank)

		if err == nil && n > 0 {
			if len(msg) < 64 {
				break
			} else {
				msg = msg[63:]
			}
		} else {
			break
		}
	}
}

func copyB(a, b []byte) []byte {
	la := len(a)
	lb := len(b)
	c := make([]byte, la+lb)
	for i := 0; i < la; i++ {
		c[i] = a[i]
	}
	for i := 0; i < lb; i++ {
		c[la+i] = b[i]
	}
	return c
}

func (t *GoUSBHID) Read() ([]byte, uint16, int, error) {
	buf := make([]byte, 64)
	bufLength, err := t.epr.Read(buf)
	var marshalled []byte

	for i := 0; i < bufLength; i++ {
		// 35 : '#' magic header
		if buf[i] == 35 && buf[i+1] == 35 {
			msgType := binary.BigEndian.Uint16(buf[i+2 : i+4])
			msgLength := int(binary.BigEndian.Uint32(buf[i+4 : i+8]))
			originalMsgLength := msgLength

			if (bufLength - i - 8) < msgLength {
				marshalled = copyB(marshalled, buf[i+8:])
				msgLength = msgLength - (len(buf) - i - 8)

				for msgLength > 0 {
					_, err = t.epr.Read(buf)
					bufLength = len(buf)
					if bufLength > 0 {
						l := int(math.Min(float64(bufLength-1), float64(msgLength)))
						marshalled = copyB(marshalled, buf[1:l+1])
						msgLength = msgLength - l
					}
				}
			} else {
				marshalled = buf[i+8 : i+8+msgLength]
			}
			return marshalled, msgType, originalMsgLength, nil
		}
	}
	var msgType uint16
	switch err.Error() {
	case "protocol error":
		msgType = ProtocolError
		break
	case "cannot send after transport endpoint shutdown":
		msgType = EndpointError
		break
	case "no such device":
		msgType = DisconnectedError
		break
	default:
		msgType = TimeoutError
		break
	}
	return marshalled, msgType, 0, err
}
