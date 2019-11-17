package model

// Choice of format for the repurchase transaction type information.
type RepurchaseType7Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType5Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RepurchaseType7Choice) SetCode(value string) {
	r.Code = (*RepurchaseType5Code)(&value)
}

func (r *RepurchaseType7Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
