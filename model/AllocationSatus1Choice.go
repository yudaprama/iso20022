package model

// Choice of format for the allocation status reason.
type AllocationSatus1Choice struct {

	// Provides the status of allocation of collateral to cover the instruction.
	Code *AllocationStatus1Code `xml:"Cd"`

	// Provides the status of allocation of collateral to cover the instruction.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AllocationSatus1Choice) SetCode(value string) {
	a.Code = (*AllocationStatus1Code)(&value)
}

func (a *AllocationSatus1Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}
