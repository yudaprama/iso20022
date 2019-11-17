package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03000104 struct {
	XMLName xml.Name                                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.04 Document"`
	Message *SecuritiesSettlementConditionsModificationRequestV04 `xml:"SctiesSttlmCondsModReq"`
}

func (d *Document03000104) AddMessage() *SecuritiesSettlementConditionsModificationRequestV04 {
	d.Message = new(SecuritiesSettlementConditionsModificationRequestV04)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesSettlementConditionsModificationRequest to an account servicer to request the modification of a processing indicator or another non-matching information.
// The account owner/servicer relationship may be:
// - a central securities depository participant which has an account with a central securities depository.
// It could also be, if agreed in a service level agreement:
// - a global custodian which has an account with its local agent (sub-custodian), or
// - an investment management institution which manage a fund account opened at a custodian, or
// - a broker which has an account with a custodian, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
//
// This message cannot be used to request the modification of trade or event details.
// The use of AdditionalInformation and its fields must be pre-agreed between account servicer and account owner. The fields in that sequence cannot be used to amend a trade or event detail unless authorised by country market practice.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementConditionsModificationRequestV04 struct {

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Details of the request.
	RequestDetails []*model.RequestDetails11 `xml:"ReqDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	AdditionalInformation []*model.AdditionalInformation7 `xml:"AddtlInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementConditionsModificationRequestV04) AddAccountOwner() *model.PartyIdentification36Choice {
	s.AccountOwner = new(model.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesSettlementConditionsModificationRequestV04) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	s.SafekeepingAccount = new(model.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementConditionsModificationRequestV04) AddRequestDetails() *model.RequestDetails11 {
	newValue := new(model.RequestDetails11)
	s.RequestDetails = append(s.RequestDetails, newValue)
	return newValue
}

func (s *SecuritiesSettlementConditionsModificationRequestV04) AddAdditionalInformation() *model.AdditionalInformation7 {
	newValue := new(model.AdditionalInformation7)
	s.AdditionalInformation = append(s.AdditionalInformation, newValue)
	return newValue
}

func (s *SecuritiesSettlementConditionsModificationRequestV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
