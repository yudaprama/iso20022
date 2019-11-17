package model

// Choice of format for the repurchase transaction type information.
type RepurchaseType13Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType6Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RepurchaseType13Choice) SetCode(value string) {
	r.Code = (*RepurchaseType6Code)(&value)
}

func (r *RepurchaseType13Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
