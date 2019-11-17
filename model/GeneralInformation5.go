package model

// Information concerning the negotiation process leading to a treasury trade.
type GeneralInformation5 struct {

	// Indicates whether the trade is a block or single trade.
	BlockIndicator *YesNoIndicator `xml:"BlckInd,omitempty"`

	// Reference to a preceding transaction, for example, an option or swap.
	RelatedTradeReference *Max35Text `xml:"RltdTradRef,omitempty"`

	// Method used by the trading parties to negotiate and/or execute a deal.
	DealingMethod *Trading1MethodCode `xml:"DealgMtd,omitempty"`

	// Specifies the broker which arranged the deal between the trading side and the counterparty side or, when two money brokers are involved, between the trading side and the other money broker.
	BrokerIdentification *PartyIdentification73Choice `xml:"BrkrId,omitempty"`

	// Counterparty's reference for the trade.
	CounterpartyReference *Max35Text `xml:"CtrPtyRef,omitempty"`

	// Brokerage fee for a broker confirmation.
	BrokersCommission *ActiveCurrencyAndAmount `xml:"BrkrsComssn,omitempty"`

	// Specifies additional information for the receiver and applies to the whole message.
	SenderToReceiverInformation *Max210Text `xml:"SndrToRcvrInf,omitempty"`

	// Specifies the branch at the trading side with which the deal was done.
	DealingBranchTradingSide *PartyIdentification73Choice `xml:"DealgBrnchTradgSd,omitempty"`

	// Specifies the branch at the counterparty side with which the deal was done.
	DealingBranchCounterpartySide *PartyIdentification73Choice `xml:"DealgBrnchCtrPtySd,omitempty"`

	// Specifies the name and/or electronic address of the receiver of the message who may be contacted for any queries concerning this trade.
	ContactInformation *ContactInformation1 `xml:"CtctInf,omitempty"`

	// Specifies the type, date and version of the agreement used in a trade.
	AgreementDetails *AgreementConditions1 `xml:"AgrmtDtls,omitempty"`

	// Specifies the year of definitions of the agreement.
	DefinitionsYear *ISOYear `xml:"DefsYr,omitempty"`

	// Specifies a reference applied to the trade instruction by a broker.
	BrokersReference *Max35Text `xml:"BrkrsRef,omitempty"`
}

func (g *GeneralInformation5) SetBlockIndicator(value string) {
	g.BlockIndicator = (*YesNoIndicator)(&value)
}

func (g *GeneralInformation5) SetRelatedTradeReference(value string) {
	g.RelatedTradeReference = (*Max35Text)(&value)
}

func (g *GeneralInformation5) SetDealingMethod(value string) {
	g.DealingMethod = (*Trading1MethodCode)(&value)
}

func (g *GeneralInformation5) AddBrokerIdentification() *PartyIdentification73Choice {
	g.BrokerIdentification = new(PartyIdentification73Choice)
	return g.BrokerIdentification
}

func (g *GeneralInformation5) SetCounterpartyReference(value string) {
	g.CounterpartyReference = (*Max35Text)(&value)
}

func (g *GeneralInformation5) SetBrokersCommission(value, currency string) {
	g.BrokersCommission = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GeneralInformation5) SetSenderToReceiverInformation(value string) {
	g.SenderToReceiverInformation = (*Max210Text)(&value)
}

func (g *GeneralInformation5) AddDealingBranchTradingSide() *PartyIdentification73Choice {
	g.DealingBranchTradingSide = new(PartyIdentification73Choice)
	return g.DealingBranchTradingSide
}

func (g *GeneralInformation5) AddDealingBranchCounterpartySide() *PartyIdentification73Choice {
	g.DealingBranchCounterpartySide = new(PartyIdentification73Choice)
	return g.DealingBranchCounterpartySide
}

func (g *GeneralInformation5) AddContactInformation() *ContactInformation1 {
	g.ContactInformation = new(ContactInformation1)
	return g.ContactInformation
}

func (g *GeneralInformation5) AddAgreementDetails() *AgreementConditions1 {
	g.AgreementDetails = new(AgreementConditions1)
	return g.AgreementDetails
}

func (g *GeneralInformation5) SetDefinitionsYear(value string) {
	g.DefinitionsYear = (*ISOYear)(&value)
}

func (g *GeneralInformation5) SetBrokersReference(value string) {
	g.BrokersReference = (*Max35Text)(&value)
}
