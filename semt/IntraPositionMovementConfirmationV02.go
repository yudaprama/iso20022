package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500102 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.015.001.02 Document"`
	Message *IntraPositionMovementConfirmationV02 `xml:"IntraPosMvmntConf"`
}

func (d *Document01500102) AddMessage() *IntraPositionMovementConfirmationV02 {
	d.Message = new(IntraPositionMovementConfirmationV02)
	return d.Message
}

// Scope
// An account servicer sends a IntraPositionMovementConfirmation to an account owner to confirm the movement of securities within its holding from one sub-balance to another, for example, blocking of securities.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type IntraPositionMovementConfirmationV02 struct {

	// Additional parameters to the transaction.
	AdditionalParameters *model.AdditionalParameters8 `xml:"AddtlParams,omitempty"`

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
	IntraPositionDetails *model.IntraPositionDetails12 `xml:"IntraPosDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementConfirmationV02) AddAdditionalParameters() *model.AdditionalParameters8 {
	i.AdditionalParameters = new(model.AdditionalParameters8)
	return i.AdditionalParameters
}

func (i *IntraPositionMovementConfirmationV02) AddAccountOwner() *model.PartyIdentification36Choice {
	i.AccountOwner = new(model.PartyIdentification36Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementConfirmationV02) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	i.SafekeepingAccount = new(model.SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementConfirmationV02) AddSafekeepingPlace() *model.SafekeepingPlaceFormat3Choice {
	i.SafekeepingPlace = new(model.SafekeepingPlaceFormat3Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionMovementConfirmationV02) AddFinancialInstrumentIdentification() *model.SecurityIdentification14 {
	i.FinancialInstrumentIdentification = new(model.SecurityIdentification14)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionMovementConfirmationV02) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes21 {
	i.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes21)
	return i.FinancialInstrumentAttributes
}

func (i *IntraPositionMovementConfirmationV02) AddIntraPositionDetails() *model.IntraPositionDetails12 {
	i.IntraPositionDetails = new(model.IntraPositionDetails12)
	return i.IntraPositionDetails
}

func (i *IntraPositionMovementConfirmationV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
