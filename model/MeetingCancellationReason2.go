package model

// Specifies the reason for cancelling a meeting.
type MeetingCancellationReason2 struct {

	// Reason for cancelling a meeting.
	CancellationReasonCode *MeetingCancellationReason1Choice `xml:"CxlRsnCd,omitempty"`

	// Provides more information on the reason for cancelling a meeting in free format form.
	CancellationReason *Max140Text `xml:"CxlRsn,omitempty"`
}

func (m *MeetingCancellationReason2) AddCancellationReasonCode() *MeetingCancellationReason1Choice {
	m.CancellationReasonCode = new(MeetingCancellationReason1Choice)
	return m.CancellationReasonCode
}

func (m *MeetingCancellationReason2) SetCancellationReason(value string) {
	m.CancellationReason = (*Max140Text)(&value)
}
