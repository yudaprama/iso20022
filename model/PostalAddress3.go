package model

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress3 struct {

	// Type of address.
	AddressType *AddressType1Code `xml:"AdrTp"`

	// Indicates whether mail should be sent to an address.
	MailingIndicator *YesNoIndicator `xml:"MlngInd"`

	// Indicates whether the address is the official address of the party.
	RegistrationAddressIndicator *YesNoIndicator `xml:"RegnAdrInd"`

	// Information that locates and identifies a specific address, as defined by postal services.
	NameAndAddress *NameAndAddress4 `xml:"NmAndAdr"`
}

func (p *PostalAddress3) SetAddressType(value string) {
	p.AddressType = (*AddressType1Code)(&value)
}

func (p *PostalAddress3) SetMailingIndicator(value string) {
	p.MailingIndicator = (*YesNoIndicator)(&value)
}

func (p *PostalAddress3) SetRegistrationAddressIndicator(value string) {
	p.RegistrationAddressIndicator = (*YesNoIndicator)(&value)
}

func (p *PostalAddress3) AddNameAndAddress() *NameAndAddress4 {
	p.NameAndAddress = new(NameAndAddress4)
	return p.NameAndAddress
}
