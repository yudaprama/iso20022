package model

// Characteristics of the goods that are certified, in the context of a commercial trade transaction.
type CertifiedCharacteristics1Choice struct {

	// Country of origin of the goods, as proven by the certificate.
	Origin *CountryCode `xml:"Orgn"`

	// Quality of the goods, as proven by the certificate.
	Quality *Max70Text `xml:"Qlty"`

	// Analysis of the goods, as proven by the certificate.
	Analysis *Max70Text `xml:"Anlys"`

	// Weight of the goods, as proven by the certificate.
	Weight *Quantity4 `xml:"Wght"`

	// Quantity of the goods, as proven by the certificate.
	Quantity *Quantity4 `xml:"Qty"`

	// Indicates if the goods have passed the health check.
	HealthIndication *YesNoIndicator `xml:"HlthIndctn"`

	// Indicates if the goods have passed the phytosanitary check.
	PhytosanitaryIndication *YesNoIndicator `xml:"PhytosntryIndctn"`
}

func (c *CertifiedCharacteristics1Choice) SetOrigin(value string) {
	c.Origin = (*CountryCode)(&value)
}

func (c *CertifiedCharacteristics1Choice) SetQuality(value string) {
	c.Quality = (*Max70Text)(&value)
}

func (c *CertifiedCharacteristics1Choice) SetAnalysis(value string) {
	c.Analysis = (*Max70Text)(&value)
}

func (c *CertifiedCharacteristics1Choice) AddWeight() *Quantity4 {
	c.Weight = new(Quantity4)
	return c.Weight
}

func (c *CertifiedCharacteristics1Choice) AddQuantity() *Quantity4 {
	c.Quantity = new(Quantity4)
	return c.Quantity
}

func (c *CertifiedCharacteristics1Choice) SetHealthIndication(value string) {
	c.HealthIndication = (*YesNoIndicator)(&value)
}

func (c *CertifiedCharacteristics1Choice) SetPhytosanitaryIndication(value string) {
	c.PhytosanitaryIndication = (*YesNoIndicator)(&value)
}
