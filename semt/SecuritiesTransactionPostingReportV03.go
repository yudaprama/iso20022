package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700103 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.03 Document"`
	Message *SecuritiesTransactionPostingReportV03 `xml:"SctiesTxPstngRpt"`
}

func (d *Document01700103) AddMessage() *SecuritiesTransactionPostingReportV03 {
	d.Message = new(SecuritiesTransactionPostingReportV03)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesTransactionPostingReport to an account owner to provide the details of increases and decreases of holdings which occurred during a specified period, for all or selected securities in the specified safekeeping account or sub-safekeeping account which the account servicer holds for the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// Usage
// This message may be used as a trade date based or a settlement date based statement.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesTransactionPostingReportV03 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *model.Statement11 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*model.FinancialInstrumentDetails9 `xml:"FinInstrmDtls,omitempty"`

	// Details at sub-account level.
	SubAccountDetails []*model.SubAccountIdentification21 `xml:"SubAcctDtls,omitempty"`
}

func (s *SecuritiesTransactionPostingReportV03) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesTransactionPostingReportV03) AddStatementGeneralDetails() *model.Statement11 {
	s.StatementGeneralDetails = new(model.Statement11)
	return s.StatementGeneralDetails
}

func (s *SecuritiesTransactionPostingReportV03) AddAccountOwner() *model.PartyIdentification36Choice {
	s.AccountOwner = new(model.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesTransactionPostingReportV03) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	s.SafekeepingAccount = new(model.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionPostingReportV03) AddFinancialInstrumentDetails() *model.FinancialInstrumentDetails9 {
	newValue := new(model.FinancialInstrumentDetails9)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}

func (s *SecuritiesTransactionPostingReportV03) AddSubAccountDetails() *model.SubAccountIdentification21 {
	newValue := new(model.SubAccountIdentification21)
	s.SubAccountDetails = append(s.SubAccountDetails, newValue)
	return newValue
}
