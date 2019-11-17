package model

// Choice of format for the failing reason.
type FailingReason8Choice struct {

	// Specifies the reason why the instruction has a failing settlement status.
	Code *FailingReason2Code `xml:"Cd"`

	// Specifies the reason why the instruction has a failing settlement status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (f *FailingReason8Choice) SetCode(value string) {
	f.Code = (*FailingReason2Code)(&value)
}

func (f *FailingReason8Choice) AddProprietary() *GenericIdentification30 {
	f.Proprietary = new(GenericIdentification30)
	return f.Proprietary
}
