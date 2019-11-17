package model

// Capabilities of the POI (Point Of Interaction) performing the transaction.
type PointOfInteractionCapabilities2 struct {

	// Card reading capabilities of the POI (Point Of Interaction) performing the transaction.
	CardReadingCapabilities []*CardDataReading1Code `xml:"CardRdngCpblties,omitempty"`

	// Cardholder verification capabilities of the POI (Point Of Interaction) performing the transaction.
	CardholderVerificationCapabilities []*CardholderVerificationCapability1Code `xml:"CrdhldrVrfctnCpblties,omitempty"`

	// On-line and off-line capabilities of the POI (Point Of Interaction).
	OnLineCapabilities *OnLineCapability1Code `xml:"OnLineCpblties,omitempty"`

	// Capabilities of the display components performing the transaction.
	DisplayCapabilities []*DisplayCapabilities2 `xml:"DispCpblties,omitempty"`

	// Number of columns of the printer component.
	PrintLineWidth *Number `xml:"PrtLineWidth,omitempty"`

	// Available language in the display and printer interface.
	AvailableLanguage []*ISO2ALanguageCode `xml:"AvlblLang,omitempty"`
}

func (p *PointOfInteractionCapabilities2) AddCardReadingCapabilities(value string) {
	p.CardReadingCapabilities = append(p.CardReadingCapabilities, (*CardDataReading1Code)(&value))
}

func (p *PointOfInteractionCapabilities2) AddCardholderVerificationCapabilities(value string) {
	p.CardholderVerificationCapabilities = append(p.CardholderVerificationCapabilities, (*CardholderVerificationCapability1Code)(&value))
}

func (p *PointOfInteractionCapabilities2) SetOnLineCapabilities(value string) {
	p.OnLineCapabilities = (*OnLineCapability1Code)(&value)
}

func (p *PointOfInteractionCapabilities2) AddDisplayCapabilities() *DisplayCapabilities2 {
	newValue := new(DisplayCapabilities2)
	p.DisplayCapabilities = append(p.DisplayCapabilities, newValue)
	return newValue
}

func (p *PointOfInteractionCapabilities2) SetPrintLineWidth(value string) {
	p.PrintLineWidth = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities2) AddAvailableLanguage(value string) {
	p.AvailableLanguage = append(p.AvailableLanguage, (*ISO2ALanguageCode)(&value))
}
