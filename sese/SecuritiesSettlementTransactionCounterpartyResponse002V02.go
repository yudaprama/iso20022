package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04000202 struct {
	XMLName xml.Name                                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.002.02 Document"`
	Message *SecuritiesSettlementTransactionCounterpartyResponse002V02 `xml:"SctiesSttlmTxCtrPtyRspn"`
}

func (d *Document04000202) AddMessage() *SecuritiesSettlementTransactionCounterpartyResponse002V02 {
	d.Message = new(SecuritiesSettlementTransactionCounterpartyResponse002V02)
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
type SecuritiesSettlementTransactionCounterpartyResponse002V02 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *model.TransactionIdentification7 `xml:"TxId"`

	// Provides the response status related to an allegement or a counterparty's instruction.
	ResponseStatus *model.ResponseStatus8Choice `xml:"RspnSts"`

	// Identifies the details of the transaction.
	TransactionDetails *model.TransactionDetails92 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionCounterpartyResponse002V02) AddTransactionIdentification() *model.TransactionIdentification7 {
	s.TransactionIdentification = new(model.TransactionIdentification7)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionCounterpartyResponse002V02) AddResponseStatus() *model.ResponseStatus8Choice {
	s.ResponseStatus = new(model.ResponseStatus8Choice)
	return s.ResponseStatus
}

func (s *SecuritiesSettlementTransactionCounterpartyResponse002V02) AddTransactionDetails() *model.TransactionDetails92 {
	s.TransactionDetails = new(model.TransactionDetails92)
	return s.TransactionDetails
}

func (s *SecuritiesSettlementTransactionCounterpartyResponse002V02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
