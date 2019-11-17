package model

// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider.
type CurrencyConversion3 struct {

	// Result of a requested currency conversion.
	Result *CurrencyConversionResponse1Code `xml:"Rslt"`

	// Plain text explaining the result of the  currency conversion request.
	ResultReason *Max35Text `xml:"RsltRsn,omitempty"`

	// Information about the conversion of currency.
	Conversion *CurrencyConversion2 `xml:"Convs,omitempty"`
}

func (c *CurrencyConversion3) SetResult(value string) {
	c.Result = (*CurrencyConversionResponse1Code)(&value)
}

func (c *CurrencyConversion3) SetResultReason(value string) {
	c.ResultReason = (*Max35Text)(&value)
}

func (c *CurrencyConversion3) AddConversion() *CurrencyConversion2 {
	c.Conversion = new(CurrencyConversion2)
	return c.Conversion
}
