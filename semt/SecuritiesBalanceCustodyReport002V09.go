package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200209 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.002.002.09 Document"`
	Message *SecuritiesBalanceCustodyReport002V09 `xml:"SctiesBalCtdyRpt"`
}

func (d *Document00200209) AddMessage() *SecuritiesBalanceCustodyReport002V09 {
	d.Message = new(SecuritiesBalanceCustodyReport002V09)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesBalanceCustodyReport to an account owner to provide, at a moment in time, the quantity and identification of the financial instruments that the account servicer holds for the account owner.
//
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants, or
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer, or
// - a transfer agent acting on behalf of a fund manager or an account owner's designated agent.
//
// Usage
// The message can also include availability and the location of holdings to facilitate trading and minimise settlement issues. The message reports all information per financial instrument, that is, when a financial instrument is held at multiple places of safekeeping, the total holdings for all locations can be provided.
// The message should be sent at a frequency agreed bi-laterally between the account servicer and the account owner. The message may be provided on a trade date, contractual or settlement date basis.
// There may be one or more intermediary parties, for example, an intermediary or a concentrator between the account owner and account servicer.
//
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesBalanceCustodyReport002V09 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *model.Statement52 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	//
	AccountServicer *model.PartyIdentification111 `xml:"AcctSvcr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount33 `xml:"SfkpgAcct"`

	// Information about the party that provides services relating to financial products to investors, for example, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*model.Intermediary37 `xml:"IntrmyInf,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*model.AggregateBalanceInformation33 `xml:"BalForAcct,omitempty"`

	// Sub-account of the safekeeping or investment account.
	SubAccountDetails []*model.SubAccountIdentification46 `xml:"SubAcctDtls,omitempty"`

	// Total valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyTotalAmounts *model.TotalValueInPageAndStatement3 `xml:"AcctBaseCcyTtlAmts,omitempty"`
}

func (s *SecuritiesBalanceCustodyReport002V09) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesBalanceCustodyReport002V09) AddStatementGeneralDetails() *model.Statement52 {
	s.StatementGeneralDetails = new(model.Statement52)
	return s.StatementGeneralDetails
}

func (s *SecuritiesBalanceCustodyReport002V09) AddAccountOwner() *model.PartyIdentification109 {
	s.AccountOwner = new(model.PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesBalanceCustodyReport002V09) AddAccountServicer() *model.PartyIdentification111 {
	s.AccountServicer = new(model.PartyIdentification111)
	return s.AccountServicer
}

func (s *SecuritiesBalanceCustodyReport002V09) AddSafekeepingAccount() *model.SecuritiesAccount33 {
	s.SafekeepingAccount = new(model.SecuritiesAccount33)
	return s.SafekeepingAccount
}

func (s *SecuritiesBalanceCustodyReport002V09) AddIntermediaryInformation() *model.Intermediary37 {
	newValue := new(model.Intermediary37)
	s.IntermediaryInformation = append(s.IntermediaryInformation, newValue)
	return newValue
}

func (s *SecuritiesBalanceCustodyReport002V09) AddBalanceForAccount() *model.AggregateBalanceInformation33 {
	newValue := new(model.AggregateBalanceInformation33)
	s.BalanceForAccount = append(s.BalanceForAccount, newValue)
	return newValue
}

func (s *SecuritiesBalanceCustodyReport002V09) AddSubAccountDetails() *model.SubAccountIdentification46 {
	newValue := new(model.SubAccountIdentification46)
	s.SubAccountDetails = append(s.SubAccountDetails, newValue)
	return newValue
}

func (s *SecuritiesBalanceCustodyReport002V09) AddAccountBaseCurrencyTotalAmounts() *model.TotalValueInPageAndStatement3 {
	s.AccountBaseCurrencyTotalAmounts = new(model.TotalValueInPageAndStatement3)
	return s.AccountBaseCurrencyTotalAmounts
}
