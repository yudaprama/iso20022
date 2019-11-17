package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300106 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.013.001.06 Document"`
	Message *PortfolioTransferConfirmationV06 `xml:"PrtflTrfConf"`
}

func (d *Document01300106) AddMessage() *PortfolioTransferConfirmationV06 {
	d.Message = new(PortfolioTransferConfirmationV06)
	return d.Message
}

// Scope
// An executing party, for example, a (old) plan manager (Transferor), sends the PortfolioTransferConfirmation message to the instructing party, for example, a (new) plan manager (Transferee), to confirm the transfer of one or more ISA or portfolio products from the client's account at the old plan manager (Transferor) to the client's account at the new plan manager (Transferee) through a nominee account.
// Usage
// The PortfolioTransferConfirmation message is used to confirm the transfer of one or more ISA or portfolio products.
// The reference of each product transfer confirmation is identified in TransferConfirmationIdentification. The reference of the original product transfer is specified in TransferInstructionReference. The message identification of the PortfolioTransferInstruction message in which the product transfers were conveyed may also be quoted in RelatedReference.
type PortfolioTransferConfirmationV06 struct {

	// Identifies the message.
	MessageReference *model.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Information identifying the primary individual investor, for example, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *model.IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, for example, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *model.IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying the other individual investors, for example, name, address, social security number and date of birth.
	OtherIndividualInvestor []*model.IndividualPerson8 `xml:"OthrIndvInvstr,omitempty"`

	// Information identifying the primary corporate investor, for example, name and address.
	PrimaryCorporateInvestor *model.Organisation4 `xml:"PmryCorpInvstr,omitempty"`

	// Information identifying the secondary corporate investor, for example, name and address.
	SecondaryCorporateInvestor *model.Organisation4 `xml:"ScndryCorpInvstr,omitempty"`

	// Information identifying the other corporate investors, for example, name and address.
	OtherCorporateInvestor []*model.Organisation4 `xml:"OthrCorpInvstr,omitempty"`

	// Identification of an account owned by the investor at the old plan manager (account servicer).
	TransferorAccount *model.Account15 `xml:"TrfrAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *model.Account16 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	Transferee *model.PartyIdentification2Choice `xml:"Trfee"`

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *model.CashAccount29 `xml:"CshAcct,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*model.ISATransfer21 `xml:"PdctTrf"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PortfolioTransferConfirmationV06) AddMessageReference() *model.MessageIdentification1 {
	p.MessageReference = new(model.MessageIdentification1)
	return p.MessageReference
}

func (p *PortfolioTransferConfirmationV06) AddPoolReference() *model.AdditionalReference3 {
	p.PoolReference = new(model.AdditionalReference3)
	return p.PoolReference
}

func (p *PortfolioTransferConfirmationV06) AddPreviousReference() *model.AdditionalReference3 {
	p.PreviousReference = new(model.AdditionalReference3)
	return p.PreviousReference
}

func (p *PortfolioTransferConfirmationV06) AddRelatedReference() *model.AdditionalReference3 {
	p.RelatedReference = new(model.AdditionalReference3)
	return p.RelatedReference
}

func (p *PortfolioTransferConfirmationV06) AddPrimaryIndividualInvestor() *model.IndividualPerson8 {
	p.PrimaryIndividualInvestor = new(model.IndividualPerson8)
	return p.PrimaryIndividualInvestor
}

func (p *PortfolioTransferConfirmationV06) AddSecondaryIndividualInvestor() *model.IndividualPerson8 {
	p.SecondaryIndividualInvestor = new(model.IndividualPerson8)
	return p.SecondaryIndividualInvestor
}

func (p *PortfolioTransferConfirmationV06) AddOtherIndividualInvestor() *model.IndividualPerson8 {
	newValue := new(model.IndividualPerson8)
	p.OtherIndividualInvestor = append(p.OtherIndividualInvestor, newValue)
	return newValue
}

func (p *PortfolioTransferConfirmationV06) AddPrimaryCorporateInvestor() *model.Organisation4 {
	p.PrimaryCorporateInvestor = new(model.Organisation4)
	return p.PrimaryCorporateInvestor
}

func (p *PortfolioTransferConfirmationV06) AddSecondaryCorporateInvestor() *model.Organisation4 {
	p.SecondaryCorporateInvestor = new(model.Organisation4)
	return p.SecondaryCorporateInvestor
}

func (p *PortfolioTransferConfirmationV06) AddOtherCorporateInvestor() *model.Organisation4 {
	newValue := new(model.Organisation4)
	p.OtherCorporateInvestor = append(p.OtherCorporateInvestor, newValue)
	return newValue
}

func (p *PortfolioTransferConfirmationV06) AddTransferorAccount() *model.Account15 {
	p.TransferorAccount = new(model.Account15)
	return p.TransferorAccount
}

func (p *PortfolioTransferConfirmationV06) AddNomineeAccount() *model.Account16 {
	p.NomineeAccount = new(model.Account16)
	return p.NomineeAccount
}

func (p *PortfolioTransferConfirmationV06) AddTransferee() *model.PartyIdentification2Choice {
	p.Transferee = new(model.PartyIdentification2Choice)
	return p.Transferee
}

func (p *PortfolioTransferConfirmationV06) AddCashAccount() *model.CashAccount29 {
	p.CashAccount = new(model.CashAccount29)
	return p.CashAccount
}

func (p *PortfolioTransferConfirmationV06) AddProductTransfer() *model.ISATransfer21 {
	newValue := new(model.ISATransfer21)
	p.ProductTransfer = append(p.ProductTransfer, newValue)
	return newValue
}

func (p *PortfolioTransferConfirmationV06) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	p.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return p.MarketPracticeVersion
}

func (p *PortfolioTransferConfirmationV06) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}
