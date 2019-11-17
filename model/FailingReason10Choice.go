package model

// Choice of format for the failing reason.
type FailingReason10Choice struct {

	// Specifies the reason why the instruction has a failing settlement status.
	Code *FailingReason3Code `xml:"Cd"`

	// Specifies the reason why the instruction has a failing settlement status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *FailingReason10Choice) SetCode(value string) {
	f.Code = (*FailingReason3Code)(&value)
}

func (f *FailingReason10Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
