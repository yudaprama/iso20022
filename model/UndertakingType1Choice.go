package model

// Choice of format for the undertaking type.
type UndertakingType1Choice struct {

	// Type of undertaking.
	//
	Code *ExternalUndertakingType1Code `xml:"Cd"`

	// Type of undertaking expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (u *UndertakingType1Choice) SetCode(value string) {
	u.Code = (*ExternalUndertakingType1Code)(&value)
}

func (u *UndertakingType1Choice) AddProprietary() *GenericIdentification1 {
	u.Proprietary = new(GenericIdentification1)
	return u.Proprietary
}
