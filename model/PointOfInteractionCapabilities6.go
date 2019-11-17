package model

// Capabilities of the POI (Point Of Interaction) performing the transaction.
type PointOfInteractionCapabilities6 struct {

	// Card reading capabilities of the POI (Point Of Interaction) performing the transaction.
	CardReadingCapabilities []*CardDataReading5Code `xml:"CardRdngCpblties,omitempty"`

	// Cardholder verification capabilities of the POI (Point Of Interaction) performing the transaction.
	CardholderVerificationCapabilities []*CardholderVerificationCapability4Code `xml:"CrdhldrVrfctnCpblties,omitempty"`

	// Maximum number of digits the POI is able to accept when the cardholder enters its PIN.
	PINLengthCapabilities *Number `xml:"PINLngthCpblties,omitempty"`

	// Maximum number of characters of the approval code the POI is able to manage.
	ApprovalCodeLength *Number `xml:"ApprvlCdLngth,omitempty"`

	// Maximum data length in bytes that a card issuer can return to the ICC at the terminal.
	MaxScriptLength *Number `xml:"MxScrptLngth,omitempty"`

	// True if the POI is able to capture card.
	CardCaptureCapable *TrueFalseIndicator `xml:"CardCaptrCpbl,omitempty"`

	// On-line and off-line capabilities of the POI (Point Of Interaction).
	OnLineCapabilities *OnLineCapability1Code `xml:"OnLineCpblties,omitempty"`

	// Capabilities of the terminal to display or print message to the cardholder and the merchant.
	MessageCapabilities []*DisplayCapabilities4 `xml:"MsgCpblties,omitempty"`
}

func (p *PointOfInteractionCapabilities6) AddCardReadingCapabilities(value string) {
	p.CardReadingCapabilities = append(p.CardReadingCapabilities, (*CardDataReading5Code)(&value))
}

func (p *PointOfInteractionCapabilities6) AddCardholderVerificationCapabilities(value string) {
	p.CardholderVerificationCapabilities = append(p.CardholderVerificationCapabilities, (*CardholderVerificationCapability4Code)(&value))
}

func (p *PointOfInteractionCapabilities6) SetPINLengthCapabilities(value string) {
	p.PINLengthCapabilities = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities6) SetApprovalCodeLength(value string) {
	p.ApprovalCodeLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities6) SetMaxScriptLength(value string) {
	p.MaxScriptLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities6) SetCardCaptureCapable(value string) {
	p.CardCaptureCapable = (*TrueFalseIndicator)(&value)
}

func (p *PointOfInteractionCapabilities6) SetOnLineCapabilities(value string) {
	p.OnLineCapabilities = (*OnLineCapability1Code)(&value)
}

func (p *PointOfInteractionCapabilities6) AddMessageCapabilities() *DisplayCapabilities4 {
	newValue := new(DisplayCapabilities4)
	p.MessageCapabilities = append(p.MessageCapabilities, newValue)
	return newValue
}
