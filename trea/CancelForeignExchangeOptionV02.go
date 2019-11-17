package trea

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100102 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:trea.011.001.02 Document"`
	Message *CancelForeignExchangeOptionV02 `xml:"CclFXOptnV02"`
}

func (d *Document01100102) AddMessage() *CancelForeignExchangeOptionV02 {
	d.Message = new(CancelForeignExchangeOptionV02)
	return d.Message
}

// Scope
// The CancelForeignExchangeOption message is sent by a participant to a central system or to a counterparty to notify the cancellation of a foreign currency option contract.
// Usage
// The message will contain a Related Reference to link it to the previously sent notification. It may contain a reason for cancellation.
// This message is only suitable for Simple (i.e. not Barrier) Vanilla (i.e. not Binary, Digital, Notouch) Foreign Exchange Options.
type CancelForeignExchangeOptionV02 struct {

	// Provides reference and date of the foreign exchange option trade which is cancelled.
	TradeInformation *model.TradeAgreement2 `xml:"TradInf"`

	// Specifies the trading side of the currency option trade which is cancelled.
	TradingSideIdentification *model.TradePartyIdentification4 `xml:"TradgSdId,omitempty"`

	// Specifies the counterparty of the currency option trade which is cancelled.
	CounterpartySideIdentification *model.TradePartyIdentification4 `xml:"CtrPtySdId,omitempty"`

	// Specifies the parameters of the currency option which is bought by the trading side.
	Option *model.Option3 `xml:"Optn,omitempty"`
}

func (c *CancelForeignExchangeOptionV02) AddTradeInformation() *model.TradeAgreement2 {
	c.TradeInformation = new(model.TradeAgreement2)
	return c.TradeInformation
}

func (c *CancelForeignExchangeOptionV02) AddTradingSideIdentification() *model.TradePartyIdentification4 {
	c.TradingSideIdentification = new(model.TradePartyIdentification4)
	return c.TradingSideIdentification
}

func (c *CancelForeignExchangeOptionV02) AddCounterpartySideIdentification() *model.TradePartyIdentification4 {
	c.CounterpartySideIdentification = new(model.TradePartyIdentification4)
	return c.CounterpartySideIdentification
}

func (c *CancelForeignExchangeOptionV02) AddOption() *model.Option3 {
	c.Option = new(model.Option3)
	return c.Option
}
