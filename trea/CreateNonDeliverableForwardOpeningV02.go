package trea

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100102 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:trea.001.001.02 Document"`
	Message *CreateNonDeliverableForwardOpeningV02 `xml:"CretNDFOpngV02"`
}

func (d *Document00100102) AddMessage() *CreateNonDeliverableForwardOpeningV02 {
	d.Message = new(CreateNonDeliverableForwardOpeningV02)
	return d.Message
}

// Scope
// The CreateNonDeliverableForwardOpening message is sent by a participant to a central system or to a counterparty to notify the opening of a non deliverable trade.
// Usage
// The trading parties will send similar messages to the central settlement system and the central settlement system will send notifications to both parties.
type CreateNonDeliverableForwardOpeningV02 struct {

	// Provides identification and date of the non deliverable trade which is created.
	TradeInformation *model.TradeAgreement1 `xml:"TradInf"`

	// Specifies the trading side of the non deliverable trade which is created.
	TradingSideIdentification *model.TradePartyIdentification3 `xml:"TradgSdId"`

	// Specifies the counterparty of the non deliverable trade which is created.
	CounterpartySideIdentification *model.TradePartyIdentification3 `xml:"CtrPtySdId"`

	// Specifies the amounts of the non deliverable trade which is created.
	TradeAmounts *model.AmountsAndValueDate1 `xml:"TradAmts"`

	// Specifies the rate agreed at the opening of a non deliverable trade.
	AgreedRate *model.AgreedRate1 `xml:"AgrdRate"`

	// Specifies the valuation conditions of the non deliverable trade which is created.
	ValuationConditions *model.NonDeliverableForwardValuationConditions2 `xml:"ValtnConds"`
}

func (c *CreateNonDeliverableForwardOpeningV02) AddTradeInformation() *model.TradeAgreement1 {
	c.TradeInformation = new(model.TradeAgreement1)
	return c.TradeInformation
}

func (c *CreateNonDeliverableForwardOpeningV02) AddTradingSideIdentification() *model.TradePartyIdentification3 {
	c.TradingSideIdentification = new(model.TradePartyIdentification3)
	return c.TradingSideIdentification
}

func (c *CreateNonDeliverableForwardOpeningV02) AddCounterpartySideIdentification() *model.TradePartyIdentification3 {
	c.CounterpartySideIdentification = new(model.TradePartyIdentification3)
	return c.CounterpartySideIdentification
}

func (c *CreateNonDeliverableForwardOpeningV02) AddTradeAmounts() *model.AmountsAndValueDate1 {
	c.TradeAmounts = new(model.AmountsAndValueDate1)
	return c.TradeAmounts
}

func (c *CreateNonDeliverableForwardOpeningV02) AddAgreedRate() *model.AgreedRate1 {
	c.AgreedRate = new(model.AgreedRate1)
	return c.AgreedRate
}

func (c *CreateNonDeliverableForwardOpeningV02) AddValuationConditions() *model.NonDeliverableForwardValuationConditions2 {
	c.ValuationConditions = new(model.NonDeliverableForwardValuationConditions2)
	return c.ValuationConditions
}
