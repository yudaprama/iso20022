package fxtr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400104 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.014.001.04 Document"`
	Message *ForeignExchangeTradeInstructionV04 `xml:"FXTradInstr"`
}

func (d *Document01400104) AddMessage() *ForeignExchangeTradeInstructionV04 {
	d.Message = new(ForeignExchangeTradeInstructionV04)
	return d.Message
}

// Scope
// The ForeignExchangeTradeInstruction message is sent by a participant to a central settlement system to notify the creation of the foreign exchange trade agreed by both trading parties.
// Usage
// The ForeignExchangeTradeInstruction message is sent from a participant to a central settlement system to advise of the creation of a foreign exchange trade.
type ForeignExchangeTradeInstructionV04 struct {

	// General information related to the trade.
	TradeInformation *model.TradeAgreement14 `xml:"TradInf"`

	// Party(ies) on the trading side of the trade.
	TradingSideIdentification *model.TradePartyIdentification6 `xml:"TradgSdId"`

	// Party(ies) on the counterparty side of the trade.
	CounterpartySideIdentification *model.TradePartyIdentification6 `xml:"CtrPtySdId"`

	// Amounts of the trade.
	TradeAmounts *model.AmountsAndValueDate1 `xml:"TradAmts"`

	// Exchange rate as agreed by the traders.
	AgreedRate *model.AgreedRate3 `xml:"AgrdRate"`

	// Specifies the conditions for a non deliverable opening or fixing confirmation.
	NonDeliverableForwardConditions *model.NonDeliverableForwardConditions1 `xml:"NDFConds,omitempty"`

	// Settlement instructions for the amounts received by the trading side.
	TradingSideSettlementInstructions *model.SettlementParties29 `xml:"TradgSdSttlmInstrs,omitempty"`

	// Settlement instructions for the amounts received by the counterparty.
	CounterpartySideSettlementInstructions *model.SettlementParties29 `xml:"CtrPtySdSttlmInstrs,omitempty"`

	// Specifies whether the trade is a block or an individual trade. It also contains supplementary information such as free format information, broker's identification, dealing branches and references.
	OptionalGeneralInformation *model.GeneralInformation5 `xml:"OptnlGnlInf,omitempty"`

	// Information that is to be provided to trade repositories in the context of the regulatory standards around over-the-counter (OTC) derivatives, central counterparties and trade repositories.
	RegulatoryReporting *model.RegulatoryReporting6 `xml:"RgltryRptg,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *ForeignExchangeTradeInstructionV04) AddTradeInformation() *model.TradeAgreement14 {
	f.TradeInformation = new(model.TradeAgreement14)
	return f.TradeInformation
}

func (f *ForeignExchangeTradeInstructionV04) AddTradingSideIdentification() *model.TradePartyIdentification6 {
	f.TradingSideIdentification = new(model.TradePartyIdentification6)
	return f.TradingSideIdentification
}

func (f *ForeignExchangeTradeInstructionV04) AddCounterpartySideIdentification() *model.TradePartyIdentification6 {
	f.CounterpartySideIdentification = new(model.TradePartyIdentification6)
	return f.CounterpartySideIdentification
}

func (f *ForeignExchangeTradeInstructionV04) AddTradeAmounts() *model.AmountsAndValueDate1 {
	f.TradeAmounts = new(model.AmountsAndValueDate1)
	return f.TradeAmounts
}

func (f *ForeignExchangeTradeInstructionV04) AddAgreedRate() *model.AgreedRate3 {
	f.AgreedRate = new(model.AgreedRate3)
	return f.AgreedRate
}

func (f *ForeignExchangeTradeInstructionV04) AddNonDeliverableForwardConditions() *model.NonDeliverableForwardConditions1 {
	f.NonDeliverableForwardConditions = new(model.NonDeliverableForwardConditions1)
	return f.NonDeliverableForwardConditions
}

func (f *ForeignExchangeTradeInstructionV04) AddTradingSideSettlementInstructions() *model.SettlementParties29 {
	f.TradingSideSettlementInstructions = new(model.SettlementParties29)
	return f.TradingSideSettlementInstructions
}

func (f *ForeignExchangeTradeInstructionV04) AddCounterpartySideSettlementInstructions() *model.SettlementParties29 {
	f.CounterpartySideSettlementInstructions = new(model.SettlementParties29)
	return f.CounterpartySideSettlementInstructions
}

func (f *ForeignExchangeTradeInstructionV04) AddOptionalGeneralInformation() *model.GeneralInformation5 {
	f.OptionalGeneralInformation = new(model.GeneralInformation5)
	return f.OptionalGeneralInformation
}

func (f *ForeignExchangeTradeInstructionV04) AddRegulatoryReporting() *model.RegulatoryReporting6 {
	f.RegulatoryReporting = new(model.RegulatoryReporting6)
	return f.RegulatoryReporting
}

func (f *ForeignExchangeTradeInstructionV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
