package model

// Card payment transaction choice between cancellation, authorisation request and authorisation response.
type CardPaymentDataSetTransaction4Choice struct {

	// Completed card payment transaction to be captured.
	Completion *CardPaymentDataSetTransaction14 `xml:"Cmpltn"`

	// Cancelled card payment transaction to be captured.
	Cancellation *CardPaymentDataSetTransaction15 `xml:"Cxl"`

	// Card payment transaction including an authorisation request.
	AuthorisationRequest *CardPaymentDataSetTransaction16 `xml:"AuthstnReq"`

	// Card payment transaction including an authorisation response.
	AuthorisationResponse *CardPaymentDataSetTransaction17 `xml:"AuthstnRspn"`
}

func (c *CardPaymentDataSetTransaction4Choice) AddCompletion() *CardPaymentDataSetTransaction14 {
	c.Completion = new(CardPaymentDataSetTransaction14)
	return c.Completion
}

func (c *CardPaymentDataSetTransaction4Choice) AddCancellation() *CardPaymentDataSetTransaction15 {
	c.Cancellation = new(CardPaymentDataSetTransaction15)
	return c.Cancellation
}

func (c *CardPaymentDataSetTransaction4Choice) AddAuthorisationRequest() *CardPaymentDataSetTransaction16 {
	c.AuthorisationRequest = new(CardPaymentDataSetTransaction16)
	return c.AuthorisationRequest
}

func (c *CardPaymentDataSetTransaction4Choice) AddAuthorisationResponse() *CardPaymentDataSetTransaction17 {
	c.AuthorisationResponse = new(CardPaymentDataSetTransaction17)
	return c.AuthorisationResponse
}
