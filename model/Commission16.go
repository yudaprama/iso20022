package model

// Amount of money due to a party as compensation for a service.
type Commission16 struct {

	// Specification of the commission type.
	Type *CommissionType2Choice `xml:"Tp"`

	// Amount of money due to a party as compensation for a service.
	Commission *AmountOrRate2Choice `xml:"Comssn"`

	// Information related to an identification, eg, party identification or account identification.
	RecipientIdentification *PartyIdentification54 `xml:"RcptId,omitempty"`

	// Date at which an operation is triggered to calculate, for instance, a commission, fee, asset values, etc.
	CalculationDate *ISODate `xml:"ClctnDt,omitempty"`

	// Total value of the commissions for a specific trade.
	TotalCommission *AmountAndDirection29 `xml:"TtlComssn,omitempty"`

	// Amount that results of the calculation of VAT on net fees, according to the transaction current tariffs.
	TotalVATAmount *ActiveCurrencyAndAmount `xml:"TtlVATAmt,omitempty"`

	// Specifies the VAT rate.
	VATRate *BaseOneRate `xml:"VATRate,omitempty"`
}

func (c *Commission16) AddType() *CommissionType2Choice {
	c.Type = new(CommissionType2Choice)
	return c.Type
}

func (c *Commission16) AddCommission() *AmountOrRate2Choice {
	c.Commission = new(AmountOrRate2Choice)
	return c.Commission
}

func (c *Commission16) AddRecipientIdentification() *PartyIdentification54 {
	c.RecipientIdentification = new(PartyIdentification54)
	return c.RecipientIdentification
}

func (c *Commission16) SetCalculationDate(value string) {
	c.CalculationDate = (*ISODate)(&value)
}

func (c *Commission16) AddTotalCommission() *AmountAndDirection29 {
	c.TotalCommission = new(AmountAndDirection29)
	return c.TotalCommission
}

func (c *Commission16) SetTotalVATAmount(value, currency string) {
	c.TotalVATAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *Commission16) SetVATRate(value string) {
	c.VATRate = (*BaseOneRate)(&value)
}
