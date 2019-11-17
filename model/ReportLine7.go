package model

// Information about a payment against a purchase order.
type ReportLine7 struct {

	// Unique identification assigned by the matching application to the transaction.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Unique identification of the purchase order, assigned by the buyer.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Specifies the adjustments applied to obtain the net amount.
	Adjustment []*Adjustment6 `xml:"Adjstmnt,omitempty"`

	// Net amount, after adjustments, intended to be paid.
	NetAmount *CurrencyAndAmount `xml:"NetAmt"`
}

func (r *ReportLine7) SetTransactionIdentification(value string) {
	r.TransactionIdentification = (*Max35Text)(&value)
}

func (r *ReportLine7) AddPurchaseOrderReference() *DocumentIdentification7 {
	r.PurchaseOrderReference = new(DocumentIdentification7)
	return r.PurchaseOrderReference
}

func (r *ReportLine7) AddAdjustment() *Adjustment6 {
	newValue := new(Adjustment6)
	r.Adjustment = append(r.Adjustment, newValue)
	return newValue
}

func (r *ReportLine7) SetNetAmount(value, currency string) {
	r.NetAmount = NewCurrencyAndAmount(value, currency)
}
