package model

// Trade settlement details for this invoice which involves the payment of an outstanding debt, account, or charge
type TradeSettlement1 struct {

	// Monetary value that is an exact amount due and payable, such as the amount due to the creditor.
	DuePayableAmount []*CurrencyAndAmount `xml:"DuePyblAmt,omitempty"`

	// Unique and unambiguous reference assigned by the creditor.
	CreditorReference []*CreditorReferenceInformation2 `xml:"CdtrRef,omitempty"`

	// Unique and unambiguous identifier for a payment transaction, as assigned by the originator. The payment transaction reference is used for reconciliation or to link tasks relating to the payment transaction.
	PaymentReference []*Max35Text `xml:"PmtRef,omitempty"`

	// Code specifying the currency of the invoice.
	InvoiceCurrencyCode *CurrencyCode `xml:"InvcCcyCd,omitempty"`

	// Organization issuing the invoice.
	Invoicer *TradeParty1 `xml:"Invcr,omitempty"`

	// Party to whom the invoice was issued.
	Invoicee *TradeParty1 `xml:"Invcee,omitempty"`

	// Party specified to receive payment for the invoice.
	Payee *TradeParty1 `xml:"Pyee,omitempty"`

	// Party specified to initiate payment for the invoice.
	Payer *TradeParty1 `xml:"Pyer,omitempty"`

	// Currency exchange applicable to a tax.
	TaxCurrencyExchange *CurrencyReference2 `xml:"TaxCcyXchg,omitempty"`

	// Currency exchange applicable to the invoice.
	InvoiceCurrencyExchange *CurrencyReference2 `xml:"InvcCcyXchg,omitempty"`

	// Currency exchange applicable to the payment.
	PaymentCurrencyExchange *CurrencyReference2 `xml:"PmtCcyXchg,omitempty"`

	// Means of payment (for example, credit transfer, cheque, money order, or credit card) specified to initiate payment of the invoice.
	PaymentMeans []*PaymentMeans1 `xml:"PmtMeans,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
	Tax []*SettlementTax1 `xml:"Tax,omitempty"`

	// Specifies the applicable billing period.
	BillingPeriod *Period1 `xml:"BllgPrd,omitempty"`

	// Allowance or charge specified.
	AllowanceCharge []*SettlementAllowanceCharge1 `xml:"AllwncChrg,omitempty"`

	// Tax subtotal calculated.
	SubTotalCalculatedTax []*SettlementSubTotalCalculatedTax1 `xml:"SubTtlClctdTax,omitempty"`

	// Logistics service charge specified.
	LogisticsCharge []*ChargesDetails2 `xml:"LogstcsChrg,omitempty"`

	// Payment terms.
	PaymentTerms []*PaymentTerms3 `xml:"PmtTerms,omitempty"`

	// Monetary totals specified for the invoice.
	MonetarySummation *SettlementMonetarySummation1 `xml:"MntrySummtn"`

	// Financial adjustment specified.
	AdjustmentAmountAndReason []*DocumentAdjustment2 `xml:"AdjstmntAmtAndRsn,omitempty"`

	// Invoice document referenced.
	InvoiceReferencedDocument *DocumentIdentification22 `xml:"InvcRefdDoc,omitempty"`

	// Pro-forma invoice document referenced.
	ProformaInvoiceReferencedDocument *DocumentIdentification22 `xml:"ProfrmInvcRefdDoc,omitempty"`

	// Letter of credit document referenced.
	LetterOfCreditReferencedDocument *DocumentIdentification7 `xml:"LttrOfCdtRefdDoc,omitempty"`

	// Financial card specified. The card is used to represent a financial account for the purpose of payment settlement.
	FinancialCard []*FinancialCard1 `xml:"FinCard,omitempty"`

	// Specific purchase account for recording debits and credits for accounting purposes.
	PurchaseAccountingAccount []*AccountingAccount1 `xml:"PurchsAcctgAcct,omitempty"`

	// Factoring list document referenced.
	IssuerFactoringListIdentification []*Max35Text `xml:"IssrFactrgListId,omitempty"`

	// Factoring agreement document referenced.
	IssuerFactoringAgreementIdentification []*Max35Text `xml:"IssrFactrgAgrmtId,omitempty"`
}

func (t *TradeSettlement1) AddDuePayableAmount(value, currency string) {
	t.DuePayableAmount = append(t.DuePayableAmount, NewCurrencyAndAmount(value, currency))
}

func (t *TradeSettlement1) AddCreditorReference() *CreditorReferenceInformation2 {
	newValue := new(CreditorReferenceInformation2)
	t.CreditorReference = append(t.CreditorReference, newValue)
	return newValue
}

func (t *TradeSettlement1) AddPaymentReference(value string) {
	t.PaymentReference = append(t.PaymentReference, (*Max35Text)(&value))
}

func (t *TradeSettlement1) SetInvoiceCurrencyCode(value string) {
	t.InvoiceCurrencyCode = (*CurrencyCode)(&value)
}

func (t *TradeSettlement1) AddInvoicer() *TradeParty1 {
	t.Invoicer = new(TradeParty1)
	return t.Invoicer
}

func (t *TradeSettlement1) AddInvoicee() *TradeParty1 {
	t.Invoicee = new(TradeParty1)
	return t.Invoicee
}

func (t *TradeSettlement1) AddPayee() *TradeParty1 {
	t.Payee = new(TradeParty1)
	return t.Payee
}

func (t *TradeSettlement1) AddPayer() *TradeParty1 {
	t.Payer = new(TradeParty1)
	return t.Payer
}

func (t *TradeSettlement1) AddTaxCurrencyExchange() *CurrencyReference2 {
	t.TaxCurrencyExchange = new(CurrencyReference2)
	return t.TaxCurrencyExchange
}

func (t *TradeSettlement1) AddInvoiceCurrencyExchange() *CurrencyReference2 {
	t.InvoiceCurrencyExchange = new(CurrencyReference2)
	return t.InvoiceCurrencyExchange
}

func (t *TradeSettlement1) AddPaymentCurrencyExchange() *CurrencyReference2 {
	t.PaymentCurrencyExchange = new(CurrencyReference2)
	return t.PaymentCurrencyExchange
}

func (t *TradeSettlement1) AddPaymentMeans() *PaymentMeans1 {
	newValue := new(PaymentMeans1)
	t.PaymentMeans = append(t.PaymentMeans, newValue)
	return newValue
}

func (t *TradeSettlement1) AddTax() *SettlementTax1 {
	newValue := new(SettlementTax1)
	t.Tax = append(t.Tax, newValue)
	return newValue
}

func (t *TradeSettlement1) AddBillingPeriod() *Period1 {
	t.BillingPeriod = new(Period1)
	return t.BillingPeriod
}

func (t *TradeSettlement1) AddAllowanceCharge() *SettlementAllowanceCharge1 {
	newValue := new(SettlementAllowanceCharge1)
	t.AllowanceCharge = append(t.AllowanceCharge, newValue)
	return newValue
}

func (t *TradeSettlement1) AddSubTotalCalculatedTax() *SettlementSubTotalCalculatedTax1 {
	newValue := new(SettlementSubTotalCalculatedTax1)
	t.SubTotalCalculatedTax = append(t.SubTotalCalculatedTax, newValue)
	return newValue
}

func (t *TradeSettlement1) AddLogisticsCharge() *ChargesDetails2 {
	newValue := new(ChargesDetails2)
	t.LogisticsCharge = append(t.LogisticsCharge, newValue)
	return newValue
}

func (t *TradeSettlement1) AddPaymentTerms() *PaymentTerms3 {
	newValue := new(PaymentTerms3)
	t.PaymentTerms = append(t.PaymentTerms, newValue)
	return newValue
}

func (t *TradeSettlement1) AddMonetarySummation() *SettlementMonetarySummation1 {
	t.MonetarySummation = new(SettlementMonetarySummation1)
	return t.MonetarySummation
}

func (t *TradeSettlement1) AddAdjustmentAmountAndReason() *DocumentAdjustment2 {
	newValue := new(DocumentAdjustment2)
	t.AdjustmentAmountAndReason = append(t.AdjustmentAmountAndReason, newValue)
	return newValue
}

func (t *TradeSettlement1) AddInvoiceReferencedDocument() *DocumentIdentification22 {
	t.InvoiceReferencedDocument = new(DocumentIdentification22)
	return t.InvoiceReferencedDocument
}

func (t *TradeSettlement1) AddProformaInvoiceReferencedDocument() *DocumentIdentification22 {
	t.ProformaInvoiceReferencedDocument = new(DocumentIdentification22)
	return t.ProformaInvoiceReferencedDocument
}

func (t *TradeSettlement1) AddLetterOfCreditReferencedDocument() *DocumentIdentification7 {
	t.LetterOfCreditReferencedDocument = new(DocumentIdentification7)
	return t.LetterOfCreditReferencedDocument
}

func (t *TradeSettlement1) AddFinancialCard() *FinancialCard1 {
	newValue := new(FinancialCard1)
	t.FinancialCard = append(t.FinancialCard, newValue)
	return newValue
}

func (t *TradeSettlement1) AddPurchaseAccountingAccount() *AccountingAccount1 {
	newValue := new(AccountingAccount1)
	t.PurchaseAccountingAccount = append(t.PurchaseAccountingAccount, newValue)
	return newValue
}

func (t *TradeSettlement1) AddIssuerFactoringListIdentification(value string) {
	t.IssuerFactoringListIdentification = append(t.IssuerFactoringListIdentification, (*Max35Text)(&value))
}

func (t *TradeSettlement1) AddIssuerFactoringAgreementIdentification(value string) {
	t.IssuerFactoringAgreementIdentification = append(t.IssuerFactoringAgreementIdentification, (*Max35Text)(&value))
}
