package trea

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400102 struct {
	XMLName xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:trea.004.001.02 Document"`
	Message *CreateNonDeliverableForwardValuationV02 `xml:"CretNDFValtnV02"`
}

func (d *Document00400102) AddMessage() *CreateNonDeliverableForwardValuationV02 {
	d.Message = new(CreateNonDeliverableForwardValuationV02)
	return d.Message
}

// Scope
// The CreateNonDeliverableForwardValuation message is sent by a participant to a central system or to a counterparty to notify the valuation of a non deliverable trade.
// Usage
// The two trading parties will both send similar notifications to the central settlement system and the central settlement system will send notifications to both.
type CreateNonDeliverableForwardValuationV02 struct {

	// Provides identification and date of the valuation of the non deliverable trade which is created.
	TradeInformation *model.TradeAgreement1 `xml:"TradInf"`

	// Specifies the trading side of the valuation of the non deliverable trade which is created.
	TradingSideIdentification *model.TradePartyIdentification3 `xml:"TradgSdId"`

	// Specifies the counterparty of the valuation of the non deliverable trade which is created.
	CounterpartySideIdentification *model.TradePartyIdentification3 `xml:"CtrPtySdId"`

	// Specifies the amounts of the valuation of the non deliverable trade which is created.
	TradeAmounts *model.AmountsAndValueDate1 `xml:"TradAmts"`

	// Specifies the valuation information of the valuation of the non deliverable trade which is created.
	ValuationInformation *model.ValuationData2 `xml:"ValtnInf"`

	// Specifies the valuation rate of the valuation of the non deliverable trade which is created.
	ValuationRate *model.AgreedRate1 `xml:"ValtnRate"`
}

func (c *CreateNonDeliverableForwardValuationV02) AddTradeInformation() *model.TradeAgreement1 {
	c.TradeInformation = new(model.TradeAgreement1)
	return c.TradeInformation
}

func (c *CreateNonDeliverableForwardValuationV02) AddTradingSideIdentification() *model.TradePartyIdentification3 {
	c.TradingSideIdentification = new(model.TradePartyIdentification3)
	return c.TradingSideIdentification
}

func (c *CreateNonDeliverableForwardValuationV02) AddCounterpartySideIdentification() *model.TradePartyIdentification3 {
	c.CounterpartySideIdentification = new(model.TradePartyIdentification3)
	return c.CounterpartySideIdentification
}

func (c *CreateNonDeliverableForwardValuationV02) AddTradeAmounts() *model.AmountsAndValueDate1 {
	c.TradeAmounts = new(model.AmountsAndValueDate1)
	return c.TradeAmounts
}

func (c *CreateNonDeliverableForwardValuationV02) AddValuationInformation() *model.ValuationData2 {
	c.ValuationInformation = new(model.ValuationData2)
	return c.ValuationInformation
}

func (c *CreateNonDeliverableForwardValuationV02) AddValuationRate() *model.AgreedRate1 {
	c.ValuationRate = new(model.AgreedRate1)
	return c.ValuationRate
}
