package model

// Payment token information.
type CardPaymentToken1 struct {

	// Additional token payment information.
	TokenCharacteristic []*Max35Text `xml:"TknChrtc,omitempty"`

	// Identifier of a token provider requestor.
	TokenRequestor *PaymentTokenIdentifiers1 `xml:"TknRqstr,omitempty"`
}

func (c *CardPaymentToken1) AddTokenCharacteristic(value string) {
	c.TokenCharacteristic = append(c.TokenCharacteristic, (*Max35Text)(&value))
}

func (c *CardPaymentToken1) AddTokenRequestor() *PaymentTokenIdentifiers1 {
	c.TokenRequestor = new(PaymentTokenIdentifiers1)
	return c.TokenRequestor
}
