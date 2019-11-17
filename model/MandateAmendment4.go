package model

// Identifies the mandate to be amended and gives details of the new mandate.
type MandateAmendment4 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the amendment reason.
	AmendmentReason *MandateAmendmentReason1 `xml:"AmdmntRsn"`

	// Provides the amended mandate data.
	Mandate *Mandate6 `xml:"Mndt"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate3Choice `xml:"OrgnlMndt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateAmendment4) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateAmendment4) AddAmendmentReason() *MandateAmendmentReason1 {
	m.AmendmentReason = new(MandateAmendmentReason1)
	return m.AmendmentReason
}

func (m *MandateAmendment4) AddMandate() *Mandate6 {
	m.Mandate = new(Mandate6)
	return m.Mandate
}

func (m *MandateAmendment4) AddOriginalMandate() *OriginalMandate3Choice {
	m.OriginalMandate = new(OriginalMandate3Choice)
	return m.OriginalMandate
}

func (m *MandateAmendment4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
