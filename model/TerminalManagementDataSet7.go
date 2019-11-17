package model

// Identification of requested data set.
type TerminalManagementDataSet7 struct {

	// Identification of the required data set.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Transport key encrypted by the TM key encryption RSA key.
	EncryptedKey *ContentInformationType5 `xml:"NcrptdKey,omitempty"`
}

func (t *TerminalManagementDataSet7) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet7) SetPOIChallenge(value string) {
	t.POIChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet7) SetTMChallenge(value string) {
	t.TMChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet7) AddEncryptedKey() *ContentInformationType5 {
	t.EncryptedKey = new(ContentInformationType5)
	return t.EncryptedKey
}
