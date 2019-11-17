package model

// Net cash movement to a fund as a result of investment funds transactions.
type NetCashForecast2 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt"`

	// Net amount of the cash flow, expressed as an amount of money.
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Net amount, expressed as a number of units.
	NetUnitsNumber *FinancialInstrumentQuantity1 `xml:"NetUnitsNb,omitempty"`

	// Specifies the direction of the cash flow from the perspective of the fund.
	FlowDirection *FlowDirectionType1Code `xml:"FlowDrctn"`
}

func (n *NetCashForecast2) SetCashSettlementDate(value string) {
	n.CashSettlementDate = (*ISODate)(&value)
}

func (n *NetCashForecast2) SetNetAmount(value, currency string) {
	n.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (n *NetCashForecast2) AddNetUnitsNumber() *FinancialInstrumentQuantity1 {
	n.NetUnitsNumber = new(FinancialInstrumentQuantity1)
	return n.NetUnitsNumber
}

func (n *NetCashForecast2) SetFlowDirection(value string) {
	n.FlowDirection = (*FlowDirectionType1Code)(&value)
}
