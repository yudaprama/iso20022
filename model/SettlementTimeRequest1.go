package model

// Information on the requested settlement time of the instruction.
type SettlementTimeRequest1 struct {

	// Time by which the funds must be credited, with confirmation, to the CLS Bank's account at the central bank, expressed in Central European Time (CET).
	CLSTime *ISOTime `xml:"CLSTm"`
}

func (s *SettlementTimeRequest1) SetCLSTime(value string) {
	s.CLSTime = (*ISOTime)(&value)
}
