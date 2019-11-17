package trea

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800102 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:trea.008.001.02 Document"`
	Message *StatusNotificationV02 `xml:"trea.008.001.02"`
}

func (d *Document00800102) AddMessage() *StatusNotificationV02 {
	d.Message = new(StatusNotificationV02)
	return d.Message
}

// Scope
// The StatusNotification message is sent by a central system to a participant to notify the current status of a trade in the system.
// Usage
// This message will be sent at specific times agreed upon by the central settlement system and a participant in a central settlement system.
type StatusNotificationV02 struct {

	// Provides information on the status of a trade in a system.
	TradeData *model.TradeData1 `xml:"TradData"`
}

func (s *StatusNotificationV02) AddTradeData() *model.TradeData1 {
	s.TradeData = new(model.TradeData1)
	return s.TradeData
}
