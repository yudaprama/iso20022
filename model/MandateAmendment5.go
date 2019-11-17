package model

// Identifies the mandate to be amended and gives details of the new mandate.
type MandateAmendment5 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the amendment reason.
	AmendmentReason *MandateAmendmentReason1 `xml:"AmdmntRsn"`

	// Provides the amended mandate data.
	Mandate *Mandate8 `xml:"Mndt"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate4Choice `xml:"OrgnlMndt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateAmendment5) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateAmendment5) AddAmendmentReason() *MandateAmendmentReason1 {
	m.AmendmentReason = new(MandateAmendmentReason1)
	return m.AmendmentReason
}

func (m *MandateAmendment5) AddMandate() *Mandate8 {
	m.Mandate = new(Mandate8)
	return m.Mandate
}

func (m *MandateAmendment5) AddOriginalMandate() *OriginalMandate4Choice {
	m.OriginalMandate = new(OriginalMandate4Choice)
	return m.OriginalMandate
}

func (m *MandateAmendment5) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
