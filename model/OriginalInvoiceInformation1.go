package model

// General information about the invoice contained in the original request.
type OriginalInvoiceInformation1 struct {

	// Unique identifier of the document.
	DocumentNumber *Max35Text `xml:"DocNb"`

	// Total amount of the invoice, being the sum of total invoice lines amounts, total invoice additional amounts (allowances and charges) and total tax amounts.
	TotalInvoiceAmount *ActiveCurrencyAndAmount `xml:"TtlInvcAmt"`

	// Issue date of the document.
	IssueDate *ISODate `xml:"IsseDt"`

	// Due date for the payment of the invoice.
	PaymentDueDate *ISODate `xml:"PmtDueDt"`
}

func (o *OriginalInvoiceInformation1) SetDocumentNumber(value string) {
	o.DocumentNumber = (*Max35Text)(&value)
}

func (o *OriginalInvoiceInformation1) SetTotalInvoiceAmount(value, currency string) {
	o.TotalInvoiceAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (o *OriginalInvoiceInformation1) SetIssueDate(value string) {
	o.IssueDate = (*ISODate)(&value)
}

func (o *OriginalInvoiceInformation1) SetPaymentDueDate(value string) {
	o.PaymentDueDate = (*ISODate)(&value)
}
