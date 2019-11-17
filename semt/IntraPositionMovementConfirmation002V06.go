package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500206 struct {
	XMLName xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.015.002.06 Document"`
	Message *IntraPositionMovementConfirmation002V06 `xml:"IntraPosMvmntConf"`
}

func (d *Document01500206) AddMessage() *IntraPositionMovementConfirmation002V06 {
	d.Message = new(IntraPositionMovementConfirmation002V06)
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
type IntraPositionMovementConfirmation002V06 struct {

	// Additional parameters to the transaction.
	AdditionalParameters *model.AdditionalParameters25 `xml:"AddtlParams,omitempty"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *model.SafekeepingPlaceFormat17Choice `xml:"SfkpgPlc,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *model.SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *model.FinancialInstrumentAttributes75 `xml:"FinInstrmAttrbts,omitempty"`

	// Intra-position movement transaction details.
	IntraPositionDetails *model.IntraPositionDetails43 `xml:"IntraPosDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementConfirmation002V06) AddAdditionalParameters() *model.AdditionalParameters25 {
	i.AdditionalParameters = new(model.AdditionalParameters25)
	return i.AdditionalParameters
}

func (i *IntraPositionMovementConfirmation002V06) AddAccountOwner() *model.PartyIdentification103Choice {
	i.AccountOwner = new(model.PartyIdentification103Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementConfirmation002V06) AddSafekeepingAccount() *model.SecuritiesAccount27 {
	i.SafekeepingAccount = new(model.SecuritiesAccount27)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementConfirmation002V06) AddSafekeepingPlace() *model.SafekeepingPlaceFormat17Choice {
	i.SafekeepingPlace = new(model.SafekeepingPlaceFormat17Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionMovementConfirmation002V06) AddFinancialInstrumentIdentification() *model.SecurityIdentification20 {
	i.FinancialInstrumentIdentification = new(model.SecurityIdentification20)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionMovementConfirmation002V06) AddFinancialInstrumentAttributes() *model.FinancialInstrumentAttributes75 {
	i.FinancialInstrumentAttributes = new(model.FinancialInstrumentAttributes75)
	return i.FinancialInstrumentAttributes
}

func (i *IntraPositionMovementConfirmation002V06) AddIntraPositionDetails() *model.IntraPositionDetails43 {
	i.IntraPositionDetails = new(model.IntraPositionDetails43)
	return i.IntraPositionDetails
}

func (i *IntraPositionMovementConfirmation002V06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
