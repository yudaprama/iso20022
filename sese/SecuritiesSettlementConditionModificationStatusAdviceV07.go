package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03100107 struct {
	XMLName xml.Name                                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.07 Document"`
	Message *SecuritiesSettlementConditionModificationStatusAdviceV07 `xml:"SctiesSttlmCondModStsAdvc"`
}

func (d *Document03100107) AddMessage() *SecuritiesSettlementConditionModificationStatusAdviceV07 {
	d.Message = new(SecuritiesSettlementConditionModificationStatusAdviceV07)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementConditionsModificationStatusAdvice to an account owner to advise the status of a modification request previously instructed by the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// A SecuritiesSettlementConditionsModificationRequest may contain requests on multiple transactions. However, one SecuritiesSettlementConditionsModificationStatusAdvice must be sent per transaction modified unless the SecuritiesSettlementConditionsModificationRequest is rejected as a whole.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementConditionModificationStatusAdviceV07 struct {

	// Identification of the SecuritiesSettlementConditionsModificationRequest.
	RequestReference *model.Identification14 `xml:"ReqRef"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Details of the request.
	RequestDetails *model.RequestDetails15 `xml:"ReqDtls,omitempty"`

	// Provides details on the processing status of the request.
	ProcessingStatus *model.ProcessingStatus50Choice `xml:"PrcgSts"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV07) AddRequestReference() *model.Identification14 {
	s.RequestReference = new(model.Identification14)
	return s.RequestReference
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV07) AddAccountOwner() *model.PartyIdentification98 {
	s.AccountOwner = new(model.PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV07) AddSafekeepingAccount() *model.SecuritiesAccount19 {
	s.SafekeepingAccount = new(model.SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV07) AddRequestDetails() *model.RequestDetails15 {
	s.RequestDetails = new(model.RequestDetails15)
	return s.RequestDetails
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV07) AddProcessingStatus() *model.ProcessingStatus50Choice {
	s.ProcessingStatus = new(model.ProcessingStatus50Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
