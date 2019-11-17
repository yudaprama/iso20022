package model

// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider. The currency conversion has to be accepted by the cardholder.
type CurrencyConversion6 struct {

	// Identification of the currency conversion operation for the service provider.
	CurrencyConversionIdentification *Max35Text `xml:"CcyConvsId,omitempty"`

	// Currency into which the amount is converted (ISO 4217, 3 alphanumeric characters).
	TargetCurrency *CurrencyDetails1 `xml:"TrgtCcy"`

	// Amount converted in the target currency, including additional charges.
	ResultingAmount *ImpliedCurrencyAndAmount `xml:"RsltgAmt"`

	// Exchange rate, expressed as a percentage, applied to convert the original amount into the resulting amount.
	ExchangeRate *PercentageRate `xml:"XchgRate"`

	// Exchange rate, expressed as a percentage, applied to convert the resulting amount into the original amount.
	InvertedExchangeRate *PercentageRate `xml:"NvrtdXchgRate,omitempty"`

	// Date and time at which the exchange rate has been quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Validity limit of the exchange rate.
	ValidUntil *ISODateTime `xml:"VldUntil,omitempty"`

	// Currency from which the amount is converted (ISO 4217, 3 alphanumeric characters).
	SourceCurrency *CurrencyDetails1 `xml:"SrcCcy"`

	// Original amount in the source currency.
	OriginalAmount *ImpliedCurrencyAndAmount `xml:"OrgnlAmt"`

	// Commission or additional charges made as part of a currency conversion.
	CommissionDetails []*Commission19 `xml:"ComssnDtls,omitempty"`

	// Markup made as part of a currency conversion.
	MarkUpDetails []*Commission18 `xml:"MrkUpDtls,omitempty"`

	// Card scheme declaration (disclaimer) to present to the cardholder.
	DeclarationDetails *ActionMessage5 `xml:"DclrtnDtls,omitempty"`
}

func (c *CurrencyConversion6) SetCurrencyConversionIdentification(value string) {
	c.CurrencyConversionIdentification = (*Max35Text)(&value)
}

func (c *CurrencyConversion6) AddTargetCurrency() *CurrencyDetails1 {
	c.TargetCurrency = new(CurrencyDetails1)
	return c.TargetCurrency
}

func (c *CurrencyConversion6) SetResultingAmount(value, currency string) {
	c.ResultingAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion6) SetExchangeRate(value string) {
	c.ExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion6) SetInvertedExchangeRate(value string) {
	c.InvertedExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion6) SetQuotationDate(value string) {
	c.QuotationDate = (*ISODateTime)(&value)
}

func (c *CurrencyConversion6) SetValidUntil(value string) {
	c.ValidUntil = (*ISODateTime)(&value)
}

func (c *CurrencyConversion6) AddSourceCurrency() *CurrencyDetails1 {
	c.SourceCurrency = new(CurrencyDetails1)
	return c.SourceCurrency
}

func (c *CurrencyConversion6) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion6) AddCommissionDetails() *Commission19 {
	newValue := new(Commission19)
	c.CommissionDetails = append(c.CommissionDetails, newValue)
	return newValue
}

func (c *CurrencyConversion6) AddMarkUpDetails() *Commission18 {
	newValue := new(Commission18)
	c.MarkUpDetails = append(c.MarkUpDetails, newValue)
	return newValue
}

func (c *CurrencyConversion6) AddDeclarationDetails() *ActionMessage5 {
	c.DeclarationDetails = new(ActionMessage5)
	return c.DeclarationDetails
}
