package model

// Provides information on the requested settlement time(s) of the payment instruction.
type SettlementTimeRequest2 struct {

	// Time by which the amount of money must be credited, with confirmation, to the CLS Bank's account at the central bank.
	// Usage: Time must be expressed in Central European Time (CET).
	CLSTime *ISOTime `xml:"CLSTm,omitempty"`

	// Time until when the payment may be settled.
	TillTime *ISOTime `xml:"TillTm,omitempty"`

	// Time as from when the payment may be settled.
	FromTime *ISOTime `xml:"FrTm,omitempty"`

	// Time by when the payment must be settled to avoid rejection.
	RejectTime *ISOTime `xml:"RjctTm,omitempty"`
}

func (s *SettlementTimeRequest2) SetCLSTime(value string) {
	s.CLSTime = (*ISOTime)(&value)
}

func (s *SettlementTimeRequest2) SetTillTime(value string) {
	s.TillTime = (*ISOTime)(&value)
}

func (s *SettlementTimeRequest2) SetFromTime(value string) {
	s.FromTime = (*ISOTime)(&value)
}

func (s *SettlementTimeRequest2) SetRejectTime(value string) {
	s.RejectTime = (*ISOTime)(&value)
}
