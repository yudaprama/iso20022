package secl

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100103 struct {
	XMLName xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:secl.001.001.03 Document"`
	Message *TradeLegNotificationV03 `xml:"TradLegNtfctn"`
}

func (d *Document00100103) AddMessage() *TradeLegNotificationV03 {
	d.Message = new(TradeLegNotificationV03)
	return d.Message
}

// Scope
// The TradeLegNotification message is sent by the central counterparty (CCP) to a clearing member to report the trade that has been executed by the trading platform.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The CCP reports both sides of the trade from the clearing member perspective. The CCP sends a message to the global clearing member of the seller and a message to the global clearing member of the buyer. Note: An individual clearing member only clear its own trades.
type TradeLegNotificationV03 struct {

	// Provides the identification of the clearing member (individual clearing member or general clearing member).
	ClearingMember *model.PartyIdentification35Choice `xml:"ClrMmb"`

	// Identifies the clearing member account at the CCP through which the trade must be cleared (sometimes called position account).
	ClearingAccount *model.SecuritiesAccount18 `xml:"ClrAcct"`

	// An account opened by the central counterparty in the name of the clearing member or its settlement agent within the account structure, for settlement purposes (gives information about the clearing member/its settlement agent account at the central securities depository).
	DeliveryAccount *model.SecuritiesAccount19 `xml:"DlvryAcct,omitempty"`

	// Provides details about the non clearing member identification and account.
	NonClearingMember *model.PartyIdentificationAndAccount31 `xml:"NonClrMmb,omitempty"`

	// Provides clearing details such as the settlement netting (or not) eligibility code or the clearing segment.
	ClearingDetails *model.Clearing4 `xml:"ClrDtls,omitempty"`

	// Provides details about the trade leg such as the trade date, the settlement date or the trading currency.
	TradeLegDetails *model.TradeLeg8 `xml:"TradLegDtls"`

	// Provides details about the settlement details of the trade leg such the settlement amount or the place of settlement.
	SettlementDetails *model.Settlement1 `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TradeLegNotificationV03) AddClearingMember() *model.PartyIdentification35Choice {
	t.ClearingMember = new(model.PartyIdentification35Choice)
	return t.ClearingMember
}

func (t *TradeLegNotificationV03) AddClearingAccount() *model.SecuritiesAccount18 {
	t.ClearingAccount = new(model.SecuritiesAccount18)
	return t.ClearingAccount
}

func (t *TradeLegNotificationV03) AddDeliveryAccount() *model.SecuritiesAccount19 {
	t.DeliveryAccount = new(model.SecuritiesAccount19)
	return t.DeliveryAccount
}

func (t *TradeLegNotificationV03) AddNonClearingMember() *model.PartyIdentificationAndAccount31 {
	t.NonClearingMember = new(model.PartyIdentificationAndAccount31)
	return t.NonClearingMember
}

func (t *TradeLegNotificationV03) AddClearingDetails() *model.Clearing4 {
	t.ClearingDetails = new(model.Clearing4)
	return t.ClearingDetails
}

func (t *TradeLegNotificationV03) AddTradeLegDetails() *model.TradeLeg8 {
	t.TradeLegDetails = new(model.TradeLeg8)
	return t.TradeLegDetails
}

func (t *TradeLegNotificationV03) AddSettlementDetails() *model.Settlement1 {
	t.SettlementDetails = new(model.Settlement1)
	return t.SettlementDetails
}

func (t *TradeLegNotificationV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}
