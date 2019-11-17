package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300107 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.003.001.07 Document"`
	Message *SecuritiesBalanceAccountingReportV07 `xml:"SctiesBalAcctgRpt"`
}

func (d *Document00300107) AddMessage() *SecuritiesBalanceAccountingReportV07 {
	d.Message = new(SecuritiesBalanceAccountingReportV07)
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
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesBalanceAccountingReportV07 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *model.Statement20 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *model.PartyIdentification49Choice `xml:"AcctSvcr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount11 `xml:"SfkpgAcct"`

	// Information about the party that provides services relating to financial products to investors, for example, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*model.Intermediary23 `xml:"IntrmyInf,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*model.AggregateBalanceInformation22 `xml:"BalForAcct,omitempty"`

	// Sub-account of the safekeeping or investment account.
	SubAccountDetails []*model.SubAccountIdentification29 `xml:"SubAcctDtls,omitempty"`

	// Total valuation amounts provided in the base currency of the account.
	AccountBaseCurrencyTotalAmounts *model.TotalValueInPageAndStatement2 `xml:"AcctBaseCcyTtlAmts,omitempty"`

	// Total valuation amounts provided in another currency than the base currency of the account.
	AlternateReportingCurrencyTotalAmounts *model.TotalValueInPageAndStatement2 `xml:"AltrnRptgCcyTtlAmts,omitempty"`
}

func (s *SecuritiesBalanceAccountingReportV07) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesBalanceAccountingReportV07) AddStatementGeneralDetails() *model.Statement20 {
	s.StatementGeneralDetails = new(model.Statement20)
	return s.StatementGeneralDetails
}

func (s *SecuritiesBalanceAccountingReportV07) AddAccountOwner() *model.PartyIdentification36Choice {
	s.AccountOwner = new(model.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesBalanceAccountingReportV07) AddAccountServicer() *model.PartyIdentification49Choice {
	s.AccountServicer = new(model.PartyIdentification49Choice)
	return s.AccountServicer
}

func (s *SecuritiesBalanceAccountingReportV07) AddSafekeepingAccount() *model.SecuritiesAccount11 {
	s.SafekeepingAccount = new(model.SecuritiesAccount11)
	return s.SafekeepingAccount
}

func (s *SecuritiesBalanceAccountingReportV07) AddIntermediaryInformation() *model.Intermediary23 {
	newValue := new(model.Intermediary23)
	s.IntermediaryInformation = append(s.IntermediaryInformation, newValue)
	return newValue
}

func (s *SecuritiesBalanceAccountingReportV07) AddBalanceForAccount() *model.AggregateBalanceInformation22 {
	newValue := new(model.AggregateBalanceInformation22)
	s.BalanceForAccount = append(s.BalanceForAccount, newValue)
	return newValue
}

func (s *SecuritiesBalanceAccountingReportV07) AddSubAccountDetails() *model.SubAccountIdentification29 {
	newValue := new(model.SubAccountIdentification29)
	s.SubAccountDetails = append(s.SubAccountDetails, newValue)
	return newValue
}

func (s *SecuritiesBalanceAccountingReportV07) AddAccountBaseCurrencyTotalAmounts() *model.TotalValueInPageAndStatement2 {
	s.AccountBaseCurrencyTotalAmounts = new(model.TotalValueInPageAndStatement2)
	return s.AccountBaseCurrencyTotalAmounts
}

func (s *SecuritiesBalanceAccountingReportV07) AddAlternateReportingCurrencyTotalAmounts() *model.TotalValueInPageAndStatement2 {
	s.AlternateReportingCurrencyTotalAmounts = new(model.TotalValueInPageAndStatement2)
	return s.AlternateReportingCurrencyTotalAmounts
}
