package model

// Specifies the matching processing change requested.
type UnilateralSplit3Choice struct {

	// Unilateral split expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType5Code `xml:"Cd"`

	// Unilateral split expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (u *UnilateralSplit3Choice) SetCode(value string) {
	u.Code = (*SecuritiesTransactionType5Code)(&value)
}

func (u *UnilateralSplit3Choice) AddProprietary() *GenericIdentification30 {
	u.Proprietary = new(GenericIdentification30)
	return u.Proprietary
}
