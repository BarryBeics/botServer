// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ActivityReport struct {
	ID               string  `json:"_id,omitempty" bson:"_id,omitempty"`
	Timestamp      int      `json:"Timestamp" bson:"Timestamp"`
	Qty            int      `json:"Qty"`
	AvgGain        float64  `json:"AvgGain"`
	TopAGain       *float64 `json:"TopAGain,omitempty"`
	TopBGain       *float64 `json:"TopBGain,omitempty"`
	TopCGain       *float64 `json:"TopCGain,omitempty"`
	FearGreedIndex int      `json:"FearGreedIndex"`
}

type HistoricPrices struct {
	Pair      []*Pair `json:"Pair,omitempty"`
	Timestamp      int      `json:"Timestamp" bson:"Timestamp"`
}

type MarkAsTestedInput struct {
	BotInstanceName string `json:"BotInstanceName"`
	Tested          bool   `json:"Tested"`
}

type NewActivityReport struct {
	Timestamp      int      `json:"Timestamp" bson:"Timestamp"`
	Qty            int      `json:"Qty"`
	AvgGain        float64  `json:"AvgGain"`
	TopAGain       *float64 `json:"TopAGain,omitempty"`
	TopBGain       *float64 `json:"TopBGain,omitempty"`
	TopCGain       *float64 `json:"TopCGain,omitempty"`
	FearGreedIndex int      `json:"FearGreedIndex"`
}

type NewHistoricPriceInput struct {
	Pairs     []*PairInput `json:"pairs"`
	Timestamp      int      `json:"Timestamp" bson:"Timestamp"`
}

type NewTradeOutcomeReport struct {
	Timestamp      int      `json:"Timestamp" bson:"Timestamp"`
	BotName          string  `json:"BotName"`
	PercentageChange float64 `json:"PercentageChange"`
	Balance          float64 `json:"Balance"`
	Symbol           string  `json:"Symbol"`
	Outcome          string  `json:"Outcome"`
	Fee              *float64 `json:"Fee,omitempty"`
	ElapsedTime      int     `json:"ElapsedTime"`
	Volume           float64 `json:"Volume"`
	FearGreedIndex   int     `json:"FearGreedIndex"`
	MarketStatus     string  `json:"MarketStatus"`
}

type Pair struct {
	Symbol string `json:"Symbol"`
	Price  string `json:"Price"`
}

type PairInput struct {
	Symbol string `json:"Symbol"`
	Price  string `json:"Price"`
}

type Strategy struct {
	BotInstanceName      string   `json:"BotInstanceName"`
	TradeDuration        int      `json:"TradeDuration"`
	IncrementsAtr        int      `json:"IncrementsATR"`
	LongSMADuration      int      `json:"LongSMADuration"`
	ShortSMADuration     int      `json:"ShortSMADuration"`
	WINCounter           *int     `json:"WINCounter,omitempty"`
	LOSSCounter          *int     `json:"LOSSCounter,omitempty"`
	TIMEOUTGainCounter   *int     `json:"TIMEOUTGainCounter,omitempty"`
	TIMEOUTLossCounter   *int     `json:"TIMEOUTLossCounter,omitempty"`
	AccountBalance       float64  `json:"AccountBalance"`
	MovingAveMomentum    float64  `json:"MovingAveMomentum"`
	TakeProfitPercentage *float64 `json:"TakeProfitPercentage,omitempty"`
	StopLossPercentage   *float64 `json:"StopLossPercentage,omitempty"`
	ATRtollerance        *float64 `json:"ATRtollerance,omitempty"`
	FeesTotal            *float64 `json:"FeesTotal,omitempty"`
	Tested               *bool    `json:"Tested,omitempty"`
	Owner                *string  `json:"Owner,omitempty"`
	CreatedOn            int      `json:"CreatedOn"`
}

type StrategyInput struct {
	BotInstanceName      string   `json:"BotInstanceName"`
	TradeDuration        int      `json:"TradeDuration"`
	IncrementsAtr        int      `json:"IncrementsATR"`
	LongSMADuration      int      `json:"LongSMADuration"`
	ShortSMADuration     int      `json:"ShortSMADuration"`
	WINCounter           *int     `json:"WINCounter,omitempty"`
	LOSSCounter          *int     `json:"LOSSCounter,omitempty"`
	TIMEOUTGainCounter   *int     `json:"TIMEOUTGainCounter,omitempty"`
	TIMEOUTLossCounter   *int     `json:"TIMEOUTLossCounter,omitempty"`
	AccountBalance       float64  `json:"AccountBalance"`
	MovingAveMomentum    float64  `json:"MovingAveMomentum"`
	TakeProfitPercentage float64  `json:"TakeProfitPercentage"`
	StopLossPercentage   float64  `json:"StopLossPercentage"`
	ATRtollerance        *float64 `json:"ATRtollerance,omitempty"`
	FeesTotal            *float64 `json:"FeesTotal,omitempty"`
	Tested               *bool    `json:"Tested,omitempty"`
	Owner                string   `json:"Owner"`
	CreatedOn            int      `json:"CreatedOn"`
}

type TradeOutcomeReport struct {
	ID               string  `json:"_id,omitempty" bson:"_id,omitempty"`
	Timestamp      int      `json:"Timestamp" bson:"Timestamp"`
	BotName          string   `json:"BotName"`
	PercentageChange float64  `json:"PercentageChange"`
	Balance          float64  `json:"Balance"`
	Symbol           string   `json:"Symbol"`
	Outcome          string   `json:"Outcome"`
	Fee              *float64 `json:"Fee,omitempty"`
	ElapsedTime      int      `json:"ElapsedTime"`
	Volume           float64  `json:"Volume"`
	FearGreedIndex   int      `json:"FearGreedIndex"`
	MarketStatus     string   `json:"MarketStatus"`
}

type UpdateCountersInput struct {
	BotInstanceName    string   `json:"BotInstanceName"`
	WINCounter         *bool    `json:"WINCounter,omitempty"`
	LOSSCounter        *bool    `json:"LOSSCounter,omitempty"`
	TIMEOUTGainCounter *bool    `json:"TIMEOUTGainCounter,omitempty"`
	TIMEOUTLossCounter *bool    `json:"TIMEOUTLossCounter,omitempty"`
	AccountBalance     float64  `json:"AccountBalance"`
	FeesTotal          *float64 `json:"FeesTotal,omitempty"`
}
