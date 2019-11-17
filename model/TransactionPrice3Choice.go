package model

// Specifies the price information related to the underlying transaction.
type TransactionPrice3Choice struct {

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric"`

	// Proprietary price specification related to the underlying transaction.
	Proprietary []*ProprietaryPrice2 `xml:"Prtry"`
}

func (t *TransactionPrice3Choice) AddDealPrice() *Price2 {
	t.DealPrice = new(Price2)
	return t.DealPrice
}

func (t *TransactionPrice3Choice) AddProprietary() *ProprietaryPrice2 {
	newValue := new(ProprietaryPrice2)
	t.Proprietary = append(t.Proprietary, newValue)
	return newValue
}
