package model

// Choice of format for the Awaiting Affirmation reason.
type AwaitingAffirmationReason1Choice struct {

	// Specifies the reason why the instruction has been alleged.
	Code *AwaitingAffirmationReason1Code `xml:"Cd"`

	// Specifies the reason why the instruction has been alleged.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (a *AwaitingAffirmationReason1Choice) SetCode(value string) {
	a.Code = (*AwaitingAffirmationReason1Code)(&value)
}

func (a *AwaitingAffirmationReason1Choice) AddProprietary() *GenericIdentification38 {
	a.Proprietary = new(GenericIdentification38)
	return a.Proprietary
}
