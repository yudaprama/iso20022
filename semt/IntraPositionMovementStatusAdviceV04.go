package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400104 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.04 Document"`
	Message *IntraPositionMovementStatusAdviceV04 `xml:"IntraPosMvmntStsAdvc"`
}

func (d *Document01400104) AddMessage() *IntraPositionMovementStatusAdviceV04 {
	d.Message = new(IntraPositionMovementStatusAdviceV04)
	return d.Message
}

// Scope
// An account servicer sends a IntraPositionMovementStatusAdvice to an account owner to advise the status of a intra-position movement instruction previously sent by the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
type IntraPositionMovementStatusAdviceV04 struct {

	// Unambiguous identification of a transaction as per the account owner (or the instructing party managing the account).
	TransactionIdentification *model.TransactionIdentifications29 `xml:"TxId"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *model.IntraPositionProcessingStatus5Choice `xml:"PrcgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *model.SettlementStatus16Choice `xml:"SttlmSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *model.IntraPositionDetails31 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementStatusAdviceV04) AddTransactionIdentification() *model.TransactionIdentifications29 {
	i.TransactionIdentification = new(model.TransactionIdentifications29)
	return i.TransactionIdentification
}

func (i *IntraPositionMovementStatusAdviceV04) AddProcessingStatus() *model.IntraPositionProcessingStatus5Choice {
	i.ProcessingStatus = new(model.IntraPositionProcessingStatus5Choice)
	return i.ProcessingStatus
}

func (i *IntraPositionMovementStatusAdviceV04) AddSettlementStatus() *model.SettlementStatus16Choice {
	i.SettlementStatus = new(model.SettlementStatus16Choice)
	return i.SettlementStatus
}

func (i *IntraPositionMovementStatusAdviceV04) AddTransactionDetails() *model.IntraPositionDetails31 {
	i.TransactionDetails = new(model.IntraPositionDetails31)
	return i.TransactionDetails
}

func (i *IntraPositionMovementStatusAdviceV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
