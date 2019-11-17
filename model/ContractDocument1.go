package model

// Document that contains the information of the contract agreed between both parties.
type ContractDocument1 struct {

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	Reference *Max35Text `xml:"Ref"`

	// Signoff date of the document.
	SignOffDate *ISODate `xml:"SgnOffDt,omitempty"`

	// Identification of the version of the contract.
	Version *Max6Text `xml:"Vrsn,omitempty"`
}

func (c *ContractDocument1) SetReference(value string) {
	c.Reference = (*Max35Text)(&value)
}

func (c *ContractDocument1) SetSignOffDate(value string) {
	c.SignOffDate = (*ISODate)(&value)
}

func (c *ContractDocument1) SetVersion(value string) {
	c.Version = (*Max6Text)(&value)
}
