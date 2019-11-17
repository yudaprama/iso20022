package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04000102 struct {
	XMLName xml.Name                                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.040.001.02 Document"`
	Message *SecuritiesSettlementTransactionCounterpartyResponseV02 `xml:"SctiesSttlmTxCtrPtyRspn"`
}

func (d *Document04000102) AddMessage() *SecuritiesSettlementTransactionCounterpartyResponseV02 {
	d.Message = new(SecuritiesSettlementTransactionCounterpartyResponseV02)
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
type SecuritiesSettlementTransactionCounterpartyResponseV02 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *model.TransactionIdentification6 `xml:"TxId"`

	// Provides the response status related to an allegement or a counterparty's instruction.
	ResponseStatus *model.ResponseStatus6Choice `xml:"RspnSts"`

	// Identifies the details of the transaction.
	TransactionDetails *model.TransactionDetails82 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionCounterpartyResponseV02) AddTransactionIdentification() *model.TransactionIdentification6 {
	s.TransactionIdentification = new(model.TransactionIdentification6)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionCounterpartyResponseV02) AddResponseStatus() *model.ResponseStatus6Choice {
	s.ResponseStatus = new(model.ResponseStatus6Choice)
	return s.ResponseStatus
}

func (s *SecuritiesSettlementTransactionCounterpartyResponseV02) AddTransactionDetails() *model.TransactionDetails82 {
	s.TransactionDetails = new(model.TransactionDetails82)
	return s.TransactionDetails
}

func (s *SecuritiesSettlementTransactionCounterpartyResponseV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
