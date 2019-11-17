package trea

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600102 struct {
	XMLName xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:trea.006.001.02 Document"`
	Message *CancelNonDeliverableForwardValuationV02 `xml:"CclNDFValtnV02"`
}

func (d *Document00600102) AddMessage() *CancelNonDeliverableForwardValuationV02 {
	d.Message = new(CancelNonDeliverableForwardValuationV02)
	return d.Message
}

// Scope
// The CancelNonDeliverableForwardValuation message is sent by a participant to a central system or to a counterparty to notify the cancellation of the valuation of a non deliverable trade previously confirmed by the sender.
// Usage
// The message will contain a Related Reference to link it to the previously sent notification. It may contain a reason for cancellation.
type CancelNonDeliverableForwardValuationV02 struct {

	// Provides references and date of the valuation of the non deliverable trade which is cancelled.
	TradeInformation *model.TradeAgreement2 `xml:"TradInf"`

	// Specifies the trading side of the non deliverable trade which is cancelled.
	TradingSideIdentification *model.TradePartyIdentification3 `xml:"TradgSdId,omitempty"`

	// Specifies the counterparty of the non deliverable trade which is cancelled.
	CounterpartySideIdentification *model.TradePartyIdentification3 `xml:"CtrPtySdId,omitempty"`

	// Specifies the amounts of the valuation of the non deliverable trade which is cancelled.
	TradeAmounts *model.AmountsAndValueDate1 `xml:"TradAmts,omitempty"`

	// Specifies the valuation rate of the valuation of the non deliverable trade which is cancelled.
	ValuationRate *model.AgreedRate1 `xml:"ValtnRate,omitempty"`

	// Specifies the valuation information of the valuation of the non deliverable trade which is cancelled.
	ValuationInformation *model.ValuationData2 `xml:"ValtnInf,omitempty"`
}

func (c *CancelNonDeliverableForwardValuationV02) AddTradeInformation() *model.TradeAgreement2 {
	c.TradeInformation = new(model.TradeAgreement2)
	return c.TradeInformation
}

func (c *CancelNonDeliverableForwardValuationV02) AddTradingSideIdentification() *model.TradePartyIdentification3 {
	c.TradingSideIdentification = new(model.TradePartyIdentification3)
	return c.TradingSideIdentification
}

func (c *CancelNonDeliverableForwardValuationV02) AddCounterpartySideIdentification() *model.TradePartyIdentification3 {
	c.CounterpartySideIdentification = new(model.TradePartyIdentification3)
	return c.CounterpartySideIdentification
}

func (c *CancelNonDeliverableForwardValuationV02) AddTradeAmounts() *model.AmountsAndValueDate1 {
	c.TradeAmounts = new(model.AmountsAndValueDate1)
	return c.TradeAmounts
}

func (c *CancelNonDeliverableForwardValuationV02) AddValuationRate() *model.AgreedRate1 {
	c.ValuationRate = new(model.AgreedRate1)
	return c.ValuationRate
}

func (c *CancelNonDeliverableForwardValuationV02) AddValuationInformation() *model.ValuationData2 {
	c.ValuationInformation = new(model.ValuationData2)
	return c.ValuationInformation
}
