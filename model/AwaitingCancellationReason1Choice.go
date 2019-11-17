package model

// Choice of format for the allegment reason.
type AwaitingCancellationReason1Choice struct {

	// Specifies the reason why the instruction has been alleged.
	Code *AwaitingCancellationReason1Code `xml:"Cd"`

	// Specifies the reason why the instruction has been alleged.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (a *AwaitingCancellationReason1Choice) SetCode(value string) {
	a.Code = (*AwaitingCancellationReason1Code)(&value)
}

func (a *AwaitingCancellationReason1Choice) AddProprietary() *GenericIdentification38 {
	a.Proprietary = new(GenericIdentification38)
	return a.Proprietary
}
