package model

// Choice of format for the failing reason.
type FailingReason9Choice struct {

	// Specifies the reason why the instruction has a failing settlement status.
	Code *FailingReason1Code `xml:"Cd"`

	// Specifies the reason why the instruction has a failing settlement status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (f *FailingReason9Choice) SetCode(value string) {
	f.Code = (*FailingReason1Code)(&value)
}

func (f *FailingReason9Choice) AddProprietary() *GenericIdentification30 {
	f.Proprietary = new(GenericIdentification30)
	return f.Proprietary
}
