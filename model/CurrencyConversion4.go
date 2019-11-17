package model

// Currency conversion accepted by the customer, either to convert the amount to dispense in the base currency of the ATM, or to convert the total requested amount in the currency of the customer (so called dynamic currency conversion).
type CurrencyConversion4 struct {

	// Identification of the currency conversion operation.
	CurrencyConversionIdentification *Max35Text `xml:"CcyConvsId,omitempty"`

	// Currency into which the amount is converted (ISO 4217, 3 alphanumeric characters).
	TargetCurrency *CurrencyDetails2 `xml:"TrgtCcy"`

	// Amount converted in the target currency, including commission and mark-up.
	ResultingAmount *ImpliedCurrencyAndAmount `xml:"RsltgAmt"`

	// Exchange rate, expressed as a percentage, applied to convert the original amount into the resulting amount.
	ExchangeRate *PercentageRate `xml:"XchgRate"`

	// Exchange rate expressed as a decimal, for example 0.7 is 7/10 and 70%.
	ExchangeRateDecimal *BaseOneRate `xml:"XchgRateDcml,omitempty"`

	// Exchange rate, expressed as a percentage, applied to convert the resulting amount into the original amount.
	InvertedExchangeRate *PercentageRate `xml:"NvrtdXchgRate,omitempty"`

	// Date and time at which the exchange rate has been quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Validity limit of the exchange rate.
	ValidUntil *ISODateTime `xml:"VldUntil,omitempty"`

	// Currency from which the amount is converted (ISO 4217, 3 alphanumeric characters).
	SourceCurrency *CurrencyDetails2 `xml:"SrcCcy"`

	// Original amount in the source currency.
	OriginalAmount *ImpliedCurrencyAndAmount `xml:"OrgnlAmt"`

	// Commission or additional charges made as part of a currency conversion.
	CommissionDetails []*Commission19 `xml:"ComssnDtls,omitempty"`

	// Mark-up made as part of a currency conversion.
	MarkUpDetails []*Commission18 `xml:"MrkUpDtls,omitempty"`

	// Card scheme declaration (disclaimer) to present to the cardholder.
	DeclarationDetails *Max2048Text `xml:"DclrtnDtls,omitempty"`
}

func (c *CurrencyConversion4) SetCurrencyConversionIdentification(value string) {
	c.CurrencyConversionIdentification = (*Max35Text)(&value)
}

func (c *CurrencyConversion4) AddTargetCurrency() *CurrencyDetails2 {
	c.TargetCurrency = new(CurrencyDetails2)
	return c.TargetCurrency
}

func (c *CurrencyConversion4) SetResultingAmount(value, currency string) {
	c.ResultingAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion4) SetExchangeRate(value string) {
	c.ExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion4) SetExchangeRateDecimal(value string) {
	c.ExchangeRateDecimal = (*BaseOneRate)(&value)
}

func (c *CurrencyConversion4) SetInvertedExchangeRate(value string) {
	c.InvertedExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion4) SetQuotationDate(value string) {
	c.QuotationDate = (*ISODateTime)(&value)
}

func (c *CurrencyConversion4) SetValidUntil(value string) {
	c.ValidUntil = (*ISODateTime)(&value)
}

func (c *CurrencyConversion4) AddSourceCurrency() *CurrencyDetails2 {
	c.SourceCurrency = new(CurrencyDetails2)
	return c.SourceCurrency
}

func (c *CurrencyConversion4) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion4) AddCommissionDetails() *Commission19 {
	newValue := new(Commission19)
	c.CommissionDetails = append(c.CommissionDetails, newValue)
	return newValue
}

func (c *CurrencyConversion4) AddMarkUpDetails() *Commission18 {
	newValue := new(Commission18)
	c.MarkUpDetails = append(c.MarkUpDetails, newValue)
	return newValue
}

func (c *CurrencyConversion4) SetDeclarationDetails(value string) {
	c.DeclarationDetails = (*Max2048Text)(&value)
}
