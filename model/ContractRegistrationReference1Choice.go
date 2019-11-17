package model

// Reference details of the contract registration.
type ContractRegistrationReference1Choice struct {

	// Unique and unambiguous identification of the registered contract.
	RegisteredContractIdentification *Max35Text `xml:"RegdCtrctId"`

	// Identification of the contract.
	Contract *DocumentIdentification28 `xml:"Ctrct"`
}

func (c *ContractRegistrationReference1Choice) SetRegisteredContractIdentification(value string) {
	c.RegisteredContractIdentification = (*Max35Text)(&value)
}

func (c *ContractRegistrationReference1Choice) AddContract() *DocumentIdentification28 {
	c.Contract = new(DocumentIdentification28)
	return c.Contract
}
