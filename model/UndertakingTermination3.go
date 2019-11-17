package model

// Information about the notification of the termination of an undertaking.
type UndertakingTermination3 struct {

	// Date on which the termination is effective.
	EffectiveDate *ISODate `xml:"FctvDt"`

	// Reason for the termination.
	Reason *TerminationReason1Choice `xml:"Rsn,omitempty"`

	// Additional information related to the termination.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *UndertakingTermination3) SetEffectiveDate(value string) {
	u.EffectiveDate = (*ISODate)(&value)
}

func (u *UndertakingTermination3) AddReason() *TerminationReason1Choice {
	u.Reason = new(TerminationReason1Choice)
	return u.Reason
}

func (u *UndertakingTermination3) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}
