package model

// Capabilities of the POI (Point Of Interaction) performing the transaction.
type PointOfInteractionCapabilities3 struct {

	// Card reading capabilities of the POI (Point Of Interaction) performing the transaction.
	CardReadingCapabilities []*CardDataReading1Code `xml:"CardRdngCpblties,omitempty"`

	// Cardholder verification capabilities of the POI (Point Of Interaction) performing the transaction.
	CardholderVerificationCapabilities []*CardholderVerificationCapability1Code `xml:"CrdhldrVrfctnCpblties,omitempty"`

	// Maximum number of digits the POI is able to accept when the cardholder enters its PIN.
	PINLengthCapabilities *Number `xml:"PINLngthCpblties,omitempty"`

	// Maximum number of characters of the approval code the POI is able to manage.
	ApprovalCodeLength *Number `xml:"ApprvlCdLngth,omitempty"`

	// True if the POI is able to capture card.
	CardCaptureCapable *TrueFalseIndicator `xml:"CardCaptrCpbl,omitempty"`

	// On-line and off-line capabilities of the POI (Point Of Interaction).
	OnLineCapabilities *OnLineCapability1Code `xml:"OnLineCpblties,omitempty"`

	// Capabilities of the display components performing the transaction.
	DisplayCapabilities []*DisplayCapabilities2 `xml:"DispCpblties,omitempty"`

	// Number of columns of the printer component.
	PrintLineWidth *Number `xml:"PrtLineWidth,omitempty"`

	// Available language in the display and printer interface.
	// Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AvailableLanguage []*LanguageCode `xml:"AvlblLang,omitempty"`
}

func (p *PointOfInteractionCapabilities3) AddCardReadingCapabilities(value string) {
	p.CardReadingCapabilities = append(p.CardReadingCapabilities, (*CardDataReading1Code)(&value))
}

func (p *PointOfInteractionCapabilities3) AddCardholderVerificationCapabilities(value string) {
	p.CardholderVerificationCapabilities = append(p.CardholderVerificationCapabilities, (*CardholderVerificationCapability1Code)(&value))
}

func (p *PointOfInteractionCapabilities3) SetPINLengthCapabilities(value string) {
	p.PINLengthCapabilities = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities3) SetApprovalCodeLength(value string) {
	p.ApprovalCodeLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities3) SetCardCaptureCapable(value string) {
	p.CardCaptureCapable = (*TrueFalseIndicator)(&value)
}

func (p *PointOfInteractionCapabilities3) SetOnLineCapabilities(value string) {
	p.OnLineCapabilities = (*OnLineCapability1Code)(&value)
}

func (p *PointOfInteractionCapabilities3) AddDisplayCapabilities() *DisplayCapabilities2 {
	newValue := new(DisplayCapabilities2)
	p.DisplayCapabilities = append(p.DisplayCapabilities, newValue)
	return newValue
}

func (p *PointOfInteractionCapabilities3) SetPrintLineWidth(value string) {
	p.PrintLineWidth = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities3) AddAvailableLanguage(value string) {
	p.AvailableLanguage = append(p.AvailableLanguage, (*LanguageCode)(&value))
}
