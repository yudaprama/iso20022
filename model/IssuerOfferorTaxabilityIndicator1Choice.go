package model

// Choice between a standard code or proprietary code to specify the issuer / offeror taxability status.
type IssuerOfferorTaxabilityIndicator1Choice struct {

	// Standard code to specify information regarding the issuer / offeror taxability status.
	Code *IssuerTaxability2Code `xml:"Cd"`

	// Proprietary identification to specify information regarding the issuer / offeror taxability status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *IssuerOfferorTaxabilityIndicator1Choice) SetCode(value string) {
	i.Code = (*IssuerTaxability2Code)(&value)
}

func (i *IssuerOfferorTaxabilityIndicator1Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
