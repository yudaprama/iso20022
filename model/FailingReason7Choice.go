package model

// Choice of format for the failing reason.
type FailingReason7Choice struct {

	// Specifies the reason why the instruction has a failing settlement status.
	Code *FailingReason3Code `xml:"Cd"`

	// Specifies the reason why the instruction has a failing settlement status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (f *FailingReason7Choice) SetCode(value string) {
	f.Code = (*FailingReason3Code)(&value)
}

func (f *FailingReason7Choice) AddProprietary() *GenericIdentification30 {
	f.Proprietary = new(GenericIdentification30)
	return f.Proprietary
}
