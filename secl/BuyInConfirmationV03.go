package secl

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900103 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.009.001.03 Document"`
	Message *BuyInConfirmationV03 `xml:"BuyInConf"`
}

func (d *Document00900103) AddMessage() *BuyInConfirmationV03 {
	d.Message = new(BuyInConfirmationV03)
	return d.Message
}

// Scope
// The Buy In Confirmation message is sent by the central counterparty (CCP) to the clearing member to confirm the details of the transaction resulting from the buy in.
//
// The message definition is intended for use with the ISO 20022 Business Application Header.
//
// Usage
// The Buy In Confirmation message is sent by the central counterparty (CCP) to confirm the details of the buy in transaction.
type BuyInConfirmationV03 struct {

	// Unambiguous identification of the transaction as known by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId,omitempty"`

	// Provides the identification of the clearing member (individual clearing member or general clearing member).
	ClearingMember *model.PartyIdentification35Choice `xml:"ClrMmb"`

	// Provides the buy-in details.
	BuyInDetails *model.BuyIn2 `xml:"BuyInDtls"`

	// Provides details about the original settlement obligation that did not settle and for which the buy in process will be launched.
	OriginalSettlementObligation *model.SettlementObligation7 `xml:"OrgnlSttlmOblgtn,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (b *BuyInConfirmationV03) SetTransactionIdentification(value string) {
	b.TransactionIdentification = (*model.Max35Text)(&value)
}

func (b *BuyInConfirmationV03) AddClearingMember() *model.PartyIdentification35Choice {
	b.ClearingMember = new(model.PartyIdentification35Choice)
	return b.ClearingMember
}

func (b *BuyInConfirmationV03) AddBuyInDetails() *model.BuyIn2 {
	b.BuyInDetails = new(model.BuyIn2)
	return b.BuyInDetails
}

func (b *BuyInConfirmationV03) AddOriginalSettlementObligation() *model.SettlementObligation7 {
	b.OriginalSettlementObligation = new(model.SettlementObligation7)
	return b.OriginalSettlementObligation
}

func (b *BuyInConfirmationV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	b.SupplementaryData = append(b.SupplementaryData, newValue)
	return newValue
}
