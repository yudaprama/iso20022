package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03100106 struct {
	XMLName xml.Name                                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.06 Document"`
	Message *SecuritiesSettlementConditionModificationStatusAdviceV06 `xml:"SctiesSttlmCondModStsAdvc"`
}

func (d *Document03100106) AddMessage() *SecuritiesSettlementConditionModificationStatusAdviceV06 {
	d.Message = new(SecuritiesSettlementConditionModificationStatusAdviceV06)
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
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementConditionModificationStatusAdviceV06 struct {

	// Identification of the SecuritiesSettlementConditionsModificationRequest.
	RequestReference *model.Identification1 `xml:"ReqRef"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Details of the request.
	RequestDetails *model.RequestDetails13 `xml:"ReqDtls,omitempty"`

	// Provides details on the processing status of the request.
	ProcessingStatus *model.ProcessingStatus44Choice `xml:"PrcgSts"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV06) AddRequestReference() *model.Identification1 {
	s.RequestReference = new(model.Identification1)
	return s.RequestReference
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV06) AddAccountOwner() *model.PartyIdentification36Choice {
	s.AccountOwner = new(model.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV06) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	s.SafekeepingAccount = new(model.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV06) AddRequestDetails() *model.RequestDetails13 {
	s.RequestDetails = new(model.RequestDetails13)
	return s.RequestDetails
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV06) AddProcessingStatus() *model.ProcessingStatus44Choice {
	s.ProcessingStatus = new(model.ProcessingStatus44Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
