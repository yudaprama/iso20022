package trea

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900102 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:trea.009.001.02 Document"`
	Message *CreateForeignExchangeOptionV02 `xml:"CretFXOptnV02"`
}

func (d *Document00900102) AddMessage() *CreateForeignExchangeOptionV02 {
	d.Message = new(CreateForeignExchangeOptionV02)
	return d.Message
}

// Scope
// The CreateForeignExchangeOption message is sent by a participant to a central system or to a counterparty to confirm a foreign currency option contract.
// Usage
// Both trading parties will send a notification to the central settlement system. The central settlement system will then send a ForeignExchangeOptionNotification to both.
// This message is only suitable for Simple (i.e. not Barrier) Vanilla (i.e. not Binary, Digital, Notouch) Foreign Exchange Options.
type CreateForeignExchangeOptionV02 struct {

	// Provides identification and date of the foreign exchange option trade which is created.
	TradeInformation *model.TradeAgreement1 `xml:"TradInf"`

	// Specifies the trading side of the currency option trade which is created.
	TradingSideIdentification *model.TradePartyIdentification4 `xml:"TradgSdId"`

	// Specifies the counterparty of the currency option trade which is created.
	CounterpartySideIdentification *model.TradePartyIdentification4 `xml:"CtrPtySdId"`

	// Specifies the parameters of the currency option which is sold by the trading side.
	Option *model.Option3 `xml:"Optn"`
}

func (c *CreateForeignExchangeOptionV02) AddTradeInformation() *model.TradeAgreement1 {
	c.TradeInformation = new(model.TradeAgreement1)
	return c.TradeInformation
}

func (c *CreateForeignExchangeOptionV02) AddTradingSideIdentification() *model.TradePartyIdentification4 {
	c.TradingSideIdentification = new(model.TradePartyIdentification4)
	return c.TradingSideIdentification
}

func (c *CreateForeignExchangeOptionV02) AddCounterpartySideIdentification() *model.TradePartyIdentification4 {
	c.CounterpartySideIdentification = new(model.TradePartyIdentification4)
	return c.CounterpartySideIdentification
}

func (c *CreateForeignExchangeOptionV02) AddOption() *model.Option3 {
	c.Option = new(model.Option3)
	return c.Option
}
