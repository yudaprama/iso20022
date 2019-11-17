package model

// Detailed description of the items that correspond to the parameters set in a request and for which a report has been generated.
type TransactionReportItems3 struct {

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *TransactionStatus4 `xml:"TxSts"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Reference to the purchase order of the underlying transaction.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Party that buys goods or services, or a financial instrument.
	Buyer *PartyIdentification26 `xml:"Buyr"`

	// Party that sells goods or services, or a financial instrument.
	Seller *PartyIdentification26 `xml:"Sellr"`

	// Financial institution of the buyer.
	BuyerBank *BICIdentification1 `xml:"BuyrBk"`

	// Country of the buyer bank.
	BuyerBankCountry *CountryCode `xml:"BuyrBkCtry"`

	// Financial institution of the seller.
	SellerBank *BICIdentification1 `xml:"SellrBk"`

	// Country of the seller bank.
	SellerBankCountry *CountryCode `xml:"SellrBkCtry"`

	// Financial institution that is an obligor bank to the transaction.
	ObligorBank []*BICIdentification1 `xml:"OblgrBk,omitempty"`

	// Financial institution that is a data set submitting bank to the transaction.
	SubmittingBank []*BICIdentification1 `xml:"SubmitgBk,omitempty"`

	// Amount of baseline not yet utilised.
	OutstandingAmount *CurrencyAndAmount `xml:"OutsdngAmt"`

	// Total net amount as specified in the baseline.
	TotalNetAmount *CurrencyAndAmount `xml:"TtlNetAmt"`

	// Next processing step required.
	PendingRequestForAction []*PendingActivity2 `xml:"PdgReqForActn,omitempty"`
}

func (t *TransactionReportItems3) SetTransactionIdentification(value string) {
	t.TransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionReportItems3) AddEstablishedBaselineIdentification() *DocumentIdentification3 {
	t.EstablishedBaselineIdentification = new(DocumentIdentification3)
	return t.EstablishedBaselineIdentification
}

func (t *TransactionReportItems3) AddTransactionStatus() *TransactionStatus4 {
	t.TransactionStatus = new(TransactionStatus4)
	return t.TransactionStatus
}

func (t *TransactionReportItems3) AddUserTransactionReference() *DocumentIdentification5 {
	newValue := new(DocumentIdentification5)
	t.UserTransactionReference = append(t.UserTransactionReference, newValue)
	return newValue
}

func (t *TransactionReportItems3) AddPurchaseOrderReference() *DocumentIdentification7 {
	t.PurchaseOrderReference = new(DocumentIdentification7)
	return t.PurchaseOrderReference
}

func (t *TransactionReportItems3) AddBuyer() *PartyIdentification26 {
	t.Buyer = new(PartyIdentification26)
	return t.Buyer
}

func (t *TransactionReportItems3) AddSeller() *PartyIdentification26 {
	t.Seller = new(PartyIdentification26)
	return t.Seller
}

func (t *TransactionReportItems3) AddBuyerBank() *BICIdentification1 {
	t.BuyerBank = new(BICIdentification1)
	return t.BuyerBank
}

func (t *TransactionReportItems3) SetBuyerBankCountry(value string) {
	t.BuyerBankCountry = (*CountryCode)(&value)
}

func (t *TransactionReportItems3) AddSellerBank() *BICIdentification1 {
	t.SellerBank = new(BICIdentification1)
	return t.SellerBank
}

func (t *TransactionReportItems3) SetSellerBankCountry(value string) {
	t.SellerBankCountry = (*CountryCode)(&value)
}

func (t *TransactionReportItems3) AddObligorBank() *BICIdentification1 {
	newValue := new(BICIdentification1)
	t.ObligorBank = append(t.ObligorBank, newValue)
	return newValue
}

func (t *TransactionReportItems3) AddSubmittingBank() *BICIdentification1 {
	newValue := new(BICIdentification1)
	t.SubmittingBank = append(t.SubmittingBank, newValue)
	return newValue
}

func (t *TransactionReportItems3) SetOutstandingAmount(value, currency string) {
	t.OutstandingAmount = NewCurrencyAndAmount(value, currency)
}

func (t *TransactionReportItems3) SetTotalNetAmount(value, currency string) {
	t.TotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (t *TransactionReportItems3) AddPendingRequestForAction() *PendingActivity2 {
	newValue := new(PendingActivity2)
	t.PendingRequestForAction = append(t.PendingRequestForAction, newValue)
	return newValue
}
