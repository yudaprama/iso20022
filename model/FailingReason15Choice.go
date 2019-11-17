package model

// Choice of format for the failing reason.
type FailingReason15Choice struct {

	// Specifies the reason why the instruction has a failing settlement status.
	Code *FailingReason1Code `xml:"Cd"`

	// Specifies the reason why the instruction has a failing settlement status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *FailingReason15Choice) SetCode(value string) {
	f.Code = (*FailingReason1Code)(&value)
}

func (f *FailingReason15Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
