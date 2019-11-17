package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300102 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.02 Document"`
	Message *IntraPositionMovementInstructionV02 `xml:"IntraPosMvmntInstr"`
}

func (d *Document01300102) AddMessage() *IntraPositionMovementInstructionV02 {
	d.Message = new(IntraPositionMovementInstructionV02)
	return d.Message
}

// Scope
// An account owner sends a IntraPositionMovementInstruction to an account servicer to instruct the movement of securities within its holding from one sub-balance to another, for example, blocking of securities.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with its local agent (sub-custodian), or
// - an investment management institution which manage a fund account opened at a custodian, or
// - broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure.
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type IntraPositionMovementInstructionV02 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *model.Identification1 `xml:"CorpActnEvtId,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages *model.Linkages8 `xml:"Lnkgs,omitempty"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *model.SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes21 `xml:"FinInstrmAttrbts,omitempty"`

	// Intra-position movement transaction details.
	IntraPositionDetails *model.IntraPositionDetails11 `xml:"IntraPosDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementInstructionV02) SetTransactionIdentification(value string) {
	i.TransactionIdentification = (*model.Max35Text)(&value)
}

func (i *IntraPositionMovementInstructionV02) AddCorporateActionEventIdentification() *model.Identification1 {
	i.CorporateActionEventIdentification = new(model.Identification1)
	return i.CorporateActionEventIdentification
}

func (i *IntraPositionMovementInstructionV02) AddLinkages() *model.Linkages8 {
	i.Linkages = new(model.Linkages8)
	return i.Linkages
}

func (i *IntraPositionMovementInstructionV02) AddAccountOwner() *model.PartyIdentification36Choice {
	i.AccountOwner = new(model.PartyIdentification36Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementInstructionV02) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	i.SafekeepingAccount = new(model.SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementInstructionV02) AddSafekeepingPlace() *model.SafekeepingPlaceFormat3Choice {
	i.SafekeepingPlace = new(model.SafekeepingPlaceFormat3Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionMovementInstructionV02) AddFinancialInstrumentIdentification() *model.SecurityIdentification14 {
	i.FinancialInstrumentIdentification = new(model.SecurityIdentification14)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionMovementInstructionV02) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes21 {
	i.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes21)
	return i.FinancialInstrumentAttributes
}

func (i *IntraPositionMovementInstructionV02) AddIntraPositionDetails() *model.IntraPositionDetails11 {
	i.IntraPositionDetails = new(model.IntraPositionDetails11)
	return i.IntraPositionDetails
}

func (i *IntraPositionMovementInstructionV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
