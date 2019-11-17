package model

// Identification of requested data set.
type TerminalManagementDataSet12 struct {

	// Identification of the required data set.
	Identification *DataSetIdentification4 `xml:"Id"`

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Temporary encryption key that the host will use for protecting keys to download.
	SessionKey *CryptographicKey5 `xml:"SsnKey,omitempty"`
}

func (t *TerminalManagementDataSet12) AddIdentification() *DataSetIdentification4 {
	t.Identification = new(DataSetIdentification4)
	return t.Identification
}

func (t *TerminalManagementDataSet12) SetPOIChallenge(value string) {
	t.POIChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet12) SetTMChallenge(value string) {
	t.TMChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet12) AddSessionKey() *CryptographicKey5 {
	t.SessionKey = new(CryptographicKey5)
	return t.SessionKey
}
