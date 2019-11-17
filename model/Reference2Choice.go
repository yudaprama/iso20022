package model

// Choice between the identification of the MarginCallRequest message, the MarginCallResponse message, the CollateralProposal message, the CollateralProposalResponse message, the CollateralSubstitutionRequest message, the CollateralSubstitutionResponse message, the CollateralSubstitutionConfirmation message, the InterestPaymentRequest message, the InterestPaymentResponse message, the InterestPaymentStatement message or the MarginCallDisputeNotification message.
type Reference2Choice struct {

	// Unique identifier of the margin call request.
	MarginCallRequestIdentification *Max35Text `xml:"MrgnCallReqId"`

	// Unique identifier of the margin call response.
	MarginCallResponseIdentification *Max35Text `xml:"MrgnCallRspnId"`

	// Unique identifier of the collateral proposal.
	CollateralProposalIdentification *Max35Text `xml:"CollPrpslId"`

	// Unique identifier of the collateral proposal response.
	CollateralProposalResponseIdentification *Max35Text `xml:"CollPrpslRspnId"`

	// Unique identifier of the dispute notification.
	DisputeNotificationIdentification *Max35Text `xml:"DsptNtfctnId"`

	// Choice between the identification of the MarginCallRequest message, the MarginCallResponse message, the CollateralProposal message, the CollateralProposalResponse message, the CollateralSubstitutionRequest message, the CollateralSubstitutionResponse message, the CollateralSubstitutionConfirmation message, the InterestPaymentRequest message, the InterestPaymentResponse message, the InterestPaymentStatement message or the MarginCallDisputeNotification message.
	CollateralSubstitutionRequestIdentification *Max35Text `xml:"CollSbstitnReqId"`

	// Unique identifier of the collateral substitution response.
	CollateralSubstitutionResponseIdentification *Max35Text `xml:"CollSbstitnRspnId"`

	// Unique identifier of the collateral substitution confirmation.
	CollateralSubstitutionConfirmationIdentification *Max35Text `xml:"CollSbstitnConfId"`

	// Unique identifier of the interest payment request.
	InterestPaymentRequestIdentification *Max35Text `xml:"IntrstPmtReqId"`

	// Unique identifier of the interest payment response.
	InterestPaymentResponseIdentification *Max35Text `xml:"IntrstPmtRspnId"`

	// Unique identifier of the interest payment statement.
	InterestPaymentStatementIdentification *Max35Text `xml:"IntrstPmtStmtId"`
}

func (r *Reference2Choice) SetMarginCallRequestIdentification(value string) {
	r.MarginCallRequestIdentification = (*Max35Text)(&value)
}

func (r *Reference2Choice) SetMarginCallResponseIdentification(value string) {
	r.MarginCallResponseIdentification = (*Max35Text)(&value)
}

func (r *Reference2Choice) SetCollateralProposalIdentification(value string) {
	r.CollateralProposalIdentification = (*Max35Text)(&value)
}

func (r *Reference2Choice) SetCollateralProposalResponseIdentification(value string) {
	r.CollateralProposalResponseIdentification = (*Max35Text)(&value)
}

func (r *Reference2Choice) SetDisputeNotificationIdentification(value string) {
	r.DisputeNotificationIdentification = (*Max35Text)(&value)
}

func (r *Reference2Choice) SetCollateralSubstitutionRequestIdentification(value string) {
	r.CollateralSubstitutionRequestIdentification = (*Max35Text)(&value)
}

func (r *Reference2Choice) SetCollateralSubstitutionResponseIdentification(value string) {
	r.CollateralSubstitutionResponseIdentification = (*Max35Text)(&value)
}

func (r *Reference2Choice) SetCollateralSubstitutionConfirmationIdentification(value string) {
	r.CollateralSubstitutionConfirmationIdentification = (*Max35Text)(&value)
}

func (r *Reference2Choice) SetInterestPaymentRequestIdentification(value string) {
	r.InterestPaymentRequestIdentification = (*Max35Text)(&value)
}

func (r *Reference2Choice) SetInterestPaymentResponseIdentification(value string) {
	r.InterestPaymentResponseIdentification = (*Max35Text)(&value)
}

func (r *Reference2Choice) SetInterestPaymentStatementIdentification(value string) {
	r.InterestPaymentStatementIdentification = (*Max35Text)(&value)
}
