package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03800204 struct {
	XMLName xml.Name                                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.038.002.04 Document"`
	Message *SecuritiesSettlementTransactionModificationRequest002V04 `xml:"SctiesSttlmTxModReq"`
}

func (d *Document03800204) AddMessage() *SecuritiesSettlementTransactionModificationRequest002V04 {
	d.Message = new(SecuritiesSettlementTransactionModificationRequest002V04)
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
type SecuritiesSettlementTransactionModificationRequest002V04 struct {

	// Identifies the details of the transaction that is being modified.
	ModifiedTransactionDetails *model.TransactionDetails85 `xml:"ModfdTxDtls"`

	// Specifies the type of update requested.
	UpdateType []*model.UpdateType22Choice `xml:"UpdTp"`
}

func (s *SecuritiesSettlementTransactionModificationRequest002V04) AddModifiedTransactionDetails() *model.TransactionDetails85 {
	s.ModifiedTransactionDetails = new(model.TransactionDetails85)
	return s.ModifiedTransactionDetails
}

func (s *SecuritiesSettlementTransactionModificationRequest002V04) AddUpdateType() *model.UpdateType22Choice {
	newValue := new(model.UpdateType22Choice)
	s.UpdateType = append(s.UpdateType, newValue)
	return newValue
}
