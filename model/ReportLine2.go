package model

// Information about a payment against a purchase order.
type ReportLine2 struct {

	// Unique identification assigned by the matching application to the transaction.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Unique identification of the purchase order, assigned by the buyer.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Specifies the adjustments applied to obtain the net amount.
	Adjustment []*Adjustment4 `xml:"Adjstmnt,omitempty"`

	// Net amount, after adjustments, intended to be paid.
	NetAmount *CurrencyAndAmount `xml:"NetAmt"`
}

func (r *ReportLine2) SetTransactionIdentification(value string) {
	r.TransactionIdentification = (*Max35Text)(&value)
}

func (r *ReportLine2) AddPurchaseOrderReference() *DocumentIdentification7 {
	r.PurchaseOrderReference = new(DocumentIdentification7)
	return r.PurchaseOrderReference
}

func (r *ReportLine2) AddAdjustment() *Adjustment4 {
	newValue := new(Adjustment4)
	r.Adjustment = append(r.Adjustment, newValue)
	return newValue
}

func (r *ReportLine2) SetNetAmount(value, currency string) {
	r.NetAmount = NewCurrencyAndAmount(value, currency)
}
