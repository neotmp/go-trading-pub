package simulation

import "time"

// all other structs would have S prefix
type Simulation struct {
	Id          uint16
	Name        string
	Memo        string
	Version     uint8 // when modified, create new version and old version status becomes canceled
	CreatedAt   time.Time
	ModifiedAt  time.Time
	From        time.Time
	To          time.Time
	Hedge       bool
	RealAccount bool
	AccountId   uint16
	Equity      float32 // initial conditions
	Margin      float32 // initial conditions
	FreeMargin  float32 // initial conditions
	MarginCall  float32 // initial conditions
	Orders      []*SOrder
	Positions   []*SPosition
	Status      uint8 // 1 - running , 2 - concluded, 0 - canceled
	Result      SResult
}

type SResult struct {
	Id           uint16
	SimulationId uint16 // fk

}

type SPosition struct {
	Id        uint16
	Price     float32
	Direction uint8 // 1 - buy, 2 - sell
}
type SOrder struct {
	Id        uint16
	Price     float32
	Direction uint8 // 1 - buy, 2 - sell
	Volume    float32
	SL        float32
	TP        float32
	TS        float32
	Type      uint8 // 1 - market, 2 - conditional??
	Status    uint8 // 1 - executed, 2 - pending, 0 - canceled
}

// the questions I seek answers to based on the past data:
// 1. How much adverse price movement can I survive?
// 2. What is the optimal hold period?
// 3.

// questions about the future, based on the past:
//

// Math Model and assumptions to test:
// 1. I enter position at or near resistance/support level expecting that trend reverse
// it may take some time which means that I'd be loosing equity. Too much equity loss may result in
// margin call. To prevent it from happening I need to have a model to calculate "safe" boundaries, i.e.
// how much capital should be allocated to possible drawdown.
// 2. What is the model to estimate the boundaries?
// Each calendar month I calculate MAX and MIN value, which is a good approximation of volatility.
// Than I do the same for, say, 3-month period. Now I have possible outliers and volatility avaraged over 3-month period.
// 3. When I open position I can expect prices to go up or down within established boundaries.
// Example, I expected price of EURUSD to go up. I buy at 1.02000. Based on the recent volatility, which is 250 pips, I also anticipate that
// the price can go down 250 pips, resulting in 25$ loss of equity, thus necessitating my having to be able to cover that drop in equity
// without triggering a margin call or having to forsefully to close my positions in order to stave off possible furthure  deterioration of
// my situation.
// This is the answer that the simulation is to provide.

// How do we do it? First run it based on historical data, 15-minute fractals.

// 1. Open position(s)
// 2. Calculate free margin
// 3. Calculate how much reserves we need to cover: worst case, best case, average case
// worst case is the MAX outlier, best case - 50% within expected range, average case - 100% within range
// 4. How can I hedge my position by taking the opposite deal, e.g. long if I'm short and vice versa.
// I don't know yet??????
// 5. Move to next point in timeframe, calculate equity, free margin, record data, create report.
// 6. If equity drops below margin call, close position(s)
// Close positions based on: losing most money, largest by volume or margin, othe considerations???
//

// How to work w/ probababilities
// Let's assume that we have time series covering one month. Within this period we can
