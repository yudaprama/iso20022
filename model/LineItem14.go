package model

// Calculation of the current situation of a baseline as a result of the submission of a commercial data set.
type LineItem14 struct {

	// Calculated information about the goods of the underlying transaction.
	LineItemDetails []*LineItemDetails12 `xml:"LineItmDtls"`

	// Line items total amount as indicated in the baseline.
	OrderedLineItemsTotalAmount *CurrencyAndAmount `xml:"OrdrdLineItmsTtlAmt"`

	// Line items total amount accepted by a data set submission(s).
	AcceptedLineItemsTotalAmount *CurrencyAndAmount `xml:"AccptdLineItmsTtlAmt"`

	// Difference between the ordered and the accepted line items total amount.
	OutstandingLineItemsTotalAmount *CurrencyAndAmount `xml:"OutsdngLineItmsTtlAmt"`

	// Line item total amount for which a mismatched data set has been submitted and has not yet been accepted or rejected.
	PendingLineItemsTotalAmount *CurrencyAndAmount `xml:"PdgLineItmsTtlAmt"`

	// Total net amount as indicated in the baseline.
	OrderedTotalNetAmount *CurrencyAndAmount `xml:"OrdrdTtlNetAmt"`

	// Total net amount accepted by a data set submission.
	AcceptedTotalNetAmount *CurrencyAndAmount `xml:"AccptdTtlNetAmt"`

	// Total net amount for which a mismatched data set has been submitted and has not yet been accepted or rejected.
	OutstandingTotalNetAmount *CurrencyAndAmount `xml:"OutsdngTtlNetAmt"`

	// Difference between the ordered and the accepted total net amount.
	PendingTotalNetAmount *CurrencyAndAmount `xml:"PdgTtlNetAmt"`
}

func (l *LineItem14) AddLineItemDetails() *LineItemDetails12 {
	newValue := new(LineItemDetails12)
	l.LineItemDetails = append(l.LineItemDetails, newValue)
	return newValue
}

func (l *LineItem14) SetOrderedLineItemsTotalAmount(value, currency string) {
	l.OrderedLineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem14) SetAcceptedLineItemsTotalAmount(value, currency string) {
	l.AcceptedLineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem14) SetOutstandingLineItemsTotalAmount(value, currency string) {
	l.OutstandingLineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem14) SetPendingLineItemsTotalAmount(value, currency string) {
	l.PendingLineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem14) SetOrderedTotalNetAmount(value, currency string) {
	l.OrderedTotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem14) SetAcceptedTotalNetAmount(value, currency string) {
	l.AcceptedTotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem14) SetOutstandingTotalNetAmount(value, currency string) {
	l.OutstandingTotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem14) SetPendingTotalNetAmount(value, currency string) {
	l.PendingTotalNetAmount = NewCurrencyAndAmount(value, currency)
}
