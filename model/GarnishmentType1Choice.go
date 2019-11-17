package model

// Specifies the type of garnishment.
type GarnishmentType1Choice struct {

	// Garnishment type in a coded form.
	// Would suggest this to be an External Code List to contain:
	// GNCS    Garnishment from a third party payer for Child Support
	// GNDP    Garnishment from a Direct Payer for Child Support
	// GTPP     Garnishment from a third party payer to taxing agency
	Code *ExternalGarnishmentType1Code `xml:"Cd"`

	// Proprietary identification of the type of garnishment.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (g *GarnishmentType1Choice) SetCode(value string) {
	g.Code = (*ExternalGarnishmentType1Code)(&value)
}

func (g *GarnishmentType1Choice) SetProprietary(value string) {
	g.Proprietary = (*Max35Text)(&value)
}
