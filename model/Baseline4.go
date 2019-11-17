package model

// Specifies the commercial details of the underlying transaction.
type Baseline4 struct {

	// Identifies the baseline provided by the submitter.
	SubmitterBaselineIdentification *DocumentIdentification1 `xml:"SubmitrBaselnId"`

	// Identifies the service requested by the submitter by means of a code.
	ServiceCode *TradeFinanceService2Code `xml:"SvcCd"`

	// Reference to the purchase order of the underlying transaction.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Party that buys goods or services, or a financial instrument.
	Buyer *PartyIdentification26 `xml:"Buyr"`

	// Party that sells goods or services, or a financial instrument.
	Seller *PartyIdentification26 `xml:"Sellr"`

	// Financial institution of the buyer, uniquely identified by its BIC.
	BuyerBank *BICIdentification1 `xml:"BuyrBk"`

	// Financial institution of the seller, uniquely identified by its BIC.
	SellerBank *BICIdentification1 `xml:"SellrBk"`

	// Financial institution on the buyer's side, uniquely identified by its BIC. As part of the transaction, it may submit data sets.
	BuyerSideSubmittingBank []*BICIdentification1 `xml:"BuyrSdSubmitgBk,omitempty"`

	// Financial institution on the seller's side, uniquely identified by its BIC. As part of the transaction, it may submit data sets.
	SellerSideSubmittingBank []*BICIdentification1 `xml:"SellrSdSubmitgBk,omitempty"`

	// Party to be invoiced for the purchase.
	BillTo *PartyIdentification26 `xml:"BllTo,omitempty"`

	// Party to whom the goods must be delivered in the end.
	ShipTo *PartyIdentification26 `xml:"ShipTo,omitempty"`

	// Party to whom the goods must be delivered.
	Consignee *PartyIdentification26 `xml:"Consgn,omitempty"`

	// Goods or services that are part of a commercial trade agreement.
	Goods *LineItem11 `xml:"Goods"`

	// Specifies the payment terms by means of a code and a limit in time.
	PaymentTerms []*PaymentTerms5 `xml:"PmtTerms"`

	// Specifies how the underlying transaction should be settled.
	SettlementTerms *SettlementTerms3 `xml:"SttlmTerms,omitempty"`

	// Specifies the details of the payment obligation between financial institutions in this transaction.
	PaymentObligation []*PaymentObligation2 `xml:"PmtOblgtn,omitempty"`

	// Specifies the latest date on which a data set must be matched with a baseline.
	LatestMatchDate *ISODate `xml:"LatstMtchDt,omitempty"`

	// Specifies that a commercial data set is required for each shipment part of the transaction.
	CommercialDataSetRequired *RequiredSubmission2 `xml:"ComrclDataSetReqrd"`

	// Specifies that a transport data set is required for each shipment part of the transaction.
	TransportDataSetRequired *RequiredSubmission2 `xml:"TrnsprtDataSetReqrd,omitempty"`

	// Specifies that an insurance data set is required for each shipment part of the transaction.
	InsuranceDataSetRequired *RequiredSubmission3 `xml:"InsrncDataSetReqrd,omitempty"`

	// Specifies that a certificate data set is required for each shipment part of the transaction.
	CertificateDataSetRequired []*RequiredSubmission4 `xml:"CertDataSetReqrd,omitempty"`

	// Specifies that another type of certificate data set is required for each shipment part of the transaction.
	OtherCertificateDataSetRequired []*RequiredSubmission5 `xml:"OthrCertDataSetReqrd,omitempty"`

	// Specifies that IntentToPayNotice message(s) are expected as part of this transaction.
	IntentToPayExpected *YesNoIndicator `xml:"InttToPayXpctd"`
}

func (b *Baseline4) AddSubmitterBaselineIdentification() *DocumentIdentification1 {
	b.SubmitterBaselineIdentification = new(DocumentIdentification1)
	return b.SubmitterBaselineIdentification
}

func (b *Baseline4) SetServiceCode(value string) {
	b.ServiceCode = (*TradeFinanceService2Code)(&value)
}

func (b *Baseline4) AddPurchaseOrderReference() *DocumentIdentification7 {
	b.PurchaseOrderReference = new(DocumentIdentification7)
	return b.PurchaseOrderReference
}

func (b *Baseline4) AddBuyer() *PartyIdentification26 {
	b.Buyer = new(PartyIdentification26)
	return b.Buyer
}

func (b *Baseline4) AddSeller() *PartyIdentification26 {
	b.Seller = new(PartyIdentification26)
	return b.Seller
}

func (b *Baseline4) AddBuyerBank() *BICIdentification1 {
	b.BuyerBank = new(BICIdentification1)
	return b.BuyerBank
}

func (b *Baseline4) AddSellerBank() *BICIdentification1 {
	b.SellerBank = new(BICIdentification1)
	return b.SellerBank
}

func (b *Baseline4) AddBuyerSideSubmittingBank() *BICIdentification1 {
	newValue := new(BICIdentification1)
	b.BuyerSideSubmittingBank = append(b.BuyerSideSubmittingBank, newValue)
	return newValue
}

func (b *Baseline4) AddSellerSideSubmittingBank() *BICIdentification1 {
	newValue := new(BICIdentification1)
	b.SellerSideSubmittingBank = append(b.SellerSideSubmittingBank, newValue)
	return newValue
}

func (b *Baseline4) AddBillTo() *PartyIdentification26 {
	b.BillTo = new(PartyIdentification26)
	return b.BillTo
}

func (b *Baseline4) AddShipTo() *PartyIdentification26 {
	b.ShipTo = new(PartyIdentification26)
	return b.ShipTo
}

func (b *Baseline4) AddConsignee() *PartyIdentification26 {
	b.Consignee = new(PartyIdentification26)
	return b.Consignee
}

func (b *Baseline4) AddGoods() *LineItem11 {
	b.Goods = new(LineItem11)
	return b.Goods
}

func (b *Baseline4) AddPaymentTerms() *PaymentTerms5 {
	newValue := new(PaymentTerms5)
	b.PaymentTerms = append(b.PaymentTerms, newValue)
	return newValue
}

func (b *Baseline4) AddSettlementTerms() *SettlementTerms3 {
	b.SettlementTerms = new(SettlementTerms3)
	return b.SettlementTerms
}

func (b *Baseline4) AddPaymentObligation() *PaymentObligation2 {
	newValue := new(PaymentObligation2)
	b.PaymentObligation = append(b.PaymentObligation, newValue)
	return newValue
}

func (b *Baseline4) SetLatestMatchDate(value string) {
	b.LatestMatchDate = (*ISODate)(&value)
}

func (b *Baseline4) AddCommercialDataSetRequired() *RequiredSubmission2 {
	b.CommercialDataSetRequired = new(RequiredSubmission2)
	return b.CommercialDataSetRequired
}

func (b *Baseline4) AddTransportDataSetRequired() *RequiredSubmission2 {
	b.TransportDataSetRequired = new(RequiredSubmission2)
	return b.TransportDataSetRequired
}

func (b *Baseline4) AddInsuranceDataSetRequired() *RequiredSubmission3 {
	b.InsuranceDataSetRequired = new(RequiredSubmission3)
	return b.InsuranceDataSetRequired
}

func (b *Baseline4) AddCertificateDataSetRequired() *RequiredSubmission4 {
	newValue := new(RequiredSubmission4)
	b.CertificateDataSetRequired = append(b.CertificateDataSetRequired, newValue)
	return newValue
}

func (b *Baseline4) AddOtherCertificateDataSetRequired() *RequiredSubmission5 {
	newValue := new(RequiredSubmission5)
	b.OtherCertificateDataSetRequired = append(b.OtherCertificateDataSetRequired, newValue)
	return newValue
}

func (b *Baseline4) SetIntentToPayExpected(value string) {
	b.IntentToPayExpected = (*YesNoIndicator)(&value)
}
