package model

// Choice of format for the generated reason.
type GeneratedReasons5Choice struct {

	// Specifies the reason why the transaction was generated.
	Code *GeneratedReason3Code `xml:"Cd"`

	// Specifies the reason why the transaction was generated.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (g *GeneratedReasons5Choice) SetCode(value string) {
	g.Code = (*GeneratedReason3Code)(&value)
}

func (g *GeneratedReasons5Choice) AddProprietary() *GenericIdentification30 {
	g.Proprietary = new(GenericIdentification30)
	return g.Proprietary
}
