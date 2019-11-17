package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02900102 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.029.001.02 Document"`
	Message *SecuritiesSettlementAllegementRemovalAdviceV02 `xml:"SctiesSttlmAllgmtRmvlAdvc"`
}

func (d *Document02900102) AddMessage() *SecuritiesSettlementAllegementRemovalAdviceV02 {
	d.Message = new(SecuritiesSettlementAllegementRemovalAdviceV02)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementAllegementRemovalAdvice to an account owner to acknowledge that a previously sent allegement is no longer outstanding, because the alleged party sent its instruction.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// using the relevant elements in the Business Application Header.
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementAllegementRemovalAdviceV02 struct {

	// Provides transaction type and identification information.
	AccountServicerTransactionIdentification *model.SettlementTypeAndIdentification3 `xml:"AcctSvcrTxId"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *model.Identification1 `xml:"MktInfrstrctrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Identifies the details of the transaction.
	TransactionDetails *model.TransactionDetails28 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV02) AddAccountServicerTransactionIdentification() *model.SettlementTypeAndIdentification3 {
	s.AccountServicerTransactionIdentification = new(model.SettlementTypeAndIdentification3)
	return s.AccountServicerTransactionIdentification
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV02) AddMarketInfrastructureTransactionIdentification() *model.Identification1 {
	s.MarketInfrastructureTransactionIdentification = new(model.Identification1)
	return s.MarketInfrastructureTransactionIdentification
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV02) AddAccountOwner() *model.PartyIdentification36Choice {
	s.AccountOwner = new(model.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV02) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	s.SafekeepingAccount = new(model.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV02) AddTransactionDetails() *model.TransactionDetails28 {
	s.TransactionDetails = new(model.TransactionDetails28)
	return s.TransactionDetails
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
