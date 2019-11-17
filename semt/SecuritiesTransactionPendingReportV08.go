package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01800108 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.018.001.08 Document"`
	Message *SecuritiesTransactionPendingReportV08 `xml:"SctiesTxPdgRpt"`
}

func (d *Document01800108) AddMessage() *SecuritiesTransactionPendingReportV08 {
	d.Message = new(SecuritiesTransactionPendingReportV08)
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
type SecuritiesTransactionPendingReportV08 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *model.Statement41 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Status information.
	Status []*model.StatusAndReason32 `xml:"Sts,omitempty"`

	// Details of the transactions reported.
	Transactions []*model.Transaction53 `xml:"Txs,omitempty"`
}

func (s *SecuritiesTransactionPendingReportV08) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesTransactionPendingReportV08) AddStatementGeneralDetails() *model.Statement41 {
	s.StatementGeneralDetails = new(model.Statement41)
	return s.StatementGeneralDetails
}

func (s *SecuritiesTransactionPendingReportV08) AddAccountOwner() *model.PartyIdentification98 {
	s.AccountOwner = new(model.PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesTransactionPendingReportV08) AddSafekeepingAccount() *model.SecuritiesAccount24 {
	s.SafekeepingAccount = new(model.SecuritiesAccount24)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionPendingReportV08) AddStatus() *model.StatusAndReason32 {
	newValue := new(model.StatusAndReason32)
	s.Status = append(s.Status, newValue)
	return newValue
}

func (s *SecuritiesTransactionPendingReportV08) AddTransactions() *model.Transaction53 {
	newValue := new(model.Transaction53)
	s.Transactions = append(s.Transactions, newValue)
	return newValue
}
