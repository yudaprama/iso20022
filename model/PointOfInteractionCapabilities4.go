package model

// Capabilities of the terminal performing the transaction.
type PointOfInteractionCapabilities4 struct {

	// Card reading capabilities of the terminal performing the transaction.
	// It correspond to the ISO 8583 field number 22-2 for the version 93, and field number 27-1 for the version 2003.
	CardReadingCapabilities []*CardDataReading2Code `xml:"CardRdngCpblties"`

	// Card writting capabilities of the terminal performing the transaction.
	// It correspond to the ISO 8583 field number 22-10 for the version 93, and field number 27-8_9 for the version 2003.
	CardWrittingCapabilities []*CardDataReading3Code `xml:"CardWrttgCpblties,omitempty"`

	// Cardholder verification capabilities by the terminal.
	// It correspond to the ISO 8583 field number 22-2 for the versions 87 and 93, and field number 27-2 for the version 2003.
	CardholderVerificationCapabilities []*CardholderVerificationCapability2Code `xml:"CrdhldrVrfctnCpblties,omitempty"`

	// Maximum number of digits the POI is able to accept when the cardholder enters its PIN.
	// It correspond to the ISO 8583, field number 25 for the version 87, 22-12 for the version 93, and field number 27-11 for the version 2003.
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
	// It correspond to the ISO 8583 field number 22-11 for the version 93, and field number 27-6 for the version 2003.
	MessageCapabilities []*DisplayCapabilities3 `xml:"MsgCpblties,omitempty"`
}

func (p *PointOfInteractionCapabilities4) AddCardReadingCapabilities(value string) {
	p.CardReadingCapabilities = append(p.CardReadingCapabilities, (*CardDataReading2Code)(&value))
}

func (p *PointOfInteractionCapabilities4) AddCardWrittingCapabilities(value string) {
	p.CardWrittingCapabilities = append(p.CardWrittingCapabilities, (*CardDataReading3Code)(&value))
}

func (p *PointOfInteractionCapabilities4) AddCardholderVerificationCapabilities(value string) {
	p.CardholderVerificationCapabilities = append(p.CardholderVerificationCapabilities, (*CardholderVerificationCapability2Code)(&value))
}

func (p *PointOfInteractionCapabilities4) SetPINLengthCapabilities(value string) {
	p.PINLengthCapabilities = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities4) SetApprovalCodeLength(value string) {
	p.ApprovalCodeLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities4) SetMaxScriptLength(value string) {
	p.MaxScriptLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities4) SetCardCaptureCapable(value string) {
	p.CardCaptureCapable = (*TrueFalseIndicator)(&value)
}

func (p *PointOfInteractionCapabilities4) SetOnLineCapabilities(value string) {
	p.OnLineCapabilities = (*OnLineCapability1Code)(&value)
}

func (p *PointOfInteractionCapabilities4) AddMessageCapabilities() *DisplayCapabilities3 {
	newValue := new(DisplayCapabilities3)
	p.MessageCapabilities = append(p.MessageCapabilities, newValue)
	return newValue
}
