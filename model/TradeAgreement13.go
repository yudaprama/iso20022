package model

// Contractual details related to the agreement between parties.
type TradeAgreement13 struct {

	// Party that is specified as the buyer for this trade agreement.
	Buyer *TradeParty3 `xml:"Buyr"`

	// Party that is specified as the seller for this trade agreement.
	Seller *TradeParty3 `xml:"Sellr"`

	// Quotation document referenced from this trade agreement.
	QuotationDocumentIdentification *DocumentIdentification22 `xml:"QtnDocId,omitempty"`

	// Contract document referenced from this trade agreement.
	ContractDocumentIdentification *DocumentIdentification22 `xml:"CtrctDocId,omitempty"`

	// Buyer order document referenced from this trade agreement.
	BuyerOrderIdentificationDocument *DocumentIdentification22 `xml:"BuyrOrdrIdDoc,omitempty"`

	// Additional document referenced from this trade agreement.
	AdditionalReferenceDocument []*DocumentGeneralInformation2 `xml:"AddtlRefDoc,omitempty"`

	// Specifies the applicable Incoterm and associated location.
	Incoterms *Incoterms3 `xml:"Incotrms,omitempty"`
}

func (t *TradeAgreement13) AddBuyer() *TradeParty3 {
	t.Buyer = new(TradeParty3)
	return t.Buyer
}

func (t *TradeAgreement13) AddSeller() *TradeParty3 {
	t.Seller = new(TradeParty3)
	return t.Seller
}

func (t *TradeAgreement13) AddQuotationDocumentIdentification() *DocumentIdentification22 {
	t.QuotationDocumentIdentification = new(DocumentIdentification22)
	return t.QuotationDocumentIdentification
}

func (t *TradeAgreement13) AddContractDocumentIdentification() *DocumentIdentification22 {
	t.ContractDocumentIdentification = new(DocumentIdentification22)
	return t.ContractDocumentIdentification
}

func (t *TradeAgreement13) AddBuyerOrderIdentificationDocument() *DocumentIdentification22 {
	t.BuyerOrderIdentificationDocument = new(DocumentIdentification22)
	return t.BuyerOrderIdentificationDocument
}

func (t *TradeAgreement13) AddAdditionalReferenceDocument() *DocumentGeneralInformation2 {
	newValue := new(DocumentGeneralInformation2)
	t.AdditionalReferenceDocument = append(t.AdditionalReferenceDocument, newValue)
	return newValue
}

func (t *TradeAgreement13) AddIncoterms() *Incoterms3 {
	t.Incoterms = new(Incoterms3)
	return t.Incoterms
}
