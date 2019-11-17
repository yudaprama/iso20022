package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04000101 struct {
	XMLName xml.Name                                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.01 Document"`
	Message *SecuritiesSettlementTransactionCounterpartyResponseV01 `xml:"SctiesSttlmTxCtrPtyRspn"`
}

func (d *Document04000101) AddMessage() *SecuritiesSettlementTransactionCounterpartyResponseV01 {
	d.Message = new(SecuritiesSettlementTransactionCounterpartyResponseV01)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesSettlementTransactionCounterpartyResponse to advise the account servicer that:
// - the allegement received is either rejected (that is counterparty's transaction is unknown) or accepted (i.e. either the allegement was passed to the client or the transaction is know with or without mismatches)
// - the modification or cancellation request sent by the counterparty for a matched transaction is affirmed or not. The account servicer will therefore proceed or not with the counterparty's request to modify or cancel the transaction.
// The account servicer may be a central securities depository or another settlement market infrastructure acting on behalf of their participants
// The account owner may be:
// - a central securities depository participant which has an account with a central securities depository or a market infrastructure
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message sent by the account owner to the account servicer,
// - provide a third party with a copy of a message being sent by the account owner for information,
// - re-send to a third party a copy of a message being sent by the account owner for information
// using the relevant elements in the Business Application Header.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionCounterpartyResponseV01 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *model.TransactionIdentification2 `xml:"TxId"`

	// Provides the response status related to an allegement or a counterparty's instruction.
	ResponseStatus *model.ResponseStatus3Choice `xml:"RspnSts"`

	// Identifies the details of the transaction.
	TransactionDetails *model.TransactionDetails40 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionCounterpartyResponseV01) AddTransactionIdentification() *model.TransactionIdentification2 {
	s.TransactionIdentification = new(model.TransactionIdentification2)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionCounterpartyResponseV01) AddResponseStatus() *model.ResponseStatus3Choice {
	s.ResponseStatus = new(model.ResponseStatus3Choice)
	return s.ResponseStatus
}

func (s *SecuritiesSettlementTransactionCounterpartyResponseV01) AddTransactionDetails() *model.TransactionDetails40 {
	s.TransactionDetails = new(model.TransactionDetails40)
	return s.TransactionDetails
}

func (s *SecuritiesSettlementTransactionCounterpartyResponseV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
