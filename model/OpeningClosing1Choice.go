package model

// Choice of format for the opening closing information.
type OpeningClosing1Choice struct {

	// Opening closing information expressed as an ISO 20022 code.
	Code *OpeningClosing1Code `xml:"Cd"`

	// Opening closing information expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (o *OpeningClosing1Choice) SetCode(value string) {
	o.Code = (*OpeningClosing1Code)(&value)
}

func (o *OpeningClosing1Choice) AddProprietary() *GenericIdentification20 {
	o.Proprietary = new(GenericIdentification20)
	return o.Proprietary
}
