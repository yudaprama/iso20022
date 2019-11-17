package model

// Choice of format for the repurchase transaction type information.
type RepurchaseType22Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType9Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RepurchaseType22Choice) SetCode(value string) {
	r.Code = (*RepurchaseType9Code)(&value)
}

func (r *RepurchaseType22Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
