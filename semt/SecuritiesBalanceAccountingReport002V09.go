package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300209 struct {
	XMLName xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.003.002.09 Document"`
	Message *SecuritiesBalanceAccountingReport002V09 `xml:"SctiesBalAcctgRpt"`
}

func (d *Document00300209) AddMessage() *SecuritiesBalanceAccountingReport002V09 {
	d.Message = new(SecuritiesBalanceAccountingReport002V09)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesBalanceAccountingReport to an account owner to provide, at a moment in time, valuations of the portfolio together with details of each financial instrument holding.
// The account servicer/owner relationship may be:
// - an accounting agent acting on behalf of an account owner, or
// - a transfer agent acting on behalf of a fund manager or an account owner's designated agent.
//
// Usage
// The message should be sent at a frequency agreed bi-laterally between the account servicer and the account owner.
// The message can be sent either audited or un-audited and may be provided on a trade date, contractual or settlement date basis.
// This message can only be used to list the holdings of a single (master) account. However, it is possible to break down these holdings into one or several sub-accounts. Therefore, the message can be used to either specify holdings at
// - the main account level, or,
// - the sub-account level.
// This message can be used to report where the financial instruments are safe-kept, physically or notionally. If a security is held in more than one safekeeping place, this can also be indicated.
// The SecuritiesBalanceAccountingReport message should not be used for trading purposes.
// There may be one or more intermediary parties, for example, an intermediary or a concentrator between the account owner and account servicer.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesBalanceAccountingReport002V09 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *model.Statement51 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *model.PartyIdentification111 `xml:"AcctSvcr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount33 `xml:"SfkpgAcct"`

	// Information about the party that provides services relating to financial products to investors, for example, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*model.Intermediary37 `xml:"IntrmyInf,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*model.AggregateBalanceInformation32 `xml:"BalForAcct,omitempty"`

	// Sub-account of the safekeeping or investment account.
	SubAccountDetails []*model.SubAccountIdentification45 `xml:"SubAcctDtls,omitempty"`

	// Total valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyTotalAmounts *model.TotalValueInPageAndStatement4 `xml:"AcctBaseCcyTtlAmts,omitempty"`

	// Total valuation amounts provided in another currency than the base currency of the account.
	AlternateReportingCurrencyTotalAmounts *model.TotalValueInPageAndStatement4 `xml:"AltrnRptgCcyTtlAmts,omitempty"`
}

func (s *SecuritiesBalanceAccountingReport002V09) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesBalanceAccountingReport002V09) AddStatementGeneralDetails() *model.Statement51 {
	s.StatementGeneralDetails = new(model.Statement51)
	return s.StatementGeneralDetails
}

func (s *SecuritiesBalanceAccountingReport002V09) AddAccountOwner() *model.PartyIdentification109 {
	s.AccountOwner = new(model.PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesBalanceAccountingReport002V09) AddAccountServicer() *model.PartyIdentification111 {
	s.AccountServicer = new(model.PartyIdentification111)
	return s.AccountServicer
}

func (s *SecuritiesBalanceAccountingReport002V09) AddSafekeepingAccount() *model.SecuritiesAccount33 {
	s.SafekeepingAccount = new(model.SecuritiesAccount33)
	return s.SafekeepingAccount
}

func (s *SecuritiesBalanceAccountingReport002V09) AddIntermediaryInformation() *model.Intermediary37 {
	newValue := new(model.Intermediary37)
	s.IntermediaryInformation = append(s.IntermediaryInformation, newValue)
	return newValue
}

func (s *SecuritiesBalanceAccountingReport002V09) AddBalanceForAccount() *model.AggregateBalanceInformation32 {
	newValue := new(model.AggregateBalanceInformation32)
	s.BalanceForAccount = append(s.BalanceForAccount, newValue)
	return newValue
}

func (s *SecuritiesBalanceAccountingReport002V09) AddSubAccountDetails() *model.SubAccountIdentification45 {
	newValue := new(model.SubAccountIdentification45)
	s.SubAccountDetails = append(s.SubAccountDetails, newValue)
	return newValue
}

func (s *SecuritiesBalanceAccountingReport002V09) AddAccountBaseCurrencyTotalAmounts() *model.TotalValueInPageAndStatement4 {
	s.AccountBaseCurrencyTotalAmounts = new(model.TotalValueInPageAndStatement4)
	return s.AccountBaseCurrencyTotalAmounts
}

func (s *SecuritiesBalanceAccountingReport002V09) AddAlternateReportingCurrencyTotalAmounts() *model.TotalValueInPageAndStatement4 {
	s.AlternateReportingCurrencyTotalAmounts = new(model.TotalValueInPageAndStatement4)
	return s.AlternateReportingCurrencyTotalAmounts
}
