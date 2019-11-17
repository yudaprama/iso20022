package model

// Choice of format for the generated reason.
type GeneratedReasons1Choice struct {

	// Specifies the reason why the transaction was generated.
	Code *GeneratedReason2Code `xml:"Cd"`

	// Specifies the reason why the transaction was generated.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (g *GeneratedReasons1Choice) SetCode(value string) {
	g.Code = (*GeneratedReason2Code)(&value)
}

func (g *GeneratedReasons1Choice) AddProprietary() *GenericIdentification20 {
	g.Proprietary = new(GenericIdentification20)
	return g.Proprietary
}
