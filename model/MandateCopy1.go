package model

// Identifies the mandate, for which a copy of the details is requested.
type MandateCopy1 struct {

	// Provides information on the original message.
	OriginalMessageInformation *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty"`

	// Provides the original mandate data.
	OriginalMandate *OriginalMandate4Choice `xml:"OrgnlMndt"`

	// Indicates the status of the mandate.
	MandateStatus *MandateStatus1Choice `xml:"MndtSts,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateCopy1) AddOriginalMessageInformation() *OriginalMessageInformation1 {
	m.OriginalMessageInformation = new(OriginalMessageInformation1)
	return m.OriginalMessageInformation
}

func (m *MandateCopy1) AddOriginalMandate() *OriginalMandate4Choice {
	m.OriginalMandate = new(OriginalMandate4Choice)
	return m.OriginalMandate
}

func (m *MandateCopy1) AddMandateStatus() *MandateStatus1Choice {
	m.MandateStatus = new(MandateStatus1Choice)
	return m.MandateStatus
}

func (m *MandateCopy1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}
