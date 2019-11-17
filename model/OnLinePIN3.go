package model

// Encrypted personal identification number (PIN) and related information.
type OnLinePIN3 struct {

	// Encrypted PIN (Personal Identification Number).
	EncryptedPINBlock *ContentInformationType7 `xml:"NcrptdPINBlck"`

	// PIN (Personal Identification Number) format before encryption.
	PINFormat *PINFormat3Code `xml:"PINFrmt"`

	// Additional information required to verify the PIN (Personal Identification Number.
	AdditionalInput *Max35Text `xml:"AddtlInpt,omitempty"`
}

func (o *OnLinePIN3) AddEncryptedPINBlock() *ContentInformationType7 {
	o.EncryptedPINBlock = new(ContentInformationType7)
	return o.EncryptedPINBlock
}

func (o *OnLinePIN3) SetPINFormat(value string) {
	o.PINFormat = (*PINFormat3Code)(&value)
}

func (o *OnLinePIN3) SetAdditionalInput(value string) {
	o.AdditionalInput = (*Max35Text)(&value)
}
