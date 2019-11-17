package trea

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200102 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:trea.012.001.02 Document"`
	Message *ForeignExchangeOptionNotificationV02 `xml:"FXOptnNtfctnV02"`
}

func (d *Document01200102) AddMessage() *ForeignExchangeOptionNotificationV02 {
	d.Message = new(ForeignExchangeOptionNotificationV02)
	return d.Message
}

// Scope
// The ForeignExchangeOptionNotification message is sent by a central system to a participant to provide details of a foreign exchange option trade.
// Usage
// The notification is sent by the central settlement system to the two trading parties after it has received Create, Amend or Cancel messages from both. The message may also contain information on the settlement of the trade and/or premium.
type ForeignExchangeOptionNotificationV02 struct {

	// Specifies the trading side of the currency option trade which is reported.
	TradingSideIdentification *model.TradePartyIdentification4 `xml:"TradgSdId"`

	// Specifies the counterparty of the non deliverable trade which is reported.
	CounterpartySideIdentification *model.TradePartyIdentification4 `xml:"CtrPtySdId"`

	// Provides information on the conditions of the option.
	OptionData *model.OptionData2 `xml:"OptnData"`

	// Provides information on the status of a trade in a settlement system.
	TradeStatus *model.TradeStatus1 `xml:"TradSts"`

	// Provides information on the settlement of a trade.
	SettlementData *model.SettlementData2 `xml:"SttlmData,omitempty"`
}

func (f *ForeignExchangeOptionNotificationV02) AddTradingSideIdentification() *model.TradePartyIdentification4 {
	f.TradingSideIdentification = new(model.TradePartyIdentification4)
	return f.TradingSideIdentification
}

func (f *ForeignExchangeOptionNotificationV02) AddCounterpartySideIdentification() *model.TradePartyIdentification4 {
	f.CounterpartySideIdentification = new(model.TradePartyIdentification4)
	return f.CounterpartySideIdentification
}

func (f *ForeignExchangeOptionNotificationV02) AddOptionData() *model.OptionData2 {
	f.OptionData = new(model.OptionData2)
	return f.OptionData
}

func (f *ForeignExchangeOptionNotificationV02) AddTradeStatus() *model.TradeStatus1 {
	f.TradeStatus = new(model.TradeStatus1)
	return f.TradeStatus
}

func (f *ForeignExchangeOptionNotificationV02) AddSettlementData() *model.SettlementData2 {
	f.SettlementData = new(model.SettlementData2)
	return f.SettlementData
}
