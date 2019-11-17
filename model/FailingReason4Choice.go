package model

// Choice of format for the failing reason.
type FailingReason4Choice struct {

	// Specifies the reason why the instruction has a failing settlement status.
	Code *FailingReason3Code `xml:"Cd"`

	// Specifies the reason why the instruction has a failing settlement status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (f *FailingReason4Choice) SetCode(value string) {
	f.Code = (*FailingReason3Code)(&value)
}

func (f *FailingReason4Choice) AddProprietary() *GenericIdentification20 {
	f.Proprietary = new(GenericIdentification20)
	return f.Proprietary
}
