package types

import (
	"strconv"

	"github.com/conejoninja/cerrojo/pb/exchange"
)

type InputScriptTyper interface {
	String() string
	UnmarshalJSON(data []byte) error
}

type TxRequestDetailsTyper interface {
	GetExtraDataLen() uint32
	GetExtraDataOffset() uint32
	Reset()
	ProtoMessage()
	GetTxHash() []byte
	SetExtraDataOffset(*uint32)
	String() string
	GetRequestIndex() uint32
	SetRequestIndex(*uint32)
	SetTxHash([]byte)
	SetExtraDataLen(*uint32)
}

type FailureTyper interface {
	String() string
	UnmarshalJSON(data []byte) error
}

type IdentityTyper interface {
	ProtoMessage()
	GetPort() string
	SetPath(*string)
	SetIndex(*uint32)
	GetProto() string
	GetPath() string
	String() string
	SetHost(*string)
	GetHost() string
	GetIndex() uint32
	SetProto(*string)
	SetUser(*string)
	SetPort(*string)
	Reset()
	GetUser() string
}

type PinMatrixRequestTyper interface {
	String() string
	UnmarshalJSON(data []byte) error
}

type RawTransactionTyper interface {
	SetPayload([]byte)
	GetPayload() []byte
	Reset()
	String() string
	ProtoMessage()
}

type RequestTyper interface {
	String() string
	UnmarshalJSON(data []byte) error
}

type HDNodeTyper interface {
	GetPrivateKey() []byte
	GetDepth() uint32
	GetChainCode() []byte
	SetChildNum(*uint32)
	ProtoMessage()
	SetDepth(*uint32)
	String() string
	Reset()
	GetFingerprint() uint32
	GetPublicKey() []byte
	GetChildNum() uint32
	SetPublicKey([]byte)
	SetFingerprint(*uint32)
	SetChainCode([]byte)
	SetPrivateKey([]byte)
}

type TxInputTyper interface {
	SetMultisig(MultisigRedeemScriptTyper)
	SetAmount(*uint64)
	GetSequence() uint32
	GetAmount() uint64
	String() string
	SetPrevIndex(*uint32)
	GetScriptType() InputScriptTyper
	SetScriptSig([]byte)
	SetSequence(*uint32)
	SetAddressN([]uint32)
	SetPrevHash([]byte)
	GetAddressN() []uint32
	ProtoMessage()
	GetPrevHash() []byte
	GetPrevIndex() uint32
	GetScriptSig() []byte
	GetMultisig() MultisigRedeemScriptTyper
	Reset()
	SetScriptType(InputScriptTyper)
}

type PolicyTyper interface {
	GetEnabled() bool
	SetPolicyName(*string)
	SetEnabled(*bool)
	Reset()
	String() string
	ProtoMessage()
	GetPolicyName() string
}

type OutputAddressTyper interface {
	String() string
	UnmarshalJSON(data []byte) error
}

type RecoveryDeviceTyper interface {
	UnmarshalJSON(data []byte) error
	String() string
}

type TxOutputTyper interface {
	GetAddressN() []uint32
	SetAddress(*string)
	SetAddressType(OutputAddressTyper)
	GetMultisig() MultisigRedeemScriptTyper
	Reset()
	GetExchangeType() ExchangeTyper
	SetScriptType(OutputScriptTyper)
	SetAddressN([]uint32)
	SetAmount(*uint64)
	String() string
	GetScriptType() OutputScriptTyper
	GetAddressType() OutputAddressTyper
	GetOpReturnData() []byte
	SetMultisig(MultisigRedeemScriptTyper)
	SetOpReturnData([]byte)
	SetExchangeType(ExchangeTyper)
	GetAddress() string
	ProtoMessage()
	GetAmount() uint64
}

type TransactionTyper interface {
	SetInputsCnt(*uint32)
	SetOutputsCnt(*uint32)
	SetExtraDataLen(*uint32)
	GetBinOutputs() []TxOutputBinTyper
	SetExtraData([]byte)
	GetOutputs() []TxOutputTyper
	GetExtraData() []byte
	GetExtraDataLen() uint32
	GetVersion() uint32
	GetInputs() []TxInputTyper
	ProtoMessage()
	SetVersion(*uint32)
	SetBinOutputs([]TxOutputBinTyper)
	GetInputsCnt() uint32
	Reset()
	GetOutputsCnt() uint32
	SetLockTime(*uint32)
	GetLockTime() uint32
	String() string
	SetOutputs([]TxOutputTyper)
	SetInputs([]TxInputTyper)
}

type ExchangeTyper interface {
	GetWithdrawalAddressN() []uint32
	GetReturnAddressN() []uint32
	ProtoMessage()
	GetSignedExchangeResponse() *exchange.SignedExchangeResponse
	SetSignedExchangeResponse(*exchange.SignedExchangeResponse)
	SetWithdrawalAddressN([]uint32)
	Reset()
	String() string
	GetWithdrawalCoinName() string
	SetWithdrawalCoinName(*string)
	SetReturnAddressN([]uint32)
}

type TxRequestSerializedTyper interface {
	String() string
	GetSignatureIndex() uint32
	GetSerializedTx() []byte
	SetSignature([]byte)
	ProtoMessage()
	GetSignature() []byte
	Reset()
	SetSignatureIndex(*uint32)
	SetSerializedTx([]byte)
}

type OutputScriptTyper interface {
	String() string
	UnmarshalJSON(data []byte) error
}

type MultisigRedeemScriptTyper interface {
	String() string
	ProtoMessage()
	SetPubkeys([]HDNodePathTyper)
	SetSignatures([][]byte)
	SetM(*uint32)
	GetPubkeys() []HDNodePathTyper
	GetSignatures() [][]byte
	GetM() uint32
	Reset()
}

type WordRequestTyper interface {
	String() string
	UnmarshalJSON(data []byte) error
}

type HDNodePathTyper interface {
	GetAddressN() []uint32
	Reset()
	String() string
	ProtoMessage()
	GetNode() HDNodeTyper
	SetNode(HDNodeTyper)
	SetAddressN([]uint32)
}

type CoinTyper interface {
	SetAddressTypeP2Sh(*uint32)
	SetAddressTypeP2Wpkh(*uint32)
	SetCoinShortcut(*string)
	Reset()
	String() string
	GetCoinShortcut() string
	GetAddressType() uint32
	GetBip44AccountPath() uint32
	SetMaxfeeKb(*uint64)
	SetSignedMessageHeader(*string)
	SetAddressType(*uint32)
	GetMaxfeeKb() uint64
	GetAddressTypeP2Sh() uint32
	GetAddressTypeP2Wpkh() uint32
	GetCoinName() string
	ProtoMessage()
	GetSignedMessageHeader() string
	SetBip44AccountPath(*uint32)
	GetAddressTypeP2Wsh() uint32
	SetAddressTypeP2Wsh(*uint32)
	SetCoinName(*string)
}

type TxOutputBinTyper interface {
	SetScriptPubkey([]byte)
	GetAmount() uint64
	GetScriptPubkey() []byte
	Reset()
	String() string
	ProtoMessage()
	SetAmount(*uint64)
}

type ButtonRequestTyper interface {
	String() string
	UnmarshalJSON(data []byte) error
}

type Typer interface {
	GetInputScriptType() InputScriptTyper
	GetTxRequestDetailsType() TxRequestDetailsTyper
	GetFailureType() FailureTyper
	GetIdentityType() IdentityTyper
	GetPinMatrixRequestType() PinMatrixRequestTyper
	GetRawTransactionType() RawTransactionTyper
	GetRequestType() RequestTyper
	GetHDNodeType() HDNodeTyper
	GetTxInputType() TxInputTyper
	GetPolicyType() PolicyTyper
	GetOutputAddressType() OutputAddressTyper
	GetRecoveryDeviceType() RecoveryDeviceTyper
	GetTxOutputType() TxOutputTyper
	GetTransactionType() TransactionTyper
	GetExchangeType() ExchangeTyper
	GetTxOutputBinType() TxOutputBinTyper
	GetButtonRequestType() ButtonRequestTyper
	GetTxRequestSerializedType() TxRequestSerializedTyper
	GetOutputScriptType() OutputScriptTyper
	GetMultisigRedeemScriptType() MultisigRedeemScriptTyper
	GetWordRequestType() WordRequestTyper
	GetHDNodePathType() HDNodePathTyper
	GetCoinType() CoinTyper
}

var InputScriptType_name = map[int32]string{
	0: "InputScriptType_SPENDADDRESS",
	1: "InputScriptType_SPENDMULTISIG",
	2: "InputScriptType_EXTERNAL",
	3: "InputScriptType_SPENDWITNESS",
	4: "SHWITNESS",
}

var RequestType_name = map[int32]string{
	0: "RequestType_TXINPUT",
	1: "RequestType_TXOUTPUT",
	2: "RequestType_TXMETA",
	3: "RequestType_TXFINISHED",
	4: "RequestType_TXEXTRADATA",
}

var RecoveryDeviceType_name = map[int32]string{
	0: "RecoveryDeviceType_RecoveryDeviceType_ScrambledWords",
	1: "RecoveryDeviceType_RecoveryDeviceType_Matrix",
}

var OutputScriptType_name = map[int32]string{
	0: "OutputScriptType_PAYTOADDRESS",
	1: "OutputScriptType_PAYTOSCRIPTHASH",
	2: "OutputScriptType_PAYTOMULTISIG",
	3: "OutputScriptType_PAYTOOPRETURN",
	4: "OutputScriptType_PAYTOWITNESS",
	5: "SHWITNESS",
}

var OutputAddressType_name = map[int32]string{
	0: "OutputAddressType_SPEND",
	1: "OutputAddressType_TRANSFER",
	2: "OutputAddressType_CHANGE",
	3: "OutputAddressType_EXCHANGE",
}

var PinMatrixRequestType_name = map[int32]string{
	1: "PinMatrixRequestType_PinMatrixRequestType_Current",
	2: "PinMatrixRequestType_PinMatrixRequestType_NewFirst",
	3: "PinMatrixRequestType_PinMatrixRequestType_NewSecond",
}

var WordRequestType_name = map[int32]string{
	0: "WordRequestType_WordRequestType_Plain",
}

var FailureType_name = map[int32]string{
	1:  "FailureType_Failure_UnexpectedMessage",
	2:  "FailureType_Failure_ButtonExpected",
	3:  "FailureType_Failure_SyntaxError",
	4:  "FailureType_Failure_ActionCancelled",
	5:  "FailureType_Failure_PinExpected",
	6:  "FailureType_Failure_PinCancelled",
	7:  "FailureType_Failure_PinInvalid",
	8:  "FailureType_Failure_InvalidSignature",
	9:  "FailureType_Failure_Other",
	10: "FailureType_Failure_NotEnoughFunds",
	11: "FailureType_Failure_NotInitialized",
	99: "FailureType_Failure_FirmwareError",
}

var PinMatrixRequestType_value = map[string]int32{
	"PinMatrixRequestType_PinMatrixRequestType_Current":   1,
	"PinMatrixRequestType_PinMatrixRequestType_NewFirst":  2,
	"PinMatrixRequestType_PinMatrixRequestType_NewSecond": 3,
}

var WordRequestType_value = map[string]int32{
	"WordRequestType_WordRequestType_Plain": 0,
}

var FailureType_value = map[string]int32{
	"FailureType_Failure_UnexpectedMessage": 1,
	"FailureType_Failure_ButtonExpected":    2,
	"FailureType_Failure_SyntaxError":       3,
	"FailureType_Failure_ActionCancelled":   4,
	"FailureType_Failure_PinExpected":       5,
	"FailureType_Failure_PinCancelled":      6,
	"FailureType_Failure_PinInvalid":        7,
	"FailureType_Failure_InvalidSignature":  8,
	"FailureType_Failure_Other":             9,
	"FailureType_Failure_NotEnoughFunds":    10,
	"FailureType_Failure_NotInitialized":    11,
	"FailureType_Failure_FirmwareError":     99,
}

var InputScriptType_value = map[string]int32{
	"InputScriptType_SPENDADDRESS":  0,
	"InputScriptType_SPENDMULTISIG": 1,
	"InputScriptType_EXTERNAL":      2,
	"InputScriptType_SPENDWITNESS":  3,
	"SHWITNESS":                     4,
}

var RequestType_value = map[string]int32{
	"RequestType_TXINPUT":     0,
	"RequestType_TXOUTPUT":    1,
	"RequestType_TXMETA":      2,
	"RequestType_TXFINISHED":  3,
	"RequestType_TXEXTRADATA": 4,
}

var RecoveryDeviceType_value = map[string]int32{
	"RecoveryDeviceType_RecoveryDeviceType_ScrambledWords": 0,
	"RecoveryDeviceType_RecoveryDeviceType_Matrix":         1,
}

var OutputScriptType_value = map[string]int32{
	"OutputScriptType_PAYTOADDRESS":    0,
	"OutputScriptType_PAYTOSCRIPTHASH": 1,
	"OutputScriptType_PAYTOMULTISIG":   2,
	"OutputScriptType_PAYTOOPRETURN":   3,
	"OutputScriptType_PAYTOWITNESS":    4,
	"SHWITNESS":                        5,
}

var OutputAddressType_value = map[string]int32{
	"OutputAddressType_SPEND":    0,
	"OutputAddressType_TRANSFER": 1,
	"OutputAddressType_CHANGE":   2,
	"OutputAddressType_EXCHANGE": 3,
}

type OutputAddressType int32

func OutputAddressTyper2Type(x OutputAddressTyper) OutputAddressType {
	value := (x).String()
	tmp, _ := strconv.Atoi(value)
	return OutputAddressType(int32(tmp))
}

type FailureType int32

func FailureTyper2Type(x FailureTyper) FailureType {
	value := (x).String()
	tmp, _ := strconv.Atoi(value)
	return FailureType(int32(tmp))
}

type OutputScriptType int32

func OutputScriptTyper2Type(x OutputScriptTyper) OutputScriptType {
	value := (x).String()
	tmp, _ := strconv.Atoi(value)
	return OutputScriptType(int32(tmp))
}

type InputScriptType int32

func InputScriptTyper2Type(x InputScriptTyper) InputScriptType {
	value := (x).String()
	tmp, _ := strconv.Atoi(value)
	return InputScriptType(int32(tmp))
}

type ButtonRequestType int32

func ButtonRequestTyper2Type(x ButtonRequestTyper) ButtonRequestType {
	value := (x).String()
	tmp, _ := strconv.Atoi(value)
	return ButtonRequestType(int32(tmp))
}

type PinMatrixRequestType int32

func PinMatrixRequestTyper2Type(x PinMatrixRequestTyper) PinMatrixRequestType {
	value := (x).String()
	tmp, _ := strconv.Atoi(value)
	return PinMatrixRequestType(int32(tmp))
}

type RecoveryDeviceType int32

func RecoveryDeviceTyper2Type(x RecoveryDeviceTyper) RecoveryDeviceType {
	value := (x).String()
	tmp, _ := strconv.Atoi(value)
	return RecoveryDeviceType(int32(tmp))
}

type WordRequestType int32

func WordRequestTyper2Type(x WordRequestTyper) WordRequestType {
	value := (x).String()
	tmp, _ := strconv.Atoi(value)
	return WordRequestType(int32(tmp))
}

type RequestType int32

func RequestTyper2Type(x RequestTyper) RequestType {
	value := (x).String()
	tmp, _ := strconv.Atoi(value)
	return RequestType(int32(tmp))
}
