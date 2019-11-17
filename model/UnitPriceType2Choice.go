package model

// Choice of formats for the price type.
type UnitPriceType2Choice struct {

	//  Type of price expressed as a code.
	Code *TypeOfPrice10Code `xml:"Cd"`

	//  Type of price expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (u *UnitPriceType2Choice) SetCode(value string) {
	u.Code = (*TypeOfPrice10Code)(&value)
}

func (u *UnitPriceType2Choice) AddProprietary() *GenericIdentification47 {
	u.Proprietary = new(GenericIdentification47)
	return u.Proprietary
}
