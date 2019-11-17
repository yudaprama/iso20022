package model

// Choice of format for the repurchase transaction type information.
type RepurchaseType15Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType5Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RepurchaseType15Choice) SetCode(value string) {
	r.Code = (*RepurchaseType5Code)(&value)
}

func (r *RepurchaseType15Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
