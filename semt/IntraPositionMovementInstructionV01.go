package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:semt.013.001.01 Document"`
	Message *IntraPositionMovementInstructionV01 `xml:"IntraPosMvmntInstr"`
}

func (d *Document01300101) AddMessage() *IntraPositionMovementInstructionV01 {
	d.Message = new(IntraPositionMovementInstructionV01)
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
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type IntraPositionMovementInstructionV01 struct {

	// Information that unambiguously identifies an IntraPositionMovementTransaction and an IntraPositionMovementInstruction message as known by the account owner (or the instructing party acting on its behalf).
	Identification *model.TransactionAndDocumentIdentification1 `xml:"Id"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *model.Identification1 `xml:"CorpActnEvtId,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages *model.Linkages2 `xml:"Lnkgs,omitempty"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *model.SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes4 `xml:"FinInstrmAttrbts,omitempty"`

	// Intra-position movement transaction details.
	IntraPositionDetails *model.IntraPositionDetails1 `xml:"IntraPosDtls"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (i *IntraPositionMovementInstructionV01) AddIdentification() *model.TransactionAndDocumentIdentification1 {
	i.Identification = new(model.TransactionAndDocumentIdentification1)
	return i.Identification
}

func (i *IntraPositionMovementInstructionV01) AddCorporateActionEventIdentification() *model.Identification1 {
	i.CorporateActionEventIdentification = new(model.Identification1)
	return i.CorporateActionEventIdentification
}

func (i *IntraPositionMovementInstructionV01) AddLinkages() *model.Linkages2 {
	i.Linkages = new(model.Linkages2)
	return i.Linkages
}

func (i *IntraPositionMovementInstructionV01) AddAccountOwner() *model.PartyIdentification13Choice {
	i.AccountOwner = new(model.PartyIdentification13Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementInstructionV01) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	i.SafekeepingAccount = new(model.SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementInstructionV01) AddSafekeepingPlace() *model.SafekeepingPlaceFormat3Choice {
	i.SafekeepingPlace = new(model.SafekeepingPlaceFormat3Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionMovementInstructionV01) AddFinancialInstrumentIdentification() *model.SecurityIdentification11 {
	i.FinancialInstrumentIdentification = new(model.SecurityIdentification11)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionMovementInstructionV01) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes4 {
	i.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes4)
	return i.FinancialInstrumentAttributes
}

func (i *IntraPositionMovementInstructionV01) AddIntraPositionDetails() *model.IntraPositionDetails1 {
	i.IntraPositionDetails = new(model.IntraPositionDetails1)
	return i.IntraPositionDetails
}

func (i *IntraPositionMovementInstructionV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	i.MessageOriginator = new(model.PartyIdentification10Choice)
	return i.MessageOriginator
}

func (i *IntraPositionMovementInstructionV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	i.MessageRecipient = new(model.PartyIdentification10Choice)
	return i.MessageRecipient
}

func (i *IntraPositionMovementInstructionV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	i.Extension = append(i.Extension, newValue)
	return newValue
}
