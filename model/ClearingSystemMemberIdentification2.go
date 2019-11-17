package model

// Unique identification, as assigned by a clearing system, to unambiguously identify a member of the clearing system.
type ClearingSystemMemberIdentification2 struct {

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystemIdentification *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty"`

	// Identification of a member of a clearing system.
	MemberIdentification *Max35Text `xml:"MmbId"`
}

func (c *ClearingSystemMemberIdentification2) AddClearingSystemIdentification() *ClearingSystemIdentification2Choice {
	c.ClearingSystemIdentification = new(ClearingSystemIdentification2Choice)
	return c.ClearingSystemIdentification
}

func (c *ClearingSystemMemberIdentification2) SetMemberIdentification(value string) {
	c.MemberIdentification = (*Max35Text)(&value)
}
