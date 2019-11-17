package model

// Set of elements identifying the price information related to the underlying transaction.
type TransactionPrice1Choice struct {

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *CurrencyAndAmount `xml:"DealPric"`

	// Proprietary price specification of the underlying transaction.
	Proprietary []*ProprietaryPrice1 `xml:"Prtry"`
}

func (t *TransactionPrice1Choice) SetDealPrice(value, currency string) {
	t.DealPrice = NewCurrencyAndAmount(value, currency)
}

func (t *TransactionPrice1Choice) AddProprietary() *ProprietaryPrice1 {
	newValue := new(ProprietaryPrice1)
	t.Proprietary = append(t.Proprietary, newValue)
	return newValue
}
