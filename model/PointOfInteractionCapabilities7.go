package model

// Capabilities of the ATM terminal.
type PointOfInteractionCapabilities7 struct {

	// Card reading capabilities of the ATM performing the transaction.
	CardReadData []*CardDataReading4Code `xml:"CardRdData,omitempty"`

	// Card writing capabilities of the terminal performing the transaction.
	CardWriteData []*CardDataReading4Code `xml:"CardWrtData,omitempty"`

	// Customer and card authentication capabilities available at the ATM.
	Authentication []*CardholderVerificationCapability3Code `xml:"Authntcn,omitempty"`

	// Maximum number of digits the ATM is able to accept when the cardholder enters its PIN.
	PINLengthCapabilities *Number `xml:"PINLngthCpblties,omitempty"`

	// Maximum number of characters of the approval code the ATM is able to manage.
	ApprovalCodeLength *Number `xml:"ApprvlCdLngth,omitempty"`

	// Maximum data length in bytes that a card issuer can return to the ICC at the terminal.
	MaxScriptLength *Number `xml:"MxScrptLngth,omitempty"`

	// True if the ATM is able to capture card.
	CardCaptureCapable *TrueFalseIndicator `xml:"CardCaptrCpbl,omitempty"`

	// Type of media the ATM is able to dispense.
	WithdrawalMedia []*ATMMediaType1Code `xml:"WdrwlMdia,omitempty"`

	// Type of media the customer is able to deposit in the ATM.
	DepositedMedia []*ATMMediaType2Code `xml:"DpstdMdia,omitempty"`

	// Capabilities of the terminal to display or print message to the cardholder and the merchant.
	MessageCapabilities []*DisplayCapabilities5 `xml:"MsgCpblties,omitempty"`
}

func (p *PointOfInteractionCapabilities7) AddCardReadData(value string) {
	p.CardReadData = append(p.CardReadData, (*CardDataReading4Code)(&value))
}

func (p *PointOfInteractionCapabilities7) AddCardWriteData(value string) {
	p.CardWriteData = append(p.CardWriteData, (*CardDataReading4Code)(&value))
}

func (p *PointOfInteractionCapabilities7) AddAuthentication(value string) {
	p.Authentication = append(p.Authentication, (*CardholderVerificationCapability3Code)(&value))
}

func (p *PointOfInteractionCapabilities7) SetPINLengthCapabilities(value string) {
	p.PINLengthCapabilities = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities7) SetApprovalCodeLength(value string) {
	p.ApprovalCodeLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities7) SetMaxScriptLength(value string) {
	p.MaxScriptLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities7) SetCardCaptureCapable(value string) {
	p.CardCaptureCapable = (*TrueFalseIndicator)(&value)
}

func (p *PointOfInteractionCapabilities7) AddWithdrawalMedia(value string) {
	p.WithdrawalMedia = append(p.WithdrawalMedia, (*ATMMediaType1Code)(&value))
}

func (p *PointOfInteractionCapabilities7) AddDepositedMedia(value string) {
	p.DepositedMedia = append(p.DepositedMedia, (*ATMMediaType2Code)(&value))
}

func (p *PointOfInteractionCapabilities7) AddMessageCapabilities() *DisplayCapabilities5 {
	newValue := new(DisplayCapabilities5)
	p.MessageCapabilities = append(p.MessageCapabilities, newValue)
	return newValue
}
