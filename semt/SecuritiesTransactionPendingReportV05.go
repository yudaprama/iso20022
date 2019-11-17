package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01800105 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.018.001.05 Document"`
	Message *SecuritiesTransactionPendingReportV05 `xml:"SctiesTxPdgRpt"`
}

func (d *Document01800105) AddMessage() *SecuritiesTransactionPendingReportV05 {
	d.Message = new(SecuritiesTransactionPendingReportV05)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesTransactionPendingReport to an account owner to provide, as at a specified time, the details of pending increases and decreases of holdings, for all or selected securities in a specified safekeeping account, for all or selected reasons why the transaction is pending.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The statement may also include future settlement or forward transactions which have become binding on the account owner.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesTransactionPendingReportV05 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *model.Statement14 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Status information.
	Status []*model.StatusAndReason18 `xml:"Sts,omitempty"`

	// Details of the transactions reported.
	Transactions []*model.Transaction34 `xml:"Txs,omitempty"`
}

func (s *SecuritiesTransactionPendingReportV05) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesTransactionPendingReportV05) AddStatementGeneralDetails() *model.Statement14 {
	s.StatementGeneralDetails = new(model.Statement14)
	return s.StatementGeneralDetails
}

func (s *SecuritiesTransactionPendingReportV05) AddAccountOwner() *model.PartyIdentification36Choice {
	s.AccountOwner = new(model.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesTransactionPendingReportV05) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	s.SafekeepingAccount = new(model.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionPendingReportV05) AddStatus() *model.StatusAndReason18 {
	newValue := new(model.StatusAndReason18)
	s.Status = append(s.Status, newValue)
	return newValue
}

func (s *SecuritiesTransactionPendingReportV05) AddTransactions() *model.Transaction34 {
	newValue := new(model.Transaction34)
	s.Transactions = append(s.Transactions, newValue)
	return newValue
}
