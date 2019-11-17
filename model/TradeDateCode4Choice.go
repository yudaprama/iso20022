package model

// Choice of format for the trade date code.
type TradeDateCode4Choice struct {

	// Trade date expressed as an ISO 20022 code.
	Code *DateType3Code `xml:"Cd"`

	// Trade date expressed as an proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TradeDateCode4Choice) SetCode(value string) {
	t.Code = (*DateType3Code)(&value)
}

func (t *TradeDateCode4Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
