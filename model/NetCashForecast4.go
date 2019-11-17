package model

// Net cash movement to a fund as a result of investment funds transactions.
type NetCashForecast4 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt"`

	// Net amount of the cash flow, expressed as an amount of money.
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Net amount, expressed as a number of units.
	NetUnitsNumber *FinancialInstrumentQuantity1 `xml:"NetUnitsNb,omitempty"`

	// Specifies the direction of the cash flow from the perspective of the fund.
	FlowDirection *FlowDirectionType1Code `xml:"FlowDrctn"`

	// Additional balances for cash amounts and number of units.
	// In an estimated report, the total cash derived from orders placed as a number of units is an estimated cash amount and the total number of units derived from orders placed as a cash amount is an estimated number of units.
	AdditionalBalance *FundBalance1 `xml:"AddtlBal,omitempty"`
}

func (n *NetCashForecast4) SetCashSettlementDate(value string) {
	n.CashSettlementDate = (*ISODate)(&value)
}

func (n *NetCashForecast4) SetNetAmount(value, currency string) {
	n.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (n *NetCashForecast4) AddNetUnitsNumber() *FinancialInstrumentQuantity1 {
	n.NetUnitsNumber = new(FinancialInstrumentQuantity1)
	return n.NetUnitsNumber
}

func (n *NetCashForecast4) SetFlowDirection(value string) {
	n.FlowDirection = (*FlowDirectionType1Code)(&value)
}

func (n *NetCashForecast4) AddAdditionalBalance() *FundBalance1 {
	n.AdditionalBalance = new(FundBalance1)
	return n.AdditionalBalance
}
