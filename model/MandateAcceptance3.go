package model

// Identifies the mandate, which is being accepted.
type MandateAcceptance3 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the acceptance result.
	AcceptanceResult *AcceptanceResult6 `xml:"AccptncRslt"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate2Choice `xml:"OrgnlMndt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateAcceptance3) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateAcceptance3) AddAcceptanceResult() *AcceptanceResult6 {
	m.AcceptanceResult = new(AcceptanceResult6)
	return m.AcceptanceResult
}

func (m *MandateAcceptance3) AddOriginalMandate() *OriginalMandate2Choice {
	m.OriginalMandate = new(OriginalMandate2Choice)
	return m.OriginalMandate
}

func (m *MandateAcceptance3) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
