package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02000205 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.002.05 Document"`
	Message *SecuritiesTransactionCancellationRequest002V05 `xml:"SctiesTxCxlReq"`
}

func (d *Document02000205) AddMessage() *SecuritiesTransactionCancellationRequest002V05 {
	d.Message = new(SecuritiesTransactionCancellationRequest002V05)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesTransactionCancellationRequest to an account servicer to request the cancellation of a securities transaction.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with a local custodian, or
// - an investment management institution which manage a fund account opened at a custodian, or - a broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure, or
// - a central counterparty or a stock exchange or a trade matching utility which need to instruct to a central securities depository ot another settlement market infrastructure.
//
// Usage
// The transaction may be:
// - a securities settlement transaction
// - an intra-position movement
// - a securities financing transaction
// The instruction cannot be:
// - a securities settlement conditions modification (another transaction processing command should be sent to reverse a processing change previously requested).
// - a securities financing modification
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesTransactionCancellationRequest002V05 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *model.References60Choice `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *model.RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *model.RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Message Reference identifying the Processor of the transaction.
	ProcessorTransactionIdentification *model.RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Identifies the details of the transaction.
	TransactionDetails *model.TransactionDetails83 `xml:"TxDtls,omitempty"`

	// Specifies the reason of the cancellation.
	CancellationReason *model.CancellationReason19 `xml:"CxlRsn,omitempty"`

	// Specifies whether an associated FX should be cancelled.
	FXCancellation *model.FXCancellation4Choice `xml:"FxCxl,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTransactionCancellationRequest002V05) AddAccountOwnerTransactionIdentification() *model.References60Choice {
	s.AccountOwnerTransactionIdentification = new(model.References60Choice)
	return s.AccountOwnerTransactionIdentification
}

func (s *SecuritiesTransactionCancellationRequest002V05) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*model.RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTransactionCancellationRequest002V05) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*model.RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTransactionCancellationRequest002V05) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*model.RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTransactionCancellationRequest002V05) AddAccountOwner() *model.PartyIdentification109 {
	s.AccountOwner = new(model.PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesTransactionCancellationRequest002V05) AddSafekeepingAccount() *model.SecuritiesAccount30 {
	s.SafekeepingAccount = new(model.SecuritiesAccount30)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionCancellationRequest002V05) AddTransactionDetails() *model.TransactionDetails83 {
	s.TransactionDetails = new(model.TransactionDetails83)
	return s.TransactionDetails
}

func (s *SecuritiesTransactionCancellationRequest002V05) AddCancellationReason() *model.CancellationReason19 {
	s.CancellationReason = new(model.CancellationReason19)
	return s.CancellationReason
}

func (s *SecuritiesTransactionCancellationRequest002V05) AddFXCancellation() *model.FXCancellation4Choice {
	s.FXCancellation = new(model.FXCancellation4Choice)
	return s.FXCancellation
}

func (s *SecuritiesTransactionCancellationRequest002V05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
