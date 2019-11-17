package model

// Choice between a code or a data source scheme to determine the closing of the securities lending contract.
type Reversible1Choice struct {

	// Closing of the securities lending contract is identified using a code.
	Code *Reversible1Code `xml:"Cd"`

	// Closing of the securities lending contract expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (r *Reversible1Choice) SetCode(value string) {
	r.Code = (*Reversible1Code)(&value)
}

func (r *Reversible1Choice) AddProprietary() *GenericIdentification38 {
	r.Proprietary = new(GenericIdentification38)
	return r.Proprietary
}
