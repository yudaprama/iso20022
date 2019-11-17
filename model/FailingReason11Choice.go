package model

// Choice of format for the failing reason.
type FailingReason11Choice struct {

	// Specifies the reason why the instruction has a failing settlement status.
	Code *FailingReason2Code `xml:"Cd"`

	// Specifies the reason why the instruction has a failing settlement status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *FailingReason11Choice) SetCode(value string) {
	f.Code = (*FailingReason2Code)(&value)
}

func (f *FailingReason11Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
