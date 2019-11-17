package model

// Choice of format for the generated reason.
type GeneratedReasons6Choice struct {

	// Specifies the reason why the transaction was generated.
	Code *GeneratedReason3Code `xml:"Cd"`

	// Specifies the reason why the transaction was generated.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (g *GeneratedReasons6Choice) SetCode(value string) {
	g.Code = (*GeneratedReason3Code)(&value)
}

func (g *GeneratedReasons6Choice) AddProprietary() *GenericIdentification47 {
	g.Proprietary = new(GenericIdentification47)
	return g.Proprietary
}
