package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500101 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.015.001.01 Document"`
	Message *IntraPositionMovementConfirmationV01 `xml:"IntraPosMvmntConf"`
}

func (d *Document01500101) AddMessage() *IntraPositionMovementConfirmationV01 {
	d.Message = new(IntraPositionMovementConfirmationV01)
	return d.Message
}

// Scope
// An account servicer sends a IntraPositionMovementConfirmation to an account owner to confirm the movement of securities within its holding from one sub-balance to another, for example, blocking of securities.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type IntraPositionMovementConfirmationV01 struct {

	// Information that unambiguously identifies an IntraPositionMovementConfirmation message as known by the account servicer.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Additional parameters to the transaction.
	AdditionalParameters *model.AdditionalParameters3 `xml:"AddtlParams,omitempty"`

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
	IntraPositionDetails *model.IntraPositionDetails2 `xml:"IntraPosDtls"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (i *IntraPositionMovementConfirmationV01) AddIdentification() *model.DocumentIdentification11 {
	i.Identification = new(model.DocumentIdentification11)
	return i.Identification
}

func (i *IntraPositionMovementConfirmationV01) AddAdditionalParameters() *model.AdditionalParameters3 {
	i.AdditionalParameters = new(model.AdditionalParameters3)
	return i.AdditionalParameters
}

func (i *IntraPositionMovementConfirmationV01) AddAccountOwner() *model.PartyIdentification13Choice {
	i.AccountOwner = new(model.PartyIdentification13Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementConfirmationV01) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	i.SafekeepingAccount = new(model.SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementConfirmationV01) AddSafekeepingPlace() *model.SafekeepingPlaceFormat3Choice {
	i.SafekeepingPlace = new(model.SafekeepingPlaceFormat3Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionMovementConfirmationV01) AddFinancialInstrumentIdentification() *model.SecurityIdentification11 {
	i.FinancialInstrumentIdentification = new(model.SecurityIdentification11)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionMovementConfirmationV01) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes4 {
	i.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes4)
	return i.FinancialInstrumentAttributes
}

func (i *IntraPositionMovementConfirmationV01) AddIntraPositionDetails() *model.IntraPositionDetails2 {
	i.IntraPositionDetails = new(model.IntraPositionDetails2)
	return i.IntraPositionDetails
}

func (i *IntraPositionMovementConfirmationV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	i.MessageOriginator = new(model.PartyIdentification10Choice)
	return i.MessageOriginator
}

func (i *IntraPositionMovementConfirmationV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	i.MessageRecipient = new(model.PartyIdentification10Choice)
	return i.MessageRecipient
}

func (i *IntraPositionMovementConfirmationV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	i.Extension = append(i.Extension, newValue)
	return newValue
}
