package model

// Trade settlement monetary summation specified for this supply chain trade settlement.
type LineItemMonetarySummation1 struct {

	// Monetary value of the line amount total being reported in this trade settlement monetary summation.
	LineTotalAmount []*CurrencyAndAmount `xml:"LineTtlAmt,omitempty"`

	// Monetary value of the total of all allowance amounts being reported in this line item monetary summation.
	AllowanceTotalAmount []*CurrencyAndAmount `xml:"AllwncTtlAmt,omitempty"`

	// Monetary value of the total of all charge amounts being reported in this line item monetary summation.
	ChargeTotalAmount []*CurrencyAndAmount `xml:"ChrgTtlAmt,omitempty"`

	// Monetary value of the total of all tax amounts being reported in this line item monetary summation.
	TaxTotalAmount []*CurrencyAndAmount `xml:"TaxTtlAmt,omitempty"`

	// Monetary value of the total of all tax basis amounts being reported in this line item monetary summation.
	TaxBasisTotalAmount []*CurrencyAndAmount `xml:"TaxBsisTtlAmt,omitempty"`

	// Monetary value of an amount being reported for information in this line item monetary summation.
	InformationAmount []*CurrencyAndAmount `xml:"InfAmt,omitempty"`
}

func (l *LineItemMonetarySummation1) AddLineTotalAmount(value, currency string) {
	l.LineTotalAmount = append(l.LineTotalAmount, NewCurrencyAndAmount(value, currency))
}

func (l *LineItemMonetarySummation1) AddAllowanceTotalAmount(value, currency string) {
	l.AllowanceTotalAmount = append(l.AllowanceTotalAmount, NewCurrencyAndAmount(value, currency))
}

func (l *LineItemMonetarySummation1) AddChargeTotalAmount(value, currency string) {
	l.ChargeTotalAmount = append(l.ChargeTotalAmount, NewCurrencyAndAmount(value, currency))
}

func (l *LineItemMonetarySummation1) AddTaxTotalAmount(value, currency string) {
	l.TaxTotalAmount = append(l.TaxTotalAmount, NewCurrencyAndAmount(value, currency))
}

func (l *LineItemMonetarySummation1) AddTaxBasisTotalAmount(value, currency string) {
	l.TaxBasisTotalAmount = append(l.TaxBasisTotalAmount, NewCurrencyAndAmount(value, currency))
}

func (l *LineItemMonetarySummation1) AddInformationAmount(value, currency string) {
	l.InformationAmount = append(l.InformationAmount, NewCurrencyAndAmount(value, currency))
}
