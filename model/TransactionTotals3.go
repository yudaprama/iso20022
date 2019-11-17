package model

// Transaction totals during the reconciliation period, for a certain type of transaction.
type TransactionTotals3 struct {

	// Identifier assigned by the merchant identifying a set of POI terminals performing some categories of transactions.
	POIGroupIdentification *Max35Text `xml:"POIGrpId,omitempty"`

	// Category of cards related the acceptance processing rules defined by the acquirer.
	CardProductProfile *Max35Text `xml:"CardPdctPrfl,omitempty"`

	// Currency associated with the transaction totals.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Identification of the type of transaction.
	Type *TypeTransactionTotals2Code `xml:"Tp"`

	// Total number of transactions during a reconciliation period.
	TotalNumber *Number `xml:"TtlNb"`

	// Total amount of a collection of transactions.
	CumulativeAmount *ImpliedCurrencyAndAmount `xml:"CmltvAmt"`
}

func (t *TransactionTotals3) SetPOIGroupIdentification(value string) {
	t.POIGroupIdentification = (*Max35Text)(&value)
}

func (t *TransactionTotals3) SetCardProductProfile(value string) {
	t.CardProductProfile = (*Max35Text)(&value)
}

func (t *TransactionTotals3) SetCurrency(value string) {
	t.Currency = (*CurrencyCode)(&value)
}

func (t *TransactionTotals3) SetType(value string) {
	t.Type = (*TypeTransactionTotals2Code)(&value)
}

func (t *TransactionTotals3) SetTotalNumber(value string) {
	t.TotalNumber = (*Number)(&value)
}

func (t *TransactionTotals3) SetCumulativeAmount(value, currency string) {
	t.CumulativeAmount = NewImpliedCurrencyAndAmount(value, currency)
}
