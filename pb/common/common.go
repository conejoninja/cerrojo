package common

import (
	"github.com/conejoninja/cerrojo/pb/types"
)

type MessageType int32

type ChangePiner interface {
	Reset()
	String() string
	SetRemove(*bool)
	ProtoMessage()
	GetRemove() bool
}

type Failurer interface {
	GetCode() types.FailureTyper
	GetMessage() string
	SetCode(types.FailureTyper)
	SetMessage(*string)
	Reset()
	String() string
	ProtoMessage()
}

type Initializer interface {
	String() string
	ProtoMessage()
	Reset()
}

type SignMessager interface {
	GetAddressN() []uint32
	GetMessage() []byte
	SetCoinName(*string)
	SetAddressN([]uint32)
	SetMessage([]byte)
	String() string
	ProtoMessage()
	GetCoinName() string
	Reset()
}

type DebugLinkFillConfiger interface {
	ProtoMessage()
	Reset()
	String() string
}

type RecoveryDevicer interface {
	GetPinProtection() bool
	GetLabel() string
	SetPinProtection(*bool)
	SetUseCharacterCipher(*bool)
	Reset()
	String() string
	GetUseCharacterCipher() bool
	SetEnforceWordlist(*bool)
	SetPassphraseProtection(*bool)
	ProtoMessage()
	GetLanguage() string
	GetType() uint32
	SetLanguage(*string)
	SetWordCount(*uint32)
	SetU2FCounter(*uint32)
	GetWordCount() uint32
	GetU2FCounter() uint32
	GetPassphraseProtection() bool
	GetEnforceWordlist() bool
	SetLabel(*string)
}

type MessageTyper interface {
	String() string
	UnmarshalJSON(data []byte) error
}

type CipheredKeyValuer interface {
	Reset()
	String() string
	ProtoMessage()
	GetValue() []byte
	SetValue([]byte)
}

type GetFeatureser interface {
	ProtoMessage()
	Reset()
	String() string
}

type WordRequester interface {
	Reset()
	String() string
	ProtoMessage()
	GetType() types.WordRequestTyper
	SetType(types.WordRequestTyper)
}

type Addresser interface {
	GetAddress() string
	SetAddress(*string)
	Reset()
	String() string
	ProtoMessage()
}

type GetECDHSessionKeyer interface {
	GetIdentity() types.IdentityTyper
	GetEcdsaCurveName() string
	SetIdentity(types.IdentityTyper)
	SetPeerPublicKey([]byte)
	SetEcdsaCurveName(*string)
	Reset()
	String() string
	ProtoMessage()
	GetPeerPublicKey() []byte
}

type SetU2FCounterer interface {
	Reset()
	String() string
	ProtoMessage()
	GetU2FCounter() uint32
	SetU2FCounter(*uint32)
}

type CipherKeyValuer interface {
	Reset()
	GetAddressN() []uint32
	SetKey(*string)
	String() string
	SetEncrypt(*bool)
	SetAskOnDecrypt(*bool)
	SetIv([]byte)
	GetIv() []byte
	GetKey() string
	SetValue([]byte)
	SetAskOnEncrypt(*bool)
	SetAddressN([]uint32)
	GetEncrypt() bool
	GetAskOnEncrypt() bool
	GetAskOnDecrypt() bool
	ProtoMessage()
	GetValue() []byte
}

type SignedIdentityer interface {
	GetSignature() []byte
	SetPublicKey([]byte)
	String() string
	GetPublicKey() []byte
	GetAddress() string
	SetSignature([]byte)
	SetAddress(*string)
	Reset()
	ProtoMessage()
}

type EthereumSignTxer interface {
	GetGasLimit() []byte
	GetTo() []byte
	SetTo([]byte)
	GetGasPrice() []byte
	GetDataLength() uint32
	GetChainId() uint32
	GetAddressType() types.OutputAddressTyper
	SetDataLength(*uint32)
	SetDataInitialChunk([]byte)
	SetToAddressN([]uint32)
	ProtoMessage()
	GetToAddressN() []uint32
	SetValue([]byte)
	SetAddressType(types.OutputAddressTyper)
	GetDataInitialChunk() []byte
	GetNonce() []byte
	SetAddressN([]uint32)
	GetAddressN() []uint32
	Reset()
	SetExchangeType(types.ExchangeTyper)
	GetValue() []byte
	SetGasLimit([]byte)
	GetExchangeType() types.ExchangeTyper
	SetGasPrice([]byte)
	SetNonce([]byte)
	SetChainId(*uint32)
	String() string
}

type DebugLinkLoger interface {
	GetText() string
	SetLevel(*uint32)
	SetBucket(*string)
	Reset()
	ProtoMessage()
	GetBucket() string
	SetText(*string)
	String() string
	GetLevel() uint32
}

type DebugLinkGetStater interface {
	Reset()
	String() string
	ProtoMessage()
}

type FirmwareUploader interface {
	Reset()
	GetHash() []byte
	GetPayloadHash() []byte
	String() string
	ProtoMessage()
	GetPayload() []byte
	SetPayload([]byte)
	SetHash([]byte)
	SetPayloadHash([]byte)
}

type PublicKeyer interface {
	GetNode() types.HDNodeTyper
	GetXpub() string
	SetXpub(*string)
	SetNode(types.HDNodeTyper)
	Reset()
	String() string
	ProtoMessage()
}

type DebugLinkMemoryReader interface {
	SetLength(*uint32)
	GetAddress() uint32
	GetLength() uint32
	Reset()
	String() string
	ProtoMessage()
	SetAddress(*uint32)
}

type EthereumAddresser interface {
	GetAddress() []byte
	SetAddress([]byte)
	Reset()
	String() string
	ProtoMessage()
}

type FirmwareRequester interface {
	String() string
	ProtoMessage()
	GetOffset() uint32
	GetLength() uint32
	SetOffset(*uint32)
	SetLength(*uint32)
	Reset()
}

type DebugLinkStater interface {
	ProtoMessage()
	GetRecoveryFakeWord() string
	GetRecoveryCipher() string
	SetMnemonic(*string)
	SetPin(*string)
	SetResetWord(*string)
	SetRecoveryCipher(*string)
	GetMatrix() string
	SetRecoveryFakeWord(*string)
	GetRecoveryWordPos() uint32
	SetNode(types.HDNodeTyper)
	GetNode() types.HDNodeTyper
	GetLayout() []byte
	GetResetWord() string
	SetResetEntropy([]byte)
	SetPassphraseProtection(*bool)
	Reset()
	GetRecoveryAutoCompletedWord() string
	SetLayout([]byte)
	SetFirmwareHash([]byte)
	SetRecoveryAutoCompletedWord(*string)
	SetStorageHash([]byte)
	String() string
	GetMnemonic() string
	GetResetEntropy() []byte
	GetFirmwareHash() []byte
	GetStorageHash() []byte
	SetMatrix(*string)
	GetPassphraseProtection() bool
	GetPin() string
	SetRecoveryWordPos(*uint32)
}

type Featureser interface {
	SetPatchVersion(*uint32)
	GetPinProtection() bool
	String() string
	SetLanguage(*string)
	SetCoins([]types.CoinTyper)
	GetBootloaderMode() bool
	SetLabel(*string)
	GetPolicies() []types.PolicyTyper
	GetInitialized() bool
	GetMinorVersion() uint32
	SetPinProtection(*bool)
	SetInitialized(*bool)
	GetCoins() []types.CoinTyper
	GetBootloaderHash() []byte
	GetImported() bool
	GetDeviceId() string
	GetVendor() string
	GetPassphraseProtection() bool
	GetRevision() []byte
	GetPinCached() bool
	GetFirmwarePresent() bool
	GetLanguage() string
	SetPassphraseCached(*bool)
	SetBootloaderHash([]byte)
	SetBootloaderMode(*bool)
	SetImported(*bool)
	SetPinCached(*bool)
	SetVendor(*string)
	SetPassphraseProtection(*bool)
	SetMajorVersion(*uint32)
	SetFirmwarePresent(*bool)
	GetPatchVersion() uint32
	Reset()
	SetPolicies([]types.PolicyTyper)
	SetRevision([]byte)
	ProtoMessage()
	SetMinorVersion(*uint32)
	GetPassphraseCached() bool
	SetDeviceId(*string)
	GetMajorVersion() uint32
	GetLabel() string
}

type WipeDevicer interface {
	String() string
	ProtoMessage()
	Reset()
}

type PassphraseAcker interface {
	ProtoMessage()
	GetPassphrase() string
	Reset()
	String() string
	SetPassphrase(*string)
}

type DebugLinkFlashEraser interface {
	GetSector() uint32
	Reset()
	String() string
	ProtoMessage()
	SetSector(*uint32)
}

type PassphraseRequester interface {
	ProtoMessage()
	Reset()
	String() string
}

type ApplySettingser interface {
	GetUsePassphrase() bool
	GetHomescreen() []byte
	SetLabel(*string)
	SetUsePassphrase(*bool)
	GetLanguage() string
	GetLabel() string
	String() string
	SetLanguage(*string)
	SetHomescreen([]byte)
	ProtoMessage()
	Reset()
}

type Pinger interface {
	SetMessage(*string)
	SetButtonProtection(*bool)
	GetPassphraseProtection() bool
	Reset()
	GetMessage() string
	GetButtonProtection() bool
	SetPassphraseProtection(*bool)
	String() string
	ProtoMessage()
	GetPinProtection() bool
	SetPinProtection(*bool)
}

type DebugLinkMemoryer interface {
	Reset()
	String() string
	ProtoMessage()
	GetMemory() []byte
	SetMemory([]byte)
}

type DebugLinkStoper interface {
	String() string
	ProtoMessage()
	Reset()
}

type RawTxAcker interface {
	Reset()
	String() string
	ProtoMessage()
	GetTx() types.RawTransactionTyper
	SetTx(types.RawTransactionTyper)
}

type EncryptedMessager interface {
	SetNonce([]byte)
	SetHmac([]byte)
	Reset()
	ProtoMessage()
	GetMessage() []byte
	GetHmac() []byte
	SetMessage([]byte)
	String() string
	GetNonce() []byte
}

type LoadDevicer interface {
	GetPassphraseProtection() bool
	GetLabel() string
	SetNode(types.HDNodeTyper)
	SetPassphraseProtection(*bool)
	GetMnemonic() string
	GetPin() string
	GetLanguage() string
	Reset()
	GetSkipChecksum() bool
	ProtoMessage()
	SetLanguage(*string)
	SetLabel(*string)
	SetMnemonic(*string)
	SetPin(*string)
	GetNode() types.HDNodeTyper
	String() string
	GetU2FCounter() uint32
	SetSkipChecksum(*bool)
	SetU2FCounter(*uint32)
}

type ButtonAcker interface {
	Reset()
	String() string
	ProtoMessage()
}

type CharacterRequester interface {
	GetWordPos() uint32
	GetCharacterPos() uint32
	Reset()
	String() string
	ProtoMessage()
	SetWordPos(*uint32)
	SetCharacterPos(*uint32)
}

type Canceler interface {
	Reset()
	String() string
	ProtoMessage()
}

type EthereumTxAcker interface {
	GetDataChunk() []byte
	Reset()
	String() string
	ProtoMessage()
	SetDataChunk([]byte)
}

type ResetDevicer interface {
	GetLanguage() string
	SetStrength(*uint32)
	SetLabel(*string)
	SetPassphraseProtection(*bool)
	SetLanguage(*string)
	GetPinProtection() bool
	ProtoMessage()
	GetLabel() string
	GetU2FCounter() uint32
	SetU2FCounter(*uint32)
	SetDisplayRandom(*bool)
	Reset()
	String() string
	GetPassphraseProtection() bool
	SetPinProtection(*bool)
	GetDisplayRandom() bool
	GetStrength() uint32
}

type TxSizer interface {
	Reset()
	String() string
	ProtoMessage()
	GetTxSize() uint32
	SetTxSize(*uint32)
}

type ClearSessioner interface {
	ProtoMessage()
	Reset()
	String() string
}

type TxAcker interface {
	GetTx() types.TransactionTyper
	SetTx(types.TransactionTyper)
	Reset()
	String() string
	ProtoMessage()
}

type DecryptedMessager interface {
	String() string
	ProtoMessage()
	GetMessage() []byte
	GetAddress() string
	SetMessage([]byte)
	SetAddress(*string)
	Reset()
}

type Entropyer interface {
	ProtoMessage()
	GetEntropy() []byte
	SetEntropy([]byte)
	Reset()
	String() string
}

type DecryptMessager interface {
	SetNonce([]byte)
	SetMessage([]byte)
	SetHmac([]byte)
	GetNonce() []byte
	GetMessage() []byte
	GetHmac() []byte
	ProtoMessage()
	GetAddressN() []uint32
	Reset()
	String() string
	SetAddressN([]uint32)
}

type EthereumTxRequester interface {
	GetSignatureR() []byte
	SetSignatureV(*uint32)
	SetSignatureS([]byte)
	SetSignatureDer([]byte)
	String() string
	GetDataLength() uint32
	GetSignatureV() uint32
	GetSignatureS() []byte
	GetSignatureDer() []byte
	Reset()
	ProtoMessage()
	GetHash() []byte
	SetSignatureR([]byte)
	SetDataLength(*uint32)
	SetHash([]byte)
}

type ECDHSessionKeyer interface {
	Reset()
	String() string
	ProtoMessage()
	SetSessionKey([]byte)
	GetSessionKey() []byte
}

type PinMatrixRequester interface {
	ProtoMessage()
	GetType() types.PinMatrixRequestTyper
	SetType(types.PinMatrixRequestTyper)
	Reset()
	String() string
}

type EstimateTxSizer interface {
	GetOutputsCount() uint32
	GetCoinName() string
	SetOutputsCount(*uint32)
	SetInputsCount(*uint32)
	SetCoinName(*string)
	Reset()
	String() string
	ProtoMessage()
	GetInputsCount() uint32
}

type FirmwareEraser interface {
	Reset()
	String() string
	ProtoMessage()
}

type VerifyMessager interface {
	GetSignature() []byte
	GetMessage() []byte
	SetCoinName(*string)
	SetMessage([]byte)
	SetAddress(*string)
	SetSignature([]byte)
	Reset()
	String() string
	ProtoMessage()
	GetAddress() string
	GetCoinName() string
}

type TxRequester interface {
	SetRequestType(types.RequestTyper)
	GetDetails() types.TxRequestDetailsTyper
	GetSerialized() types.TxRequestSerializedTyper
	Reset()
	SetDetails(types.TxRequestDetailsTyper)
	SetSerialized(types.TxRequestSerializedTyper)
	ProtoMessage()
	GetRequestType() types.RequestTyper
	String() string
}

type GetPublicKeyer interface {
	Reset()
	ProtoMessage()
	GetAddressN() []uint32
	SetEcdsaCurveName(*string)
	SetShowDisplay(*bool)
	String() string
	GetEcdsaCurveName() string
	GetShowDisplay() bool
	SetAddressN([]uint32)
}

type DebugLinkDecisioner interface {
	Reset()
	String() string
	ProtoMessage()
	GetYesNo() bool
	SetYesNo(*bool)
}

type PinMatrixAcker interface {
	GetPin() string
	Reset()
	String() string
	ProtoMessage()
	SetPin(*string)
}

type EntropyAcker interface {
	Reset()
	String() string
	ProtoMessage()
	GetEntropy() []byte
	SetEntropy([]byte)
}

type EncryptMessager interface {
	GetAddressN() []uint32
	Reset()
	String() string
	GetPubkey() []byte
	SetAddressN([]uint32)
	GetDisplayOnly() bool
	ProtoMessage()
	GetMessage() []byte
	SetPubkey([]byte)
	SetMessage([]byte)
	SetDisplayOnly(*bool)
	SetCoinName(*string)
	GetCoinName() string
}

type SimpleSignTxer interface {
	GetVersion() uint32
	GetLockTime() uint32
	SetVersion(*uint32)
	SetTransactions([]types.TransactionTyper)
	GetOutputs() []types.TxOutputTyper
	GetTransactions() []types.TransactionTyper
	SetLockTime(*uint32)
	SetInputs([]types.TxInputTyper)
	SetOutputs([]types.TxOutputTyper)
	ProtoMessage()
	GetInputs() []types.TxInputTyper
	Reset()
	SetCoinName(*string)
	GetCoinName() string
	String() string
}

type WordAcker interface {
	SetWord(*string)
	Reset()
	String() string
	ProtoMessage()
	GetWord() string
}

type EthereumGetAddresser interface {
	String() string
	SetAddressN([]uint32)
	SetShowDisplay(*bool)
	ProtoMessage()
	GetAddressN() []uint32
	GetShowDisplay() bool
	Reset()
}

type MessageSignaturer interface {
	ProtoMessage()
	GetAddress() string
	GetSignature() []byte
	Reset()
	String() string
	SetAddress(*string)
	SetSignature([]byte)
}

type SignTxer interface {
	GetLockTime() uint32
	Reset()
	SetInputsCount(*uint32)
	SetCoinName(*string)
	SetVersion(*uint32)
	SetLockTime(*uint32)
	GetInputsCount() uint32
	GetCoinName() string
	GetVersion() uint32
	String() string
	ProtoMessage()
	GetOutputsCount() uint32
	SetOutputsCount(*uint32)
}

type GetEntropyer interface {
	GetSize() uint32
	SetSize(*uint32)
	Reset()
	String() string
	ProtoMessage()
}

type CharacterAcker interface {
	SetDone(*bool)
	GetDelete() bool
	String() string
	ProtoMessage()
	GetCharacter() string
	GetDone() bool
	SetCharacter(*string)
	SetDelete(*bool)
	Reset()
}

type EntropyRequester interface {
	Reset()
	String() string
	ProtoMessage()
}

type ApplyPolicieser interface {
	ProtoMessage()
	GetPolicy() []types.PolicyTyper
	Reset()
	String() string
	SetPolicy([]types.PolicyTyper)
}

type ButtonRequester interface {
	String() string
	ProtoMessage()
	GetCode() types.ButtonRequestTyper
	GetData() string
	SetCode(types.ButtonRequestTyper)
	SetData(*string)
	Reset()
}

type Successer interface {
	GetMessage() string
	Reset()
	String() string
	ProtoMessage()
	SetMessage(*string)
}

type GetAddresser interface {
	GetCoinName() string
	GetScriptType() types.InputScriptTyper
	ProtoMessage()
	SetMultisig(types.MultisigRedeemScriptTyper)
	SetScriptType(types.InputScriptTyper)
	SetCoinName(*string)
	SetShowDisplay(*bool)
	SetAddressN([]uint32)
	GetAddressN() []uint32
	GetShowDisplay() bool
	GetMultisig() types.MultisigRedeemScriptTyper
	Reset()
	String() string
}

type DebugLinkMemoryWriter interface {
	GetMemory() []byte
	GetFlash() bool
	ProtoMessage()
	SetAddress(*uint32)
	GetAddress() uint32
	Reset()
	String() string
	SetMemory([]byte)
	SetFlash(*bool)
}

type SignIdentityer interface {
	SetIdentity(types.IdentityTyper)
	SetEcdsaCurveName(*string)
	ProtoMessage()
	GetChallengeVisual() string
	GetEcdsaCurveName() string
	GetChallengeHidden() []byte
	SetChallengeHidden([]byte)
	SetChallengeVisual(*string)
	Reset()
	String() string
	GetIdentity() types.IdentityTyper
}

type Messager interface {
	GetResetDevice() ResetDevicer
	GetEthereumTxAck() EthereumTxAcker
	GetTxAck() TxAcker
	GetTxSize() TxSizer
	GetClearSession() ClearSessioner
	GetPinMatrixRequest() PinMatrixRequester
	GetEstimateTxSize() EstimateTxSizer
	GetFirmwareErase() FirmwareEraser
	GetDecryptedMessage() DecryptedMessager
	GetEntropy() Entropyer
	GetDecryptMessage() DecryptMessager
	GetEthereumTxRequest() EthereumTxRequester
	GetECDHSessionKey() ECDHSessionKeyer
	GetTxRequest() TxRequester
	GetVerifyMessage() VerifyMessager
	GetEncryptMessage() EncryptMessager
	GetGetPublicKey() GetPublicKeyer
	GetDebugLinkDecision() DebugLinkDecisioner
	GetPinMatrixAck() PinMatrixAcker
	GetEntropyAck() EntropyAcker
	GetGetEntropy() GetEntropyer
	GetCharacterAck() CharacterAcker
	GetEntropyRequest() EntropyRequester
	GetSimpleSignTx() SimpleSignTxer
	GetWordAck() WordAcker
	GetEthereumGetAddress() EthereumGetAddresser
	GetMessageSignature() MessageSignaturer
	GetSignTx() SignTxer
	GetApplyPolicies() ApplyPolicieser
	GetGetAddress() GetAddresser
	GetButtonRequest() ButtonRequester
	GetSuccess() Successer
	GetSignIdentity() SignIdentityer
	GetDebugLinkMemoryWrite() DebugLinkMemoryWriter
	GetInitialize() Initializer
	GetChangePin() ChangePiner
	GetFailure() Failurer
	GetRecoveryDevice() RecoveryDevicer
	GetSignMessage() SignMessager
	GetDebugLinkFillConfig() DebugLinkFillConfiger
	GetWordRequest() WordRequester
	GetMessageType() MessageTyper
	GetCipheredKeyValue() CipheredKeyValuer
	GetGetFeatures() GetFeatureser
	GetEthereumSignTx() EthereumSignTxer
	GetDebugLinkLog() DebugLinkLoger
	GetAddress() Addresser
	GetGetECDHSessionKey() GetECDHSessionKeyer
	GetSetU2FCounter() SetU2FCounterer
	GetCipherKeyValue() CipherKeyValuer
	GetSignedIdentity() SignedIdentityer
	GetFirmwareUpload() FirmwareUploader
	GetDebugLinkGetState() DebugLinkGetStater
	GetDebugLinkState() DebugLinkStater
	GetPublicKey() PublicKeyer
	GetDebugLinkMemoryRead() DebugLinkMemoryReader
	GetEthereumAddress() EthereumAddresser
	GetFirmwareRequest() FirmwareRequester
	GetPassphraseAck() PassphraseAcker
	GetFeatures() Featureser
	GetWipeDevice() WipeDevicer
	GetDebugLinkStop() DebugLinkStoper
	GetRawTxAck() RawTxAcker
	GetEncryptedMessage() EncryptedMessager
	GetDebugLinkFlashErase() DebugLinkFlashEraser
	GetPassphraseRequest() PassphraseRequester
	GetApplySettings() ApplySettingser
	GetPing() Pinger
	GetDebugLinkMemory() DebugLinkMemoryer
	GetCancel() Canceler
	GetLoadDevice() LoadDevicer
	GetButtonAck() ButtonAcker
	GetCharacterRequest() CharacterRequester
}

var MessageType_name = map[MessageType]string{
	0:   "MessageType_MessageType_Initialize",
	1:   "MessageType_MessageType_Ping",
	2:   "MessageType_MessageType_Success",
	3:   "MessageType_MessageType_Failure",
	4:   "MessageType_MessageType_ChangePin",
	5:   "MessageType_MessageType_WipeDevice",
	6:   "MessageType_MessageType_FirmwareErase",
	7:   "MessageType_MessageType_FirmwareUpload",
	8:   "MessageType_MessageType_FirmwareRequest",
	9:   "MessageType_MessageType_GetEntropy",
	10:  "MessageType_MessageType_Entropy",
	11:  "MessageType_MessageType_GetPublicKey",
	12:  "MessageType_MessageType_PublicKey",
	13:  "MessageType_MessageType_LoadDevice",
	14:  "MessageType_MessageType_ResetDevice",
	15:  "MessageType_MessageType_SignTx",
	16:  "MessageType_MessageType_SimpleSignTx",
	17:  "MessageType_MessageType_Features",
	18:  "MessageType_MessageType_PinMatrixRequest",
	19:  "MessageType_MessageType_PinMatrixAck",
	20:  "MessageType_MessageType_Cancel",
	21:  "MessageType_MessageType_TxRequest",
	22:  "MessageType_MessageType_TxAck",
	23:  "MessageType_MessageType_CipherKeyValue",
	24:  "MessageType_MessageType_ClearSession",
	25:  "MessageType_MessageType_ApplySettings",
	26:  "MessageType_MessageType_ButtonRequest",
	27:  "MessageType_MessageType_ButtonAck",
	29:  "MessageType_MessageType_GetAddress",
	30:  "MessageType_MessageType_Address",
	35:  "MessageType_MessageType_EntropyRequest",
	36:  "MessageType_MessageType_EntropyAck",
	38:  "MessageType_MessageType_SignMessage",
	39:  "MessageType_MessageType_VerifyMessage",
	40:  "MessageType_MessageType_MessageSignature",
	41:  "MessageType_MessageType_PassphraseRequest",
	42:  "MessageType_MessageType_PassphraseAck",
	43:  "MessageType_MessageType_EstimateTxSize",
	44:  "MessageType_MessageType_TxSize",
	45:  "MessageType_MessageType_RecoveryDevice",
	46:  "MessageType_MessageType_WordRequest",
	47:  "MessageType_MessageType_WordAck",
	48:  "MessageType_MessageType_CipheredKeyValue",
	49:  "MessageType_MessageType_EncryptMessage",
	50:  "MessageType_MessageType_EncryptedMessage",
	51:  "MessageType_MessageType_DecryptMessage",
	52:  "MessageType_MessageType_DecryptedMessage",
	53:  "MessageType_MessageType_SignIdentity",
	54:  "MessageType_MessageType_SignedIdentity",
	55:  "MessageType_MessageType_GetFeatures",
	56:  "MessageType_MessageType_EthereumGetAddress",
	57:  "MessageType_MessageType_EthereumAddress",
	58:  "MessageType_MessageType_EthereumSignTx",
	59:  "MessageType_MessageType_EthereumTxRequest",
	60:  "MessageType_MessageType_EthereumTxAck",
	61:  "MessageType_MessageType_GetECDHSessionKey",
	62:  "MessageType_MessageType_ECDHSessionKey",
	63:  "FCounter",
	80:  "MessageType_MessageType_CharacterRequest",
	81:  "MessageType_MessageType_CharacterAck",
	82:  "MessageType_MessageType_RawTxAck",
	83:  "MessageType_MessageType_ApplyPolicies",
	100: "MessageType_MessageType_DebugLinkDecision",
	101: "MessageType_MessageType_DebugLinkGetState",
	102: "MessageType_MessageType_DebugLinkState",
	103: "MessageType_MessageType_DebugLinkStop",
	104: "MessageType_MessageType_DebugLinkLog",
	105: "MessageType_MessageType_DebugLinkFillConfig",
	110: "MessageType_MessageType_DebugLinkMemoryRead",
	111: "MessageType_MessageType_DebugLinkMemory",
	112: "MessageType_MessageType_DebugLinkMemoryWrite",
	113: "MessageType_MessageType_DebugLinkFlashErase",
}

var MessageType_value = map[string]MessageType{
	"MessageType_MessageType_Initialize":         0,
	"MessageType_MessageType_Ping":               1,
	"MessageType_MessageType_Success":            2,
	"MessageType_MessageType_Failure":            3,
	"MessageType_MessageType_ChangePin":          4,
	"MessageType_MessageType_WipeDevice":         5,
	"MessageType_MessageType_FirmwareErase":      6,
	"MessageType_MessageType_FirmwareUpload":     7,
	"MessageType_MessageType_FirmwareRequest":    8,
	"MessageType_MessageType_GetEntropy":         9,
	"MessageType_MessageType_Entropy":            10,
	"MessageType_MessageType_GetPublicKey":       11,
	"MessageType_MessageType_PublicKey":          12,
	"MessageType_MessageType_LoadDevice":         13,
	"MessageType_MessageType_ResetDevice":        14,
	"MessageType_MessageType_SignTx":             15,
	"MessageType_MessageType_SimpleSignTx":       16,
	"MessageType_MessageType_Features":           17,
	"MessageType_MessageType_PinMatrixRequest":   18,
	"MessageType_MessageType_PinMatrixAck":       19,
	"MessageType_MessageType_Cancel":             20,
	"MessageType_MessageType_TxRequest":          21,
	"MessageType_MessageType_TxAck":              22,
	"MessageType_MessageType_CipherKeyValue":     23,
	"MessageType_MessageType_ClearSession":       24,
	"MessageType_MessageType_ApplySettings":      25,
	"MessageType_MessageType_ButtonRequest":      26,
	"MessageType_MessageType_ButtonAck":          27,
	"MessageType_MessageType_GetAddress":         29,
	"MessageType_MessageType_Address":            30,
	"MessageType_MessageType_EntropyRequest":     35,
	"MessageType_MessageType_EntropyAck":         36,
	"MessageType_MessageType_SignMessage":        38,
	"MessageType_MessageType_VerifyMessage":      39,
	"MessageType_MessageType_MessageSignature":   40,
	"MessageType_MessageType_PassphraseRequest":  41,
	"MessageType_MessageType_PassphraseAck":      42,
	"MessageType_MessageType_EstimateTxSize":     43,
	"MessageType_MessageType_TxSize":             44,
	"MessageType_MessageType_RecoveryDevice":     45,
	"MessageType_MessageType_WordRequest":        46,
	"MessageType_MessageType_WordAck":            47,
	"MessageType_MessageType_CipheredKeyValue":   48,
	"MessageType_MessageType_EncryptMessage":     49,
	"MessageType_MessageType_EncryptedMessage":   50,
	"MessageType_MessageType_DecryptMessage":     51,
	"MessageType_MessageType_DecryptedMessage":   52,
	"MessageType_MessageType_SignIdentity":       53,
	"MessageType_MessageType_SignedIdentity":     54,
	"MessageType_MessageType_GetFeatures":        55,
	"MessageType_MessageType_EthereumGetAddress": 56,
	"MessageType_MessageType_EthereumAddress":    57,
	"MessageType_MessageType_EthereumSignTx":     58,
	"MessageType_MessageType_EthereumTxRequest":  59,
	"MessageType_MessageType_EthereumTxAck":      60,
	"MessageType_MessageType_GetECDHSessionKey":  61,
	"MessageType_MessageType_ECDHSessionKey":     62,
	"FCounter": 63,
	"MessageType_MessageType_CharacterRequest":     80,
	"MessageType_MessageType_CharacterAck":         81,
	"MessageType_MessageType_RawTxAck":             82,
	"MessageType_MessageType_ApplyPolicies":        83,
	"MessageType_MessageType_DebugLinkDecision":    100,
	"MessageType_MessageType_DebugLinkGetState":    101,
	"MessageType_MessageType_DebugLinkState":       102,
	"MessageType_MessageType_DebugLinkStop":        103,
	"MessageType_MessageType_DebugLinkLog":         104,
	"MessageType_MessageType_DebugLinkFillConfig":  105,
	"MessageType_MessageType_DebugLinkMemoryRead":  110,
	"MessageType_MessageType_DebugLinkMemory":      111,
	"MessageType_MessageType_DebugLinkMemoryWrite": 112,
	"MessageType_MessageType_DebugLinkFlashErase":  113,
}
