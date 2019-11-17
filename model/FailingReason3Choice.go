package model

// Choice of format for the failing reason.
type FailingReason3Choice struct {

	// Specifies the reason why the instruction has a failing settlement status.
	Code *FailingReason2Code `xml:"Cd"`

	// Specifies the reason why the instruction has a failing settlement status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (f *FailingReason3Choice) SetCode(value string) {
	f.Code = (*FailingReason2Code)(&value)
}

func (f *FailingReason3Choice) AddProprietary() *GenericIdentification20 {
	f.Proprietary = new(GenericIdentification20)
	return f.Proprietary
}
