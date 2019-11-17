package model

// Identifies the mandate, which is being suspended.
type MandateSuspension1 struct {

	// Unique identification, as assigned by the initiating party, to unambiguously identify the suspension request.
	SuspensionRequestIdentification *Max35Text `xml:"SspnsnReqId"`

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides detailed information on the suspension reason.
	SuspensionReason *MandateSuspensionReason1 `xml:"SspnsnRsn"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate4Choice `xml:"OrgnlMndt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateSuspension1) SetSuspensionRequestIdentification(value string) {
	m.SuspensionRequestIdentification = (*Max35Text)(&value)
}

func (m *MandateSuspension1) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateSuspension1) AddSuspensionReason() *MandateSuspensionReason1 {
	m.SuspensionReason = new(MandateSuspensionReason1)
	return m.SuspensionReason
}

func (m *MandateSuspension1) AddOriginalMandate() *OriginalMandate4Choice {
	m.OriginalMandate = new(OriginalMandate4Choice)
	return m.OriginalMandate
}

func (m *MandateSuspension1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
