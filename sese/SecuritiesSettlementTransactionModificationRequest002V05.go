package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03800205 struct {
	XMLName xml.Name                                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.038.002.05 Document"`
	Message *SecuritiesSettlementTransactionModificationRequest002V05 `xml:"SctiesSttlmTxModReq"`
}

func (d *Document03800205) AddMessage() *SecuritiesSettlementTransactionModificationRequest002V05 {
	d.Message = new(SecuritiesSettlementTransactionModificationRequest002V05)
	return d.Message
}

// Scope
// This message is sent by an account owner to an account servicer.
//
// The account owner will generally be:
// - a central securities depository participant which has an account with a central securities depository or a market infrastructure
// - an investment manager which has an account with a custodian acting as accounting and/or settlement agent.
//
// It is used to request the modification of non core business data (matching or non-matching) information in a pending or settled instruction. It can also be used for the enrichment of an incomplete transaction.
//
// Usage
// The modification must only contain the data to be modified.
// The message may also be used to:
// - re-send a message sent by the account owner to the account servicer,
// - provide a third party with a copy of a message being sent by the account owner for information,
// - re-send to a third party a copy of a message being sent by the account owner for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionModificationRequest002V05 struct {

	// Identifies the details of the transaction that is being modified.
	ModifiedTransactionDetails *model.TransactionDetails101 `xml:"ModfdTxDtls"`

	// Specifies the type of update requested.
	UpdateType []*model.UpdateType26Choice `xml:"UpdTp"`
}

func (s *SecuritiesSettlementTransactionModificationRequest002V05) AddModifiedTransactionDetails() *model.TransactionDetails101 {
	s.ModifiedTransactionDetails = new(model.TransactionDetails101)
	return s.ModifiedTransactionDetails
}

func (s *SecuritiesSettlementTransactionModificationRequest002V05) AddUpdateType() *model.UpdateType26Choice {
	newValue := new(model.UpdateType26Choice)
	s.UpdateType = append(s.UpdateType, newValue)
	return newValue
}
