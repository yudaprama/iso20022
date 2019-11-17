package model

// Specifies the choice of the two leg transaction type.
type TwoLegTransactionType1Choice struct {

	// Parameters for contracts which obligate the buyer to receive and the seller to deliver in the future the assets specified at an agreed price or contracts which grant to the holder either the privilege to purchase or the privilege to sell the assets specified at a predetermined price or formula at or within a time in the future.
	FutureOrOptionDetails *FutureOrOptionDetails1 `xml:"FutrOrOptnDtls"`

	// Provides details about the two leg transaction.
	SecuritiesFinancingDetails *SecuritiesFinancing10 `xml:"SctiesFincgDtls"`
}

func (t *TwoLegTransactionType1Choice) AddFutureOrOptionDetails() *FutureOrOptionDetails1 {
	t.FutureOrOptionDetails = new(FutureOrOptionDetails1)
	return t.FutureOrOptionDetails
}

func (t *TwoLegTransactionType1Choice) AddSecuritiesFinancingDetails() *SecuritiesFinancing10 {
	t.SecuritiesFinancingDetails = new(SecuritiesFinancing10)
	return t.SecuritiesFinancingDetails
}
