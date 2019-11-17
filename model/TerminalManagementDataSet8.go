package model

// Identification of requested data set.
type TerminalManagementDataSet8 struct {

	// Identification of the required data set.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Transport key encrypted by the TM (Terminal manager) key encryption RSA key.
	EncryptedKey *ContentInformationType7 `xml:"NcrptdKey,omitempty"`
}

func (t *TerminalManagementDataSet8) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet8) SetPOIChallenge(value string) {
	t.POIChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet8) SetTMChallenge(value string) {
	t.TMChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet8) AddEncryptedKey() *ContentInformationType7 {
	t.EncryptedKey = new(ContentInformationType7)
	return t.EncryptedKey
}
