package model

// Provides the identification of the collateral message cancellation request.
type Reference16 struct {

	// Identification of the collateral message cancellation request.
	CollateralMessageCancellationRequestIdentification *Max35Text `xml:"CollMsgCxlReqId"`
}

func (r *Reference16) SetCollateralMessageCancellationRequestIdentification(value string) {
	r.CollateralMessageCancellationRequestIdentification = (*Max35Text)(&value)
}
