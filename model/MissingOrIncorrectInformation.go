package model

// Indicates the reason for the UnableToApply. It can be missing and/or incorrect information.
type MissingOrIncorrectInformation struct {

	// Indicates the missing information.
	MissingInformation []*UnableToApplyMissingInfo1Code `xml:"MssngInf,omitempty"`

	// Indicates the incorrect information.
	IncorrectInformation []*UnableToApplyIncorrectInfo1Code `xml:"IncrrctInf,omitempty"`
}

func (m *MissingOrIncorrectInformation) AddMissingInformation(value string) {
	m.MissingInformation = append(m.MissingInformation, (*UnableToApplyMissingInfo1Code)(&value))
}

func (m *MissingOrIncorrectInformation) AddIncorrectInformation(value string) {
	m.IncorrectInformation = append(m.IncorrectInformation, (*UnableToApplyIncorrectInfo1Code)(&value))
}
