package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700101 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.001.01 Document"`
	Message *SecuritiesTransactionPostingReportV01 `xml:"SctiesTxPstngRpt"`
}

func (d *Document01700101) AddMessage() *SecuritiesTransactionPostingReportV01 {
	d.Message = new(SecuritiesTransactionPostingReportV01)
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
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesTransactionPostingReportV01 struct {

	// Information that uniquely identifies the SecuritiesTransactionPostingReport message as known by the account servicer. When the report has multiple pages, one message equals one page. Therefore, Identification uniquely identifies the page.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *model.Statement11 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*model.FinancialInstrumentDetails2 `xml:"FinInstrmDtls,omitempty"`

	// Details at sub-account level.
	SubAccountDetails []*model.SubAccountIdentification9 `xml:"SubAcctDtls,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`
}

func (s *SecuritiesTransactionPostingReportV01) AddIdentification() *model.DocumentIdentification11 {
	s.Identification = new(model.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesTransactionPostingReportV01) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesTransactionPostingReportV01) AddStatementGeneralDetails() *model.Statement11 {
	s.StatementGeneralDetails = new(model.Statement11)
	return s.StatementGeneralDetails
}

func (s *SecuritiesTransactionPostingReportV01) AddAccountOwner() *model.PartyIdentification13Choice {
	s.AccountOwner = new(model.PartyIdentification13Choice)
	return s.AccountOwner
}

func (s *SecuritiesTransactionPostingReportV01) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	s.SafekeepingAccount = new(model.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionPostingReportV01) AddFinancialInstrumentDetails() *model.FinancialInstrumentDetails2 {
	newValue := new(model.FinancialInstrumentDetails2)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}

func (s *SecuritiesTransactionPostingReportV01) AddSubAccountDetails() *model.SubAccountIdentification9 {
	newValue := new(model.SubAccountIdentification9)
	s.SubAccountDetails = append(s.SubAccountDetails, newValue)
	return newValue
}

func (s *SecuritiesTransactionPostingReportV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	s.MessageOriginator = new(model.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesTransactionPostingReportV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	s.MessageRecipient = new(model.PartyIdentification10Choice)
	return s.MessageRecipient
}
