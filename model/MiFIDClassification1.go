package model

// Details about the MIFID classification of the account owner.
type MiFIDClassification1 struct {

	// MiFID classification of the account owner.
	Classification *OrderOriginatorEligibility1Code `xml:"Clssfctn"`

	// Additional information about the source of classification.
	Narrative *Max350Text `xml:"Nrrtv,omitempty"`
}

func (m *MiFIDClassification1) SetClassification(value string) {
	m.Classification = (*OrderOriginatorEligibility1Code)(&value)
}

func (m *MiFIDClassification1) SetNarrative(value string) {
	m.Narrative = (*Max350Text)(&value)
}
