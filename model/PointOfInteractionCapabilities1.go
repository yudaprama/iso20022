package model

// Capabilities of the POI performing the transaction.
type PointOfInteractionCapabilities1 struct {

	// Card reading capabilities of the POI performing the transaction.
	CardReadingCapabilities []*CardDataReading1Code `xml:"CardRdngCpblties,omitempty"`

	// Cardholder verification capabilities of the POI performing the transaction.
	CardholderVerificationCapabilities []*CardholderVerificationCapability1Code `xml:"CrdhldrVrfctnCpblties,omitempty"`

	// On-line and off-line capabilities of the POI.
	OnLineCapabilities *OnLineCapability1Code `xml:"OnLineCpblties,omitempty"`

	// Capabilities of the display components performing the transaction.
	DisplayCapabilities []*DisplayCapabilities1 `xml:"DispCpblties,omitempty"`

	// Number of columns of the printer component.
	PrintLineWidth *Max3NumericText `xml:"PrtLineWidth,omitempty"`
}

func (p *PointOfInteractionCapabilities1) AddCardReadingCapabilities(value string) {
	p.CardReadingCapabilities = append(p.CardReadingCapabilities, (*CardDataReading1Code)(&value))
}

func (p *PointOfInteractionCapabilities1) AddCardholderVerificationCapabilities(value string) {
	p.CardholderVerificationCapabilities = append(p.CardholderVerificationCapabilities, (*CardholderVerificationCapability1Code)(&value))
}

func (p *PointOfInteractionCapabilities1) SetOnLineCapabilities(value string) {
	p.OnLineCapabilities = (*OnLineCapability1Code)(&value)
}

func (p *PointOfInteractionCapabilities1) AddDisplayCapabilities() *DisplayCapabilities1 {
	newValue := new(DisplayCapabilities1)
	p.DisplayCapabilities = append(p.DisplayCapabilities, newValue)
	return newValue
}

func (p *PointOfInteractionCapabilities1) SetPrintLineWidth(value string) {
	p.PrintLineWidth = (*Max3NumericText)(&value)
}
