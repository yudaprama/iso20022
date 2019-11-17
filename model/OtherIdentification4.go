package model

// Identification of a financial instrument using an accepted format other than an ISIN.
type OtherIdentification4 struct {

	// Identification of the fund/sub fund.
	Identification *Max35Text `xml:"Id"`

	// Identification source
	Type *IdentificationSource5Choice `xml:"Tp"`
}

func (o *OtherIdentification4) SetIdentification(value string) {
	o.Identification = (*Max35Text)(&value)
}

func (o *OtherIdentification4) AddType() *IdentificationSource5Choice {
	o.Type = new(IdentificationSource5Choice)
	return o.Type
}
