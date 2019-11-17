package model

// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider.
type CurrencyConversion7 struct {

	// Result of a requested currency conversion.
	Result *CurrencyConversionResponse1Code `xml:"Rslt"`

	// Plain text explaining the result of the  currency conversion request.
	ResultReason *Max35Text `xml:"RsltRsn,omitempty"`

	// Information about the conversion of currency.
	ConversionDetails *CurrencyConversion6 `xml:"ConvsDtls,omitempty"`
}

func (c *CurrencyConversion7) SetResult(value string) {
	c.Result = (*CurrencyConversionResponse1Code)(&value)
}

func (c *CurrencyConversion7) SetResultReason(value string) {
	c.ResultReason = (*Max35Text)(&value)
}

func (c *CurrencyConversion7) AddConversionDetails() *CurrencyConversion6 {
	c.ConversionDetails = new(CurrencyConversion6)
	return c.ConversionDetails
}
