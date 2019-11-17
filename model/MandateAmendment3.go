package model

// Identifies the mandate to be amended and gives details of the new mandate.
type MandateAmendment3 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the amendment reason.
	AmendmentReason *MandateAmendmentReason1 `xml:"AmdmntRsn"`

	// Provides the amended mandate data.
	Mandate *Mandate3 `xml:"Mndt"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate2Choice `xml:"OrgnlMndt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateAmendment3) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateAmendment3) AddAmendmentReason() *MandateAmendmentReason1 {
	m.AmendmentReason = new(MandateAmendmentReason1)
	return m.AmendmentReason
}

func (m *MandateAmendment3) AddMandate() *Mandate3 {
	m.Mandate = new(Mandate3)
	return m.Mandate
}

func (m *MandateAmendment3) AddOriginalMandate() *OriginalMandate2Choice {
	m.OriginalMandate = new(OriginalMandate2Choice)
	return m.OriginalMandate
}

func (m *MandateAmendment3) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
