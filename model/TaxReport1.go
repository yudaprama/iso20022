package model

// Contains all needed party details for tax agency (sender of the TaxReport) and tax authority (receiver of the TaxReport) and the details of the reported sales transaction and calculated tax related that specific business transaction.
type TaxReport1 struct {

	// Basic report details.
	TaxReportHeader *GroupHeader69 `xml:"TaxRptHdr"`

	// Tax reporting agent, for example seller.
	// Responsible to issue tax reporting to tax authority.
	Seller *PartyIdentification72 `xml:"Sellr"`

	// Specifies the buyer of goods/service reported in this message.
	Buyer *PartyIdentification72 `xml:"Buyr,omitempty"`

	// Contains the details of the business transactions reported in the message.
	TradeSettlement *TradeSettlement2 `xml:"TradSttlm"`

	// Reserved for parties that may be required by a specific tax authority.
	OtherParty []*PartyIdentification72 `xml:"OthrPty,omitempty"`

	// Additional reference like site key or identifier.
	AdditionalInformation []*AdditionalInformation1 `xml:"AddtlInf,omitempty"`

	// Structure to deliver link to external attachment or deliver base64-coded attachment inside message.
	AdditionalReference []*DocumentGeneralInformation2 `xml:"AddtlRef,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TaxReport1) AddTaxReportHeader() *GroupHeader69 {
	t.TaxReportHeader = new(GroupHeader69)
	return t.TaxReportHeader
}

func (t *TaxReport1) AddSeller() *PartyIdentification72 {
	t.Seller = new(PartyIdentification72)
	return t.Seller
}

func (t *TaxReport1) AddBuyer() *PartyIdentification72 {
	t.Buyer = new(PartyIdentification72)
	return t.Buyer
}

func (t *TaxReport1) AddTradeSettlement() *TradeSettlement2 {
	t.TradeSettlement = new(TradeSettlement2)
	return t.TradeSettlement
}

func (t *TaxReport1) AddOtherParty() *PartyIdentification72 {
	newValue := new(PartyIdentification72)
	t.OtherParty = append(t.OtherParty, newValue)
	return newValue
}

func (t *TaxReport1) AddAdditionalInformation() *AdditionalInformation1 {
	newValue := new(AdditionalInformation1)
	t.AdditionalInformation = append(t.AdditionalInformation, newValue)
	return newValue
}

func (t *TaxReport1) AddAdditionalReference() *DocumentGeneralInformation2 {
	newValue := new(DocumentGeneralInformation2)
	t.AdditionalReference = append(t.AdditionalReference, newValue)
	return newValue
}

func (t *TaxReport1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}
