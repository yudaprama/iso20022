package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200102 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.012.001.02 Document"`
	Message *PEPOrISAOrPortfolioTransferInstructionV02 `xml:"PEPOrISAOrPrtflTrfInstrV02"`
}

func (d *Document01200102) AddMessage() *PEPOrISAOrPortfolioTransferInstructionV02 {
	d.Message = new(PEPOrISAOrPortfolioTransferInstructionV02)
	return d.Message
}

// Scope
// An instructing party, eg, a (new) plan manager, sends the PEPOrISAOrPortfolioTransferInstruction message to the executing party, eg, a (old) plan manager, on behalf of the initiating party, eg, an investor (client), to instruct the transfer of financial instruments from the clients account at the old plan manager to the clients account at the new plan manager through a nominee account.
// Usage
// The PEPOrISAOrPortfolioTransferInstruction message is used to instruct the withdrawal of one or more PEP or ISA or portfolio products from one account and deliver them to another account.
// The PEPOrISAOrPortfolioTransferInstruction message is used to instruct one or more transfers for one client. Each transfer is for delivery to the same account. The account may be owned by one or more individual investors or one or more corporate investors. Each transfer is identified in TransferIdentification.
// If the instructing party does not have enough information to instruct the transfer, then it must first send a RequestForPEPOrISAOrPortfolioInformation message to the executing party in order to receive a PEPOrISAOrPortfolioInformation message.
type PEPOrISAOrPortfolioTransferInstructionV02 struct {

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
	ProductTransfer []*model.PEPISATransfer3 `xml:"PdctTrf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddMessageReference() *model.MessageIdentification1 {
	p.MessageReference = new(model.MessageIdentification1)
	return p.MessageReference
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddPoolReference() *model.AdditionalReference3 {
	p.PoolReference = new(model.AdditionalReference3)
	return p.PoolReference
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddPreviousReference() *model.AdditionalReference3 {
	p.PreviousReference = new(model.AdditionalReference3)
	return p.PreviousReference
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddRelatedReference() *model.AdditionalReference3 {
	p.RelatedReference = new(model.AdditionalReference3)
	return p.RelatedReference
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddPrimaryIndividualInvestor() *model.IndividualPerson8 {
	p.PrimaryIndividualInvestor = new(model.IndividualPerson8)
	return p.PrimaryIndividualInvestor
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddSecondaryIndividualInvestor() *model.IndividualPerson8 {
	p.SecondaryIndividualInvestor = new(model.IndividualPerson8)
	return p.SecondaryIndividualInvestor
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddOtherIndividualInvestor() *model.IndividualPerson8 {
	newValue := new(model.IndividualPerson8)
	p.OtherIndividualInvestor = append(p.OtherIndividualInvestor, newValue)
	return newValue
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddPrimaryCorporateInvestor() *model.Organisation4 {
	p.PrimaryCorporateInvestor = new(model.Organisation4)
	return p.PrimaryCorporateInvestor
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddSecondaryCorporateInvestor() *model.Organisation4 {
	p.SecondaryCorporateInvestor = new(model.Organisation4)
	return p.SecondaryCorporateInvestor
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddOtherCorporateInvestor() *model.Organisation4 {
	newValue := new(model.Organisation4)
	p.OtherCorporateInvestor = append(p.OtherCorporateInvestor, newValue)
	return newValue
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddClientAccount() *model.Account5 {
	p.ClientAccount = new(model.Account5)
	return p.ClientAccount
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddNomineeAccount() *model.Account6 {
	p.NomineeAccount = new(model.Account6)
	return p.NomineeAccount
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddNewPlanManager() *model.PartyIdentification2Choice {
	p.NewPlanManager = new(model.PartyIdentification2Choice)
	return p.NewPlanManager
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddCashAccount() *model.CashAccount11 {
	p.CashAccount = new(model.CashAccount11)
	return p.CashAccount
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddProductTransfer() *model.PEPISATransfer3 {
	newValue := new(model.PEPISATransfer3)
	p.ProductTransfer = append(p.ProductTransfer, newValue)
	return newValue
}

func (p *PEPOrISAOrPortfolioTransferInstructionV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}
