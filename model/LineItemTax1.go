package model

// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
type LineItemTax1 struct {

	// Amount of money resulting from the calculation of the tax.
	CalculatedAmount []*CurrencyAndAmount `xml:"ClctdAmt,omitempty"`

	// Type of tax applied.
	TypeCode *TaxTypeFormat1Choice `xml:"TpCd,omitempty"`

	// Date of the tax point  date when this tax, levy or duty becomes applicable.
	TaxPointDate *ISODate `xml:"TaxPtDt,omitempty"`

	// Rate used to calculate the amount of this tax, levy or duty.
	CalculatedRate *PercentageRate `xml:"ClctdRate,omitempty"`

	// Code specifying the category to which this tax, levy or duty applies, such as codes for 'exempt from tax', 'standard rate', "free export item - tax not charged'.
	CategoryCode *Max4Text `xml:"CtgyCd,omitempty"`

	// Category name, expressed as text, of the tax, levy or duty.
	CategoryName []*Max35Text `xml:"CtgyNm,omitempty"`
}

func (l *LineItemTax1) AddCalculatedAmount(value, currency string) {
	l.CalculatedAmount = append(l.CalculatedAmount, NewCurrencyAndAmount(value, currency))
}

func (l *LineItemTax1) AddTypeCode() *TaxTypeFormat1Choice {
	l.TypeCode = new(TaxTypeFormat1Choice)
	return l.TypeCode
}

func (l *LineItemTax1) SetTaxPointDate(value string) {
	l.TaxPointDate = (*ISODate)(&value)
}

func (l *LineItemTax1) SetCalculatedRate(value string) {
	l.CalculatedRate = (*PercentageRate)(&value)
}

func (l *LineItemTax1) SetCategoryCode(value string) {
	l.CategoryCode = (*Max4Text)(&value)
}

func (l *LineItemTax1) AddCategoryName(value string) {
	l.CategoryName = append(l.CategoryName, (*Max35Text)(&value))
}
