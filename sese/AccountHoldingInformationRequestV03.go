package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01900103 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.019.001.03 Document"`
	Message *AccountHoldingInformationRequestV03 `xml:"AcctHldgInfReq"`
}

func (d *Document01900103) AddMessage() *AccountHoldingInformationRequestV03 {
	d.Message = new(AccountHoldingInformationRequestV03)
	return d.Message
}

// Scope
// An instructing party, for example, a (new) plan manager (Transferee) sends the AccountHoldingInformationRequest message to the executing party, for example, a (old) plan manager (Transferor), on behalf of the initiating party, for example, an investor (client), to request information about financial instruments held on behalf of the client.
// Usage
// The AccountHoldingInformationRequest message is used to request information about one or more ISA or portfolio products held in a client's account for which it intends to instruct a transfer at a later time.
type AccountHoldingInformationRequestV03 struct {

	// Identifies the message.
	MessageReference *model.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Identifies the business flow type (assets to be delivered or received).
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
	TransferorAccount *model.Account15 `xml:"TrfrAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *model.Account16 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	Transferee *model.PartyIdentification2Choice `xml:"Trfee"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*model.ISATransfer15 `xml:"PdctTrf"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountHoldingInformationRequestV03) AddMessageReference() *model.MessageIdentification1 {
	a.MessageReference = new(model.MessageIdentification1)
	return a.MessageReference
}

func (a *AccountHoldingInformationRequestV03) AddPoolReference() *model.AdditionalReference3 {
	a.PoolReference = new(model.AdditionalReference3)
	return a.PoolReference
}

func (a *AccountHoldingInformationRequestV03) AddPreviousReference() *model.AdditionalReference3 {
	a.PreviousReference = new(model.AdditionalReference3)
	return a.PreviousReference
}

func (a *AccountHoldingInformationRequestV03) AddRelatedReference() *model.AdditionalReference3 {
	a.RelatedReference = new(model.AdditionalReference3)
	return a.RelatedReference
}

func (a *AccountHoldingInformationRequestV03) SetBusinessFlowDirectionType(value string) {
	a.BusinessFlowDirectionType = (*model.BusinessFlowDirectionType1Code)(&value)
}

func (a *AccountHoldingInformationRequestV03) AddPrimaryIndividualInvestor() *model.IndividualPerson8 {
	a.PrimaryIndividualInvestor = new(model.IndividualPerson8)
	return a.PrimaryIndividualInvestor
}

func (a *AccountHoldingInformationRequestV03) AddSecondaryIndividualInvestor() *model.IndividualPerson8 {
	a.SecondaryIndividualInvestor = new(model.IndividualPerson8)
	return a.SecondaryIndividualInvestor
}

func (a *AccountHoldingInformationRequestV03) AddOtherIndividualInvestor() *model.IndividualPerson8 {
	newValue := new(model.IndividualPerson8)
	a.OtherIndividualInvestor = append(a.OtherIndividualInvestor, newValue)
	return newValue
}

func (a *AccountHoldingInformationRequestV03) AddPrimaryCorporateInvestor() *model.Organisation4 {
	a.PrimaryCorporateInvestor = new(model.Organisation4)
	return a.PrimaryCorporateInvestor
}

func (a *AccountHoldingInformationRequestV03) AddSecondaryCorporateInvestor() *model.Organisation4 {
	a.SecondaryCorporateInvestor = new(model.Organisation4)
	return a.SecondaryCorporateInvestor
}

func (a *AccountHoldingInformationRequestV03) AddOtherCorporateInvestor() *model.Organisation4 {
	newValue := new(model.Organisation4)
	a.OtherCorporateInvestor = append(a.OtherCorporateInvestor, newValue)
	return newValue
}

func (a *AccountHoldingInformationRequestV03) AddTransferorAccount() *model.Account15 {
	a.TransferorAccount = new(model.Account15)
	return a.TransferorAccount
}

func (a *AccountHoldingInformationRequestV03) AddNomineeAccount() *model.Account16 {
	a.NomineeAccount = new(model.Account16)
	return a.NomineeAccount
}

func (a *AccountHoldingInformationRequestV03) AddTransferee() *model.PartyIdentification2Choice {
	a.Transferee = new(model.PartyIdentification2Choice)
	return a.Transferee
}

func (a *AccountHoldingInformationRequestV03) AddProductTransfer() *model.ISATransfer15 {
	newValue := new(model.ISATransfer15)
	a.ProductTransfer = append(a.ProductTransfer, newValue)
	return newValue
}

func (a *AccountHoldingInformationRequestV03) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountHoldingInformationRequestV03) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
