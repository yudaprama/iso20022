package model

// Identifies the mandate to be cancelled.
type MandateCancellation1 struct {

	// Set of elements used to provide information on the original messsage.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Set of elements used to provide detailed information on the cancellation reason.
	CancellationReason *CancellationReasonInformation2 `xml:"CxlRsn"`

	// Set of elements used to provide the original mandate data.
	OriginalMandate *OriginalMandate1Choice `xml:"OrgnlMndt"`
}

func (m *MandateCancellation1) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateCancellation1) AddCancellationReason() *CancellationReasonInformation2 {
	m.CancellationReason = new(CancellationReasonInformation2)
	return m.CancellationReason
}

func (m *MandateCancellation1) AddOriginalMandate() *OriginalMandate1Choice {
	m.OriginalMandate = new(OriginalMandate1Choice)
	return m.OriginalMandate
}
