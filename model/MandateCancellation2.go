package model

// Identifies the mandate to be cancelled.
type MandateCancellation2 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReason *PaymentCancellationReason1 `xml:"CxlRsn"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate2Choice `xml:"OrgnlMndt"`
}

func (m *MandateCancellation2) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateCancellation2) AddCancellationReason() *PaymentCancellationReason1 {
	m.CancellationReason = new(PaymentCancellationReason1)
	return m.CancellationReason
}

func (m *MandateCancellation2) AddOriginalMandate() *OriginalMandate2Choice {
	m.OriginalMandate = new(OriginalMandate2Choice)
	return m.OriginalMandate
}
