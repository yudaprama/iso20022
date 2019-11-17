package model

// Choice of format for the opening closing information.
type OpeningClosing4Choice struct {

	// Opening closing information expressed as an ISO 20022 code.
	Code *OpeningClosing1Code `xml:"Cd"`

	// Opening closing information expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OpeningClosing4Choice) SetCode(value string) {
	o.Code = (*OpeningClosing1Code)(&value)
}

func (o *OpeningClosing4Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
