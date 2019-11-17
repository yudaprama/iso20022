package model

// Net cash movement to a fund as a result of investment funds transactions.
type NetCashForecast3 struct {

	// Net amount of the cash flow, expressed as an amount of money.
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Net amount, expressed as a number of units.
	NetUnitsNumber *FinancialInstrumentQuantity1 `xml:"NetUnitsNb,omitempty"`

	// Specifies the direction of the cash flow from the perspective of the fund.
	FlowDirection *FlowDirectionType1Code `xml:"FlowDrctn"`
}

func (n *NetCashForecast3) SetNetAmount(value, currency string) {
	n.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (n *NetCashForecast3) AddNetUnitsNumber() *FinancialInstrumentQuantity1 {
	n.NetUnitsNumber = new(FinancialInstrumentQuantity1)
	return n.NetUnitsNumber
}

func (n *NetCashForecast3) SetFlowDirection(value string) {
	n.FlowDirection = (*FlowDirectionType1Code)(&value)
}
