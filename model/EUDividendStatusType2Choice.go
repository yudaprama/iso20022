package model

// Choice of formats for the EU dividend status type.
type EUDividendStatusType2Choice struct {

	// EU dividend status expressed as a code.
	Code *EUDividendStatus1Code `xml:"Cd"`

	// EU dividend status expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (e *EUDividendStatusType2Choice) SetCode(value string) {
	e.Code = (*EUDividendStatus1Code)(&value)
}

func (e *EUDividendStatusType2Choice) AddProprietary() *GenericIdentification47 {
	e.Proprietary = new(GenericIdentification47)
	return e.Proprietary
}
