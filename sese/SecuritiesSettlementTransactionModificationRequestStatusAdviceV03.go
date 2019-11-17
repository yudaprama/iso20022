package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03900103 struct {
	XMLName xml.Name                                                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.03 Document"`
	Message *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03 `xml:"SctiesSttlmTxModReqStsAdvc"`
}

func (d *Document03900103) AddMessage() *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03 {
	d.Message = new(SecuritiesSettlementTransactionModificationRequestStatusAdviceV03)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionModificationRequestStatusAdvice to an account owner to advise the status of a SecuritiesSettlementModificationRequest message previously sent by the account owner.
// The account servicer may be:
// - a central securities depository or another settlement market infrastructure managing securities settlement transactions on behalf of their participants
// - an custodian acting as an accounting and/or settlement agent.
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
type SecuritiesSettlementTransactionModificationRequestStatusAdviceV03 struct {

	// Reference to the unambiguous identification of the cancellation request as per the account owner.
	ModificationRequestReference *model.Identification1 `xml:"ModReqRef"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Provides unambiguous transaction identification information.
	TransactionIdentification *model.TransactionIdentifications25 `xml:"TxId,omitempty"`

	// Provides details on the processing status of the request.
	ModificationProcessingStatus *model.ModificationProcessingStatus4Choice `xml:"ModPrcgSts"`

	// Identifies the details of the transaction.
	TransactionDetails *model.TransactionDetails45 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03) AddModificationRequestReference() *model.Identification1 {
	s.ModificationRequestReference = new(model.Identification1)
	return s.ModificationRequestReference
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03) AddAccountOwner() *model.PartyIdentification36Choice {
	s.AccountOwner = new(model.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	s.SafekeepingAccount = new(model.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03) AddTransactionIdentification() *model.TransactionIdentifications25 {
	s.TransactionIdentification = new(model.TransactionIdentifications25)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03) AddModificationProcessingStatus() *model.ModificationProcessingStatus4Choice {
	s.ModificationProcessingStatus = new(model.ModificationProcessingStatus4Choice)
	return s.ModificationProcessingStatus
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03) AddTransactionDetails() *model.TransactionDetails45 {
	s.TransactionDetails = new(model.TransactionDetails45)
	return s.TransactionDetails
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
