PROBLEM

So, you have multiple Forex accounts with multiple brokers*. There's no way you can aggregate the trading data from all sources unless you’ re willing to subject yourself to tedious operations of downloading individual P&L reports and trades history, 
then loading them into Excel spreadsheet or, perhaps, Python’s pandas Dataframe for further analysis.

SOLUTION

This software allows you to mirror all your trading activity in one place. It’s not intended for intraday trading strategies or HFT with high volume of orders. It’s more suitable for relatively small number of operations that you can easily enter manually**.

You have multiple brokers with multiple accounts. You have leverages, account currencies, pairs, swaps and other conditions that correspond with high fidelity to those of your real account terms.

The software is not intended as a trading terminal***.

In my case I need to be able to run multiple risk management scenarios in addition to being able to compare performance of my investments with different brokers. It’s much closer in this regard to traditional portfolio management.

ROADMAP

STAGE ONE
Create backend that allows you to do:
Open new broker
See all brokers
Close Broker
Edit broker
Open new account
See all accounts
Edit account
Close account
Place order
See orders
Modify order
Close order
See open positions
See P&L 
See pairs
Add pairs
STAGE TWO
Cover most important functionality with tests.

STAGE THREE
Add slog

STAGE FOUR
Create API (Web Server, Fiber) that allows you to do:
(all functionality listed in above)

STAGE FIVE
Create Frontend that allows you to do:
(all functionality listed in above)

STAGE SIX
Add OpenTelemetry for tracing.

Stage SEVEN:
Implement simulation block
This is where you run multiple scenarios based on your risk assessment models.

*Currently there’s no intent of expanding functionality to other types of brokers, including crypto exchanges.

**Currently there’s no API available solutions, so manually entering every transaction will soon become extremely impractical with high volumes of trades.

***Due to lack of API support from Forex brokers there’s no viable path toward creating trading terminal at the time. It is possible for other financial instruments, such as crypto assets. 
I’ll keep this in mind since placing order functionality may be easily extended to interact with external APIs, e.g. crypto exchanges.
