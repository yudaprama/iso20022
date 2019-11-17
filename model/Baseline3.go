package model

// Specifies the commercial details of the underlying transaction.
type Baseline3 struct {

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
	Goods *LineItem7 `xml:"Goods"`

	// Specifies the payment terms by means of a code and a limit in time.
	PaymentTerms []*PaymentTerms1 `xml:"PmtTerms"`

	// Specifies how the underlying transaction should be settled.
	SettlementTerms *SettlementTerms2 `xml:"SttlmTerms,omitempty"`

	// Specifies the details of the payment obligation between financial institutions in this transaction.
	PaymentObligation []*PaymentObligation1 `xml:"PmtOblgtn,omitempty"`

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

func (b *Baseline3) AddSubmitterBaselineIdentification() *DocumentIdentification1 {
	b.SubmitterBaselineIdentification = new(DocumentIdentification1)
	return b.SubmitterBaselineIdentification
}

func (b *Baseline3) SetServiceCode(value string) {
	b.ServiceCode = (*TradeFinanceService2Code)(&value)
}

func (b *Baseline3) AddPurchaseOrderReference() *DocumentIdentification7 {
	b.PurchaseOrderReference = new(DocumentIdentification7)
	return b.PurchaseOrderReference
}

func (b *Baseline3) AddBuyer() *PartyIdentification26 {
	b.Buyer = new(PartyIdentification26)
	return b.Buyer
}

func (b *Baseline3) AddSeller() *PartyIdentification26 {
	b.Seller = new(PartyIdentification26)
	return b.Seller
}

func (b *Baseline3) AddBuyerBank() *BICIdentification1 {
	b.BuyerBank = new(BICIdentification1)
	return b.BuyerBank
}

func (b *Baseline3) AddSellerBank() *BICIdentification1 {
	b.SellerBank = new(BICIdentification1)
	return b.SellerBank
}

func (b *Baseline3) AddBuyerSideSubmittingBank() *BICIdentification1 {
	newValue := new(BICIdentification1)
	b.BuyerSideSubmittingBank = append(b.BuyerSideSubmittingBank, newValue)
	return newValue
}

func (b *Baseline3) AddSellerSideSubmittingBank() *BICIdentification1 {
	newValue := new(BICIdentification1)
	b.SellerSideSubmittingBank = append(b.SellerSideSubmittingBank, newValue)
	return newValue
}

func (b *Baseline3) AddBillTo() *PartyIdentification26 {
	b.BillTo = new(PartyIdentification26)
	return b.BillTo
}

func (b *Baseline3) AddShipTo() *PartyIdentification26 {
	b.ShipTo = new(PartyIdentification26)
	return b.ShipTo
}

func (b *Baseline3) AddConsignee() *PartyIdentification26 {
	b.Consignee = new(PartyIdentification26)
	return b.Consignee
}

func (b *Baseline3) AddGoods() *LineItem7 {
	b.Goods = new(LineItem7)
	return b.Goods
}

func (b *Baseline3) AddPaymentTerms() *PaymentTerms1 {
	newValue := new(PaymentTerms1)
	b.PaymentTerms = append(b.PaymentTerms, newValue)
	return newValue
}

func (b *Baseline3) AddSettlementTerms() *SettlementTerms2 {
	b.SettlementTerms = new(SettlementTerms2)
	return b.SettlementTerms
}

func (b *Baseline3) AddPaymentObligation() *PaymentObligation1 {
	newValue := new(PaymentObligation1)
	b.PaymentObligation = append(b.PaymentObligation, newValue)
	return newValue
}

func (b *Baseline3) SetLatestMatchDate(value string) {
	b.LatestMatchDate = (*ISODate)(&value)
}

func (b *Baseline3) AddCommercialDataSetRequired() *RequiredSubmission2 {
	b.CommercialDataSetRequired = new(RequiredSubmission2)
	return b.CommercialDataSetRequired
}

func (b *Baseline3) AddTransportDataSetRequired() *RequiredSubmission2 {
	b.TransportDataSetRequired = new(RequiredSubmission2)
	return b.TransportDataSetRequired
}

func (b *Baseline3) AddInsuranceDataSetRequired() *RequiredSubmission3 {
	b.InsuranceDataSetRequired = new(RequiredSubmission3)
	return b.InsuranceDataSetRequired
}

func (b *Baseline3) AddCertificateDataSetRequired() *RequiredSubmission4 {
	newValue := new(RequiredSubmission4)
	b.CertificateDataSetRequired = append(b.CertificateDataSetRequired, newValue)
	return newValue
}

func (b *Baseline3) AddOtherCertificateDataSetRequired() *RequiredSubmission5 {
	newValue := new(RequiredSubmission5)
	b.OtherCertificateDataSetRequired = append(b.OtherCertificateDataSetRequired, newValue)
	return newValue
}

func (b *Baseline3) SetIntentToPayExpected(value string) {
	b.IntentToPayExpected = (*YesNoIndicator)(&value)
}
