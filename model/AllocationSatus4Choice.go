package model

// Choice of format for the allocation status reason.
type AllocationSatus4Choice struct {

	// Provides the status of allocation of collateral to cover the instruction.
	Code *AllocationStatus1Code `xml:"Cd"`

	// Provides the status of allocation of collateral to cover the instruction.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AllocationSatus4Choice) SetCode(value string) {
	a.Code = (*AllocationStatus1Code)(&value)
}

func (a *AllocationSatus4Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
