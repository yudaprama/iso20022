package model

// Choice of formats for the signature type.
type SignatureType1Choice struct {

	// Signature type expressed as a code.
	Code *SignatureType2Code `xml:"Cd"`

	// Signature type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SignatureType1Choice) SetCode(value string) {
	s.Code = (*SignatureType2Code)(&value)
}

func (s *SignatureType1Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
