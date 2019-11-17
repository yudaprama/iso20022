package model

// Payment token information.
type CardPaymentToken3 struct {

	// Additional payment token information.
	TokenCharacteristic []*Max35Text `xml:"TknChrtc,omitempty"`

	// Identifier of a token provider requestor.
	TokenRequestor *PaymentTokenIdentifiers1 `xml:"TknRqstr,omitempty"`

	// Level of confidence resulting of the identification and authentication of the cardholder performed and the entity that performed it.
	TokenAssuranceLevel *Number `xml:"TknAssrncLvl,omitempty"`
}

func (c *CardPaymentToken3) AddTokenCharacteristic(value string) {
	c.TokenCharacteristic = append(c.TokenCharacteristic, (*Max35Text)(&value))
}

func (c *CardPaymentToken3) AddTokenRequestor() *PaymentTokenIdentifiers1 {
	c.TokenRequestor = new(PaymentTokenIdentifiers1)
	return c.TokenRequestor
}

func (c *CardPaymentToken3) SetTokenAssuranceLevel(value string) {
	c.TokenAssuranceLevel = (*Number)(&value)
}
