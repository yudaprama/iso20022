package model

// Choice of format for the generated reason.
type GeneratedReasons3Choice struct {

	// Specifies the reason why the transaction was generated.
	Code *GeneratedReason3Code `xml:"Cd"`

	// Specifies the reason why the transaction was generated.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (g *GeneratedReasons3Choice) SetCode(value string) {
	g.Code = (*GeneratedReason3Code)(&value)
}

func (g *GeneratedReasons3Choice) AddProprietary() *GenericIdentification20 {
	g.Proprietary = new(GenericIdentification20)
	return g.Proprietary
}
