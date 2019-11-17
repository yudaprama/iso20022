package trea

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500102 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:trea.005.001.02 Document"`
	Message *AmendNonDeliverableForwardValuationV02 `xml:"AmdNDFValtnV02"`
}

func (d *Document00500102) AddMessage() *AmendNonDeliverableForwardValuationV02 {
	d.Message = new(AmendNonDeliverableForwardValuationV02)
	return d.Message
}

// Scope
// The AmendNonDeliverableForwardValuation message is sent by a participant to a central system or to a counterparty to notify the amendment of the valuation of a non deliverable trade previously confirmed by the sender.
// Usage
// The message is sent from a participant to a central settlement system to advise of the update of a previously sent notification and it contains a "Related Reference" to link it to the previous notification.
type AmendNonDeliverableForwardValuationV02 struct {

	// Provides references and date of the valuation of the non deliverable trade which is amended.
	TradeInformation *model.TradeAgreement2 `xml:"TradInf"`

	// Specifies the trading side of the valuation of the non deliverable trade which is amended.
	TradingSideIdentification *model.TradePartyIdentification3 `xml:"TradgSdId"`

	// Specifies the counterparty of the valuation of the non deliverable trade which is amended.
	CounterpartySideIdentification *model.TradePartyIdentification3 `xml:"CtrPtySdId"`

	// Specifies the amounts of the valuation of the non deliverable trade which is amended.
	TradeAmounts *model.AmountsAndValueDate1 `xml:"TradAmts"`

	// Specifies the valuation rate of the valuation of the non deliverable trade which is amended.
	ValuationRate *model.AgreedRate1 `xml:"ValtnRate"`

	// Specifies the valuation information of the valuation of the non deliverable trade which is amended.
	ValuationInformation *model.ValuationData2 `xml:"ValtnInf"`
}

func (a *AmendNonDeliverableForwardValuationV02) AddTradeInformation() *model.TradeAgreement2 {
	a.TradeInformation = new(model.TradeAgreement2)
	return a.TradeInformation
}

func (a *AmendNonDeliverableForwardValuationV02) AddTradingSideIdentification() *model.TradePartyIdentification3 {
	a.TradingSideIdentification = new(model.TradePartyIdentification3)
	return a.TradingSideIdentification
}

func (a *AmendNonDeliverableForwardValuationV02) AddCounterpartySideIdentification() *model.TradePartyIdentification3 {
	a.CounterpartySideIdentification = new(model.TradePartyIdentification3)
	return a.CounterpartySideIdentification
}

func (a *AmendNonDeliverableForwardValuationV02) AddTradeAmounts() *model.AmountsAndValueDate1 {
	a.TradeAmounts = new(model.AmountsAndValueDate1)
	return a.TradeAmounts
}

func (a *AmendNonDeliverableForwardValuationV02) AddValuationRate() *model.AgreedRate1 {
	a.ValuationRate = new(model.AgreedRate1)
	return a.ValuationRate
}

func (a *AmendNonDeliverableForwardValuationV02) AddValuationInformation() *model.ValuationData2 {
	a.ValuationInformation = new(model.ValuationData2)
	return a.ValuationInformation
}
