package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200107 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.012.001.07 Document"`
	Message *PortfolioTransferInstructionV07 `xml:"PrtflTrfInstr"`
}

func (d *Document01200107) AddMessage() *PortfolioTransferInstructionV07 {
	d.Message = new(PortfolioTransferInstructionV07)
	return d.Message
}

// Scope
// An instructing party, for example, a (new) plan manager (Transferee), sends the PortfolioTransferInstruction message to the executing party, for example, a (old) plan manager (Transferor), on behalf of the initiating party, for example, an investor (client), to instruct the transfer of financial instruments from the clients account at the old plan manager (Transferor) to the clients account at the new plan manager (Transferee) through a nominee account.
// Usage
// The PortfolioTransferInstruction message is used to instruct the withdrawal of one or more ISA or portfolio products from one account and deliver them to another account.
// The PortfolioTransferInstruction message is used to instruct one or more transfers for one client. Each transfer is for delivery to the same account. The account may be owned by one or more individual investors or one or more corporate investors. Each transfer is identified in TransferIdentification.
// If the instructing party does not have enough information to instruct the transfer, then it must first send a AccountHoldingInformationRequest message to the executing party in order to receive a AccountHoldingInformation message.
type PortfolioTransferInstructionV07 struct {

	// Identifies the message.
	MessageReference *model.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference6 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference6 `xml:"RltdRef,omitempty"`

	// Information identifying the primary individual investor, for example, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *model.IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, for example, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *model.IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying the other individual investors, for example, name, address, social security number and date of birth.
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

	// Identification of a related party or intermediary.
	IntermediaryInformation []*model.Intermediary34 `xml:"IntrmyInf,omitempty"`

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *model.CashAccount34 `xml:"CshAcct,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*model.ISATransfer22 `xml:"PdctTrf"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PortfolioTransferInstructionV07) AddMessageReference() *model.MessageIdentification1 {
	p.MessageReference = new(model.MessageIdentification1)
	return p.MessageReference
}

func (p *PortfolioTransferInstructionV07) AddPoolReference() *model.AdditionalReference6 {
	p.PoolReference = new(model.AdditionalReference6)
	return p.PoolReference
}

func (p *PortfolioTransferInstructionV07) AddPreviousReference() *model.AdditionalReference6 {
	p.PreviousReference = new(model.AdditionalReference6)
	return p.PreviousReference
}

func (p *PortfolioTransferInstructionV07) AddRelatedReference() *model.AdditionalReference6 {
	p.RelatedReference = new(model.AdditionalReference6)
	return p.RelatedReference
}

func (p *PortfolioTransferInstructionV07) AddPrimaryIndividualInvestor() *model.IndividualPerson8 {
	p.PrimaryIndividualInvestor = new(model.IndividualPerson8)
	return p.PrimaryIndividualInvestor
}

func (p *PortfolioTransferInstructionV07) AddSecondaryIndividualInvestor() *model.IndividualPerson8 {
	p.SecondaryIndividualInvestor = new(model.IndividualPerson8)
	return p.SecondaryIndividualInvestor
}

func (p *PortfolioTransferInstructionV07) AddOtherIndividualInvestor() *model.IndividualPerson8 {
	newValue := new(model.IndividualPerson8)
	p.OtherIndividualInvestor = append(p.OtherIndividualInvestor, newValue)
	return newValue
}

func (p *PortfolioTransferInstructionV07) AddPrimaryCorporateInvestor() *model.Organisation21 {
	p.PrimaryCorporateInvestor = new(model.Organisation21)
	return p.PrimaryCorporateInvestor
}

func (p *PortfolioTransferInstructionV07) AddSecondaryCorporateInvestor() *model.Organisation21 {
	p.SecondaryCorporateInvestor = new(model.Organisation21)
	return p.SecondaryCorporateInvestor
}

func (p *PortfolioTransferInstructionV07) AddOtherCorporateInvestor() *model.Organisation21 {
	newValue := new(model.Organisation21)
	p.OtherCorporateInvestor = append(p.OtherCorporateInvestor, newValue)
	return newValue
}

func (p *PortfolioTransferInstructionV07) AddTransferorAccount() *model.Account19 {
	p.TransferorAccount = new(model.Account19)
	return p.TransferorAccount
}

func (p *PortfolioTransferInstructionV07) AddNomineeAccount() *model.Account19 {
	p.NomineeAccount = new(model.Account19)
	return p.NomineeAccount
}

func (p *PortfolioTransferInstructionV07) AddTransferee() *model.PartyIdentification70Choice {
	p.Transferee = new(model.PartyIdentification70Choice)
	return p.Transferee
}

func (p *PortfolioTransferInstructionV07) AddIntermediaryInformation() *model.Intermediary34 {
	newValue := new(model.Intermediary34)
	p.IntermediaryInformation = append(p.IntermediaryInformation, newValue)
	return newValue
}

func (p *PortfolioTransferInstructionV07) AddCashAccount() *model.CashAccount34 {
	p.CashAccount = new(model.CashAccount34)
	return p.CashAccount
}

func (p *PortfolioTransferInstructionV07) AddProductTransfer() *model.ISATransfer22 {
	newValue := new(model.ISATransfer22)
	p.ProductTransfer = append(p.ProductTransfer, newValue)
	return newValue
}

func (p *PortfolioTransferInstructionV07) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	p.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return p.MarketPracticeVersion
}

func (p *PortfolioTransferInstructionV07) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}
