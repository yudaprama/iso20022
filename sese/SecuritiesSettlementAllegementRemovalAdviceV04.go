package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02900104 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.029.001.04 Document"`
	Message *SecuritiesSettlementAllegementRemovalAdviceV04 `xml:"SctiesSttlmAllgmtRmvlAdvc"`
}

func (d *Document02900104) AddMessage() *SecuritiesSettlementAllegementRemovalAdviceV04 {
	d.Message = new(SecuritiesSettlementAllegementRemovalAdviceV04)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementAllegementRemovalAdvice to an account owner to acknowledge that a previously sent allegement is no longer outstanding, because the alleged party sent its instruction.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementAllegementRemovalAdviceV04 struct {

	// Provides transaction type and identification information.
	AccountServicerTransactionIdentification *model.SettlementTypeAndIdentification18 `xml:"AcctSvcrTxId"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *model.Identification14 `xml:"MktInfrstrctrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Identifies the details of the transaction.
	TransactionDetails *model.TransactionDetails74 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV04) AddAccountServicerTransactionIdentification() *model.SettlementTypeAndIdentification18 {
	s.AccountServicerTransactionIdentification = new(model.SettlementTypeAndIdentification18)
	return s.AccountServicerTransactionIdentification
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV04) AddMarketInfrastructureTransactionIdentification() *model.Identification14 {
	s.MarketInfrastructureTransactionIdentification = new(model.Identification14)
	return s.MarketInfrastructureTransactionIdentification
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV04) AddAccountOwner() *model.PartyIdentification98 {
	s.AccountOwner = new(model.PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV04) AddSafekeepingAccount() *model.SecuritiesAccount19 {
	s.SafekeepingAccount = new(model.SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV04) AddTransactionDetails() *model.TransactionDetails74 {
	s.TransactionDetails = new(model.TransactionDetails74)
	return s.TransactionDetails
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
