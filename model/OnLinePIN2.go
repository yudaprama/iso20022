package model

// Encrypted personal identification number (PIN) and related information.
type OnLinePIN2 struct {

	// Encrypted PIN (Personal Identification Number).
	EncryptedPINBlock *ContentInformationType5 `xml:"NcrptdPINBlck"`

	// PIN format before encryption.
	PINFormat *PINFormat2Code `xml:"PINFrmt"`

	// Additional information required to verify the PIN.
	AdditionalInput *Max35Text `xml:"AddtlInpt,omitempty"`
}

func (o *OnLinePIN2) AddEncryptedPINBlock() *ContentInformationType5 {
	o.EncryptedPINBlock = new(ContentInformationType5)
	return o.EncryptedPINBlock
}

func (o *OnLinePIN2) SetPINFormat(value string) {
	o.PINFormat = (*PINFormat2Code)(&value)
}

func (o *OnLinePIN2) SetAdditionalInput(value string) {
	o.AdditionalInput = (*Max35Text)(&value)
}
