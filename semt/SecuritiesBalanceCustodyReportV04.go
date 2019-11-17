package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200104 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:semt.002.001.04 Document"`
	Message *SecuritiesBalanceCustodyReportV04 `xml:"SctiesBalCtdyRpt"`
}

func (d *Document00200104) AddMessage() *SecuritiesBalanceCustodyReportV04 {
	d.Message = new(SecuritiesBalanceCustodyReportV04)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesBalanceCustodyReport to an account owner to provide, at a moment in time, the quantity and identification of the financial instruments that the account servicer holds for the account owner
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants, or
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer, or
// - a transfer agent acting on behalf of a fund manager or an account owner's designated agent.
// Usage
// The message can also include availability and the location of holdings to facilitate trading and minimise settlement issues. The message reports all information per financial instrument, that is, when a financial instrument is held at multiple places of safekeeping, the total holdings for all locations can be provided.
// The message should be sent at a frequency agreed bi-laterally between the account servicer and the account owner. The message may be provided on a trade date, contractual or settlement date basis.
// There may be one or more intermediary parties, for example, an intermediary or a concentrator between the account owner and account servicer.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesBalanceCustodyReportV04 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *model.Statement21 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	//
	AccountServicer *model.PartyIdentification49Choice `xml:"AcctSvcr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount11 `xml:"SfkpgAcct"`

	// Information about the party that provides services relating to financial products to investors, for example, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*model.Intermediary21 `xml:"IntrmyInf,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*model.AggregateBalanceInformation12 `xml:"BalForAcct,omitempty"`

	// Sub-account of the safekeeping or investment account.
	SubAccountDetails []*model.SubAccountIdentification17 `xml:"SubAcctDtls,omitempty"`

	// Total valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyTotalAmounts *model.TotalValueInPageAndStatement1 `xml:"AcctBaseCcyTtlAmts,omitempty"`
}

func (s *SecuritiesBalanceCustodyReportV04) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesBalanceCustodyReportV04) AddStatementGeneralDetails() *model.Statement21 {
	s.StatementGeneralDetails = new(model.Statement21)
	return s.StatementGeneralDetails
}

func (s *SecuritiesBalanceCustodyReportV04) AddAccountOwner() *model.PartyIdentification36Choice {
	s.AccountOwner = new(model.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesBalanceCustodyReportV04) AddAccountServicer() *model.PartyIdentification49Choice {
	s.AccountServicer = new(model.PartyIdentification49Choice)
	return s.AccountServicer
}

func (s *SecuritiesBalanceCustodyReportV04) AddSafekeepingAccount() *model.SecuritiesAccount11 {
	s.SafekeepingAccount = new(model.SecuritiesAccount11)
	return s.SafekeepingAccount
}

func (s *SecuritiesBalanceCustodyReportV04) AddIntermediaryInformation() *model.Intermediary21 {
	newValue := new(model.Intermediary21)
	s.IntermediaryInformation = append(s.IntermediaryInformation, newValue)
	return newValue
}

func (s *SecuritiesBalanceCustodyReportV04) AddBalanceForAccount() *model.AggregateBalanceInformation12 {
	newValue := new(model.AggregateBalanceInformation12)
	s.BalanceForAccount = append(s.BalanceForAccount, newValue)
	return newValue
}

func (s *SecuritiesBalanceCustodyReportV04) AddSubAccountDetails() *model.SubAccountIdentification17 {
	newValue := new(model.SubAccountIdentification17)
	s.SubAccountDetails = append(s.SubAccountDetails, newValue)
	return newValue
}

func (s *SecuritiesBalanceCustodyReportV04) AddAccountBaseCurrencyTotalAmounts() *model.TotalValueInPageAndStatement1 {
	s.AccountBaseCurrencyTotalAmounts = new(model.TotalValueInPageAndStatement1)
	return s.AccountBaseCurrencyTotalAmounts
}
