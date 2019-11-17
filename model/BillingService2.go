package model

// Specifies the values used for every line item service in the statement.
type BillingService2 struct {

	// Specifies further detailed values for this service.
	ServiceDetail *BillingServiceParameters3 `xml:"SvcDtl"`

	// Price applied to the service, expressed in the pricing currency.
	Price *BillingPrice1 `xml:"Pric,omitempty"`

	// Code identifying the disposition of the calculated charge.
	PaymentMethod *ServicePaymentMethod1Code `xml:"PmtMtd"`

	// Amount of the calculated charge, expressed in the pricing currency. This value does not include any tax on the service.
	OriginalChargePrice *AmountAndDirection34 `xml:"OrgnlChrgPric"`

	// Amount of the calculated charge, expressed in the settlement currency.
	OriginalChargeSettlementAmount *AmountAndDirection34 `xml:"OrgnlChrgSttlmAmt,omitempty"`

	// Average daily collected balance required to offset a balance compensable service charge, exclusive of taxes, expressed in the account currency.
	BalanceRequiredAccountAmount *AmountAndDirection34 `xml:"BalReqrdAcctAmt,omitempty"`

	// Provides the details of the taxable status of a service.
	TaxDesignation *ServiceTaxDesignation1 `xml:"TaxDsgnt"`

	// Provides tax related values for  tax calculation methods A, B or D.
	TaxCalculation *BillingMethod1Choice `xml:"TaxClctn,omitempty"`
}

func (b *BillingService2) AddServiceDetail() *BillingServiceParameters3 {
	b.ServiceDetail = new(BillingServiceParameters3)
	return b.ServiceDetail
}

func (b *BillingService2) AddPrice() *BillingPrice1 {
	b.Price = new(BillingPrice1)
	return b.Price
}

func (b *BillingService2) SetPaymentMethod(value string) {
	b.PaymentMethod = (*ServicePaymentMethod1Code)(&value)
}

func (b *BillingService2) AddOriginalChargePrice() *AmountAndDirection34 {
	b.OriginalChargePrice = new(AmountAndDirection34)
	return b.OriginalChargePrice
}

func (b *BillingService2) AddOriginalChargeSettlementAmount() *AmountAndDirection34 {
	b.OriginalChargeSettlementAmount = new(AmountAndDirection34)
	return b.OriginalChargeSettlementAmount
}

func (b *BillingService2) AddBalanceRequiredAccountAmount() *AmountAndDirection34 {
	b.BalanceRequiredAccountAmount = new(AmountAndDirection34)
	return b.BalanceRequiredAccountAmount
}

func (b *BillingService2) AddTaxDesignation() *ServiceTaxDesignation1 {
	b.TaxDesignation = new(ServiceTaxDesignation1)
	return b.TaxDesignation
}

func (b *BillingService2) AddTaxCalculation() *BillingMethod1Choice {
	b.TaxCalculation = new(BillingMethod1Choice)
	return b.TaxCalculation
}
