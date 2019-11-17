package fxtr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01600104 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.016.001.04 Document"`
	Message *ForeignExchangeTradeInstructionCancellationV04 `xml:"FXTradInstrCxl"`
}

func (d *Document01600104) AddMessage() *ForeignExchangeTradeInstructionCancellationV04 {
	d.Message = new(ForeignExchangeTradeInstructionCancellationV04)
	return d.Message
}

// Scope
// The ForeignExchangeTradeInstructionCancellation message is sent by a participant to a central settlement system to notify the cancellation of the foreign exchange trade previously confirmed by the sender.
// Usage
// The ForeignExchangeTradeInstructionCancellation message is sent from a participant to a central settlement system to advise of the cancellation of a previously sent notification. The "Related Reference" must be used to link it to the previous notification.
type ForeignExchangeTradeInstructionCancellationV04 struct {

	// General information related to the trade.
	TradeInformation *model.TradeAgreement15 `xml:"TradInf"`

	// Party(ies) on the trading side of the trade.
	TradingSideIdentification *model.TradePartyIdentification6 `xml:"TradgSdId"`

	// Party(ies) on the counterparty side of the trade.
	CounterpartySideIdentification *model.TradePartyIdentification6 `xml:"CtrPtySdId"`

	// Exchange rate as agreed by the traders.
	AgreedRate *model.AgreedRate3 `xml:"AgrdRate,omitempty"`

	// Specifies the conditions for a non deliverable opening or fixing confirmation.
	NonDeliverableForwardConditions *model.NonDeliverableForwardConditions1 `xml:"NDFConds,omitempty"`

	// Settlement instructions for the amounts received by the trading side.
	TradingSideSettlementInstructions *model.SettlementParties29 `xml:"TradgSdSttlmInstrs,omitempty"`

	// Settlement instructions for the amounts received by the counterparty.
	CounterpartySideSettlementInstructions *model.SettlementParties29 `xml:"CtrPtySdSttlmInstrs,omitempty"`

	// Specifies whether the trade is a block or an individual trade. It also contains supplementary information such as free format information, broker's identification, dealing branches and references.
	OptionalGeneralInformation *model.GeneralInformation5 `xml:"OptnlGnlInf,omitempty"`

	// Amounts of the trade.
	TradeAmounts *model.AmountsAndValueDate1 `xml:"TradAmts"`

	// Information that is to be provided to trade repositories in the context of the regulatory standards around over-the-counter (OTC) derivatives, central counterparties and trade repositories.
	RegulatoryReporting *model.RegulatoryReporting6 `xml:"RgltryRptg,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *ForeignExchangeTradeInstructionCancellationV04) AddTradeInformation() *model.TradeAgreement15 {
	f.TradeInformation = new(model.TradeAgreement15)
	return f.TradeInformation
}

func (f *ForeignExchangeTradeInstructionCancellationV04) AddTradingSideIdentification() *model.TradePartyIdentification6 {
	f.TradingSideIdentification = new(model.TradePartyIdentification6)
	return f.TradingSideIdentification
}

func (f *ForeignExchangeTradeInstructionCancellationV04) AddCounterpartySideIdentification() *model.TradePartyIdentification6 {
	f.CounterpartySideIdentification = new(model.TradePartyIdentification6)
	return f.CounterpartySideIdentification
}

func (f *ForeignExchangeTradeInstructionCancellationV04) AddAgreedRate() *model.AgreedRate3 {
	f.AgreedRate = new(model.AgreedRate3)
	return f.AgreedRate
}

func (f *ForeignExchangeTradeInstructionCancellationV04) AddNonDeliverableForwardConditions() *model.NonDeliverableForwardConditions1 {
	f.NonDeliverableForwardConditions = new(model.NonDeliverableForwardConditions1)
	return f.NonDeliverableForwardConditions
}

func (f *ForeignExchangeTradeInstructionCancellationV04) AddTradingSideSettlementInstructions() *model.SettlementParties29 {
	f.TradingSideSettlementInstructions = new(model.SettlementParties29)
	return f.TradingSideSettlementInstructions
}

func (f *ForeignExchangeTradeInstructionCancellationV04) AddCounterpartySideSettlementInstructions() *model.SettlementParties29 {
	f.CounterpartySideSettlementInstructions = new(model.SettlementParties29)
	return f.CounterpartySideSettlementInstructions
}

func (f *ForeignExchangeTradeInstructionCancellationV04) AddOptionalGeneralInformation() *model.GeneralInformation5 {
	f.OptionalGeneralInformation = new(model.GeneralInformation5)
	return f.OptionalGeneralInformation
}

func (f *ForeignExchangeTradeInstructionCancellationV04) AddTradeAmounts() *model.AmountsAndValueDate1 {
	f.TradeAmounts = new(model.AmountsAndValueDate1)
	return f.TradeAmounts
}

func (f *ForeignExchangeTradeInstructionCancellationV04) AddRegulatoryReporting() *model.RegulatoryReporting6 {
	f.RegulatoryReporting = new(model.RegulatoryReporting6)
	return f.RegulatoryReporting
}

func (f *ForeignExchangeTradeInstructionCancellationV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
