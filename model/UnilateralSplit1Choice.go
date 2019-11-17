package model

// Specifies the matching processing change requested.
type UnilateralSplit1Choice struct {

	// Unilateral split expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType5Code `xml:"Cd"`

	// Unilateral split expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (u *UnilateralSplit1Choice) SetCode(value string) {
	u.Code = (*SecuritiesTransactionType5Code)(&value)
}

func (u *UnilateralSplit1Choice) AddProprietary() *GenericIdentification20 {
	u.Proprietary = new(GenericIdentification20)
	return u.Proprietary
}
