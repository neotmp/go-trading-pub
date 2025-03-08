package portfolio

type Swap struct {
	Id     uint16
	PairId uint16
	Short  float32
	Long   float32
	// translate pips into value

}

// one portfolio per broker
// later I may attempt at merging multiple portfolios into one
type Portfolio struct {
	Id     uint16
	Profit float32
	Assets []*Asset
	Swap   float32 // sum of all swaps -- method
	// add a few more standard metrics for
	// portfolio performance evaluation
}

// get transactions
type Asset struct {
	Id           uint16
	PairId       uint16
	Volume       float32 // 0.01 - min contract
	Direction    uint8   // 0 - buy/long, 1 - sell/short
	Price        float32
	PriceCurrent float32
	Profit       float32 // method
	Swap         float32
}
