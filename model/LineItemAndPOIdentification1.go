package model

// Identifies the line item number and the purchase order.
type LineItemAndPOIdentification1 struct {

	// Identification assigned to a line item.
	LineItemIdentification []*Max70Text `xml:"LineItmId"`

	// Reference to the purchase order containing the line item.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`
}

func (l *LineItemAndPOIdentification1) AddLineItemIdentification(value string) {
	l.LineItemIdentification = append(l.LineItemIdentification, (*Max70Text)(&value))
}

func (l *LineItemAndPOIdentification1) AddPurchaseOrderReference() *DocumentIdentification7 {
	l.PurchaseOrderReference = new(DocumentIdentification7)
	return l.PurchaseOrderReference
}
