package model

// Balances of units and cash derived from investment fund orders.
type FundBalance1 struct {

	// Total number of units from orders placed in units.
	TotalUnitsFromUnitOrders *FinancialInstrumentQuantity1 `xml:"TtlUnitsFrUnitOrdrs,omitempty"`

	// Number of units derived from orders placed in cash.
	TotalUnitsFromCashOrders *FinancialInstrumentQuantity1 `xml:"TtlUnitsFrCshOrdrs,omitempty"`

	// Total amount of cash derived from orders placed as units.
	TotalCashFromUnitOrders *ActiveOrHistoricCurrencyAndAmount `xml:"TtlCshFrUnitOrdrs,omitempty"`

	// Total amount of cash from orders placed in cash.
	TotalCashFromCashOrders *ActiveOrHistoricCurrencyAndAmount `xml:"TtlCshFrCshOrdrs,omitempty"`
}

func (f *FundBalance1) AddTotalUnitsFromUnitOrders() *FinancialInstrumentQuantity1 {
	f.TotalUnitsFromUnitOrders = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsFromUnitOrders
}

func (f *FundBalance1) AddTotalUnitsFromCashOrders() *FinancialInstrumentQuantity1 {
	f.TotalUnitsFromCashOrders = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsFromCashOrders
}

func (f *FundBalance1) SetTotalCashFromUnitOrders(value, currency string) {
	f.TotalCashFromUnitOrders = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundBalance1) SetTotalCashFromCashOrders(value, currency string) {
	f.TotalCashFromCashOrders = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
