package model

// Reason for the closure of the contract.
type ContractClosureReason1Choice struct {

	// Reason in a coded form.
	Code *ExternalContractClosureReason1Code `xml:"Cd"`

	// Reason in a proprietary format.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *ContractClosureReason1Choice) SetCode(value string) {
	c.Code = (*ExternalContractClosureReason1Code)(&value)
}

func (c *ContractClosureReason1Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
