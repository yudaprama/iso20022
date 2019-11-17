package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03900204 struct {
	XMLName xml.Name                                                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.039.002.04 Document"`
	Message *SecuritiesSettlementTransactionModificationRequestStatusAdvice002V04 `xml:"SctiesSttlmTxModReqStsAdvc"`
}

func (d *Document03900204) AddMessage() *SecuritiesSettlementTransactionModificationRequestStatusAdvice002V04 {
	d.Message = new(SecuritiesSettlementTransactionModificationRequestStatusAdvice002V04)
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
type SecuritiesSettlementTransactionModificationRequestStatusAdvice002V04 struct {

	// Reference to the unambiguous identification of the cancellation request as per the account owner.
	ModificationRequestReference *model.Identification16 `xml:"ModReqRef"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Provides unambiguous transaction identification information.
	TransactionIdentification *model.TransactionIdentifications37 `xml:"TxId,omitempty"`

	// Provides details on the processing status of the request.
	ModificationProcessingStatus *model.ModificationProcessingStatus8Choice `xml:"ModPrcgSts"`

	// Identifies the details of the transaction.
	TransactionDetails *model.TransactionDetails84 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdvice002V04) AddModificationRequestReference() *model.Identification16 {
	s.ModificationRequestReference = new(model.Identification16)
	return s.ModificationRequestReference
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdvice002V04) AddAccountOwner() *model.PartyIdentification109 {
	s.AccountOwner = new(model.PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdvice002V04) AddSafekeepingAccount() *model.SecuritiesAccount30 {
	s.SafekeepingAccount = new(model.SecuritiesAccount30)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdvice002V04) AddTransactionIdentification() *model.TransactionIdentifications37 {
	s.TransactionIdentification = new(model.TransactionIdentifications37)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdvice002V04) AddModificationProcessingStatus() *model.ModificationProcessingStatus8Choice {
	s.ModificationProcessingStatus = new(model.ModificationProcessingStatus8Choice)
	return s.ModificationProcessingStatus
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdvice002V04) AddTransactionDetails() *model.TransactionDetails84 {
	s.TransactionDetails = new(model.TransactionDetails84)
	return s.TransactionDetails
}

func (s *SecuritiesSettlementTransactionModificationRequestStatusAdvice002V04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
