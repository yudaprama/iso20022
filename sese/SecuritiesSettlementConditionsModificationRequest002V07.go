package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03000207 struct {
	XMLName xml.Name                                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.002.07 Document"`
	Message *SecuritiesSettlementConditionsModificationRequest002V07 `xml:"SctiesSttlmCondsModReq"`
}

func (d *Document03000207) AddMessage() *SecuritiesSettlementConditionsModificationRequest002V07 {
	d.Message = new(SecuritiesSettlementConditionsModificationRequest002V07)
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
type SecuritiesSettlementConditionsModificationRequest002V07 struct {

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Details of the request.
	RequestDetails []*model.RequestDetails16 `xml:"ReqDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	AdditionalInformation []*model.AdditionalInformation14 `xml:"AddtlInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementConditionsModificationRequest002V07) AddAccountOwner() *model.PartyIdentification119 {
	s.AccountOwner = new(model.PartyIdentification119)
	return s.AccountOwner
}

func (s *SecuritiesSettlementConditionsModificationRequest002V07) AddSafekeepingAccount() *model.SecuritiesAccount30 {
	s.SafekeepingAccount = new(model.SecuritiesAccount30)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementConditionsModificationRequest002V07) AddRequestDetails() *model.RequestDetails16 {
	newValue := new(model.RequestDetails16)
	s.RequestDetails = append(s.RequestDetails, newValue)
	return newValue
}

func (s *SecuritiesSettlementConditionsModificationRequest002V07) AddAdditionalInformation() *model.AdditionalInformation14 {
	newValue := new(model.AdditionalInformation14)
	s.AdditionalInformation = append(s.AdditionalInformation, newValue)
	return newValue
}

func (s *SecuritiesSettlementConditionsModificationRequest002V07) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
