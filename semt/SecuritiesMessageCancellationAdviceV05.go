package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02000105 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 Document"`
	Message *SecuritiesMessageCancellationAdviceV05 `xml:"SctiesMsgCxlAdvc"`
}

func (d *Document02000105) AddMessage() *SecuritiesMessageCancellationAdviceV05 {
	d.Message = new(SecuritiesMessageCancellationAdviceV05)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesMessageCancellationAdvice to an account owner to inform of the cancellation of a securities message previously sent by an account servicer.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The previously sent message may be:
// - a securities settlement transaction confirmation
// - a report (transactions, pending transactions, allegements, accounting and custody securities balance)
// - a allegement notification (when sent by mistake or because the counterparty cancelled its instruction)
// - a portfolio transfer notification
// - an intra-position movement confirmation
// - a transaction generation notification
//
// The previously sent message cannot be a status advice message (any). If a status advice should not have been sent, a new status advice with the correct status should be sent, not a cancellation advice.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesMessageCancellationAdviceV05 struct {

	// Reference to the message advised to be cancelled by the account servicer
	Reference *model.References43Choice `xml:"Ref"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesMessageCancellationAdviceV05) AddReference() *model.References43Choice {
	s.Reference = new(model.References43Choice)
	return s.Reference
}

func (s *SecuritiesMessageCancellationAdviceV05) AddAccountOwner() *model.PartyIdentification98 {
	s.AccountOwner = new(model.PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesMessageCancellationAdviceV05) AddSafekeepingAccount() *model.SecuritiesAccount24 {
	s.SafekeepingAccount = new(model.SecuritiesAccount24)
	return s.SafekeepingAccount
}

func (s *SecuritiesMessageCancellationAdviceV05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
