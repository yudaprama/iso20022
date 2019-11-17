package model

// Identifies the mandate to be cancelled.
type MandateCancellation4 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReason *PaymentCancellationReason1 `xml:"CxlRsn"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate3Choice `xml:"OrgnlMndt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateCancellation4) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateCancellation4) AddCancellationReason() *PaymentCancellationReason1 {
	m.CancellationReason = new(PaymentCancellationReason1)
	return m.CancellationReason
}

func (m *MandateCancellation4) AddOriginalMandate() *OriginalMandate3Choice {
	m.OriginalMandate = new(OriginalMandate3Choice)
	return m.OriginalMandate
}

func (m *MandateCancellation4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
