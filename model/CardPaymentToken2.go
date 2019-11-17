package model

// Payment token information.
type CardPaymentToken2 struct {

	// Additional token payment information.
	TokenCharacteristic []*Max35Text `xml:"TknChrtc,omitempty"`

	// Level of confidence resulting of the identification and authentication of the cardholder performed and the entity that performed it.
	TokenAssuranceLevel *Number `xml:"TknAssrncLvl,omitempty"`
}

func (c *CardPaymentToken2) AddTokenCharacteristic(value string) {
	c.TokenCharacteristic = append(c.TokenCharacteristic, (*Max35Text)(&value))
}

func (c *CardPaymentToken2) SetTokenAssuranceLevel(value string) {
	c.TokenAssuranceLevel = (*Number)(&value)
}
