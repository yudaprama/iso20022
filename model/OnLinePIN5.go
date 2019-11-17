package model

// Encrypted personal identification number (PIN) and related information.
type OnLinePIN5 struct {

	// Encrypted PIN (Personal Identification Number).
	EncryptedPINBlock *ContentInformationType10 `xml:"NcrptdPINBlck"`

	// PIN (Personal Identification Number) format before encryption.
	PINFormat *PINFormat4Code `xml:"PINFrmt"`

	// Additional information required to verify the PIN (Personal Identification Number).
	AdditionalInput *Max35Text `xml:"AddtlInpt,omitempty"`
}

func (o *OnLinePIN5) AddEncryptedPINBlock() *ContentInformationType10 {
	o.EncryptedPINBlock = new(ContentInformationType10)
	return o.EncryptedPINBlock
}

func (o *OnLinePIN5) SetPINFormat(value string) {
	o.PINFormat = (*PINFormat4Code)(&value)
}

func (o *OnLinePIN5) SetAdditionalInput(value string) {
	o.AdditionalInput = (*Max35Text)(&value)
}
