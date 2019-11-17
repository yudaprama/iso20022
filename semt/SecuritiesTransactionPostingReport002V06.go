package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700206 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.017.002.06 Document"`
	Message *SecuritiesTransactionPostingReport002V06 `xml:"SctiesTxPstngRpt"`
}

func (d *Document01700206) AddMessage() *SecuritiesTransactionPostingReport002V06 {
	d.Message = new(SecuritiesTransactionPostingReport002V06)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesTransactionPostingReport to an account owner to provide the details of increases and decreases of holdings which occurred during a specified period, for all or selected securities in the specified safekeeping account or sub-safekeeping account which the account servicer holds for the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// This message may be used as a trade date based or a settlement date based statement.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesTransactionPostingReport002V06 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *model.Statement56 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Reporting per financial instrument.
	FinancialInstrumentDetails []*model.FinancialInstrumentDetails23 `xml:"FinInstrmDtls,omitempty"`

	// Details at sub-account level.
	SubAccountDetails []*model.SubAccountIdentification47 `xml:"SubAcctDtls,omitempty"`
}

func (s *SecuritiesTransactionPostingReport002V06) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesTransactionPostingReport002V06) AddStatementGeneralDetails() *model.Statement56 {
	s.StatementGeneralDetails = new(model.Statement56)
	return s.StatementGeneralDetails
}

func (s *SecuritiesTransactionPostingReport002V06) AddAccountOwner() *model.PartyIdentification109 {
	s.AccountOwner = new(model.PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesTransactionPostingReport002V06) AddSafekeepingAccount() *model.SecuritiesAccount27 {
	s.SafekeepingAccount = new(model.SecuritiesAccount27)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionPostingReport002V06) AddFinancialInstrumentDetails() *model.FinancialInstrumentDetails23 {
	newValue := new(model.FinancialInstrumentDetails23)
	s.FinancialInstrumentDetails = append(s.FinancialInstrumentDetails, newValue)
	return newValue
}

func (s *SecuritiesTransactionPostingReport002V06) AddSubAccountDetails() *model.SubAccountIdentification47 {
	newValue := new(model.SubAccountIdentification47)
	s.SubAccountDetails = append(s.SubAccountDetails, newValue)
	return newValue
}
