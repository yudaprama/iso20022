package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400103 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.03 Document"`
	Message *IntraPositionMovementStatusAdviceV03 `xml:"IntraPosMvmntStsAdvc"`
}

func (d *Document01400103) AddMessage() *IntraPositionMovementStatusAdviceV03 {
	d.Message = new(IntraPositionMovementStatusAdviceV03)
	return d.Message
}

// Scope
// An account servicer sends a IntraPositionMovementStatusAdvice to an account owner to advise the status of a intra-position movement instruction previously sent by the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type IntraPositionMovementStatusAdviceV03 struct {

	// Unambiguous identification of a transaction as per the account owner (or the instructing party managing the account).
	TransactionIdentification *model.TransactionIdentifications15 `xml:"TxId"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *model.IntraPositionProcessingStatus3Choice `xml:"PrcgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *model.SettlementStatus9Choice `xml:"SttlmSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *model.IntraPositionDetails19 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementStatusAdviceV03) AddTransactionIdentification() *model.TransactionIdentifications15 {
	i.TransactionIdentification = new(model.TransactionIdentifications15)
	return i.TransactionIdentification
}

func (i *IntraPositionMovementStatusAdviceV03) AddProcessingStatus() *model.IntraPositionProcessingStatus3Choice {
	i.ProcessingStatus = new(model.IntraPositionProcessingStatus3Choice)
	return i.ProcessingStatus
}

func (i *IntraPositionMovementStatusAdviceV03) AddSettlementStatus() *model.SettlementStatus9Choice {
	i.SettlementStatus = new(model.SettlementStatus9Choice)
	return i.SettlementStatus
}

func (i *IntraPositionMovementStatusAdviceV03) AddTransactionDetails() *model.IntraPositionDetails19 {
	i.TransactionDetails = new(model.IntraPositionDetails19)
	return i.TransactionDetails
}

func (i *IntraPositionMovementStatusAdviceV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
