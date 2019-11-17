package model

// Specifies the parameters for which a transaction report must be generated.
type ReportSpecification4 struct {

	// Unique identification assigned by the matching application to a transaction, for which the matching application must generate a report.
	TransactionIdentification []*Max35Text `xml:"TxId,omitempty"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus []*TransactionStatus4 `xml:"TxSts,omitempty"`

	// Reference to the identification of a transaction of a user, for which the matching application must generate a report.
	SubmitterTransactionReference []*Max35Text `xml:"SubmitrTxRef,omitempty"`

	// Specifies a list of entities for which the matching application must generate a report.
	EntitiesToBeReported []*BICIdentification1 `xml:"NttiesToBeRptd,omitempty"`

	// Financial institution that is the counterparty to the trade transaction.
	Correspondent []*BICIdentification1 `xml:"Crspdt,omitempty"`

	// Financial institution that is a data set submitting bank to the transaction.
	SubmittingBank []*BICIdentification1 `xml:"SubmitgBk,omitempty"`

	// Financial institution that is an obligor bank to the transaction.
	ObligorBank []*BICIdentification1 `xml:"OblgrBk,omitempty"`

	// Party that buys goods or services, or a financial instrument.
	Buyer []*PartyIdentification28 `xml:"Buyr,omitempty"`

	// Party that sells goods or services, or a financial instrument.
	Seller []*PartyIdentification28 `xml:"Sellr,omitempty"`

	// Country of the buyer.
	BuyerCountry []*CountryCode `xml:"BuyrCtry,omitempty"`

	// Country of the seller.
	SellerCountry []*CountryCode `xml:"SellrCtry,omitempty"`

	// Country of the financial institution which is the other party to the trade.
	CorrespondentCountry []*CountryCode `xml:"CrspdtCtry,omitempty"`

	// Specifies a pending request for action for which the matching application must generate a report.
	PendingRequestForAction []*PendingActivity1 `xml:"PdgReqForActn,omitempty"`
}

func (r *ReportSpecification4) AddTransactionIdentification(value string) {
	r.TransactionIdentification = append(r.TransactionIdentification, (*Max35Text)(&value))
}

func (r *ReportSpecification4) AddTransactionStatus() *TransactionStatus4 {
	newValue := new(TransactionStatus4)
	r.TransactionStatus = append(r.TransactionStatus, newValue)
	return newValue
}

func (r *ReportSpecification4) AddSubmitterTransactionReference(value string) {
	r.SubmitterTransactionReference = append(r.SubmitterTransactionReference, (*Max35Text)(&value))
}

func (r *ReportSpecification4) AddEntitiesToBeReported() *BICIdentification1 {
	newValue := new(BICIdentification1)
	r.EntitiesToBeReported = append(r.EntitiesToBeReported, newValue)
	return newValue
}

func (r *ReportSpecification4) AddCorrespondent() *BICIdentification1 {
	newValue := new(BICIdentification1)
	r.Correspondent = append(r.Correspondent, newValue)
	return newValue
}

func (r *ReportSpecification4) AddSubmittingBank() *BICIdentification1 {
	newValue := new(BICIdentification1)
	r.SubmittingBank = append(r.SubmittingBank, newValue)
	return newValue
}

func (r *ReportSpecification4) AddObligorBank() *BICIdentification1 {
	newValue := new(BICIdentification1)
	r.ObligorBank = append(r.ObligorBank, newValue)
	return newValue
}

func (r *ReportSpecification4) AddBuyer() *PartyIdentification28 {
	newValue := new(PartyIdentification28)
	r.Buyer = append(r.Buyer, newValue)
	return newValue
}

func (r *ReportSpecification4) AddSeller() *PartyIdentification28 {
	newValue := new(PartyIdentification28)
	r.Seller = append(r.Seller, newValue)
	return newValue
}

func (r *ReportSpecification4) AddBuyerCountry(value string) {
	r.BuyerCountry = append(r.BuyerCountry, (*CountryCode)(&value))
}

func (r *ReportSpecification4) AddSellerCountry(value string) {
	r.SellerCountry = append(r.SellerCountry, (*CountryCode)(&value))
}

func (r *ReportSpecification4) AddCorrespondentCountry(value string) {
	r.CorrespondentCountry = append(r.CorrespondentCountry, (*CountryCode)(&value))
}

func (r *ReportSpecification4) AddPendingRequestForAction() *PendingActivity1 {
	newValue := new(PendingActivity1)
	r.PendingRequestForAction = append(r.PendingRequestForAction, newValue)
	return newValue
}
