package model

// Choice of format for the repurchase transaction type information.
type RepurchaseType21Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType8Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RepurchaseType21Choice) SetCode(value string) {
	r.Code = (*RepurchaseType8Code)(&value)
}

func (r *RepurchaseType21Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
