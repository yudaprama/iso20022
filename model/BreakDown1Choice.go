package model

// Specifies if breakdown is by purchase order or commercial invoice.
type BreakDown1Choice struct {

	// The intention to pay is based on a purchase order.
	ByPurchaseOrder *ReportLine5 `xml:"ByPurchsOrdr"`

	// The intention to pay is based on a commercial invoice.
	ByCommercialInvoice *ReportLine6 `xml:"ByComrclInvc"`
}

func (b *BreakDown1Choice) AddByPurchaseOrder() *ReportLine5 {
	b.ByPurchaseOrder = new(ReportLine5)
	return b.ByPurchaseOrder
}

func (b *BreakDown1Choice) AddByCommercialInvoice() *ReportLine6 {
	b.ByCommercialInvoice = new(ReportLine6)
	return b.ByCommercialInvoice
}
