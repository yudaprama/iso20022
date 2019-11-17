package trea

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200102 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:trea.002.001.02 Document"`
	Message *AmendNonDeliverableForwardOpeningV02 `xml:"AmdNDFOpngV02"`
}

func (d *Document00200102) AddMessage() *AmendNonDeliverableForwardOpeningV02 {
	d.Message = new(AmendNonDeliverableForwardOpeningV02)
	return d.Message
}

// Scope
// The AmendNonDeliverableForwardOpening message is sent by a participant to a central system or to a counterparty to notify the amendment of the opening of a non deliverable trade previously confirmed by the sender.
// Usage
// The message is sent from a participant to a central settlement system to advise of the update of a previously sent notification and it contains a "Related Reference" to link it to the previous notification.
type AmendNonDeliverableForwardOpeningV02 struct {

	// Provides references and date of the non deliverable trade which is amended.
	TradeInformation *model.TradeAgreement2 `xml:"TradInf"`

	// Specifies the trading side of the non deliverable trade which is amended.
	TradingSideIdentification *model.TradePartyIdentification3 `xml:"TradgSdId"`

	// Specifies the counterparty of the non deliverable trade which is amended.
	CounterpartySideIdentification *model.TradePartyIdentification3 `xml:"CtrPtySdId"`

	// Specifies the amounts of the non deliverable trade which is amended.
	TradeAmounts *model.AmountsAndValueDate1 `xml:"TradAmts"`

	// Specifies the rate of the non deliverable trade which is amended.
	AgreedRate *model.AgreedRate1 `xml:"AgrdRate"`

	// Specifies the valuation conditions of the non deliverable trade which is amended.
	ValuationConditions *model.NonDeliverableForwardValuationConditions2 `xml:"ValtnConds"`
}

func (a *AmendNonDeliverableForwardOpeningV02) AddTradeInformation() *model.TradeAgreement2 {
	a.TradeInformation = new(model.TradeAgreement2)
	return a.TradeInformation
}

func (a *AmendNonDeliverableForwardOpeningV02) AddTradingSideIdentification() *model.TradePartyIdentification3 {
	a.TradingSideIdentification = new(model.TradePartyIdentification3)
	return a.TradingSideIdentification
}

func (a *AmendNonDeliverableForwardOpeningV02) AddCounterpartySideIdentification() *model.TradePartyIdentification3 {
	a.CounterpartySideIdentification = new(model.TradePartyIdentification3)
	return a.CounterpartySideIdentification
}

func (a *AmendNonDeliverableForwardOpeningV02) AddTradeAmounts() *model.AmountsAndValueDate1 {
	a.TradeAmounts = new(model.AmountsAndValueDate1)
	return a.TradeAmounts
}

func (a *AmendNonDeliverableForwardOpeningV02) AddAgreedRate() *model.AgreedRate1 {
	a.AgreedRate = new(model.AgreedRate1)
	return a.AgreedRate
}

func (a *AmendNonDeliverableForwardOpeningV02) AddValuationConditions() *model.NonDeliverableForwardValuationConditions2 {
	a.ValuationConditions = new(model.NonDeliverableForwardValuationConditions2)
	return a.ValuationConditions
}
