package model

// Reason for a suspended status.
type SuspendedStatusReason1 struct {

	// Reason for a suspended status in structured form.
	Structured []*SuspendedStatusReason2Code `xml:"Strd"`

	// Reason for a suspended status in free format text.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (s *SuspendedStatusReason1) AddStructured(value string) {
	s.Structured = append(s.Structured, (*SuspendedStatusReason2Code)(&value))
}

func (s *SuspendedStatusReason1) SetAdditionalInformation(value string) {
	s.AdditionalInformation = (*Max350Text)(&value)
}
