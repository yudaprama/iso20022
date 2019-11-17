package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03900104 struct {
	XMLName xml.Name                                                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.001.04 Document"`
	Message *SecuritiesSettlementTransactionModificationRequestStatusAdviceV04 `xml:"SctiesSttlmTxModReqStsAdvc"`
}

func (d *Document03900104) AddMessage() *SecuritiesSettlementTransactionModificationRequestStatusAdviceV04 {
	d.Message = new(SecuritiesSettlementTransactionModificationRequestStatusAdviceV04)
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
type SecuritiesSettlementTransactionModificationRequestStatusAdviceV04 struct {

	// Reference to the unambiguous identification of the cancellation request as per the account owner.
	ModificationRequestReference *model.Identification14 `xml:"ModReqRef"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Provides unambiguous transaction identification information.
	TransactionIdentification *model.TransactionIdentifications33 `xml:"TxId,omitempty"`

	// Provides details on the processing status of the request.
	ModificationProcessingStatus *model.ModificationProcessingStatus7Choice `xml:"ModPrcgSts"`

	// Identifies the details of the transaction.
	TransactionDetails *model.TransactionDetails81 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV04) AddModificationRequestReference() *model.Identification14 {
	s.ModificationRequestReference = new(model.Identification14)
	return s.ModificationRequestReference
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV04) AddAccountOwner() *model.PartyIdentification98 {
	s.AccountOwner = new(model.PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV04) AddSafekeepingAccount() *model.SecuritiesAccount19 {
	s.SafekeepingAccount = new(model.SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV04) AddTransactionIdentification() *model.TransactionIdentifications33 {
	s.TransactionIdentification = new(model.TransactionIdentifications33)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV04) AddModificationProcessingStatus() *model.ModificationProcessingStatus7Choice {
	s.ModificationProcessingStatus = new(model.ModificationProcessingStatus7Choice)
	return s.ModificationProcessingStatus
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV04) AddTransactionDetails() *model.TransactionDetails81 {
	s.TransactionDetails = new(model.TransactionDetails81)
	return s.TransactionDetails
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdviceV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
