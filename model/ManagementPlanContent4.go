package model

// Content of the management plan.
type ManagementPlanContent4 struct {

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Certificate chain of an asymmetric encryption keys for the encryption of temporary transport key of the key to inject.
	KeyEnciphermentCertificate []*Max10KBinary `xml:"KeyNcphrmntCert,omitempty"`

	// Terminal management action to be performed by the point of interaction (POI).
	Action []*TMSAction4 `xml:"Actn"`
}

func (m *ManagementPlanContent4) SetTMChallenge(value string) {
	m.TMChallenge = (*Max140Binary)(&value)
}

func (m *ManagementPlanContent4) AddKeyEnciphermentCertificate(value string) {
	m.KeyEnciphermentCertificate = append(m.KeyEnciphermentCertificate, (*Max10KBinary)(&value))
}

func (m *ManagementPlanContent4) AddAction() *TMSAction4 {
	newValue := new(TMSAction4)
	m.Action = append(m.Action, newValue)
	return newValue
}
