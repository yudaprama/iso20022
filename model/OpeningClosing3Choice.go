package model

// Choice of format for the opening closing information.
type OpeningClosing3Choice struct {

	// Opening closing information expressed as an ISO 20022 code.
	Code *OpeningClosing1Code `xml:"Cd"`

	// Opening closing information expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (o *OpeningClosing3Choice) SetCode(value string) {
	o.Code = (*OpeningClosing1Code)(&value)
}

func (o *OpeningClosing3Choice) AddProprietary() *GenericIdentification30 {
	o.Proprietary = new(GenericIdentification30)
	return o.Proprietary
}
