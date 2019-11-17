package model

// Set of elements used to provide further information on the reason for the unable to apply investigation.
type MissingOrIncorrectInformation2 struct {

	// Indicates whether the request is related to an AML (Anti Money Laundering) investigation or not.
	AntiMoneyLaunderingRequest *AMLIndicator `xml:"AMLReq,omitempty"`

	// Indicates the missing information.
	MissingInformation []*UnableToApplyMissingInformation2Code `xml:"MssngInf,omitempty"`

	// Indicates, in a coded form, the incorrect information.
	IncorrectInformation []*UnableToApplyIncorrectInformation3Code `xml:"IncrrctInf,omitempty"`
}

func (m *MissingOrIncorrectInformation2) SetAntiMoneyLaunderingRequest(value string) {
	m.AntiMoneyLaunderingRequest = (*AMLIndicator)(&value)
}

func (m *MissingOrIncorrectInformation2) AddMissingInformation(value string) {
	m.MissingInformation = append(m.MissingInformation, (*UnableToApplyMissingInformation2Code)(&value))
}

func (m *MissingOrIncorrectInformation2) AddIncorrectInformation(value string) {
	m.IncorrectInformation = append(m.IncorrectInformation, (*UnableToApplyIncorrectInformation3Code)(&value))
}
