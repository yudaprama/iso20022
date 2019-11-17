package model

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader69 struct {

	// Report identification, for example invoice number or report number from point of sales system.
	Identification *Max35Text `xml:"Id"`

	// Date at which the status report was created.
	IssuedDate *ISODate `xml:"IssdDt"`

	// Specifies if the report is based on debit invoice, credit invoice, card transaction or cash transaction.
	ReportCategory *ExternalDocumentType1Code `xml:"RptCtgy"`

	// Specifies if the TaxReport is new, correction or remove.
	TaxReportPurpose *ExternalDocumentType1Code `xml:"TaxRptPurp"`

	// Original tax report identification, used for example original invoice number with credit notes.
	OriginalIdentification *Max35Text `xml:"OrgnlId,omitempty"`

	// Details of tax representative. The corporate (seller) is allowed to use a tax representative for value added tax responsibilities in case the seller is not registered in a specific value added tax registry.
	SellerTaxRepresentative *PartyIdentification116 `xml:"SellrTaxRprtv,omitempty"`

	// Details of tax representative. The corporate (buyer) is allowed to use a tax representative for value added tax responsibilities in  case the buyer is not registered in a specific value added tax registry.
	BuyerTaxRepresentative *PartyIdentification116 `xml:"BuyrTaxRprtv,omitempty"`

	// Specifies the language used in the message.
	LanguageCode *LanguageCode `xml:"LangCd,omitempty"`
}

func (g *GroupHeader69) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GroupHeader69) SetIssuedDate(value string) {
	g.IssuedDate = (*ISODate)(&value)
}

func (g *GroupHeader69) SetReportCategory(value string) {
	g.ReportCategory = (*ExternalDocumentType1Code)(&value)
}

func (g *GroupHeader69) SetTaxReportPurpose(value string) {
	g.TaxReportPurpose = (*ExternalDocumentType1Code)(&value)
}

func (g *GroupHeader69) SetOriginalIdentification(value string) {
	g.OriginalIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader69) AddSellerTaxRepresentative() *PartyIdentification116 {
	g.SellerTaxRepresentative = new(PartyIdentification116)
	return g.SellerTaxRepresentative
}

func (g *GroupHeader69) AddBuyerTaxRepresentative() *PartyIdentification116 {
	g.BuyerTaxRepresentative = new(PartyIdentification116)
	return g.BuyerTaxRepresentative
}

func (g *GroupHeader69) SetLanguageCode(value string) {
	g.LanguageCode = (*LanguageCode)(&value)
}
