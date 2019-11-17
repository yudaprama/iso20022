package model

// Choice of format for the repurchase transaction type information.
type RepurchaseType3Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType3Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RepurchaseType3Choice) SetCode(value string) {
	r.Code = (*RepurchaseType3Code)(&value)
}

func (r *RepurchaseType3Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
