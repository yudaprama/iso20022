package model

// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider.
type CurrencyConversion10 struct {

	// Result of a requested currency conversion.
	Result *CurrencyConversionResponse2Code `xml:"Rslt"`

	// Plain text explaining the result of the  currency conversion request.
	ResultReason *Max35Text `xml:"RsltRsn,omitempty"`

	// Information about the conversion of currency.
	Conversion *CurrencyConversion9 `xml:"Convs,omitempty"`
}

func (c *CurrencyConversion10) SetResult(value string) {
	c.Result = (*CurrencyConversionResponse2Code)(&value)
}

func (c *CurrencyConversion10) SetResultReason(value string) {
	c.ResultReason = (*Max35Text)(&value)
}

func (c *CurrencyConversion10) AddConversion() *CurrencyConversion9 {
	c.Conversion = new(CurrencyConversion9)
	return c.Conversion
}
