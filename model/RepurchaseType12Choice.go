package model

// Choice of format for the repurchase transaction type information.
type RepurchaseType12Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType3Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RepurchaseType12Choice) SetCode(value string) {
	r.Code = (*RepurchaseType3Code)(&value)
}

func (r *RepurchaseType12Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
