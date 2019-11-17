package model

// Provide information about the reason for the modification and about the application request which triggered this modification.
type InvestmentAccountModificationDetails struct {

	// Reason for the modification brought to the investment account information.
	ModificationReason *Max350Text `xml:"ModRsn,omitempty"`

	// Unique and unambiguous identifier of the account modification request.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`
}

func (i *InvestmentAccountModificationDetails) SetModificationReason(value string) {
	i.ModificationReason = (*Max350Text)(&value)
}

func (i *InvestmentAccountModificationDetails) SetAccountApplicationIdentification(value string) {
	i.AccountApplicationIdentification = (*Max35Text)(&value)
}
