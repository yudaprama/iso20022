package model

// Encrypted personal identification number (PIN) and related information.
type OnLinePIN4 struct {

	// Encrypted PIN (Personal Identification Number).
	EncryptedPINBlock *ContentInformationType10 `xml:"NcrptdPINBlck"`

	// PIN (Personal Identification Number) format before encryption.
	PINFormat *PINFormat3Code `xml:"PINFrmt"`

	// Additional information required to verify the PIN (Personal Identification Number).
	AdditionalInput *Max35Text `xml:"AddtlInpt,omitempty"`
}

func (o *OnLinePIN4) AddEncryptedPINBlock() *ContentInformationType10 {
	o.EncryptedPINBlock = new(ContentInformationType10)
	return o.EncryptedPINBlock
}

func (o *OnLinePIN4) SetPINFormat(value string) {
	o.PINFormat = (*PINFormat3Code)(&value)
}

func (o *OnLinePIN4) SetAdditionalInput(value string) {
	o.AdditionalInput = (*Max35Text)(&value)
}
