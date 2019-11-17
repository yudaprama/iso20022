package model

// Choice of meeting cancellation reason.
type MeetingCancellationReason1Choice struct {

	// Specifies the reason for cancelling a meeting in coded form.
	Code *MeetingCancellationReason2Code `xml:"Cd"`

	// Specifies the reason for cancelling a meeting in free text form.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (m *MeetingCancellationReason1Choice) SetCode(value string) {
	m.Code = (*MeetingCancellationReason2Code)(&value)
}

func (m *MeetingCancellationReason1Choice) AddProprietary() *GenericIdentification13 {
	m.Proprietary = new(GenericIdentification13)
	return m.Proprietary
}
