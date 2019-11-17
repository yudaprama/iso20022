package model

// Card payment transaction choice between cancellation, authorisation request and authorisation response.
type CardPaymentDataSetTransaction2Choice struct {

	// Completed card payment transaction to be captured.
	Completion *CardPaymentDataSetTransaction6 `xml:"Cmpltn"`

	// Cancelled card payment transaction to be captured.
	Cancellation *CardPaymentDataSetTransaction7 `xml:"Cxl"`

	// Card payment transaction including an authorisation request.
	AuthorisationRequest *CardPaymentDataSetTransaction8 `xml:"AuthstnReq"`

	// Card payment transaction including an authorisation response.
	AuthorisationResponse *CardPaymentDataSetTransaction9 `xml:"AuthstnRspn"`
}

func (c *CardPaymentDataSetTransaction2Choice) AddCompletion() *CardPaymentDataSetTransaction6 {
	c.Completion = new(CardPaymentDataSetTransaction6)
	return c.Completion
}

func (c *CardPaymentDataSetTransaction2Choice) AddCancellation() *CardPaymentDataSetTransaction7 {
	c.Cancellation = new(CardPaymentDataSetTransaction7)
	return c.Cancellation
}

func (c *CardPaymentDataSetTransaction2Choice) AddAuthorisationRequest() *CardPaymentDataSetTransaction8 {
	c.AuthorisationRequest = new(CardPaymentDataSetTransaction8)
	return c.AuthorisationRequest
}

func (c *CardPaymentDataSetTransaction2Choice) AddAuthorisationResponse() *CardPaymentDataSetTransaction9 {
	c.AuthorisationResponse = new(CardPaymentDataSetTransaction9)
	return c.AuthorisationResponse
}
