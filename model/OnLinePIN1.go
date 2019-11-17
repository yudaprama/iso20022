package model

// Encrypted personal identification number (PIN) and related information.
type OnLinePIN1 struct {

	// Encrypted PIN (Personal Identification Number).
	EncryptedPINBlock *ContentInformationType2 `xml:"NcrptdPINBlck"`

	// PIN format before encryption.
	PINFormat *PINFormat1Code `xml:"PINFrmt"`

	// Additional information required to verify the PIN.
	AdditionalInput *Max35Text `xml:"AddtlInpt,omitempty"`
}

func (o *OnLinePIN1) AddEncryptedPINBlock() *ContentInformationType2 {
	o.EncryptedPINBlock = new(ContentInformationType2)
	return o.EncryptedPINBlock
}

func (o *OnLinePIN1) SetPINFormat(value string) {
	o.PINFormat = (*PINFormat1Code)(&value)
}

func (o *OnLinePIN1) SetAdditionalInput(value string) {
	o.AdditionalInput = (*Max35Text)(&value)
}
