package model

// Choice of formats for the specification of other identification.
type OtherIdentification1Choice struct {

	// Other identification expressed as a code.
	Code *PersonIdentificationType5Code `xml:"Cd"`

	// Other identification expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OtherIdentification1Choice) SetCode(value string) {
	o.Code = (*PersonIdentificationType5Code)(&value)
}

func (o *OtherIdentification1Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
