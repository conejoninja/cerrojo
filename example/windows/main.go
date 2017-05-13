// Copyright 2013 Google Inc.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// lsusb lists attached USB devices.
package main

import (
	"fmt"

	"bufio"
	"encoding/hex"
	"os"

	"github.com/conejoninja/cerrojo"
	"github.com/conejoninja/cerrojo/devices"
	"github.com/conejoninja/cerrojo/transport"
	"github.com/karalabe/hid"
)

var client cerrojo.Client
var dconf devices.Device

func main() {
	devicesConf := devices.GetDevices()

	var devInfo hid.DeviceInfo
	for _, dev := range hid.Enumerate(21324, 1) {
		fmt.Println("TREZOR device found")
		devInfo = dev
		break
	}

	var t transport.HIDAPI
	t.SetDevice(devInfo)
	client.SetTransport(&t, devicesConf["trezor"])

	fmt.Println("Please SET your device to NOT have a PIN NOR A PASSPHRASE")
	fmt.Println("Sending PING to device, it should appear \"PONG\"")
	str, msgType := call(client.Ping("PONG", false, false, false))
	fmt.Println(str, msgType)

	fmt.Println("\n\n\n")

	fmt.Println("Getting features from devices, you will see a long JSON string")
	str, msgType = call(client.GetFeatures())
	fmt.Println(str, msgType)

	fmt.Println("\n\n\n")

	fmt.Println("We are going to Activate TREZOR Password Manager, please CLICK [CONFIRM] on your device")
	fmt.Println("You will see some strange characters next, don't worry")
	str, msgType = call(client.GetMasterKey())
	fmt.Println(str, msgType)
	if msgType == 48 {
		masterKey := hex.EncodeToString([]byte(str))
		filename, _, encKey := cerrojo.GetFileEncKey(masterKey)

		// OPEN FILE
		fmt.Println("Reading file:", filename)
		contentByte, err := readFile("./" + filename)
		fmt.Println("Bytes read:", len(contentByte))
		fmt.Println("Errors:", err)
		content := string(contentByte)
		if err == nil {
			// DECRYPT STORAGE
			fmt.Println("Decrypting Storage")
			data, err := cerrojo.DecryptStorage(content, encKey)
			fmt.Println("Errors found:", err)
			fmt.Println("Version:", data.Version)
			fmt.Println("Config:", data.Config)
			fmt.Println("Tags:", data.Tags)
			fmt.Println("Number of entries found:", len(data.Entries))
			fmt.Println("We are going to try to decrypt all the entries, you need to click [CONFIRM] several times")

			for k, e := range data.Entries {
				str, msgType = call(client.GetEntryNonce(e.Title, e.Username, e.Nonce))
				pswd, epswd := cerrojo.DecryptEntry(string(e.Password.Data), str)
				fmt.Println("Decryp PSWD ERROR:", epswd)
				note, enote := cerrojo.DecryptEntry(string(e.SafeNote.Data), str)
				fmt.Println("Decryp NOTE ERROR:", enote)
				pswdStr := ""
				noteStr := ""
				if len(pswd) > 2 {
					pswdStr = pswd[1 : len(pswd)-1]
				}
				if len(note) > 2 {
					noteStr = note[1 : len(note)-1]
				}
				printEntryExtended(k, e, pswdStr, noteStr)

			}

			str = ""
		} else {
			str = "Error opening file " + filename
		}

	}
}

func call(msg []byte) (string, uint16) {
	str, msgType := client.Call(msg)

	if msgType == 18 {
		/*fmt.Println(str)
		line, err := prompt.Readline()
		if err != nil {
			fmt.Println("ERR", err)
		}
		str, msgType = call(client.PinMatrixAck(line))*/
	} else if msgType == 26 {
		fmt.Println(str)
		str, msgType = call(client.ButtonAck())
	} else if msgType == 41 {
		/*fmt.Println(str)
		line, err := prompt.Readline()
		if err != nil {
			fmt.Println("ERR", err)
		}    */
		str, msgType = call(client.PassphraseAck(""))
	} else if msgType == 46 {
		/*fmt.Println(str)
		line, err := prompt.Readline()
		if err != nil {
			fmt.Println("ERR", err)
		}
		str, msgType = call(client.WordAck(line))*/
	} else if msgType == transport.EndpointError || msgType == transport.ProtocolError || msgType == transport.DisconnectedError {
		/*fmt.Println("Device disconnected, trying to reconnect")
		if connect() > 0 {
			return call(msg)
		} */
	}

	return str, msgType
}

func printStorage(s cerrojo.Storage) {
	fmt.Println("Password Entries")
	fmt.Println("================")
	fmt.Println("")

	for id, e := range s.Entries {
		printEntry(id, e)
	}

	fmt.Println("")
	fmt.Println("Select entry number to decrypt: ")

}

func printEntry(id string, e cerrojo.Entry) {
	fmt.Printf("Entry id: #%s\n", id)
	for i := 0; i < (11 + len(id)); i++ {
		fmt.Print("-")
	}
	fmt.Println("")
	fmt.Println("* title : ", e.Note)
	fmt.Println("* item/url : ", e.Title)
	fmt.Println("* username : ", e.Username)
	fmt.Println("* tags : ", e.Tags)
	fmt.Println("")
}

func printEntryExtended(id string, e cerrojo.Entry, password, safenote string) {
	fmt.Printf("Entry id: #%s\n", id)
	for i := 0; i < (11 + len(id)); i++ {
		fmt.Print("-")
	}
	fmt.Println("")
	fmt.Println("* title : ", e.Note)
	fmt.Println("* item/url : ", e.Title)
	fmt.Println("* username : ", e.Username)
	fmt.Println("* password : ", password)
	fmt.Println("* safe note : ", safenote)
	fmt.Println("* tags : ", e.Tags)
	fmt.Println("")
}

func readFile(filename string) ([]byte, error) {
	var empty []byte

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return empty, err
	}

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return empty, statsErr
	}
	var size int64 = stats.Size()
	fw := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(fw)
	return fw, err
}
