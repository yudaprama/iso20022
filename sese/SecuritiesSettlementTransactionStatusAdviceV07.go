package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02400107 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.024.001.07 Document"`
	Message *SecuritiesSettlementTransactionStatusAdviceV07 `xml:"SctiesSttlmTxStsAdvc"`
}

func (d *Document02400107) AddMessage() *SecuritiesSettlementTransactionStatusAdviceV07 {
	d.Message = new(SecuritiesSettlementTransactionStatusAdviceV07)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionStatusAdvice to an account owner to advise the status of a securities settlement transaction instruction previously sent by the account owner or the status of a settlement transaction existing in the books of the servicer for the account of the owner. The status may be a processing, pending processing, internal matching, matching and/or settlement status.
// The status advice may be sent as a response to the request of the account owner or not.
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
type SecuritiesSettlementTransactionStatusAdviceV07 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *model.TransactionIdentifications31 `xml:"TxId"`

	// Link to another transaction - provided for information only.
	Linkages *model.Linkages41 `xml:"Lnkgs,omitempty"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *model.ProcessingStatus48Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of an instruction as per the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *model.MatchingStatus24Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *model.MatchingStatus24Choice `xml:"MtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *model.SettlementStatus17Choice `xml:"SttlmSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *model.TransactionDetails79 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionStatusAdviceV07) AddTransactionIdentification() *model.TransactionIdentifications31 {
	s.TransactionIdentification = new(model.TransactionIdentifications31)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionStatusAdviceV07) AddLinkages() *model.Linkages41 {
	s.Linkages = new(model.Linkages41)
	return s.Linkages
}

func (s *SecuritiesSettlementTransactionStatusAdviceV07) AddProcessingStatus() *model.ProcessingStatus48Choice {
	s.ProcessingStatus = new(model.ProcessingStatus48Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesSettlementTransactionStatusAdviceV07) AddInferredMatchingStatus() *model.MatchingStatus24Choice {
	s.InferredMatchingStatus = new(model.MatchingStatus24Choice)
	return s.InferredMatchingStatus
}

func (s *SecuritiesSettlementTransactionStatusAdviceV07) AddMatchingStatus() *model.MatchingStatus24Choice {
	s.MatchingStatus = new(model.MatchingStatus24Choice)
	return s.MatchingStatus
}

func (s *SecuritiesSettlementTransactionStatusAdviceV07) AddSettlementStatus() *model.SettlementStatus17Choice {
	s.SettlementStatus = new(model.SettlementStatus17Choice)
	return s.SettlementStatus
}

func (s *SecuritiesSettlementTransactionStatusAdviceV07) AddTransactionDetails() *model.TransactionDetails79 {
	s.TransactionDetails = new(model.TransactionDetails79)
	return s.TransactionDetails
}

func (s *SecuritiesSettlementTransactionStatusAdviceV07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
