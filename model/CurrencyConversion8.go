package model

// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider.
type CurrencyConversion8 struct {

	// True if the cardholder has accepted the currency conversion that the acquirer has proposed.
	AcceptedByCardholder *TrueFalseIndicator `xml:"AccptdByCrdhldr,omitempty"`

	// Conversion between the currency of a card acceptor and the currency of a cardholder, provided by a dedicated service provider.
	Conversion *CurrencyConversion6 `xml:"Convs,omitempty"`
}

func (c *CurrencyConversion8) SetAcceptedByCardholder(value string) {
	c.AcceptedByCardholder = (*TrueFalseIndicator)(&value)
}

func (c *CurrencyConversion8) AddConversion() *CurrencyConversion6 {
	c.Conversion = new(CurrencyConversion6)
	return c.Conversion
}
