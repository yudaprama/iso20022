package model

// Specifies the balance details.
type ContractBalanceType1Choice struct {

	// Specifies the nature of a balance, in a coded form.
	Code *ExternalContractBalanceType1Code `xml:"Cd"`

	// Specifies the nature of a balance, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *ContractBalanceType1Choice) SetCode(value string) {
	c.Code = (*ExternalContractBalanceType1Code)(&value)
}

func (c *ContractBalanceType1Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
