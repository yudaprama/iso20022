package model

// Set of characteristics that unambiguously identify the single invoice financing request.
type InvoiceRequestInformation1 struct {

	// General information that unambiguously identify the invoice to be financed, such as invoice type, invoice number and issue date.
	InvoiceGeneralInformation *DocumentGeneralInformation1 `xml:"InvcGnlInf"`

	// Specifies totals related to the invoice, such as total invoice amount and total tax amount.
	InvoiceTotalsInformation *InvoiceTotals1 `xml:"InvcTtlsInf"`

	// Amount of credit/debit note related to the invoice to be financed.
	CreditDebitNoteAmount *ActiveCurrencyAndAmount `xml:"CdtDbtNoteAmt,omitempty"`

	// Details of a single instalment to be financed, related to an invoice settlement (amount, payment due date).
	InstalmentInformation []*Instalment1 `xml:"InstlmtInf,omitempty"`

	// Amount requested by the requestor party, related to a single invoice to be financed.
	RequestedAmount *FinancingRateOrAmountChoice `xml:"ReqdAmt,omitempty"`

	// Person or organization that represents the creditor for the invoice to be financed.
	Supplier *PartyAndAccountIdentificationAndContactInformation1 `xml:"Spplr"`

	// Person or organization that represents the debtor for the invoice to be financed.
	Buyer *PartyIdentificationAndContactInformation1 `xml:"Buyr"`

	// Specifies payment terms and conditions related to a single invoice to be financed, including identifier of possible account used for payment.
	InvoicePaymentInformation *PaymentInformation15 `xml:"InvcPmtInf"`

	// Information about a document related to the invoice to be financed, in structured form.
	ReferredDocument []*ReferredDocumentInformation2 `xml:"RfrdDoc,omitempty"`
}

func (i *InvoiceRequestInformation1) AddInvoiceGeneralInformation() *DocumentGeneralInformation1 {
	i.InvoiceGeneralInformation = new(DocumentGeneralInformation1)
	return i.InvoiceGeneralInformation
}

func (i *InvoiceRequestInformation1) AddInvoiceTotalsInformation() *InvoiceTotals1 {
	i.InvoiceTotalsInformation = new(InvoiceTotals1)
	return i.InvoiceTotalsInformation
}

func (i *InvoiceRequestInformation1) SetCreditDebitNoteAmount(value, currency string) {
	i.CreditDebitNoteAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InvoiceRequestInformation1) AddInstalmentInformation() *Instalment1 {
	newValue := new(Instalment1)
	i.InstalmentInformation = append(i.InstalmentInformation, newValue)
	return newValue
}

func (i *InvoiceRequestInformation1) AddRequestedAmount() *FinancingRateOrAmountChoice {
	i.RequestedAmount = new(FinancingRateOrAmountChoice)
	return i.RequestedAmount
}

func (i *InvoiceRequestInformation1) AddSupplier() *PartyAndAccountIdentificationAndContactInformation1 {
	i.Supplier = new(PartyAndAccountIdentificationAndContactInformation1)
	return i.Supplier
}

func (i *InvoiceRequestInformation1) AddBuyer() *PartyIdentificationAndContactInformation1 {
	i.Buyer = new(PartyIdentificationAndContactInformation1)
	return i.Buyer
}

func (i *InvoiceRequestInformation1) AddInvoicePaymentInformation() *PaymentInformation15 {
	i.InvoicePaymentInformation = new(PaymentInformation15)
	return i.InvoicePaymentInformation
}

func (i *InvoiceRequestInformation1) AddReferredDocument() *ReferredDocumentInformation2 {
	newValue := new(ReferredDocumentInformation2)
	i.ReferredDocument = append(i.ReferredDocument, newValue)
	return newValue
}
