package model

// Identifies the mandate, which is being accepted.
type MandateAcceptance1 struct {

	// Set of elements used to provide information on the original messsage.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Set of elements used to provide detailed information on the acceptance result.
	AcceptanceResult *AcceptanceResult6 `xml:"AccptncRslt"`

	// Set of elements used to provide the original mandate data.
	OriginalMandate *OriginalMandate1Choice `xml:"OrgnlMndt"`
}

func (m *MandateAcceptance1) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateAcceptance1) AddAcceptanceResult() *AcceptanceResult6 {
	m.AcceptanceResult = new(AcceptanceResult6)
	return m.AcceptanceResult
}

func (m *MandateAcceptance1) AddOriginalMandate() *OriginalMandate1Choice {
	m.OriginalMandate = new(OriginalMandate1Choice)
	return m.OriginalMandate
}
