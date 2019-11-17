package model

// Specifies the payment terms of the underlying transaction.
type EarlyPayment1 struct {

	// Date before which the early payment discount is valid for payment.
	EarlyPaymentDate *ISODate `xml:"EarlyPmtDt"`

	// Discount percent for early payment.
	DiscountPercent *PercentageRate `xml:"DscntPct"`

	// Early payment discount with tax, with currency.
	DiscountAmount *CurrencyAndAmount `xml:"DscntAmt"`

	// In tax specification for early payment discount one defined the applied tax rates for specific early payment. VAT stands for value added tax.
	EarlyPaymentTaxSpecification []*EarlyPaymentsVAT1 `xml:"EarlyPmtTaxSpcfctn,omitempty"`

	// Tax total in early payment, with currency.
	EarlyPaymentTaxTotal *CurrencyAndAmount `xml:"EarlyPmtTaxTtl,omitempty"`

	// Payable amount with discount of early payment, with currency.
	DuePayableAmountWithEarlyPayment *CurrencyAndAmount `xml:"DuePyblAmtWthEarlyPmt,omitempty"`
}

func (e *EarlyPayment1) SetEarlyPaymentDate(value string) {
	e.EarlyPaymentDate = (*ISODate)(&value)
}

func (e *EarlyPayment1) SetDiscountPercent(value string) {
	e.DiscountPercent = (*PercentageRate)(&value)
}

func (e *EarlyPayment1) SetDiscountAmount(value, currency string) {
	e.DiscountAmount = NewCurrencyAndAmount(value, currency)
}

func (e *EarlyPayment1) AddEarlyPaymentTaxSpecification() *EarlyPaymentsVAT1 {
	newValue := new(EarlyPaymentsVAT1)
	e.EarlyPaymentTaxSpecification = append(e.EarlyPaymentTaxSpecification, newValue)
	return newValue
}

func (e *EarlyPayment1) SetEarlyPaymentTaxTotal(value, currency string) {
	e.EarlyPaymentTaxTotal = NewCurrencyAndAmount(value, currency)
}

func (e *EarlyPayment1) SetDuePayableAmountWithEarlyPayment(value, currency string) {
	e.DuePayableAmountWithEarlyPayment = NewCurrencyAndAmount(value, currency)
}
