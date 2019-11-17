package model

// Total quantity of securities to be transferred, expressed in a number of units or a percentage rate.
type Quantity13Choice struct {

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Total quantity of securities to be settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate"`
}

func (q *Quantity13Choice) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	q.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return q.TotalUnitsNumber
}

func (q *Quantity13Choice) SetPortfolioTransferOutRate(value string) {
	q.PortfolioTransferOutRate = (*PercentageRate)(&value)
}
