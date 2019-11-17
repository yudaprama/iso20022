package model

// Identifies the mandate to be amended and gives details of the new mandate.
type MandateAmendment2 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the amendment reason.
	AmendmentReason *MandateAmendmentReason1 `xml:"AmdmntRsn"`

	// Provides the amended mandate data.
	Mandate *Mandate3 `xml:"Mndt"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate2Choice `xml:"OrgnlMndt"`
}

func (m *MandateAmendment2) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateAmendment2) AddAmendmentReason() *MandateAmendmentReason1 {
	m.AmendmentReason = new(MandateAmendmentReason1)
	return m.AmendmentReason
}

func (m *MandateAmendment2) AddMandate() *Mandate3 {
	m.Mandate = new(Mandate3)
	return m.Mandate
}

func (m *MandateAmendment2) AddOriginalMandate() *OriginalMandate2Choice {
	m.OriginalMandate = new(OriginalMandate2Choice)
	return m.OriginalMandate
}
