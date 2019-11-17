package model

// Specifies the matching processing change requested.
type UnilateralSplit4Choice struct {

	// Unilateral split expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType5Code `xml:"Cd"`

	// Unilateral split expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (u *UnilateralSplit4Choice) SetCode(value string) {
	u.Code = (*SecuritiesTransactionType5Code)(&value)
}

func (u *UnilateralSplit4Choice) AddProprietary() *GenericIdentification47 {
	u.Proprietary = new(GenericIdentification47)
	return u.Proprietary
}
