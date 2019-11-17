package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01900101 struct {
	XMLName xml.Name                                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.019.001.01 Document"`
	Message *RequestForPEPOrISAOrPortfolioInformationV01 `xml:"ReqForPEPOrISAOrPrtflInfV01"`
}

func (d *Document01900101) AddMessage() *RequestForPEPOrISAOrPortfolioInformationV01 {
	d.Message = new(RequestForPEPOrISAOrPortfolioInformationV01)
	return d.Message
}

// Scope
// An instructing party, eg, a (new) plan manager sends the RequestForPEPorISAOrPortfolioInformation message to the executing party, eg, a (old) plan manager, on behalf of the initiating party, eg, an investor (client), to request information about financial instruments held on behalf of the client.
// Usage
// The RequestForPEPOrISAOrPortfolioInformation message is used to request information about one or more PEP or ISA or portfolio products held in a client's account for which it intends to instruct a transfer at a later time.
type RequestForPEPOrISAOrPortfolioInformationV01 struct {

	// Identifies the message.
	MessageReference *model.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

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
	ClientAccount *model.Account5 `xml:"ClntAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *model.Account6 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	NewPlanManager *model.PartyIdentification2Choice `xml:"NewPlanMgr"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*model.PEPISATransfer5 `xml:"PdctTrf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddMessageReference() *model.MessageIdentification1 {
	r.MessageReference = new(model.MessageIdentification1)
	return r.MessageReference
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddPoolReference() *model.AdditionalReference3 {
	r.PoolReference = new(model.AdditionalReference3)
	return r.PoolReference
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddPreviousReference() *model.AdditionalReference3 {
	r.PreviousReference = new(model.AdditionalReference3)
	return r.PreviousReference
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddRelatedReference() *model.AdditionalReference3 {
	r.RelatedReference = new(model.AdditionalReference3)
	return r.RelatedReference
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddPrimaryIndividualInvestor() *model.IndividualPerson8 {
	r.PrimaryIndividualInvestor = new(model.IndividualPerson8)
	return r.PrimaryIndividualInvestor
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddSecondaryIndividualInvestor() *model.IndividualPerson8 {
	r.SecondaryIndividualInvestor = new(model.IndividualPerson8)
	return r.SecondaryIndividualInvestor
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddOtherIndividualInvestor() *model.IndividualPerson8 {
	newValue := new(model.IndividualPerson8)
	r.OtherIndividualInvestor = append(r.OtherIndividualInvestor, newValue)
	return newValue
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddPrimaryCorporateInvestor() *model.Organisation4 {
	r.PrimaryCorporateInvestor = new(model.Organisation4)
	return r.PrimaryCorporateInvestor
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddSecondaryCorporateInvestor() *model.Organisation4 {
	r.SecondaryCorporateInvestor = new(model.Organisation4)
	return r.SecondaryCorporateInvestor
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddOtherCorporateInvestor() *model.Organisation4 {
	newValue := new(model.Organisation4)
	r.OtherCorporateInvestor = append(r.OtherCorporateInvestor, newValue)
	return newValue
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddClientAccount() *model.Account5 {
	r.ClientAccount = new(model.Account5)
	return r.ClientAccount
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddNomineeAccount() *model.Account6 {
	r.NomineeAccount = new(model.Account6)
	return r.NomineeAccount
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddNewPlanManager() *model.PartyIdentification2Choice {
	r.NewPlanManager = new(model.PartyIdentification2Choice)
	return r.NewPlanManager
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddProductTransfer() *model.PEPISATransfer5 {
	newValue := new(model.PEPISATransfer5)
	r.ProductTransfer = append(r.ProductTransfer, newValue)
	return newValue
}

func (r *RequestForPEPOrISAOrPortfolioInformationV01) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}
