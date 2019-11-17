package fxtr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700104 struct {
	XMLName xml.Name                                             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.017.001.04 Document"`
	Message *ForeignExchangeTradeStatusAndDetailsNotificationV04 `xml:"FXTradStsAndDtlsNtfctn"`
}

func (d *Document01700104) AddMessage() *ForeignExchangeTradeStatusAndDetailsNotificationV04 {
	d.Message = new(ForeignExchangeTradeStatusAndDetailsNotificationV04)
	return d.Message
}

// Scope
// The ForeignExchangeTradeStatusAndDetails message is sent by a central system to the participant to provide notification of the status and details of a foreign exchange trade.
// Usage
// The notification is sent by a central settlement system to the two trading parties after it has received foreign exchange trade instructions from both.
type ForeignExchangeTradeStatusAndDetailsNotificationV04 struct {

	// Provides information on the status of a foreign exchange trade in the central system.
	StatusDetails *model.TradeData14 `xml:"StsDtls"`

	// General information related to the foreign exchange trade.
	TradeInformation *model.TradeAgreement12 `xml:"TradInf"`

	// Party(ies) on the trading side of the foreign exchange trade.
	TradingSideIdentification *model.TradePartyIdentification6 `xml:"TradgSdId"`

	// Party(ies) on the counterparty side of the foreign exchange trade.
	CounterpartySideIdentification *model.TradePartyIdentification6 `xml:"CtrPtySdId"`

	// Amounts of the foreign exchange trade.
	TradeAmounts *model.AmountsAndValueDate1 `xml:"TradAmts"`

	// Exchange rate as agreed by the traders.
	AgreedRate *model.AgreedRate3 `xml:"AgrdRate"`

	// Provides the opening and fixing information for an NDF trade.
	NonDeliverableForwardConditions *model.NonDeliverableForwardConditions2 `xml:"NDFConds,omitempty"`

	// Settlement instructions for the amounts received by the trading side.
	TradingSideSettlementInstructions *model.SettlementParties29 `xml:"TradgSdSttlmInstrs,omitempty"`

	// Settlement instructions for the amounts received by the counterparty.
	CounterpartySideSettlementInstructions *model.SettlementParties29 `xml:"CtrPtySdSttlmInstrs,omitempty"`

	// Additional Information about the foreign exchange trade.
	GeneralInformation *model.GeneralInformation5 `xml:"GnlInf,omitempty"`

	// Details of the split trade.
	SplitTradeInformation []*model.SplitTradeDetails3 `xml:"SpltTradInf,omitempty"`

	// Information that is to be provided to trade repositories in the context of the regulatory standards around over-the-counter (OTC) derivatives, central counterparties and trade repositories.
	RegulatoryReporting *model.RegulatoryReporting6 `xml:"RgltryRptg,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddStatusDetails() *model.TradeData14 {
	f.StatusDetails = new(model.TradeData14)
	return f.StatusDetails
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddTradeInformation() *model.TradeAgreement12 {
	f.TradeInformation = new(model.TradeAgreement12)
	return f.TradeInformation
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddTradingSideIdentification() *model.TradePartyIdentification6 {
	f.TradingSideIdentification = new(model.TradePartyIdentification6)
	return f.TradingSideIdentification
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddCounterpartySideIdentification() *model.TradePartyIdentification6 {
	f.CounterpartySideIdentification = new(model.TradePartyIdentification6)
	return f.CounterpartySideIdentification
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddTradeAmounts() *model.AmountsAndValueDate1 {
	f.TradeAmounts = new(model.AmountsAndValueDate1)
	return f.TradeAmounts
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddAgreedRate() *model.AgreedRate3 {
	f.AgreedRate = new(model.AgreedRate3)
	return f.AgreedRate
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddNonDeliverableForwardConditions() *model.NonDeliverableForwardConditions2 {
	f.NonDeliverableForwardConditions = new(model.NonDeliverableForwardConditions2)
	return f.NonDeliverableForwardConditions
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddTradingSideSettlementInstructions() *model.SettlementParties29 {
	f.TradingSideSettlementInstructions = new(model.SettlementParties29)
	return f.TradingSideSettlementInstructions
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddCounterpartySideSettlementInstructions() *model.SettlementParties29 {
	f.CounterpartySideSettlementInstructions = new(model.SettlementParties29)
	return f.CounterpartySideSettlementInstructions
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddGeneralInformation() *model.GeneralInformation5 {
	f.GeneralInformation = new(model.GeneralInformation5)
	return f.GeneralInformation
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddSplitTradeInformation() *model.SplitTradeDetails3 {
	newValue := new(model.SplitTradeDetails3)
	f.SplitTradeInformation = append(f.SplitTradeInformation, newValue)
	return newValue
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddRegulatoryReporting() *model.RegulatoryReporting6 {
	f.RegulatoryReporting = new(model.RegulatoryReporting6)
	return f.RegulatoryReporting
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
