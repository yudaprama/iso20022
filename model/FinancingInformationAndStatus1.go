package model

// Financing information and status.
type FinancingInformationAndStatus1 struct {

	// Specifies summary information about invoices/instalments financed, such as total amount financed, number of single requests accepted.
	FinancingAllowedSummary *FinancingAllowedSummary1 `xml:"FincgAllwdSummry"`

	// Specifies detailed information about single invoice/instalment financing result, such as result of request (financed or not financed), amount, percentage applied.
	InvoiceFinancingDetails []*InvoiceFinancingDetails1 `xml:"InvcFincgDtls"`
}

func (f *FinancingInformationAndStatus1) AddFinancingAllowedSummary() *FinancingAllowedSummary1 {
	f.FinancingAllowedSummary = new(FinancingAllowedSummary1)
	return f.FinancingAllowedSummary
}

func (f *FinancingInformationAndStatus1) AddInvoiceFinancingDetails() *InvoiceFinancingDetails1 {
	newValue := new(InvoiceFinancingDetails1)
	f.InvoiceFinancingDetails = append(f.InvoiceFinancingDetails, newValue)
	return newValue
}
