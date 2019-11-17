package model

// Regroups identification parameters for trade items.
type FinancialItemParameters1 struct {

	// Unique identification of this item relative to the issuing party.
	Identifier *Max35Text `xml:"Idr"`

	// Date of creation of the item.
	IssueDate *ISODate `xml:"IsseDt"`

	// Identifier of related items, for example an assignment or an advice.
	RelatedItem []*QualifiedDocumentInformation1 `xml:"RltdItm,omitempty"`

	// Specifies the function of the document related to the item.
	DocumentPurpose *ExternalDocumentPurpose1Code `xml:"DocPurp,omitempty"`

	// Language used for textual information in item.
	LanguageCode *LanguageCode `xml:"LangCd,omitempty"`

	// Party that issued this list of items.
	Issuer *QualifiedPartyIdentification1 `xml:"Issr,omitempty"`

	// Receiving party of this list of items.
	Recipient *QualifiedPartyIdentification1 `xml:"Rcpt,omitempty"`

	// Party that acts as buyer of the goods or services referred to by the financial item.
	Buyer *QualifiedPartyIdentification1 `xml:"Buyr,omitempty"`

	// Party that acts as seller of the goods or services referred to by the financial item.
	Seller *QualifiedPartyIdentification1 `xml:"Sellr,omitempty"`

	// Financial agent for the seller.
	SellerFinancialAgent *QualifiedPartyIdentification1 `xml:"SellrFinAgt,omitempty"`

	// Financial agent for the buyer.
	BuyerFinancialAgent *QualifiedPartyIdentification1 `xml:"BuyrFinAgt,omitempty"`

	// Reference to contract that governs the exchange of the message.
	GoverningContract []*QualifiedDocumentInformation1 `xml:"GovngCtrct,omitempty"`

	// Rules and laws governing the item.
	LegalContext *GovernanceRules2 `xml:"LglCntxt,omitempty"`

	// Currency of the item.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Defines the account debited for charges (or credited for reimbursement).
	DebitAccount *AccountIdentification4Choice `xml:"DbtAcct,omitempty"`

	// Defines the account credited for charges (or debited for reimbursement).
	CreditAccount *AccountIdentification4Choice `xml:"CdtAcct,omitempty"`

	// Identification of the geographical environment of the trade market.
	TradeMarket *TradeMarket1Choice `xml:"TradMkt,omitempty"`
}

func (f *FinancialItemParameters1) SetIdentifier(value string) {
	f.Identifier = (*Max35Text)(&value)
}

func (f *FinancialItemParameters1) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialItemParameters1) AddRelatedItem() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	f.RelatedItem = append(f.RelatedItem, newValue)
	return newValue
}

func (f *FinancialItemParameters1) SetDocumentPurpose(value string) {
	f.DocumentPurpose = (*ExternalDocumentPurpose1Code)(&value)
}

func (f *FinancialItemParameters1) SetLanguageCode(value string) {
	f.LanguageCode = (*LanguageCode)(&value)
}

func (f *FinancialItemParameters1) AddIssuer() *QualifiedPartyIdentification1 {
	f.Issuer = new(QualifiedPartyIdentification1)
	return f.Issuer
}

func (f *FinancialItemParameters1) AddRecipient() *QualifiedPartyIdentification1 {
	f.Recipient = new(QualifiedPartyIdentification1)
	return f.Recipient
}

func (f *FinancialItemParameters1) AddBuyer() *QualifiedPartyIdentification1 {
	f.Buyer = new(QualifiedPartyIdentification1)
	return f.Buyer
}

func (f *FinancialItemParameters1) AddSeller() *QualifiedPartyIdentification1 {
	f.Seller = new(QualifiedPartyIdentification1)
	return f.Seller
}

func (f *FinancialItemParameters1) AddSellerFinancialAgent() *QualifiedPartyIdentification1 {
	f.SellerFinancialAgent = new(QualifiedPartyIdentification1)
	return f.SellerFinancialAgent
}

func (f *FinancialItemParameters1) AddBuyerFinancialAgent() *QualifiedPartyIdentification1 {
	f.BuyerFinancialAgent = new(QualifiedPartyIdentification1)
	return f.BuyerFinancialAgent
}

func (f *FinancialItemParameters1) AddGoverningContract() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	f.GoverningContract = append(f.GoverningContract, newValue)
	return newValue
}

func (f *FinancialItemParameters1) AddLegalContext() *GovernanceRules2 {
	f.LegalContext = new(GovernanceRules2)
	return f.LegalContext
}

func (f *FinancialItemParameters1) SetCurrency(value string) {
	f.Currency = (*CurrencyCode)(&value)
}

func (f *FinancialItemParameters1) AddDebitAccount() *AccountIdentification4Choice {
	f.DebitAccount = new(AccountIdentification4Choice)
	return f.DebitAccount
}

func (f *FinancialItemParameters1) AddCreditAccount() *AccountIdentification4Choice {
	f.CreditAccount = new(AccountIdentification4Choice)
	return f.CreditAccount
}

func (f *FinancialItemParameters1) AddTradeMarket() *TradeMarket1Choice {
	f.TradeMarket = new(TradeMarket1Choice)
	return f.TradeMarket
}
