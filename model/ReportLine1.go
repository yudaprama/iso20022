package model

// Information about a payment against a purchase order.
type ReportLine1 struct {

	// Unique identification assigned by the matching application to the transaction.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Identifies the status of the transaction.
	TransactionStatus *TransactionStatus4 `xml:"TxSts"`

	// Unique identification of the purchase order, assigned by the buyer.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Total amount of the purchase order after taxes, adjustments and charges.
	PurchaseOrderTotalNetAmount *CurrencyAndAmount `xml:"PurchsOrdrTtlNetAmt"`

	// Accumulated net amount, after adjustments, intended to be paid.
	AccumulatedNetAmount *CurrencyAndAmount `xml:"AcmltdNetAmt"`
}

func (r *ReportLine1) SetTransactionIdentification(value string) {
	r.TransactionIdentification = (*Max35Text)(&value)
}

func (r *ReportLine1) AddTransactionStatus() *TransactionStatus4 {
	r.TransactionStatus = new(TransactionStatus4)
	return r.TransactionStatus
}

func (r *ReportLine1) AddPurchaseOrderReference() *DocumentIdentification7 {
	r.PurchaseOrderReference = new(DocumentIdentification7)
	return r.PurchaseOrderReference
}

func (r *ReportLine1) SetPurchaseOrderTotalNetAmount(value, currency string) {
	r.PurchaseOrderTotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (r *ReportLine1) SetAccumulatedNetAmount(value, currency string) {
	r.AccumulatedNetAmount = NewCurrencyAndAmount(value, currency)
}
