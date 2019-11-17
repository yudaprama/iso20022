package model

// Specifies a collection of monetary totals for this settlement.
type SettlementMonetarySummation1 struct {

	// Monetary value of the line amount total being reported for this settlement.
	LineTotalAmount []*CurrencyAndAmount `xml:"LineTtlAmt,omitempty"`

	// Monetary value of the allowance total being reported for this settlement.
	AllowanceTotalAmount []*CurrencyAndAmount `xml:"AllwncTtlAmt,omitempty"`

	// Monetary value of the total discount being reported for this settlement.
	TotalDiscountAmount []*CurrencyAndAmount `xml:"TtlDscntAmt,omitempty"`

	// Monetary value of the charge amount total being reported for this settlement.
	ChargeTotalAmount []*CurrencyAndAmount `xml:"ChrgTtlAmt,omitempty"`

	// Monetary value of the total prepaid amount being reported for this settlement.
	TotalPrepaidAmount []*CurrencyAndAmount `xml:"TtlPrepdAmt,omitempty"`

	// Monetary value of the total of all tax basis amounts being reported for this settlement.
	TaxTotalAmount []*CurrencyAndAmount `xml:"TaxTtlAmt,omitempty"`

	// Monetary value of the total of all tax basis amounts being reported for this settlement.
	TaxBasisAmount []*CurrencyAndAmount `xml:"TaxBsisAmt,omitempty"`

	// Monetary value of a rounding amount being applied and reported for this settlement.
	RoundingAmount []*CurrencyAndAmount `xml:"RndgAmt,omitempty"`

	// Monetary value of the grand total being reported for this settlement, to include addition and subtraction of individual summation amounts.
	GrandTotalAmount []*CurrencyAndAmount `xml:"GrdTtlAmt,omitempty"`

	// Monetary value of an amount being reported as information for this settlement.
	InformationAmount []*CurrencyAndAmount `xml:"InfAmt,omitempty"`
}

func (s *SettlementMonetarySummation1) AddLineTotalAmount(value, currency string) {
	s.LineTotalAmount = append(s.LineTotalAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementMonetarySummation1) AddAllowanceTotalAmount(value, currency string) {
	s.AllowanceTotalAmount = append(s.AllowanceTotalAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementMonetarySummation1) AddTotalDiscountAmount(value, currency string) {
	s.TotalDiscountAmount = append(s.TotalDiscountAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementMonetarySummation1) AddChargeTotalAmount(value, currency string) {
	s.ChargeTotalAmount = append(s.ChargeTotalAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementMonetarySummation1) AddTotalPrepaidAmount(value, currency string) {
	s.TotalPrepaidAmount = append(s.TotalPrepaidAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementMonetarySummation1) AddTaxTotalAmount(value, currency string) {
	s.TaxTotalAmount = append(s.TaxTotalAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementMonetarySummation1) AddTaxBasisAmount(value, currency string) {
	s.TaxBasisAmount = append(s.TaxBasisAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementMonetarySummation1) AddRoundingAmount(value, currency string) {
	s.RoundingAmount = append(s.RoundingAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementMonetarySummation1) AddGrandTotalAmount(value, currency string) {
	s.GrandTotalAmount = append(s.GrandTotalAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementMonetarySummation1) AddInformationAmount(value, currency string) {
	s.InformationAmount = append(s.InformationAmount, NewCurrencyAndAmount(value, currency))
}
