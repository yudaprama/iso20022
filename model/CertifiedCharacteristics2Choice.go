package model

// Characteristics of the goods that are certified, in the context of a commercial trade transaction.
type CertifiedCharacteristics2Choice struct {

	// Country of origin of the goods, as proven by the certificate.
	Origin *CountryCode `xml:"Orgn"`

	// Quality of the goods, as proven by the certificate.
	Quality *Max70Text `xml:"Qlty"`

	// Analysis of the goods, as proven by the certificate.
	Analysis *Max70Text `xml:"Anlys"`

	// Weight of the goods, as proven by the certificate.
	Weight *Quantity9 `xml:"Wght"`

	// Quantity of the goods, as proven by the certificate.
	Quantity *Quantity9 `xml:"Qty"`

	// Indicates if the goods have passed the health check.
	HealthIndication *YesNoIndicator `xml:"HlthIndctn"`

	// Indicates if the goods have passed the phytosanitary check.
	PhytosanitaryIndication *YesNoIndicator `xml:"PhytosntryIndctn"`
}

func (c *CertifiedCharacteristics2Choice) SetOrigin(value string) {
	c.Origin = (*CountryCode)(&value)
}

func (c *CertifiedCharacteristics2Choice) SetQuality(value string) {
	c.Quality = (*Max70Text)(&value)
}

func (c *CertifiedCharacteristics2Choice) SetAnalysis(value string) {
	c.Analysis = (*Max70Text)(&value)
}

func (c *CertifiedCharacteristics2Choice) AddWeight() *Quantity9 {
	c.Weight = new(Quantity9)
	return c.Weight
}

func (c *CertifiedCharacteristics2Choice) AddQuantity() *Quantity9 {
	c.Quantity = new(Quantity9)
	return c.Quantity
}

func (c *CertifiedCharacteristics2Choice) SetHealthIndication(value string) {
	c.HealthIndication = (*YesNoIndicator)(&value)
}

func (c *CertifiedCharacteristics2Choice) SetPhytosanitaryIndication(value string) {
	c.PhytosanitaryIndication = (*YesNoIndicator)(&value)
}
