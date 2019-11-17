package model

// Choice of format for the repurchase transaction type information.
type RepurchaseType26Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType10Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RepurchaseType26Choice) SetCode(value string) {
	r.Code = (*RepurchaseType10Code)(&value)
}

func (r *RepurchaseType26Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
