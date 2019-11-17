package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01900104 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.019.001.04 Document"`
	Message *AccountHoldingInformationRequestV04 `xml:"AcctHldgInfReq"`
}

func (d *Document01900104) AddMessage() *AccountHoldingInformationRequestV04 {
	d.Message = new(AccountHoldingInformationRequestV04)
	return d.Message
}

// Scope
// An instructing party, for example, a (new) plan manager (Transferee) sends the AccountHoldingInformationRequest message to the executing party, for example, a (old) plan manager (Transferor), on behalf of the initiating party, for example, an investor (client), to request information about financial instruments held on behalf of the client.
// Usage
// The AccountHoldingInformationRequest message is used to request information about one or more ISA or portfolio products held in a client's account for which it intends to instruct a transfer at a later time.
type AccountHoldingInformationRequestV04 struct {

	// Identifies the message.
	MessageReference *model.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference6 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference6 `xml:"RltdRef,omitempty"`

	// Identifies the business flow type (assets to be delivered or received).
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
	ProductTransfer []*model.ISATransfer27 `xml:"PdctTrf"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountHoldingInformationRequestV04) AddMessageReference() *model.MessageIdentification1 {
	a.MessageReference = new(model.MessageIdentification1)
	return a.MessageReference
}

func (a *AccountHoldingInformationRequestV04) AddPoolReference() *model.AdditionalReference6 {
	a.PoolReference = new(model.AdditionalReference6)
	return a.PoolReference
}

func (a *AccountHoldingInformationRequestV04) AddPreviousReference() *model.AdditionalReference6 {
	a.PreviousReference = new(model.AdditionalReference6)
	return a.PreviousReference
}

func (a *AccountHoldingInformationRequestV04) AddRelatedReference() *model.AdditionalReference6 {
	a.RelatedReference = new(model.AdditionalReference6)
	return a.RelatedReference
}

func (a *AccountHoldingInformationRequestV04) SetBusinessFlowDirectionType(value string) {
	a.BusinessFlowDirectionType = (*model.BusinessFlowDirectionType1Code)(&value)
}

func (a *AccountHoldingInformationRequestV04) AddPrimaryIndividualInvestor() *model.IndividualPerson8 {
	a.PrimaryIndividualInvestor = new(model.IndividualPerson8)
	return a.PrimaryIndividualInvestor
}

func (a *AccountHoldingInformationRequestV04) AddSecondaryIndividualInvestor() *model.IndividualPerson8 {
	a.SecondaryIndividualInvestor = new(model.IndividualPerson8)
	return a.SecondaryIndividualInvestor
}

func (a *AccountHoldingInformationRequestV04) AddOtherIndividualInvestor() *model.IndividualPerson8 {
	newValue := new(model.IndividualPerson8)
	a.OtherIndividualInvestor = append(a.OtherIndividualInvestor, newValue)
	return newValue
}

func (a *AccountHoldingInformationRequestV04) AddPrimaryCorporateInvestor() *model.Organisation21 {
	a.PrimaryCorporateInvestor = new(model.Organisation21)
	return a.PrimaryCorporateInvestor
}

func (a *AccountHoldingInformationRequestV04) AddSecondaryCorporateInvestor() *model.Organisation21 {
	a.SecondaryCorporateInvestor = new(model.Organisation21)
	return a.SecondaryCorporateInvestor
}

func (a *AccountHoldingInformationRequestV04) AddOtherCorporateInvestor() *model.Organisation21 {
	newValue := new(model.Organisation21)
	a.OtherCorporateInvestor = append(a.OtherCorporateInvestor, newValue)
	return newValue
}

func (a *AccountHoldingInformationRequestV04) AddTransferorAccount() *model.Account19 {
	a.TransferorAccount = new(model.Account19)
	return a.TransferorAccount
}

func (a *AccountHoldingInformationRequestV04) AddNomineeAccount() *model.Account19 {
	a.NomineeAccount = new(model.Account19)
	return a.NomineeAccount
}

func (a *AccountHoldingInformationRequestV04) AddTransferee() *model.PartyIdentification70Choice {
	a.Transferee = new(model.PartyIdentification70Choice)
	return a.Transferee
}

func (a *AccountHoldingInformationRequestV04) AddProductTransfer() *model.ISATransfer27 {
	newValue := new(model.ISATransfer27)
	a.ProductTransfer = append(a.ProductTransfer, newValue)
	return newValue
}

func (a *AccountHoldingInformationRequestV04) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountHoldingInformationRequestV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
