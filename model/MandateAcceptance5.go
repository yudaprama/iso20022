package model

// Identifies the mandate, which is being accepted.
type MandateAcceptance5 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the acceptance result.
	AcceptanceResult *AcceptanceResult6 `xml:"AccptncRslt"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate5Choice `xml:"OrgnlMndt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateAcceptance5) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateAcceptance5) AddAcceptanceResult() *AcceptanceResult6 {
	m.AcceptanceResult = new(AcceptanceResult6)
	return m.AcceptanceResult
}

func (m *MandateAcceptance5) AddOriginalMandate() *OriginalMandate5Choice {
	m.OriginalMandate = new(OriginalMandate5Choice)
	return m.OriginalMandate
}

func (m *MandateAcceptance5) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
