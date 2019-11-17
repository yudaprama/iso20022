package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03800104 struct {
	XMLName xml.Name                                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.038.001.04 Document"`
	Message *SecuritiesSettlementTransactionModificationRequestV04 `xml:"SctiesSttlmTxModReq"`
}

func (d *Document03800104) AddMessage() *SecuritiesSettlementTransactionModificationRequestV04 {
	d.Message = new(SecuritiesSettlementTransactionModificationRequestV04)
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
type SecuritiesSettlementTransactionModificationRequestV04 struct {

	// Identifies the details of the transaction that is being modified.
	ModifiedTransactionDetails *model.TransactionDetails76 `xml:"ModfdTxDtls"`

	// Specifies the type of update requested.
	UpdateType []*model.UpdateType14Choice `xml:"UpdTp"`
}

func (s *SecuritiesSettlementTransactionModificationRequestV04) AddModifiedTransactionDetails() *model.TransactionDetails76 {
	s.ModifiedTransactionDetails = new(model.TransactionDetails76)
	return s.ModifiedTransactionDetails
}

func (s *SecuritiesSettlementTransactionModificationRequestV04) AddUpdateType() *model.UpdateType14Choice {
	newValue := new(model.UpdateType14Choice)
	s.UpdateType = append(s.UpdateType, newValue)
	return newValue
}
