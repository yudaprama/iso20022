package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02100106 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.021.001.06 Document"`
	Message *SecuritiesStatementQueryV06 `xml:"SctiesStmtQry"`
}

func (d *Document02100106) AddMessage() *SecuritiesStatementQueryV06 {
	d.Message = new(SecuritiesStatementQueryV06)
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
type SecuritiesStatementQueryV06 struct {

	// Description of the statement requested.
	StatementRequested *model.DocumentNumber13 `xml:"StmtReqd"`

	// General information related to report.
	StatementGeneralDetails *model.Statement42 `xml:"StmtGnlDtls,omitempty"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Additional specific query criteria.
	AdditionalQueryParameters []*model.AdditionalQueryParameters11 `xml:"AddtlQryParams,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesStatementQueryV06) AddStatementRequested() *model.DocumentNumber13 {
	s.StatementRequested = new(model.DocumentNumber13)
	return s.StatementRequested
}

func (s *SecuritiesStatementQueryV06) AddStatementGeneralDetails() *model.Statement42 {
	s.StatementGeneralDetails = new(model.Statement42)
	return s.StatementGeneralDetails
}

func (s *SecuritiesStatementQueryV06) AddAccountOwner() *model.PartyIdentification98 {
	s.AccountOwner = new(model.PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesStatementQueryV06) AddSafekeepingAccount() *model.SecuritiesAccount24 {
	s.SafekeepingAccount = new(model.SecuritiesAccount24)
	return s.SafekeepingAccount
}

func (s *SecuritiesStatementQueryV06) AddAdditionalQueryParameters() *model.AdditionalQueryParameters11 {
	newValue := new(model.AdditionalQueryParameters11)
	s.AdditionalQueryParameters = append(s.AdditionalQueryParameters, newValue)
	return newValue
}

func (s *SecuritiesStatementQueryV06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
