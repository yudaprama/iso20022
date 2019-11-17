package model

// Choice of format for the failing reason.
type FailingReason1Choice struct {

	// Specifies the reason why the instruction has a failing settlement status.
	Code *FailingReason1Code `xml:"Cd"`

	// Specifies the reason why the instruction has a failing settlement status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (f *FailingReason1Choice) SetCode(value string) {
	f.Code = (*FailingReason1Code)(&value)
}

func (f *FailingReason1Choice) AddProprietary() *GenericIdentification20 {
	f.Proprietary = new(GenericIdentification20)
	return f.Proprietary
}
