package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03800101 struct {
	XMLName xml.Name                                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.038.001.01 Document"`
	Message *SecuritiesSettlementTransactionModificationRequestV01 `xml:"SctiesSttlmTxModReq"`
}

func (d *Document03800101) AddMessage() *SecuritiesSettlementTransactionModificationRequestV01 {
	d.Message = new(SecuritiesSettlementTransactionModificationRequestV01)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesSettlementTransactionModificationRequest to an account servicer to request the modification of non core business data (matching or non-matching) information in a pending or settled instruction. It can also be used for the enrichment of an incomplete transaction.
// The account owner will generally be:
// - a central securities depository participant which has an account with a central securities depository or a market infrastructure
// - an investment manager which has an account with a custodian acting as accounting and/or settlement agent.
// Usage
// The modification must only contain the data to be modified.
// The message may also be used to:
// - re-send a message sent by the account owner to the account servicer,
// - provide a third party with a copy of a message being sent by the account owner for information,
// - re-send to a third party a copy of a message being sent by the account owner for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionModificationRequestV01 struct {

	// Identifies the details of the transaction that is being modified.
	ModifiedTransactionDetails *model.TransactionDetails41 `xml:"ModfdTxDtls"`

	// Specifies the type of update requested.
	UpdateType []*model.UpdateType5Choice `xml:"UpdTp"`
}

func (s *SecuritiesSettlementTransactionModificationRequestV01) AddModifiedTransactionDetails() *model.TransactionDetails41 {
	s.ModifiedTransactionDetails = new(model.TransactionDetails41)
	return s.ModifiedTransactionDetails
}

func (s *SecuritiesSettlementTransactionModificationRequestV01) AddUpdateType() *model.UpdateType5Choice {
	newValue := new(model.UpdateType5Choice)
	s.UpdateType = append(s.UpdateType, newValue)
	return newValue
}
