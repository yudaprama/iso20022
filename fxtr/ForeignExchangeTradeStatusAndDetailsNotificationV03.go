package fxtr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700103 struct {
	XMLName xml.Name                                             `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.017.001.03 Document"`
	Message *ForeignExchangeTradeStatusAndDetailsNotificationV03 `xml:"FXTradStsAndDtlsNtfctn"`
}

func (d *Document01700103) AddMessage() *ForeignExchangeTradeStatusAndDetailsNotificationV03 {
	d.Message = new(ForeignExchangeTradeStatusAndDetailsNotificationV03)
	return d.Message
}

// Scope
// The ForeignExchangeTradeStatusAndDetails message is sent by a central system to the participant to provide notification of the status and details of a foreign exchange trade.
// Usage
// The notification is sent by a central settlement system to the two trading parties after it has received foreign exchange trade instructions from both.
type ForeignExchangeTradeStatusAndDetailsNotificationV03 struct {

	// Provides information on the status of a foreign exchange trade in the central system.
	StatusDetails *model.TradeData9 `xml:"StsDtls"`

	// General information related to the foreign exchange trade.
	TradeInformation *model.TradeAgreement12 `xml:"TradInf"`

	// Party(ies) on the trading side of the foreign exchange trade.
	TradingSideIdentification *model.TradePartyIdentification6 `xml:"TradgSdId"`

	// Party(ies) on the counterparty side of the foreign exchange trade.
	CounterpartySideIdentification *model.TradePartyIdentification6 `xml:"CtrPtySdId"`

	// Amounts of the foreign exchange trade.
	TradeAmounts *model.AmountsAndValueDate1 `xml:"TradAmts"`

	// Exchange rate as agreed by the traders.
	AgreedRate *model.AgreedRate1 `xml:"AgrdRate"`

	// Settlement instructions for the amounts received by the trading side.
	TradingSideSettlementInstructions *model.SettlementParties29 `xml:"TradgSdSttlmInstrs,omitempty"`

	// Settlement instructions for the amounts received by the counterparty.
	CounterpartySideSettlementInstructions *model.SettlementParties29 `xml:"CtrPtySdSttlmInstrs,omitempty"`

	// Additional Information about the foreign exchange trade.
	GeneralInformation *model.GeneralInformation4 `xml:"GnlInf,omitempty"`

	// Details of the split trade.
	SplitTradeInformation []*model.SplitTradeDetails1 `xml:"SpltTradInf,omitempty"`

	// Information that is to be provided to trade repositories in the context of the regulatory standards around over-the-counter (OTC) derivatives, central counterparties and trade repositories.
	RegulatoryReporting *model.RegulatoryReporting4 `xml:"RgltryRptg,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV03) AddStatusDetails() *model.TradeData9 {
	f.StatusDetails = new(model.TradeData9)
	return f.StatusDetails
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV03) AddTradeInformation() *model.TradeAgreement12 {
	f.TradeInformation = new(model.TradeAgreement12)
	return f.TradeInformation
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV03) AddTradingSideIdentification() *model.TradePartyIdentification6 {
	f.TradingSideIdentification = new(model.TradePartyIdentification6)
	return f.TradingSideIdentification
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV03) AddCounterpartySideIdentification() *model.TradePartyIdentification6 {
	f.CounterpartySideIdentification = new(model.TradePartyIdentification6)
	return f.CounterpartySideIdentification
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV03) AddTradeAmounts() *model.AmountsAndValueDate1 {
	f.TradeAmounts = new(model.AmountsAndValueDate1)
	return f.TradeAmounts
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV03) AddAgreedRate() *model.AgreedRate1 {
	f.AgreedRate = new(model.AgreedRate1)
	return f.AgreedRate
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV03) AddTradingSideSettlementInstructions() *model.SettlementParties29 {
	f.TradingSideSettlementInstructions = new(model.SettlementParties29)
	return f.TradingSideSettlementInstructions
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV03) AddCounterpartySideSettlementInstructions() *model.SettlementParties29 {
	f.CounterpartySideSettlementInstructions = new(model.SettlementParties29)
	return f.CounterpartySideSettlementInstructions
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV03) AddGeneralInformation() *model.GeneralInformation4 {
	f.GeneralInformation = new(model.GeneralInformation4)
	return f.GeneralInformation
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV03) AddSplitTradeInformation() *model.SplitTradeDetails1 {
	newValue := new(model.SplitTradeDetails1)
	f.SplitTradeInformation = append(f.SplitTradeInformation, newValue)
	return newValue
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV03) AddRegulatoryReporting() *model.RegulatoryReporting4 {
	f.RegulatoryReporting = new(model.RegulatoryReporting4)
	return f.RegulatoryReporting
}

func (f *ForeignExchangeTradeStatusAndDetailsNotificationV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
