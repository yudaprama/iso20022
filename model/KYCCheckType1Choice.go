package model

// Choice of formats for the specification of the Know Your Customer (KYC) check type.
type KYCCheckType1Choice struct {

	// Type of Know Your Customer (KYC) check type expressed as a code.
	Code *KnowYourCustomerCheckType1Code `xml:"Cd"`

	// Type of Know Your Customer (KYC) check type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (k *KYCCheckType1Choice) SetCode(value string) {
	k.Code = (*KnowYourCustomerCheckType1Code)(&value)
}

func (k *KYCCheckType1Choice) AddProprietary() *GenericIdentification47 {
	k.Proprietary = new(GenericIdentification47)
	return k.Proprietary
}
