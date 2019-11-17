package model

// Set of elements used to provide further information on the reason for the unable to apply investigation.
type MissingOrIncorrectInformation3 struct {

	// Indicates whether the request is related to an AML (Anti Money Laundering) investigation or not.
	AntiMoneyLaunderingRequest *AMLIndicator `xml:"AMLReq,omitempty"`

	// Indicates the missing information.
	MissingInformation []*UnableToApplyMissing1 `xml:"MssngInf,omitempty"`

	// Indicates, in a coded form, the incorrect information.
	IncorrectInformation []*UnableToApplyIncorrect1 `xml:"IncrrctInf,omitempty"`
}

func (m *MissingOrIncorrectInformation3) SetAntiMoneyLaunderingRequest(value string) {
	m.AntiMoneyLaunderingRequest = (*AMLIndicator)(&value)
}

func (m *MissingOrIncorrectInformation3) AddMissingInformation() *UnableToApplyMissing1 {
	newValue := new(UnableToApplyMissing1)
	m.MissingInformation = append(m.MissingInformation, newValue)
	return newValue
}

func (m *MissingOrIncorrectInformation3) AddIncorrectInformation() *UnableToApplyIncorrect1 {
	newValue := new(UnableToApplyIncorrect1)
	m.IncorrectInformation = append(m.IncorrectInformation, newValue)
	return newValue
}
