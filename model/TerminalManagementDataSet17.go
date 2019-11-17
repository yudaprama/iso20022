package model

// Identification of requested data set.
type TerminalManagementDataSet17 struct {

	// Identification of the required data set.
	Identification *DataSetIdentification6 `xml:"Id"`

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Temporary encryption key that the host will use for protecting keys to download.
	SessionKey *CryptographicKey5 `xml:"SsnKey,omitempty"`

	// Proof of delegation to be validated by the terminal manager receiving a status report from a new POI.
	DelegationProof *Max5000Binary `xml:"DlgtnProof,omitempty"`

	// Protected proof of delegation.
	ProtectedDelegationProof *ContentInformationType12 `xml:"PrtctdDlgtnProof,omitempty"`
}

func (t *TerminalManagementDataSet17) AddIdentification() *DataSetIdentification6 {
	t.Identification = new(DataSetIdentification6)
	return t.Identification
}

func (t *TerminalManagementDataSet17) SetPOIChallenge(value string) {
	t.POIChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet17) SetTMChallenge(value string) {
	t.TMChallenge = (*Max140Binary)(&value)
}

func (t *TerminalManagementDataSet17) AddSessionKey() *CryptographicKey5 {
	t.SessionKey = new(CryptographicKey5)
	return t.SessionKey
}

func (t *TerminalManagementDataSet17) SetDelegationProof(value string) {
	t.DelegationProof = (*Max5000Binary)(&value)
}

func (t *TerminalManagementDataSet17) AddProtectedDelegationProof() *ContentInformationType12 {
	t.ProtectedDelegationProof = new(ContentInformationType12)
	return t.ProtectedDelegationProof
}
