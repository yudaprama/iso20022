package trea

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300102 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:trea.003.001.02 Document"`
	Message *CancelNonDeliverableForwardOpeningV02 `xml:"CclNDFOpngV02"`
}

func (d *Document00300102) AddMessage() *CancelNonDeliverableForwardOpeningV02 {
	d.Message = new(CancelNonDeliverableForwardOpeningV02)
	return d.Message
}

// Scope
// The CancelNonDeliverableForwardOpening message is sent by a participant to a central system or to a counterparty to notify the cancellation of the opening of a non deliverable trade previously confirmed by the sender.
// Usage
// The message will contain a Related Reference to link it to the previously sent notification. It may contain a reason for cancellation.
type CancelNonDeliverableForwardOpeningV02 struct {

	// Provides references and date of the non deliverable trade which is cancelled.
	TradeInformation *model.TradeAgreement2 `xml:"TradInf"`

	// Specifies the trading side of the non deliverable trade which is cancelled.
	TradingSideIdentification *model.TradePartyIdentification3 `xml:"TradgSdId,omitempty"`

	// Specifies the counterparty of the non deliverable trade which is cancelled.
	CounterpartySideIdentification *model.TradePartyIdentification3 `xml:"CtrPtySdId,omitempty"`

	// Specifies the amounts of the non deliverable trade which is cancelled.
	TradeAmounts *model.AmountsAndValueDate1 `xml:"TradAmts,omitempty"`

	// Specifies the rate of the non deliverable trade which is cancelled.
	AgreedRate *model.AgreedRate1 `xml:"AgrdRate,omitempty"`

	// Specifies the valuation conditions of the non deliverable trade which is cancelled.
	ValuationConditions *model.NonDeliverableForwardValuationConditions2 `xml:"ValtnConds,omitempty"`
}

func (c *CancelNonDeliverableForwardOpeningV02) AddTradeInformation() *model.TradeAgreement2 {
	c.TradeInformation = new(model.TradeAgreement2)
	return c.TradeInformation
}

func (c *CancelNonDeliverableForwardOpeningV02) AddTradingSideIdentification() *model.TradePartyIdentification3 {
	c.TradingSideIdentification = new(model.TradePartyIdentification3)
	return c.TradingSideIdentification
}

func (c *CancelNonDeliverableForwardOpeningV02) AddCounterpartySideIdentification() *model.TradePartyIdentification3 {
	c.CounterpartySideIdentification = new(model.TradePartyIdentification3)
	return c.CounterpartySideIdentification
}

func (c *CancelNonDeliverableForwardOpeningV02) AddTradeAmounts() *model.AmountsAndValueDate1 {
	c.TradeAmounts = new(model.AmountsAndValueDate1)
	return c.TradeAmounts
}

func (c *CancelNonDeliverableForwardOpeningV02) AddAgreedRate() *model.AgreedRate1 {
	c.AgreedRate = new(model.AgreedRate1)
	return c.AgreedRate
}

func (c *CancelNonDeliverableForwardOpeningV02) AddValuationConditions() *model.NonDeliverableForwardValuationConditions2 {
	c.ValuationConditions = new(model.NonDeliverableForwardValuationConditions2)
	return c.ValuationConditions
}
