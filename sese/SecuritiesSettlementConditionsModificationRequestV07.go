package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03000107 struct {
	XMLName xml.Name                                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.07 Document"`
	Message *SecuritiesSettlementConditionsModificationRequestV07 `xml:"SctiesSttlmCondsModReq"`
}

func (d *Document03000107) AddMessage() *SecuritiesSettlementConditionsModificationRequestV07 {
	d.Message = new(SecuritiesSettlementConditionsModificationRequestV07)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesSettlementConditionsModificationRequest to an account servicer to request the modification of a processing indicator or another non-matching information.
//
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
// In markets where this applies (eg, securities market infrastructures with no pre-settlement matching process), it is used by a party to approve, cancel or reject a transaction instructed by the counterparty.
//
// This message cannot be used to request the modification of trade or event details.
// The use of AdditionalInformation and its fields must be pre-agreed between account servicer and account owner. The fields in that sequence cannot be used to amend a trade or event detail unless authorised by country market practice.
type SecuritiesSettlementConditionsModificationRequestV07 struct {

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Details of the request.
	RequestDetails []*model.RequestDetails15 `xml:"ReqDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	AdditionalInformation []*model.AdditionalInformation13 `xml:"AddtlInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementConditionsModificationRequestV07) AddAccountOwner() *model.PartyIdentification98 {
	s.AccountOwner = new(model.PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesSettlementConditionsModificationRequestV07) AddSafekeepingAccount() *model.SecuritiesAccount19 {
	s.SafekeepingAccount = new(model.SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementConditionsModificationRequestV07) AddRequestDetails() *model.RequestDetails15 {
	newValue := new(model.RequestDetails15)
	s.RequestDetails = append(s.RequestDetails, newValue)
	return newValue
}

func (s *SecuritiesSettlementConditionsModificationRequestV07) AddAdditionalInformation() *model.AdditionalInformation13 {
	newValue := new(model.AdditionalInformation13)
	s.AdditionalInformation = append(s.AdditionalInformation, newValue)
	return newValue
}

func (s *SecuritiesSettlementConditionsModificationRequestV07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
