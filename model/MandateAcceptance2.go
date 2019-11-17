package model

// Identifies the mandate, which is being accepted.
type MandateAcceptance2 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the acceptance result.
	AcceptanceResult *AcceptanceResult6 `xml:"AccptncRslt"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate2Choice `xml:"OrgnlMndt"`
}

func (m *MandateAcceptance2) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateAcceptance2) AddAcceptanceResult() *AcceptanceResult6 {
	m.AcceptanceResult = new(AcceptanceResult6)
	return m.AcceptanceResult
}

func (m *MandateAcceptance2) AddOriginalMandate() *OriginalMandate2Choice {
	m.OriginalMandate = new(OriginalMandate2Choice)
	return m.OriginalMandate
}
