package model

// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider. The currency conversion has to be accepted by the cardholder.
type CurrencyConversion1 struct {

	// Identification of the currency conversion operation for the service provider.
	CurrencyConversionIdentification *Max35Text `xml:"CcyConvsId,omitempty"`

	// Result of a requested currency conversion.
	Result *CurrencyConversionResponse1Code `xml:"Rslt"`

	// Plain text explaining the result of the  currency conversion request.
	ResponseReason *Max35Text `xml:"RspnRsn,omitempty"`

	// Currency into which the amount is converted (ISO 4217, 3 alphanumeric characters).
	TargetCurrency *CurrencyCode `xml:"TrgtCcy"`

	// Currency into which the amount is converted (ISO 4217, 3 numeric characters).
	TargetCurrencyNumeric *Exact3NumericText `xml:"TrgtCcyNmrc"`

	// Maximal number of digits after the decimal separator for target currency.
	TargetCurrencyDecimal *Number `xml:"TrgtCcyDcml"`

	// Full name of the target currency.
	TargetCurrencyName *Max35Text `xml:"TrgtCcyNm,omitempty"`

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
	SourceCurrency *CurrencyCode `xml:"SrcCcy"`

	// Currency from which the amount is converted (ISO 4217, 3 numeric characters).
	SourceCurrencyNumeric *CurrencyCode `xml:"SrcCcyNmrc,omitempty"`

	// Maximal number of digits after the decimal separator for source currency.
	SourceCurrencyDecimal *Number `xml:"SrcCcyDcml"`

	// Full name of the source currency.
	SourceCurrencyName *Max35Text `xml:"SrcCcyNm,omitempty"`

	// Original amount in the source currency.
	OriginalAmount *ImpliedCurrencyAndAmount `xml:"OrgnlAmt"`

	// Commission or additional charges made as part of a currency conversion.
	CommissionDetails []*Commission19 `xml:"ComssnDtls,omitempty"`

	// Markup made as part of a currency conversion.
	MarkUpDetails []*Commission18 `xml:"MrkUpDtls,omitempty"`

	// Card scheme declaration (disclaimer) to present to the cardholder.
	DeclarationDetails *Max2048Text `xml:"DclrtnDtls,omitempty"`
}

func (c *CurrencyConversion1) SetCurrencyConversionIdentification(value string) {
	c.CurrencyConversionIdentification = (*Max35Text)(&value)
}

func (c *CurrencyConversion1) SetResult(value string) {
	c.Result = (*CurrencyConversionResponse1Code)(&value)
}

func (c *CurrencyConversion1) SetResponseReason(value string) {
	c.ResponseReason = (*Max35Text)(&value)
}

func (c *CurrencyConversion1) SetTargetCurrency(value string) {
	c.TargetCurrency = (*CurrencyCode)(&value)
}

func (c *CurrencyConversion1) SetTargetCurrencyNumeric(value string) {
	c.TargetCurrencyNumeric = (*Exact3NumericText)(&value)
}

func (c *CurrencyConversion1) SetTargetCurrencyDecimal(value string) {
	c.TargetCurrencyDecimal = (*Number)(&value)
}

func (c *CurrencyConversion1) SetTargetCurrencyName(value string) {
	c.TargetCurrencyName = (*Max35Text)(&value)
}

func (c *CurrencyConversion1) SetResultingAmount(value, currency string) {
	c.ResultingAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion1) SetExchangeRate(value string) {
	c.ExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion1) SetInvertedExchangeRate(value string) {
	c.InvertedExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion1) SetQuotationDate(value string) {
	c.QuotationDate = (*ISODateTime)(&value)
}

func (c *CurrencyConversion1) SetValidUntil(value string) {
	c.ValidUntil = (*ISODateTime)(&value)
}

func (c *CurrencyConversion1) SetSourceCurrency(value string) {
	c.SourceCurrency = (*CurrencyCode)(&value)
}

func (c *CurrencyConversion1) SetSourceCurrencyNumeric(value string) {
	c.SourceCurrencyNumeric = (*CurrencyCode)(&value)
}

func (c *CurrencyConversion1) SetSourceCurrencyDecimal(value string) {
	c.SourceCurrencyDecimal = (*Number)(&value)
}

func (c *CurrencyConversion1) SetSourceCurrencyName(value string) {
	c.SourceCurrencyName = (*Max35Text)(&value)
}

func (c *CurrencyConversion1) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion1) AddCommissionDetails() *Commission19 {
	newValue := new(Commission19)
	c.CommissionDetails = append(c.CommissionDetails, newValue)
	return newValue
}

func (c *CurrencyConversion1) AddMarkUpDetails() *Commission18 {
	newValue := new(Commission18)
	c.MarkUpDetails = append(c.MarkUpDetails, newValue)
	return newValue
}

func (c *CurrencyConversion1) SetDeclarationDetails(value string) {
	c.DeclarationDetails = (*Max2048Text)(&value)
}
