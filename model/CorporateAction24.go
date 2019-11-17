package model

// Provides information about the corporate action event.
type CorporateAction24 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate41 `xml:"DtDtls,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage *CorporateActionEventStageFormat6Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat1Choice `xml:"LtryTp,omitempty"`
}

func (c *CorporateAction24) AddDateDetails() *CorporateActionDate41 {
	c.DateDetails = new(CorporateActionDate41)
	return c.DateDetails
}

func (c *CorporateAction24) AddEventStage() *CorporateActionEventStageFormat6Choice {
	c.EventStage = new(CorporateActionEventStageFormat6Choice)
	return c.EventStage
}

func (c *CorporateAction24) AddLotteryType() *LotteryTypeFormat1Choice {
	c.LotteryType = new(LotteryTypeFormat1Choice)
	return c.LotteryType
}
