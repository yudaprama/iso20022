package model

// Unique identifier of a document, message or transaction.
type IdentificationReference11Choice struct {

	// Unambiguous identification of the confirmation transaction as known by the instructing party.
	InstructingPartyTransactionIdentification *Max35Text `xml:"InstgPtyTxId"`

	// Unambiguous identification of the confirmation transaction as known by the executing party.
	ExecutingPartyTransactionIdentification *Max35Text `xml:"ExctgPtyTxId"`

	// Unambiguous identification of the confirmation transaction as known by the market infrastructure.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId"`

	// It permits order originators to tie together groups of orders in which trades resulting from orders are associated for a specific purpose, for example the calculation of average execution price for a customer.
	ClientOrderLinkIdentification *Max35Text `xml:"ClntOrdrLkId"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId"`

	// Reference of the linked message at the trade/block level which identifies a centrally matched transaction.
	BlockIdentification *Max35Text `xml:"BlckId"`

	// Reference of the linked message at the allocation level which identifies a centrally matched transaction.
	AllocationIdentification *Max35Text `xml:"AllcnId"`

	// Reference of the linked message at the individual allocation level which identifies a centrally matched transaction.
	IndividualAllocationIdentification *Max35Text `xml:"IndvAllcnId"`

	// Reference that can be shared across a number of allocation instruction or allocation report messages, thereby making it possible to pass an identifier for an original allocation message on multiple messages (for example from one party to a second to a third, across cancel and replace messages etc).
	SecondaryAllocationIdentification *Max35Text `xml:"ScndryAllcnId"`

	// Reference identifying a index trade.
	IndexIdentification *Max35Text `xml:"IndxId"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId"`

	// Identification used to represent this transaction for compliance purposes.
	ComplianceIdentification *Max35Text `xml:"CmplcId"`

	// Unambiguous identification of the cancellation request as known by the instructing party.
	CancellationRequestIdentification *Max35Text `xml:"CxlReqId"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification *Max35Text `xml:"CollTxId"`
}

func (i *IdentificationReference11Choice) SetInstructingPartyTransactionIdentification(value string) {
	i.InstructingPartyTransactionIdentification = (*Max35Text)(&value)
}

func (i *IdentificationReference11Choice) SetExecutingPartyTransactionIdentification(value string) {
	i.ExecutingPartyTransactionIdentification = (*Max35Text)(&value)
}

func (i *IdentificationReference11Choice) SetMarketInfrastructureTransactionIdentification(value string) {
	i.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (i *IdentificationReference11Choice) SetClientOrderLinkIdentification(value string) {
	i.ClientOrderLinkIdentification = (*Max35Text)(&value)
}

func (i *IdentificationReference11Choice) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *IdentificationReference11Choice) SetBlockIdentification(value string) {
	i.BlockIdentification = (*Max35Text)(&value)
}

func (i *IdentificationReference11Choice) SetAllocationIdentification(value string) {
	i.AllocationIdentification = (*Max35Text)(&value)
}

func (i *IdentificationReference11Choice) SetIndividualAllocationIdentification(value string) {
	i.IndividualAllocationIdentification = (*Max35Text)(&value)
}

func (i *IdentificationReference11Choice) SetSecondaryAllocationIdentification(value string) {
	i.SecondaryAllocationIdentification = (*Max35Text)(&value)
}

func (i *IdentificationReference11Choice) SetIndexIdentification(value string) {
	i.IndexIdentification = (*Max35Text)(&value)
}

func (i *IdentificationReference11Choice) SetCommonIdentification(value string) {
	i.CommonIdentification = (*Max35Text)(&value)
}

func (i *IdentificationReference11Choice) SetComplianceIdentification(value string) {
	i.ComplianceIdentification = (*Max35Text)(&value)
}

func (i *IdentificationReference11Choice) SetCancellationRequestIdentification(value string) {
	i.CancellationRequestIdentification = (*Max35Text)(&value)
}

func (i *IdentificationReference11Choice) SetCollateralTransactionIdentification(value string) {
	i.CollateralTransactionIdentification = (*Max35Text)(&value)
}
