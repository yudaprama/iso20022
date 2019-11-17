package model

// Choice of format for the repurchase transaction type information.
type RepurchaseType11Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType7Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (r *RepurchaseType11Choice) SetCode(value string) {
	r.Code = (*RepurchaseType7Code)(&value)
}

func (r *RepurchaseType11Choice) AddProprietary() *GenericIdentification38 {
	r.Proprietary = new(GenericIdentification38)
	return r.Proprietary
}
