package update

// GetData is entry point for retrieving currency pairs quotes
// from ext APIs
// Once data is in the DB, we run update on positions
// Then we run update on accounts
func GetData() {
	// 1. Which pairs should be requested
	// 2. What is the timeframe
	// 3. What is our approach to data rate limits and # of monthly requests?
	// 4. What's plan B in case one API is down, what's our redundancy plan?
}
