package model

// Choice of format for the repurchase transaction type information.
type RepurchaseType8Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType6Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RepurchaseType8Choice) SetCode(value string) {
	r.Code = (*RepurchaseType6Code)(&value)
}

func (r *RepurchaseType8Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
