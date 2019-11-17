package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01800101 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.018.001.01 Document"`
	Message *SecuritiesTransactionPendingReportV01 `xml:"SctiesTxPdgRpt"`
}

func (d *Document01800101) AddMessage() *SecuritiesTransactionPendingReportV01 {
	d.Message = new(SecuritiesTransactionPendingReportV01)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesTransactionPendingReport to an account owner to provide, as at a specified time, the details of pending increases and decreases of holdings, for all or selected securities in a specified safekeeping account, for all or selected reasons why the transaction is pending.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// Usage
// The statement may also include future settlement or forward transactions which have become binding on the account owner.
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesTransactionPendingReportV01 struct {

	// Information that uniquely identifies the SecuritiesTransactionPendingReport message as known by the account servicer. When the report has multiple pages, one message equals one page. Therefore, Identification uniquely identifies the page.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *model.Statement14 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Status information.
	Status []*model.StatusAndReason1 `xml:"Sts,omitempty"`

	// Details of the transactions reported.
	Transactions []*model.Transaction8 `xml:"Txs,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`
}

func (s *SecuritiesTransactionPendingReportV01) AddIdentification() *model.DocumentIdentification11 {
	s.Identification = new(model.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesTransactionPendingReportV01) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesTransactionPendingReportV01) AddStatementGeneralDetails() *model.Statement14 {
	s.StatementGeneralDetails = new(model.Statement14)
	return s.StatementGeneralDetails
}

func (s *SecuritiesTransactionPendingReportV01) AddAccountOwner() *model.PartyIdentification13Choice {
	s.AccountOwner = new(model.PartyIdentification13Choice)
	return s.AccountOwner
}

func (s *SecuritiesTransactionPendingReportV01) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	s.SafekeepingAccount = new(model.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionPendingReportV01) AddStatus() *model.StatusAndReason1 {
	newValue := new(model.StatusAndReason1)
	s.Status = append(s.Status, newValue)
	return newValue
}

func (s *SecuritiesTransactionPendingReportV01) AddTransactions() *model.Transaction8 {
	newValue := new(model.Transaction8)
	s.Transactions = append(s.Transactions, newValue)
	return newValue
}

func (s *SecuritiesTransactionPendingReportV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	s.MessageOriginator = new(model.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesTransactionPendingReportV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	s.MessageRecipient = new(model.PartyIdentification10Choice)
	return s.MessageRecipient
}
