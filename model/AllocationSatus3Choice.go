package model

// Choice of format for the allocation status reason.
type AllocationSatus3Choice struct {

	// Provides the status of allocation of collateral to cover the instruction.
	Code *AllocationStatus1Code `xml:"Cd"`

	// Provides the status of allocation of collateral to cover the instruction.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (a *AllocationSatus3Choice) SetCode(value string) {
	a.Code = (*AllocationStatus1Code)(&value)
}

func (a *AllocationSatus3Choice) AddProprietary() *GenericIdentification30 {
	a.Proprietary = new(GenericIdentification30)
	return a.Proprietary
}
