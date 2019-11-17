package model

// Specifies the reason for cancelling a meeting.
type MeetingCancellationReason1 struct {

	// Specifies the reason for cancelling a meeting in coded form.
	Code *MeetingCancellationReason2Code `xml:"Cd"`

	// Specifies the reason for cancelling a meeting in free text form.
	ExtendedCode *Extended350Code `xml:"XtndedCd"`

	// Provides more information on the reason for cancelling a meeting in free format form.
	CancellationReason *Max140Text `xml:"CxlRsn,omitempty"`
}

func (m *MeetingCancellationReason1) SetCode(value string) {
	m.Code = (*MeetingCancellationReason2Code)(&value)
}

func (m *MeetingCancellationReason1) SetExtendedCode(value string) {
	m.ExtendedCode = (*Extended350Code)(&value)
}

func (m *MeetingCancellationReason1) SetCancellationReason(value string) {
	m.CancellationReason = (*Max140Text)(&value)
}
