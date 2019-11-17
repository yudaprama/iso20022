package model

// Choice of meeting type classification.
type MeetingTypeClassification1Choice struct {

	// Classifies the type of meeting.
	Code *MeetingTypeClassification1Code `xml:"Cd"`

	// Specifies the reason for cancelling a meeting in free text form.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (m *MeetingTypeClassification1Choice) SetCode(value string) {
	m.Code = (*MeetingTypeClassification1Code)(&value)
}

func (m *MeetingTypeClassification1Choice) AddProprietary() *GenericIdentification13 {
	m.Proprietary = new(GenericIdentification13)
	return m.Proprietary
}
