package model

// Choice of format for the repurchase transaction type information.
type RepurchaseType17Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType6Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RepurchaseType17Choice) SetCode(value string) {
	r.Code = (*RepurchaseType6Code)(&value)
}

func (r *RepurchaseType17Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
