package model

// Card payment transaction choice between cancellation, authorisation request and authorisation response.
type CardPaymentDataSetTransaction3Choice struct {

	// Completed card payment transaction to be captured.
	Completion *CardPaymentDataSetTransaction10 `xml:"Cmpltn"`

	// Cancelled card payment transaction to be captured.
	Cancellation *CardPaymentDataSetTransaction11 `xml:"Cxl"`

	// Card payment transaction including an authorisation request.
	AuthorisationRequest *CardPaymentDataSetTransaction12 `xml:"AuthstnReq"`

	// Card payment transaction including an authorisation response.
	AuthorisationResponse *CardPaymentDataSetTransaction13 `xml:"AuthstnRspn"`
}

func (c *CardPaymentDataSetTransaction3Choice) AddCompletion() *CardPaymentDataSetTransaction10 {
	c.Completion = new(CardPaymentDataSetTransaction10)
	return c.Completion
}

func (c *CardPaymentDataSetTransaction3Choice) AddCancellation() *CardPaymentDataSetTransaction11 {
	c.Cancellation = new(CardPaymentDataSetTransaction11)
	return c.Cancellation
}

func (c *CardPaymentDataSetTransaction3Choice) AddAuthorisationRequest() *CardPaymentDataSetTransaction12 {
	c.AuthorisationRequest = new(CardPaymentDataSetTransaction12)
	return c.AuthorisationRequest
}

func (c *CardPaymentDataSetTransaction3Choice) AddAuthorisationResponse() *CardPaymentDataSetTransaction13 {
	c.AuthorisationResponse = new(CardPaymentDataSetTransaction13)
	return c.AuthorisationResponse
}
