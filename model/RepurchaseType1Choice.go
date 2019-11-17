package model

// Choice of format for the repurchase transaction type information.
type RepurchaseType1Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType2Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RepurchaseType1Choice) SetCode(value string) {
	r.Code = (*RepurchaseType2Code)(&value)
}

func (r *RepurchaseType1Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
