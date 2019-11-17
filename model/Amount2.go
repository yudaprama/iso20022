package model

// Provides the amount in the reporting currency and optionally in the original currency.
type Amount2 struct {

	// Amount expressed in the original currency.
	OriginalCurrencyAmount *ActiveCurrencyAndAmount `xml:"OrgnlCcyAmt,omitempty"`

	// Amount expressed in the reporting currency.
	ReportingAmount *ImpliedCurrencyAndAmount `xml:"RptgAmt"`
}

func (a *Amount2) SetOriginalCurrencyAmount(value, currency string) {
	a.OriginalCurrencyAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *Amount2) SetReportingAmount(value, currency string) {
	a.ReportingAmount = NewImpliedCurrencyAndAmount(value, currency)
}
