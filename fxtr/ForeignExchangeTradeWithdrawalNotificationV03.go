package fxtr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300103 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:fxtr.013.001.03 Document"`
	Message *ForeignExchangeTradeWithdrawalNotificationV03 `xml:"FXTradWdrwlNtfctn"`
}

func (d *Document01300103) AddMessage() *ForeignExchangeTradeWithdrawalNotificationV03 {
	d.Message = new(ForeignExchangeTradeWithdrawalNotificationV03)
	return d.Message
}

// Scope
// The ForeignExchangeTradeWithdrawalNotification message is sent by a central system to notify the withdrawal of a foreign exchange trade which was previously notified to the receiver as an alleged trade.
// Usage
// The ForeignExchangeTradeWithdrawalNotification message is used to confirm the cancellation of a previously notified trade.
type ForeignExchangeTradeWithdrawalNotificationV03 struct {

	// Identification of the present message assigned by the party issuing the message. This identification must be unique amongst all messages of same type sent by the same party.
	MessageIdentification *model.Max35Text `xml:"MsgId"`

	// Reference to the unique system identification assigned to the trade  by the central matching system.
	MatchingSystemUniqueReference *model.Max35Text `xml:"MtchgSysUnqRef"`

	// Reason for the withdrawal notification.
	WithdrawalReason *model.WithdrawalReason1 `xml:"WdrwlRsn,omitempty"`

	// To indicate the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *model.Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *ForeignExchangeTradeWithdrawalNotificationV03) SetMessageIdentification(value string) {
	f.MessageIdentification = (*model.Max35Text)(&value)
}

func (f *ForeignExchangeTradeWithdrawalNotificationV03) SetMatchingSystemUniqueReference(value string) {
	f.MatchingSystemUniqueReference = (*model.Max35Text)(&value)
}

func (f *ForeignExchangeTradeWithdrawalNotificationV03) AddWithdrawalReason() *model.WithdrawalReason1 {
	f.WithdrawalReason = new(model.WithdrawalReason1)
	return f.WithdrawalReason
}

func (f *ForeignExchangeTradeWithdrawalNotificationV03) SetSettlementSessionIdentifier(value string) {
	f.SettlementSessionIdentifier = (*model.Exact4AlphaNumericText)(&value)
}

func (f *ForeignExchangeTradeWithdrawalNotificationV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}
