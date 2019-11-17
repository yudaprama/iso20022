package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02700104 struct {
	XMLName xml.Name                                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.04 Document"`
	Message *SecuritiesTransactionCancellationRequestStatusAdviceV04 `xml:"SctiesTxCxlReqStsAdvc"`
}

func (d *Document02700104) AddMessage() *SecuritiesTransactionCancellationRequestStatusAdviceV04 {
	d.Message = new(SecuritiesTransactionCancellationRequestStatusAdviceV04)
	return d.Message
}

// Scope
// An account servicer sends an SecuritiesTransactionCancellationRequestStatusAdvice to an account owner to advise the status of a securities transaction cancellation request previously sent by the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesTransactionCancellationRequestStatusAdviceV04 struct {

	// Reference to the unambiguous identification of the cancellation request as per the account owner.
	CancellationRequestReference *model.Identification1 `xml:"CxlReqRef"`

	// Unambiguous identification of the transaction as known by the account servicer.
	TransactionIdentification *model.TransactionIdentifications17 `xml:"TxId,omitempty"`

	// Provides details on the processing status of the request.
	ProcessingStatus *model.ProcessingStatus45Choice `xml:"PrcgSts"`

	// Identifies the details of the transaction.
	TransactionDetails *model.TransactionDetails30 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV04) AddCancellationRequestReference() *model.Identification1 {
	s.CancellationRequestReference = new(model.Identification1)
	return s.CancellationRequestReference
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV04) AddTransactionIdentification() *model.TransactionIdentifications17 {
	s.TransactionIdentification = new(model.TransactionIdentifications17)
	return s.TransactionIdentification
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV04) AddProcessingStatus() *model.ProcessingStatus45Choice {
	s.ProcessingStatus = new(model.ProcessingStatus45Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV04) AddTransactionDetails() *model.TransactionDetails30 {
	s.TransactionDetails = new(model.TransactionDetails30)
	return s.TransactionDetails
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
