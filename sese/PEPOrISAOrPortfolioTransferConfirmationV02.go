package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300102 struct {
	XMLName xml.Name                                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.02 Document"`
	Message *PEPOrISAOrPortfolioTransferConfirmationV02 `xml:"PEPOrISAOrPrtflTrfConfV02"`
}

func (d *Document01300102) AddMessage() *PEPOrISAOrPortfolioTransferConfirmationV02 {
	d.Message = new(PEPOrISAOrPortfolioTransferConfirmationV02)
	return d.Message
}

// Scope
// An executing party, eg, a (old) plan manager, sends the PEPOrISAOrPortfolioTransferConfirmation message to the instructing party, eg, a (new) plan manager, to confirm the transfer of one or more PEP or ISA or portfolio products from the client's account at the old plan manager to the client's account at the new plan manager through a nominee account.
// Usage
// The PEPOrISAOrPortfolioTransferConfirmation message is used to confirm the transfer of one or more PEP or ISA or portfolio products.
// The reference of each product transfer confirmation is identified in TransferConfirmationIdentification. The reference of the original product transfer is specified in TransferInstructionReference. The message identification of the PEPOrISAPOrPortfolioTransferInstruction message in which the product transfers were conveyed may also be quoted in RelatedReference.
type PEPOrISAOrPortfolioTransferConfirmationV02 struct {

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

	// Information identifying the other individual investors, eg, name, address, social security number and date of birth.
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

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *model.CashAccount11 `xml:"CshAcct,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*model.PEPISATransfer4 `xml:"PdctTrf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddMessageReference() *model.MessageIdentification1 {
	p.MessageReference = new(model.MessageIdentification1)
	return p.MessageReference
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddPoolReference() *model.AdditionalReference3 {
	p.PoolReference = new(model.AdditionalReference3)
	return p.PoolReference
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddPreviousReference() *model.AdditionalReference3 {
	p.PreviousReference = new(model.AdditionalReference3)
	return p.PreviousReference
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddRelatedReference() *model.AdditionalReference3 {
	p.RelatedReference = new(model.AdditionalReference3)
	return p.RelatedReference
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddPrimaryIndividualInvestor() *model.IndividualPerson8 {
	p.PrimaryIndividualInvestor = new(model.IndividualPerson8)
	return p.PrimaryIndividualInvestor
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddSecondaryIndividualInvestor() *model.IndividualPerson8 {
	p.SecondaryIndividualInvestor = new(model.IndividualPerson8)
	return p.SecondaryIndividualInvestor
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddOtherIndividualInvestor() *model.IndividualPerson8 {
	newValue := new(model.IndividualPerson8)
	p.OtherIndividualInvestor = append(p.OtherIndividualInvestor, newValue)
	return newValue
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddPrimaryCorporateInvestor() *model.Organisation4 {
	p.PrimaryCorporateInvestor = new(model.Organisation4)
	return p.PrimaryCorporateInvestor
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddSecondaryCorporateInvestor() *model.Organisation4 {
	p.SecondaryCorporateInvestor = new(model.Organisation4)
	return p.SecondaryCorporateInvestor
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddOtherCorporateInvestor() *model.Organisation4 {
	newValue := new(model.Organisation4)
	p.OtherCorporateInvestor = append(p.OtherCorporateInvestor, newValue)
	return newValue
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddClientAccount() *model.Account5 {
	p.ClientAccount = new(model.Account5)
	return p.ClientAccount
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddNomineeAccount() *model.Account6 {
	p.NomineeAccount = new(model.Account6)
	return p.NomineeAccount
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddNewPlanManager() *model.PartyIdentification2Choice {
	p.NewPlanManager = new(model.PartyIdentification2Choice)
	return p.NewPlanManager
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddCashAccount() *model.CashAccount11 {
	p.CashAccount = new(model.CashAccount11)
	return p.CashAccount
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddProductTransfer() *model.PEPISATransfer4 {
	newValue := new(model.PEPISATransfer4)
	p.ProductTransfer = append(p.ProductTransfer, newValue)
	return newValue
}

func (p *PEPOrISAOrPortfolioTransferConfirmationV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}
