package model

// Identifies the mandate to be amended and gives details of the new mandate.
type MandateAmendment1 struct {

	// Set of elements used to provide information on the original messsage.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Set of elements used to provide detailed information on the amendment reason.
	AmendmentReason *AmendmentReasonInformation1 `xml:"AmdmntRsn"`

	// Set of elements used to provide the amended mandate data.
	Mandate *MandateInformation3 `xml:"Mndt"`

	// Set of elements used to provide the original mandate data.
	OriginalMandate *OriginalMandate1Choice `xml:"OrgnlMndt"`
}

func (m *MandateAmendment1) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateAmendment1) AddAmendmentReason() *AmendmentReasonInformation1 {
	m.AmendmentReason = new(AmendmentReasonInformation1)
	return m.AmendmentReason
}

func (m *MandateAmendment1) AddMandate() *MandateInformation3 {
	m.Mandate = new(MandateInformation3)
	return m.Mandate
}

func (m *MandateAmendment1) AddOriginalMandate() *OriginalMandate1Choice {
	m.OriginalMandate = new(OriginalMandate1Choice)
	return m.OriginalMandate
}
