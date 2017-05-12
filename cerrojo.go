package cerrojo

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"io"
	"log"
	"math"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/conejoninja/cerrojo/devices"
	"github.com/conejoninja/cerrojo/pb/common"
	"github.com/conejoninja/cerrojo/pb/types"
	"github.com/conejoninja/cerrojo/transport"
	"github.com/golang/protobuf/proto"
	"golang.org/x/text/unicode/norm"
)

const hardkey uint32 = 2147483648

type Client struct {
	t    transport.Transport
	m    common.Messager
	tp   types.Typer
	info devices.Info
}

type Storage struct {
	Version string           `json:"version"`
	Config  Config           `json:"config"`
	Tags    map[string]Tag   `json:"tags"`
	Entries map[string]Entry `json:"entries"`
}

type Config struct {
	OrderType string `json:"orderType"`
}

type Tag struct {
	Title  string `json:"title"`
	Icon   string `json:"icon"`
	Active string `json:"active"`
}

type Entry struct {
	Title    string        `json:"title"`
	Username string        `json:"username"`
	Nonce    string        `json:"nonce"`
	Note     string        `json:"note"`
	Password EncryptedData `json:"password"`
	SafeNote EncryptedData `json:"safe_note"`
	Tags     []int         `json:"tags"`
}

type EncryptedData struct {
	Type string `json:"type"`
	Data []byte `json:"data"`
}

type TxRequest struct {
	Details *types.TxRequestDetailsTyper `json:"details,omitempty"`
	Type    types.RequestType            `json:"type,omitempty"`
}

func (c *Client) SetTransport(t transport.Transport, d devices.Device) {
	c.t = t
	c.m = d.Messages
	c.tp = d.Types
	c.info = d.Info
}

func (c *Client) CloseTransport() {
	c.t.Close()
}

func (c *Client) Header(msgType common.MessageType, msg []byte) []byte {

	typebuf := make([]byte, 2)
	binary.BigEndian.PutUint16(typebuf, uint16(msgType))

	msgbuf := make([]byte, 4)
	binary.BigEndian.PutUint32(msgbuf, uint32(len(msg)))

	return append(typebuf, msgbuf...)
}

func (c *Client) Initialize() []byte {
	m := c.m.GetInitialize()
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_Initialize"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) Ping(str string, pinProtection, passphraseProtection, buttonProtection bool) []byte {
	m := c.m.GetPing()
	m.SetMessage(&str)
	m.SetButtonProtection(&buttonProtection)
	m.SetPinProtection(&pinProtection)
	m.SetPassphraseProtection(&passphraseProtection)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_Ping"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) ChangePin() []byte {
	m := c.m.GetChangePin()
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_ChangePin"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) GetEntropy(size uint32) []byte {
	m := c.m.GetGetEntropy()
	m.SetSize(&size)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_GetEntropy"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) GetFeatures() []byte {
	m := c.m.GetGetFeatures()
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_GetFeatures"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) PinMatrixAck(str string) []byte {
	m := c.m.GetPinMatrixAck()
	m.SetPin(&str)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_PinMatrixAck"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) PassphraseAck(str string) []byte {
	m := c.m.GetPassphraseAck()
	m.SetPassphrase(&str)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_PassphraseAck"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}
func (c *Client) WordAck(str string) []byte {
	m := c.m.GetWordAck()
	m.SetWord(&str)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_WordAck"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) GetAddress(addressN []uint32, showDisplay bool, coinName string) []byte {
	m := c.m.GetGetAddress()
	m.SetAddressN(addressN)
	m.SetCoinName(&coinName)
	m.SetShowDisplay(&showDisplay)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_GetAddress"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) GetPublicKey(address []uint32) []byte {
	m := c.m.GetGetPublicKey()
	m.SetAddressN(address)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_GetPublicKey"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) SignMessage(message []byte) []byte {
	m := c.m.GetSignMessage()
	m.SetMessage(norm.NFC.Bytes(message))
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_SignMessage"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) SignIdentity(uri string, challengeHidden []byte, challengeVisual string, index uint32) []byte {
	m := c.m.GetSignIdentity()
	identity := URIToIdentity(uri)
	identity.SetIndex(&index)
	m.SetIdentity(identity)
	m.SetChallengeHidden(challengeHidden)
	m.SetChallengeVisual(&challengeVisual)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_SignIdentity"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) SetLabel(label string) []byte {
	m := c.m.GetApplySettings()
	m.SetLabel(&label)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_ApplySettings"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) WipeDevice() []byte {
	m := c.m.GetWipeDevice()
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_WipeDevice"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) EntropyAck(entropy []byte) []byte {
	m := c.m.GetEntropyAck()
	m.SetEntropy(entropy)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_EntropyAck"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) ResetDevice(displayRandom bool, strength uint32, passphraseProtection, pinProtection bool, label string, U2FCounter uint32) []byte {
	m := c.m.GetResetDevice()
	m.SetDisplayRandom(&displayRandom)
	m.SetStrength(&strength)
	m.SetPassphraseProtection(&passphraseProtection)
	m.SetPinProtection(&pinProtection)
	m.SetU2FCounter(&U2FCounter)
	if label != "" {
		m.SetLabel(&label)
	}
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_ResetDevice"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) LoadDevice(mnemonic string, passphraseProtection bool, label, pin string, SkipChecksum bool, U2FCounter uint32) []byte {
	m := c.m.GetLoadDevice()
	m.SetMnemonic(&mnemonic)
	m.SetPassphraseProtection(&passphraseProtection)
	if label != "" {
		m.SetLabel(&label)
	}
	if pin != "" {
		m.SetPin(&pin)
	}
	m.SetSkipChecksum(&SkipChecksum)
	m.SetU2FCounter(&U2FCounter)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_LoadDevice"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) EncryptMessage(pubkey, message string, displayOnly bool, path, coinName string) []byte {
	m := c.m.GetEncryptMessage()
	m.SetPubkey([]byte(pubkey))
	m.SetMessage([]byte(message))
	m.SetDisplayOnly(&displayOnly)
	m.SetAddressN(StringToBIP32Path(path))
	m.SetCoinName(&coinName)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_EncryptMessage"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) DecryptMessage(path string, nonce, message, hmac []byte) []byte {
	m := c.m.GetDecryptMessage()
	m.SetAddressN(StringToBIP32Path(path))
	m.SetNonce(nonce)
	m.SetMessage(message)
	m.SetHmac(hmac)
	marshalled, err := proto.Marshal(m)
	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_DecryptMessage"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) RecoveryDevice(wordCount uint32, passphraseProtection, pinProtection bool, label string, EnforceWordList bool, U2FCounter uint32) []byte {
	m := c.m.GetRecoveryDevice()
	m.SetWordCount(&wordCount)
	m.SetPassphraseProtection(&passphraseProtection)
	m.SetPinProtection(&pinProtection)
	m.SetLabel(&label)
	m.SetEnforceWordlist(&EnforceWordList)
	m.SetU2FCounter(&U2FCounter)

	if label != "" {
		m.SetLabel(&label)
	}
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_RecoveryDevice"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) SetHomescreen(homescreen []byte) []byte {
	m := c.m.GetApplySettings()
	m.SetHomescreen(homescreen)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_ApplySettings"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) VerifyMessage(address, signature string, message []byte) []byte {

	sign, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return []byte("Wrong signature")
	}

	m := c.m.GetVerifyMessage()
	m.SetAddress(&address)
	m.SetSignature(sign)
	m.SetMessage(norm.NFC.Bytes(message))
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_VerifyMessage"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) EstimateTxSize(outputsCount, inputsCount uint32, coinName string) []byte {
	m := c.m.GetEstimateTxSize()
	m.SetOutputsCount(&outputsCount)
	m.SetInputsCount(&inputsCount)
	m.SetCoinName(&coinName)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_EstimateTxSize"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) ButtonAck() []byte {
	m := c.m.GetButtonAck()
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_ButtonAck"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) GetMasterKey() []byte {
	masterKey, _ := hex.DecodeString(c.info.MasterKey)
	return c.CipherKeyValue(
		true,
		"Activate "+c.info.Name+" Password Manager?",
		masterKey,
		StringToBIP32Path("m/10016'/0"),
		[]byte{},
		true,
		true,
	)
}

func (c *Client) GetEntryNonce(title, username, nonce string) []byte {
	return c.CipherKeyValue(
		false,
		"Unlock "+title+" for user "+username+"?",
		[]byte(nonce),
		StringToBIP32Path("m/10016'/0"),
		[]byte{},
		false,
		true,
	)
}

func (c *Client) SetEntryNonce(title, username, nonce string) []byte {
	return c.CipherKeyValue(
		true,
		"Unlock "+title+" for user "+username+"?",
		[]byte(nonce),
		StringToBIP32Path("m/10016'/0"),
		[]byte{},
		false,
		true,
	)
}

func (c *Client) ClearSession() []byte {
	m := c.m.GetClearSession()
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_ClearSession"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) SetU2FCounter(U2FCounter uint32) []byte {
	m := c.m.GetSetU2FCounter()
	m.SetU2FCounter(&U2FCounter)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_SetU2FCounter"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) GetECDHSessionKey(uri string, index uint32, peerPublicKey []byte, ecdsaCurveName string) []byte {
	m := c.m.GetGetECDHSessionKey()
	identity := URIToIdentity(uri)
	identity.SetIndex(&index)
	m.SetIdentity(identity)
	m.SetPeerPublicKey(peerPublicKey)
	m.SetEcdsaCurveName(&ecdsaCurveName)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_GetECDHSessionKey"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) FirmwareErase() []byte {
	m := c.m.GetFirmwareErase()
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_FirmwareErase"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) FirmwareUpload(payload []byte) []byte {
	m := c.m.GetFirmwareUpload()
	m.SetPayload(payload)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_FirmwareUpload"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) SignTx(outputsCount, inputsCount uint32, coinName string, version, lockTime uint32) []byte {
	m := c.m.GetSignTx()
	m.SetOutputsCount(&outputsCount)
	m.SetInputsCount(&inputsCount)
	m.SetCoinName(&coinName)
	if version != 0 {
		m.SetVersion(&version)
	}
	if lockTime != 0 {
		m.SetLockTime(&lockTime)
	}
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_SignTx"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) TxAck(tx types.TransactionTyper) []byte {
	m := c.m.GetTxAck()
	m.SetTx(tx)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_TxAck"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) CipherKeyValue(encrypt bool, key string, value []byte, address []uint32, iv []byte, askOnEncrypt, askOnDecrypt bool) []byte {
	m := c.m.GetCipherKeyValue()
	m.SetKey(&key)
	if encrypt {
		paddedValue := make([]byte, 16*int(math.Ceil(float64(len(value))/16)))
		copy(paddedValue, value)
		m.SetValue(paddedValue)
	} else {
		val, err := hex.DecodeString(string(value))
		m.SetValue(val)
		if err != nil {
			fmt.Println("ERROR Decoding string")
		}
	}
	m.SetAddressN(address)
	if len(iv) > 0 {
		m.SetIv(iv)
	}
	m.SetEncrypt(&encrypt)
	m.SetAskOnEncrypt(&askOnEncrypt)
	m.SetAskOnDecrypt(&askOnDecrypt)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_CipherKeyValue"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) EthereumGetAddress(addressN []uint32, showDisplay bool) []byte {
	m := c.m.GetEthereumGetAddress()
	m.SetAddressN(addressN)
	m.SetShowDisplay(&showDisplay)
	marshalled, err := proto.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshalling")
	}

	magicHeader := append([]byte{35, 35}, c.Header(common.MessageType_value["MessageType_MessageType_EthereumGetAddress"], marshalled)...)
	msg := append(magicHeader, marshalled...)

	return msg
}

func (c *Client) Call(msg []byte) (string, uint16) {
	c.t.Write(msg)
	return c.ReadUntil()
}

func (c *Client) ReadUntil() (string, uint16) {
	var str string
	var msgType uint16
	for {
		str, msgType = c.Read()
		if msgType != 999 { //timeout
			break
		}
	}

	return str, msgType
}

func (c *Client) Read() (string, uint16) {
	marshalled, msgType, _, err := c.t.Read()
	if err != nil {
		return "Error reading", msgType
	}

	str := "Uncaught message type " + strconv.Itoa(int(msgType))
	switch common.MessageType(msgType) {
	case common.MessageType_value["MessageType_MessageType_Success"]:
		msg := c.m.GetSuccess()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (2)"
		} else {
			str = msg.GetMessage()
		}
		break
	case common.MessageType_value["MessageType_MessageType_Failure"]:
		msg := c.m.GetFailure()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (3)"
		} else {
			str = msg.GetMessage()
		}
		break
	case common.MessageType_value["MessageType_MessageType_Entropy"]:
		msg := c.m.GetEntropy()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (10)"
		} else {
			str = hex.EncodeToString(msg.GetEntropy())
		}
		break
	case common.MessageType_value["MessageType_MessageType_PublicKey"]:
		msg := c.m.GetPublicKey()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (12)"
		} else {
			smJSON, _ := json.Marshal(msg)
			str = string(smJSON)
		}
		break
	case common.MessageType_value["MessageType_MessageType_Features"]:
		msg := c.m.GetFeatures()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (17)"
		} else {
			ftsJSON, _ := json.Marshal(msg)
			str = string(ftsJSON)
		}
		break
	case common.MessageType_value["MessageType_MessageType_PinMatrixRequest"]:
		msg := c.m.GetPinMatrixRequest()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (18)"
		} else {
			msgSubType := msg.GetType()
			if msgSubType.String() == "1" {
				str = "Please enter current PIN:"
			} else if msgSubType.String() == "2" {
				str = "Please enter new PIN:"
			} else {
				str = "Please re-enter new PIN:"
			}
		}
		break
	case common.MessageType_value["MessageType_MessageType_TxRequest"]:
		msg := c.m.GetTxRequest()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (21)"
		} else {
			txreq := c.m.GetTxRequest()
			d := msg.GetDetails()
			txreq.SetDetails(d)
			txreq.SetRequestType(msg.GetRequestType())
			smJSON, _ := json.Marshal(msg)
			str = string(smJSON)
		}
		break
	case common.MessageType_value["MessageType_MessageType_ButtonRequest"]:
		msg := c.m.GetButtonRequest()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (26)"
		} else {
			str = "Confirm action on device"
		}
		break
	case common.MessageType_value["MessageType_MessageType_Address"]:
		msg := c.m.GetAddress()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (30)"
		} else {
			str = msg.GetAddress()
		}
		break
	case common.MessageType_value["MessageType_MessageType_EntropyRequest"]:
		externalEntropy, _ := GenerateRandomBytes(32)
		str, msgType = c.Call(c.EntropyAck(externalEntropy))
		break
	case common.MessageType_value["MessageType_MessageType_MessageSignature"]:
		msg := c.m.GetMessageSignature()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (40)"
		} else {
			smJSON, _ := json.Marshal(msg)
			str = string(smJSON)
		}
		break
	case common.MessageType_value["MessageType_MessageType_PassphraseRequest"]:
		msg := c.m.GetPassphraseRequest()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (41)"
		} else {
			str = "Enter your passphrase"
		}
		break
	case common.MessageType_value["MessageType_MessageType_TxSize"]:
		msg := c.m.GetTxSize()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (44)"
		} else {
			str = strconv.Itoa(int(msg.GetTxSize()))
		}
		break
	case common.MessageType_value["MessageType_MessageType_WordRequest"]:
		msg := c.m.GetWordRequest()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (46)"
		} else {
			str = "Enter the word"
		}
		break
	case common.MessageType_value["MessageType_MessageType_CipheredKeyValue"]:
		msg := c.m.GetCipheredKeyValue()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (48)"
		} else {
			str = string(msg.GetValue())
		}
		break
	case common.MessageType_value["MessageType_MessageType_EncryptedMessage"]:
		msg := c.m.GetEncryptedMessage()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (50)"
		} else {
			smJSON, _ := json.Marshal(msg)
			str = string(smJSON)
		}
		break
	case common.MessageType_value["MessageType_MessageType_DecryptedMessage"]:
		msg := c.m.GetDecryptedMessage()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (52)"
		} else {
			str = string(msg.GetMessage())
		}
		break
	case common.MessageType_value["MessageType_MessageType_SignedIdentity"]:
		msg := c.m.GetSignedIdentity()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (54)"
		} else {
			smJSON, _ := json.Marshal(msg)
			str = string(smJSON)
		}
		break
	case common.MessageType_value["MessageType_MessageType_EthereumAddress"]:
		msg := c.m.GetEthereumAddress()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (57)"
		} else {
			str = hex.EncodeToString(msg.GetAddress())
		}
		break
	case common.MessageType_value["MessageType_MessageType_ECDHSessionKey"]:
		msg := c.m.GetECDHSessionKey()
		err = proto.Unmarshal(marshalled, msg)
		if err != nil {
			str = "Error unmarshalling (62)"
		} else {
			str = string(msg.GetSessionKey())
		}
		break
	default:
		break
	}
	return str, msgType
}

func BIP32Path(keys []uint32) string {
	path := "m"
	for _, key := range keys {
		path += "/"
		if key < hardkey {
			path += string(key)
		} else {

			path += string(key-hardkey) + "'"
		}
	}
	return path
}

func StringToBIP32Path(str string) []uint32 {

	if !ValidBIP32(str) {
		return []uint32{}
	}

	re := regexp.MustCompile("([/]+)")
	str = re.ReplaceAllString(str, "/")

	keys := strings.Split(str, "/")
	path := make([]uint32, len(keys)-1)
	for k := 1; k < len(keys); k++ {
		i, _ := strconv.Atoi(strings.Replace(keys[k], "'", "", -1))
		if strings.Contains(keys[k], "'") {
			path[k-1] = hardened(uint32(i))
		} else {
			path[k-1] = uint32(i)
		}
	}
	return path
}

func ValidBIP32(path string) bool {
	re := regexp.MustCompile("([/]+)")
	path = re.ReplaceAllString(path, "/")

	re = regexp.MustCompile("^m/")
	path = re.ReplaceAllString(path, "")

	re = regexp.MustCompile("'/")
	path = re.ReplaceAllString(path+"/", "")

	re = regexp.MustCompile("[0-9/]+")
	path = re.ReplaceAllString(path, "")

	return path == ""
}

func PNGToString(filename string) ([]byte, error) {
	img := make([]byte, 1024)
	infile, err := os.Open(filename)
	if err != nil {
		return img, err
	}
	defer infile.Close()

	src, _, err := image.Decode(infile)
	if err != nil {
		return img, err
	}

	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	if w != 128 || h != 64 {
		err = errors.New("Wrong homescreen size")
		return img, err
	}

	imgBin := ""
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			color := src.At(i, j)
			r, g, b, _ := color.RGBA()
			if (r + g + b) > 0 {
				imgBin += "1"
			} else {
				imgBin += "0"
			}
		}
	}
	k := 0
	for i := 0; i < len(imgBin); i += 8 {
		if s, err := strconv.ParseUint(imgBin[i:i+8], 2, 32); err == nil {
			img[k] = byte(s)
			k++
		}
	}
	return img, nil
}

func URIToIdentity(uri string) types.IdentityTyper {
	var identity types.IdentityTyper
	u, err := url.Parse(uri)
	if err != nil {
		return identity
	}

	defaultPort := ""
	identity.SetProto(&u.Scheme)
	user := ""
	if u.User != nil {
		user = u.User.String()
	}
	identity.SetUser(&user)
	tmp := strings.Split(u.Host, ":")
	identity.SetHost(&tmp[0])
	if len(tmp) > 1 {
		identity.SetPort(&tmp[1])
	} else {
		identity.SetPort(&defaultPort)
	}
	identity.SetPath(&u.Path)
	return identity
}

func hardened(key uint32) uint32 {
	return hardkey + key
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func AES256GCMMEncrypt(plainText, key []byte) ([]byte, []byte) {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	cipheredText := aesgcm.Seal(nil, nonce, plainText, nil)
	return cipheredText, nonce
}

func AES256GCMDecrypt(cipheredText, key, nonce, tag []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	plainText, err := aesgcm.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		return []byte{}, err
	}

	return plainText, nil
}

func GetFileEncKey(masterKey string) (string, string, string) {
	fileKey := masterKey[:len(masterKey)/2]
	encKey := masterKey[len(masterKey)/2:]
	filename_mess := []byte("5f91add3fa1c3c76e90c90a3bd0999e2bd7833d06a483fe884ee60397aca277a")
	mac := hmac.New(sha256.New, []byte(fileKey))
	mac.Write(filename_mess)
	tmpMac := mac.Sum(nil)
	digest := hex.EncodeToString(tmpMac)
	filename := digest + ".pswd"
	return filename, fileKey, encKey
}

func DecryptStorage(content, key string) (Storage, error) {
	cipherKey, _ := hex.DecodeString(key)
	plainText, err := AES256GCMDecrypt([]byte(content[28:]+content[12:28]), cipherKey, []byte(content[:12]), []byte(content[12:28]))

	if err != nil {
		fmt.Println("ERROR DECRYPT STORAGE", err)
		log.Panic("Error decrypting")
	}

	var pc Storage
	fmt.Println("\n\n\nPLAIN JSON FROM SLIP-0016")
	fmt.Println(string(plainText))
	fmt.Println("\n\n\n")
	err = json.Unmarshal(plainText, &pc)
	return pc, err
}

func DecryptEntry(content, key string) (string, error) {
	cipherKey := []byte(key)
	value, err := AES256GCMDecrypt([]byte(content[28:]+content[12:28]), cipherKey, []byte(content[:12]), []byte(content[12:28]))
	return string(value), err
}

func EncryptEntry(content, key string) []byte {
	ciphered, nonce := AES256GCMMEncrypt([]byte(content), []byte(key))
	cipheredText := string(ciphered)
	l := len(ciphered)
	return []byte(string(nonce) + cipheredText[l-16:] + cipheredText[:l-16])
}

func EncryptStorage(s Storage, key string) []byte {
	cipherKey, _ := hex.DecodeString(key)
	content, err := json.Marshal(s)
	if err != nil {
		log.Panic("Error encrypting")
	}

	ciphered, nonce := AES256GCMMEncrypt(content, cipherKey)
	cipheredText := string(ciphered)
	l := len(ciphered)
	return []byte(string(nonce) + cipheredText[l-16:] + cipheredText[:l-16])
}

// TODO : Work on this
func (e *Entry) Equal(entry Entry) bool {
	if e.Title == entry.Title &&
		e.Username == entry.Username &&
		e.Nonce == entry.Nonce &&
		e.Note == entry.Note &&
		reflect.DeepEqual(e.Password.Data, entry.Password.Data) &&
		e.Password.Type == entry.Password.Type &&
		reflect.DeepEqual(e.SafeNote.Data, entry.SafeNote.Data) &&
		e.SafeNote.Type == entry.SafeNote.Type &&
		reflect.DeepEqual(e.Tags, entry.Tags) {
		return true
	}

	return false
}

// TPM uses []int instead of []byte
func (e EncryptedData) MarshalJSON() ([]byte, error) {

	l := len(e.Data)
	dataInt := make([]int, l)
	for i := 0; i < l; i++ {
		dataInt[i] = int(e.Data[i])
	}
	return json.Marshal(&struct {
		Type string `json:"type"`
		Data []int  `json:"data"`
	}{
		Type: e.Type,
		Data: dataInt,
	})
}
