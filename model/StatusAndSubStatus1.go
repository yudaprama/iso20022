package model

// Specifies the status and optionally the sub status.
type StatusAndSubStatus1 struct {

	// Status expressed as a code.
	StatusCode *Status13Choice `xml:"StsCd"`

	// Sub status expressed as a code.
	SubStatusCode *Exact4AlphaNumericText `xml:"SubStsCd,omitempty"`
}

func (s *StatusAndSubStatus1) AddStatusCode() *Status13Choice {
	s.StatusCode = new(Status13Choice)
	return s.StatusCode
}

func (s *StatusAndSubStatus1) SetSubStatusCode(value string) {
	s.SubStatusCode = (*Exact4AlphaNumericText)(&value)
}
