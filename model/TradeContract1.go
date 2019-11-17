package model

// Agreement between the parties, stipulating the terms and conditions of the delivery of goods or services.
type TradeContract1 struct {

	// Contract document referenced from this trade agreement.
	ContractDocumentIdentification *DocumentIdentification22 `xml:"CtrctDocId,omitempty"`

	// Amount of the trade contract.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Party that is specified as the buyer for this trade agreement.
	Buyer []*TradeParty2 `xml:"Buyr"`

	// Party that is specified as the seller for this trade agreement.
	Seller []*TradeParty2 `xml:"Sellr"`

	// Planned final payment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt"`

	// Indicates whether the contract duration is extended or not.
	ProlongationFlag *TrueFalseIndicator `xml:"PrlngtnFlg"`

	// Start date of the trade contract.
	StartDate *ISODate `xml:"StartDt"`

	// Currency in which the trade is being settled.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Provides details on the currency exchange rate and contract.
	ExchangeRateInformation *ExchangeRate1 `xml:"XchgRateInf,omitempty"`

	// Schedule of the payments defined for the trade contract.
	PaymentSchedule *InterestPaymentDateRange1 `xml:"PmtSchdl,omitempty"`

	// Schedule of the shipment.
	ShipmentSchedule *ShipmentSchedule2Choice `xml:"ShipmntSchdl,omitempty"`

	// Documents provided as attachments to the trade contract.
	Attachment []*DocumentGeneralInformation3 `xml:"Attchmnt,omitempty"`
}

func (t *TradeContract1) AddContractDocumentIdentification() *DocumentIdentification22 {
	t.ContractDocumentIdentification = new(DocumentIdentification22)
	return t.ContractDocumentIdentification
}

func (t *TradeContract1) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TradeContract1) AddBuyer() *TradeParty2 {
	newValue := new(TradeParty2)
	t.Buyer = append(t.Buyer, newValue)
	return newValue
}

func (t *TradeContract1) AddSeller() *TradeParty2 {
	newValue := new(TradeParty2)
	t.Seller = append(t.Seller, newValue)
	return newValue
}

func (t *TradeContract1) SetMaturityDate(value string) {
	t.MaturityDate = (*ISODate)(&value)
}

func (t *TradeContract1) SetProlongationFlag(value string) {
	t.ProlongationFlag = (*TrueFalseIndicator)(&value)
}

func (t *TradeContract1) SetStartDate(value string) {
	t.StartDate = (*ISODate)(&value)
}

func (t *TradeContract1) SetSettlementCurrency(value string) {
	t.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (t *TradeContract1) AddExchangeRateInformation() *ExchangeRate1 {
	t.ExchangeRateInformation = new(ExchangeRate1)
	return t.ExchangeRateInformation
}

func (t *TradeContract1) AddPaymentSchedule() *InterestPaymentDateRange1 {
	t.PaymentSchedule = new(InterestPaymentDateRange1)
	return t.PaymentSchedule
}

func (t *TradeContract1) AddShipmentSchedule() *ShipmentSchedule2Choice {
	t.ShipmentSchedule = new(ShipmentSchedule2Choice)
	return t.ShipmentSchedule
}

func (t *TradeContract1) AddAttachment() *DocumentGeneralInformation3 {
	newValue := new(DocumentGeneralInformation3)
	t.Attachment = append(t.Attachment, newValue)
	return newValue
}
