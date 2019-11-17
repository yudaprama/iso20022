package model

// Choice of format for the repurchase transaction type information.
type RepurchaseType2Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType4Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RepurchaseType2Choice) SetCode(value string) {
	r.Code = (*RepurchaseType4Code)(&value)
}

func (r *RepurchaseType2Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
