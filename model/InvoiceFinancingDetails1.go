package model

// Detailed information about single invoice/instalment financing result, such as result of request (financed or not financed), amount, percentage applied.
type InvoiceFinancingDetails1 struct {

	// General information that unambiguously identifies the invoice contained in the original request.
	OriginalInvoiceInformation *OriginalInvoiceInformation1 `xml:"OrgnlInvcInf"`

	// Person or organization that represents the creditor for the invoice to be financed.
	Supplier *PartyIdentification8 `xml:"Spplr,omitempty"`

	// Information about result of invoice financing request.
	InvoiceFinancingResult *FinancingResult1 `xml:"InvcFincgRslt"`

	// Includes details about a single instalment within an invoice, such as identification and amount.
	InstalmentFinancingInformation []*InstalmentFinancingInformation1 `xml:"InstlmtFincgInf,omitempty"`
}

func (i *InvoiceFinancingDetails1) AddOriginalInvoiceInformation() *OriginalInvoiceInformation1 {
	i.OriginalInvoiceInformation = new(OriginalInvoiceInformation1)
	return i.OriginalInvoiceInformation
}

func (i *InvoiceFinancingDetails1) AddSupplier() *PartyIdentification8 {
	i.Supplier = new(PartyIdentification8)
	return i.Supplier
}

func (i *InvoiceFinancingDetails1) AddInvoiceFinancingResult() *FinancingResult1 {
	i.InvoiceFinancingResult = new(FinancingResult1)
	return i.InvoiceFinancingResult
}

func (i *InvoiceFinancingDetails1) AddInstalmentFinancingInformation() *InstalmentFinancingInformation1 {
	newValue := new(InstalmentFinancingInformation1)
	i.InstalmentFinancingInformation = append(i.InstalmentFinancingInformation, newValue)
	return newValue
}
