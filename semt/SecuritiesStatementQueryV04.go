package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02100104 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.021.001.04 Document"`
	Message *SecuritiesStatementQueryV04 `xml:"SctiesStmtQry"`
}

func (d *Document02100104) AddMessage() *SecuritiesStatementQueryV04 {
	d.Message = new(SecuritiesStatementQueryV04)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesStatementQuery to an account servicer to request any existing securities statement.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with a local custodian, or
// - an investment management institution which manage a fund account opened at a custodian, or
// - a broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure, or
// - a central counterparty or a stock exchange or a trade matching utility which need to instruct to a central securities depository or another settlement market infrastructure.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesStatementQueryV04 struct {

	// Description of the statement requested.
	StatementRequested *model.DocumentNumber1 `xml:"StmtReqd"`

	// General information related to report.
	StatementGeneralDetails *model.Statement16 `xml:"StmtGnlDtls,omitempty"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Additional specific query criteria.
	AdditionalQueryParameters []*model.AdditionalQueryParameters7 `xml:"AddtlQryParams,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesStatementQueryV04) AddStatementRequested() *model.DocumentNumber1 {
	s.StatementRequested = new(model.DocumentNumber1)
	return s.StatementRequested
}

func (s *SecuritiesStatementQueryV04) AddStatementGeneralDetails() *model.Statement16 {
	s.StatementGeneralDetails = new(model.Statement16)
	return s.StatementGeneralDetails
}

func (s *SecuritiesStatementQueryV04) AddAccountOwner() *model.PartyIdentification36Choice {
	s.AccountOwner = new(model.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesStatementQueryV04) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	s.SafekeepingAccount = new(model.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesStatementQueryV04) AddAdditionalQueryParameters() *model.AdditionalQueryParameters7 {
	newValue := new(model.AdditionalQueryParameters7)
	s.AdditionalQueryParameters = append(s.AdditionalQueryParameters, newValue)
	return newValue
}

func (s *SecuritiesStatementQueryV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
