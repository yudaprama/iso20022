package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500105 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.015.001.05 Document"`
	Message *IntraPositionMovementConfirmationV05 `xml:"IntraPosMvmntConf"`
}

func (d *Document01500105) AddMessage() *IntraPositionMovementConfirmationV05 {
	d.Message = new(IntraPositionMovementConfirmationV05)
	return d.Message
}

// Scope
// An account servicer sends a IntraPositionMovementConfirmation to an account owner to confirm the movement of securities within its holding from one sub-balance to another, for example, blocking of securities.
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
type IntraPositionMovementConfirmationV05 struct {

	// Additional parameters to the transaction.
	AdditionalParameters *model.AdditionalParameters21 `xml:"AddtlParams,omitempty"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *model.SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes63 `xml:"FinInstrmAttrbts,omitempty"`

	// Intra-position movement transaction details.
	IntraPositionDetails *model.IntraPositionDetails34 `xml:"IntraPosDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementConfirmationV05) AddAdditionalParameters() *model.AdditionalParameters21 {
	i.AdditionalParameters = new(model.AdditionalParameters21)
	return i.AdditionalParameters
}

func (i *IntraPositionMovementConfirmationV05) AddAccountOwner() *model.PartyIdentification92Choice {
	i.AccountOwner = new(model.PartyIdentification92Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementConfirmationV05) AddSafekeepingAccount() *model.SecuritiesAccount24 {
	i.SafekeepingAccount = new(model.SecuritiesAccount24)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementConfirmationV05) AddSafekeepingPlace() *model.SafekeepingPlaceFormat10Choice {
	i.SafekeepingPlace = new(model.SafekeepingPlaceFormat10Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionMovementConfirmationV05) AddFinancialInstrumentIdentification() *model.SecurityIdentification19 {
	i.FinancialInstrumentIdentification = new(model.SecurityIdentification19)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionMovementConfirmationV05) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes63 {
	i.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes63)
	return i.FinancialInstrumentAttributes
}

func (i *IntraPositionMovementConfirmationV05) AddIntraPositionDetails() *model.IntraPositionDetails34 {
	i.IntraPositionDetails = new(model.IntraPositionDetails34)
	return i.IntraPositionDetails
}

func (i *IntraPositionMovementConfirmationV05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
