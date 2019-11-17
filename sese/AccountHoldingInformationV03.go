package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01800103 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.03 Document"`
	Message *AccountHoldingInformationV03 `xml:"AcctHldgInf"`
}

func (d *Document01800103) AddMessage() *AccountHoldingInformationV03 {
	d.Message = new(AccountHoldingInformationV03)
	return d.Message
}

// Scope
// An executing party, for example, a (old) plan manager (Transferor), sends the AccountHoldingInformation message to the instructing party, for example, a (new) plan manager (Transferee), to provide information about financial instruments held on behalf of a client.
// Usage
// The AccountHoldingInformation message is used to provide information about one or more ISA or portfolio products held in a client's account.
type AccountHoldingInformationV03 struct {

	// Identifies the message.
	MessageReference *model.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Identifies the business flow direction type (assets to be delivered or received).
	BusinessFlowDirectionType *model.BusinessFlowDirectionType1Code `xml:"BizFlowDrctnTp,omitempty"`

	// Information identifying the primary individual investor, eg, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *model.IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, eg, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *model.IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying other individual investors, eg, name, address, social security number and date of birth.
	OtherIndividualInvestor []*model.IndividualPerson8 `xml:"OthrIndvInvstr,omitempty"`

	// Information identifying the primary corporate investor, eg, name and address.
	PrimaryCorporateInvestor *model.Organisation4 `xml:"PmryCorpInvstr,omitempty"`

	// Information identifying the secondary corporate investor, eg, name and address.
	SecondaryCorporateInvestor *model.Organisation4 `xml:"ScndryCorpInvstr,omitempty"`

	// Information identifying the other corporate investors, eg, name and address.
	OtherCorporateInvestor []*model.Organisation4 `xml:"OthrCorpInvstr,omitempty"`

	// Identification of an account owned by the investor at the old plan manager (account servicer).
	TransferorAccount *model.Account5 `xml:"TrfrAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *model.Account6 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	Transferee *model.PartyIdentification2Choice `xml:"Trfee"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*model.ISATransfer6 `xml:"PdctTrf"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountHoldingInformationV03) AddMessageReference() *model.MessageIdentification1 {
	a.MessageReference = new(model.MessageIdentification1)
	return a.MessageReference
}

func (a *AccountHoldingInformationV03) AddPoolReference() *model.AdditionalReference3 {
	a.PoolReference = new(model.AdditionalReference3)
	return a.PoolReference
}

func (a *AccountHoldingInformationV03) AddPreviousReference() *model.AdditionalReference3 {
	a.PreviousReference = new(model.AdditionalReference3)
	return a.PreviousReference
}

func (a *AccountHoldingInformationV03) AddRelatedReference() *model.AdditionalReference3 {
	a.RelatedReference = new(model.AdditionalReference3)
	return a.RelatedReference
}

func (a *AccountHoldingInformationV03) SetBusinessFlowDirectionType(value string) {
	a.BusinessFlowDirectionType = (*model.BusinessFlowDirectionType1Code)(&value)
}

func (a *AccountHoldingInformationV03) AddPrimaryIndividualInvestor() *model.IndividualPerson8 {
	a.PrimaryIndividualInvestor = new(model.IndividualPerson8)
	return a.PrimaryIndividualInvestor
}

func (a *AccountHoldingInformationV03) AddSecondaryIndividualInvestor() *model.IndividualPerson8 {
	a.SecondaryIndividualInvestor = new(model.IndividualPerson8)
	return a.SecondaryIndividualInvestor
}

func (a *AccountHoldingInformationV03) AddOtherIndividualInvestor() *model.IndividualPerson8 {
	newValue := new(model.IndividualPerson8)
	a.OtherIndividualInvestor = append(a.OtherIndividualInvestor, newValue)
	return newValue
}

func (a *AccountHoldingInformationV03) AddPrimaryCorporateInvestor() *model.Organisation4 {
	a.PrimaryCorporateInvestor = new(model.Organisation4)
	return a.PrimaryCorporateInvestor
}

func (a *AccountHoldingInformationV03) AddSecondaryCorporateInvestor() *model.Organisation4 {
	a.SecondaryCorporateInvestor = new(model.Organisation4)
	return a.SecondaryCorporateInvestor
}

func (a *AccountHoldingInformationV03) AddOtherCorporateInvestor() *model.Organisation4 {
	newValue := new(model.Organisation4)
	a.OtherCorporateInvestor = append(a.OtherCorporateInvestor, newValue)
	return newValue
}

func (a *AccountHoldingInformationV03) AddTransferorAccount() *model.Account5 {
	a.TransferorAccount = new(model.Account5)
	return a.TransferorAccount
}

func (a *AccountHoldingInformationV03) AddNomineeAccount() *model.Account6 {
	a.NomineeAccount = new(model.Account6)
	return a.NomineeAccount
}

func (a *AccountHoldingInformationV03) AddTransferee() *model.PartyIdentification2Choice {
	a.Transferee = new(model.PartyIdentification2Choice)
	return a.Transferee
}

func (a *AccountHoldingInformationV03) AddProductTransfer() *model.ISATransfer6 {
	newValue := new(model.ISATransfer6)
	a.ProductTransfer = append(a.ProductTransfer, newValue)
	return newValue
}

func (a *AccountHoldingInformationV03) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
