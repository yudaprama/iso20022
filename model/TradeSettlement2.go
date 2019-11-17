package model

// Trade settlement details for this invoice which involves the payment of an outstanding debt, account, or charge
type TradeSettlement2 struct {

	// Payment or creditor reference.
	PaymentReference *CreditorReferenceInformation2 `xml:"PmtRef,omitempty"`

	// Date when invoice should be paid.
	DueDate *ISODate `xml:"DueDt,omitempty"`

	// Payable amount with currency code.
	DuePayableAmount *CurrencyAndAmount `xml:"DuePyblAmt"`

	// If invoice currency is different from local tax reporting currency, then applied exchange rate is given in this message structure.
	InvoiceCurrencyExchange *CurrencyReference3 `xml:"InvcCcyXchg,omitempty"`

	// Date when goods/services are delivered to buyer.
	DeliveryDate *ISODate `xml:"DlvryDt,omitempty"`

	// Period during which delivery executed or agreed invoicing period.
	BillingPeriod *Period2 `xml:"BllgPrd,omitempty"`

	// Tax total amount with currency code.
	TaxTotalAmount *CurrencyAndAmount `xml:"TaxTtlAmt"`

	// Reason for tax exemption expressed as a code,  if invoice or card transaction is out of tax processing.
	ExemptionReasonCode *Max4Text `xml:"XmptnRsnCd,omitempty"`

	// Reason for a tax exemption, if invoice or card transaction is out of tax processing.
	ExemptionReason *Max500Text `xml:"XmptnRsn,omitempty"`

	// Calculated tax subtotal
	SubTotalCalculatedTax []*SettlementSubTotalCalculatedTax2 `xml:"SubTtlClctdTax,omitempty"`

	// Details of each early payment discount.
	EarlyPayments []*EarlyPayment1 `xml:"EarlyPmts,omitempty"`
}

func (t *TradeSettlement2) AddPaymentReference() *CreditorReferenceInformation2 {
	t.PaymentReference = new(CreditorReferenceInformation2)
	return t.PaymentReference
}

func (t *TradeSettlement2) SetDueDate(value string) {
	t.DueDate = (*ISODate)(&value)
}

func (t *TradeSettlement2) SetDuePayableAmount(value, currency string) {
	t.DuePayableAmount = NewCurrencyAndAmount(value, currency)
}

func (t *TradeSettlement2) AddInvoiceCurrencyExchange() *CurrencyReference3 {
	t.InvoiceCurrencyExchange = new(CurrencyReference3)
	return t.InvoiceCurrencyExchange
}

func (t *TradeSettlement2) SetDeliveryDate(value string) {
	t.DeliveryDate = (*ISODate)(&value)
}

func (t *TradeSettlement2) AddBillingPeriod() *Period2 {
	t.BillingPeriod = new(Period2)
	return t.BillingPeriod
}

func (t *TradeSettlement2) SetTaxTotalAmount(value, currency string) {
	t.TaxTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (t *TradeSettlement2) SetExemptionReasonCode(value string) {
	t.ExemptionReasonCode = (*Max4Text)(&value)
}

func (t *TradeSettlement2) SetExemptionReason(value string) {
	t.ExemptionReason = (*Max500Text)(&value)
}

func (t *TradeSettlement2) AddSubTotalCalculatedTax() *SettlementSubTotalCalculatedTax2 {
	newValue := new(SettlementSubTotalCalculatedTax2)
	t.SubTotalCalculatedTax = append(t.SubTotalCalculatedTax, newValue)
	return newValue
}

func (t *TradeSettlement2) AddEarlyPayments() *EarlyPayment1 {
	newValue := new(EarlyPayment1)
	t.EarlyPayments = append(t.EarlyPayments, newValue)
	return newValue
}
