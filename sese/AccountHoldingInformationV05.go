package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01800105 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.05 Document"`
	Message *AccountHoldingInformationV05 `xml:"AcctHldgInf"`
}

func (d *Document01800105) AddMessage() *AccountHoldingInformationV05 {
	d.Message = new(AccountHoldingInformationV05)
	return d.Message
}

// Scope
// An executing party, for example, a (old) plan manager (Transferor), sends the AccountHoldingInformation message to the instructing party, for example, a (new) plan manager (Transferee), to provide information about financial instruments held on behalf of a client.
// Usage
// The AccountHoldingInformation message is used to provide information about one or more ISA or portfolio products held in a client's account.
type AccountHoldingInformationV05 struct {

	// Identifies the message.
	MessageReference *model.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference6 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference6 `xml:"RltdRef,omitempty"`

	// Identifies the business flow direction type (assets to be delivered or received).
	BusinessFlowDirectionType *model.BusinessFlowDirectionType1Code `xml:"BizFlowDrctnTp,omitempty"`

	// Information identifying the primary individual investor, for example, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *model.IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, for example, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *model.IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying other individual investors, for example, name, address, social security number and date of birth.
	OtherIndividualInvestor []*model.IndividualPerson8 `xml:"OthrIndvInvstr,omitempty"`

	// Information identifying the primary corporate investor, for example, name and address.
	PrimaryCorporateInvestor *model.Organisation21 `xml:"PmryCorpInvstr,omitempty"`

	// Information identifying the secondary corporate investor, for example, name and address.
	SecondaryCorporateInvestor *model.Organisation21 `xml:"ScndryCorpInvstr,omitempty"`

	// Information identifying the other corporate investors, for example, name and address.
	OtherCorporateInvestor []*model.Organisation21 `xml:"OthrCorpInvstr,omitempty"`

	// Identification of an account owned by the investor at the old plan manager (account servicer).
	TransferorAccount *model.Account19 `xml:"TrfrAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *model.Account19 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	Transferee *model.PartyIdentification70Choice `xml:"Trfee"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*model.ISATransfer23 `xml:"PdctTrf"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountHoldingInformationV05) AddMessageReference() *model.MessageIdentification1 {
	a.MessageReference = new(model.MessageIdentification1)
	return a.MessageReference
}

func (a *AccountHoldingInformationV05) AddPoolReference() *model.AdditionalReference6 {
	a.PoolReference = new(model.AdditionalReference6)
	return a.PoolReference
}

func (a *AccountHoldingInformationV05) AddPreviousReference() *model.AdditionalReference6 {
	a.PreviousReference = new(model.AdditionalReference6)
	return a.PreviousReference
}

func (a *AccountHoldingInformationV05) AddRelatedReference() *model.AdditionalReference6 {
	a.RelatedReference = new(model.AdditionalReference6)
	return a.RelatedReference
}

func (a *AccountHoldingInformationV05) SetBusinessFlowDirectionType(value string) {
	a.BusinessFlowDirectionType = (*model.BusinessFlowDirectionType1Code)(&value)
}

func (a *AccountHoldingInformationV05) AddPrimaryIndividualInvestor() *model.IndividualPerson8 {
	a.PrimaryIndividualInvestor = new(model.IndividualPerson8)
	return a.PrimaryIndividualInvestor
}

func (a *AccountHoldingInformationV05) AddSecondaryIndividualInvestor() *model.IndividualPerson8 {
	a.SecondaryIndividualInvestor = new(model.IndividualPerson8)
	return a.SecondaryIndividualInvestor
}

func (a *AccountHoldingInformationV05) AddOtherIndividualInvestor() *model.IndividualPerson8 {
	newValue := new(model.IndividualPerson8)
	a.OtherIndividualInvestor = append(a.OtherIndividualInvestor, newValue)
	return newValue
}

func (a *AccountHoldingInformationV05) AddPrimaryCorporateInvestor() *model.Organisation21 {
	a.PrimaryCorporateInvestor = new(model.Organisation21)
	return a.PrimaryCorporateInvestor
}

func (a *AccountHoldingInformationV05) AddSecondaryCorporateInvestor() *model.Organisation21 {
	a.SecondaryCorporateInvestor = new(model.Organisation21)
	return a.SecondaryCorporateInvestor
}

func (a *AccountHoldingInformationV05) AddOtherCorporateInvestor() *model.Organisation21 {
	newValue := new(model.Organisation21)
	a.OtherCorporateInvestor = append(a.OtherCorporateInvestor, newValue)
	return newValue
}

func (a *AccountHoldingInformationV05) AddTransferorAccount() *model.Account19 {
	a.TransferorAccount = new(model.Account19)
	return a.TransferorAccount
}

func (a *AccountHoldingInformationV05) AddNomineeAccount() *model.Account19 {
	a.NomineeAccount = new(model.Account19)
	return a.NomineeAccount
}

func (a *AccountHoldingInformationV05) AddTransferee() *model.PartyIdentification70Choice {
	a.Transferee = new(model.PartyIdentification70Choice)
	return a.Transferee
}

func (a *AccountHoldingInformationV05) AddProductTransfer() *model.ISATransfer23 {
	newValue := new(model.ISATransfer23)
	a.ProductTransfer = append(a.ProductTransfer, newValue)
	return newValue
}

func (a *AccountHoldingInformationV05) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountHoldingInformationV05) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
