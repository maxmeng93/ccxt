package ccxt

type Btcmarkets struct {
   *btcmarkets
   Core *btcmarkets
}

func NewBtcmarkets(userConfig map[string]interface{}) Btcmarkets {
   p := &btcmarkets{}
   p.Init(userConfig)
   return Btcmarkets{
       btcmarkets: p,
       Core:  p,
   }
}

// PLEASE DO NOT EDIT THIS FILE, IT IS GENERATED AND WILL BE OVERWRITTEN:
// https://github.com/ccxt/ccxt/blob/master/CONTRIBUTING.md#how-to-contribute-code


func (this *Btcmarkets) FetchTransactionsWithMethod(method interface{}, options ...FetchTransactionsWithMethodOptions) ([]Transaction, error) {

    opts := FetchTransactionsWithMethodOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var code interface{} = nil
    if opts.Code != nil {
        code = *opts.Code
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchTransactionsWithMethod(method, code, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewTransactionArray(res), nil
}
/**
 * @method
 * @name btcmarkets#fetchDepositsWithdrawals
 * @description fetch history of deposits and withdrawals
 * @see https://docs.btcmarkets.net/v3/#tag/Fund-Management-APIs/paths/~1v3~1transfers/get
 * @param {string} [code] unified currency code for the currency of the deposit/withdrawals, default is undefined
 * @param {int} [since] timestamp in ms of the earliest deposit/withdrawal, default is undefined
 * @param {int} [limit] max number of deposit/withdrawals to return, default is undefined
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} a list of [transaction structure]{@link https://docs.ccxt.com/#/?id=transaction-structure}
 */
func (this *Btcmarkets) FetchDepositsWithdrawals(options ...FetchDepositsWithdrawalsOptions) ([]Transaction, error) {

    opts := FetchDepositsWithdrawalsOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var code interface{} = nil
    if opts.Code != nil {
        code = *opts.Code
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchDepositsWithdrawals(code, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewTransactionArray(res), nil
}
/**
 * @method
 * @name btcmarkets#fetchDeposits
 * @description fetch all deposits made to an account
 * @see https://docs.btcmarkets.net/v3/#tag/Fund-Management-APIs/paths/~1v3~1deposits/get
 * @param {string} code unified currency code
 * @param {int} [since] the earliest time in ms to fetch deposits for
 * @param {int} [limit] the maximum number of deposits structures to retrieve
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object[]} a list of [transaction structures]{@link https://docs.ccxt.com/#/?id=transaction-structure}
 */
func (this *Btcmarkets) FetchDeposits(options ...FetchDepositsOptions) ([]Transaction, error) {

    opts := FetchDepositsOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var code interface{} = nil
    if opts.Code != nil {
        code = *opts.Code
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchDeposits(code, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewTransactionArray(res), nil
}
/**
 * @method
 * @name btcmarkets#fetchWithdrawals
 * @description fetch all withdrawals made from an account
 * @see https://docs.btcmarkets.net/v3/#tag/Fund-Management-APIs/paths/~1v3~1withdrawals/get
 * @param {string} code unified currency code
 * @param {int} [since] the earliest time in ms to fetch withdrawals for
 * @param {int} [limit] the maximum number of withdrawals structures to retrieve
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object[]} a list of [transaction structures]{@link https://docs.ccxt.com/#/?id=transaction-structure}
 */
func (this *Btcmarkets) FetchWithdrawals(options ...FetchWithdrawalsOptions) ([]Transaction, error) {

    opts := FetchWithdrawalsOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var code interface{} = nil
    if opts.Code != nil {
        code = *opts.Code
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchWithdrawals(code, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewTransactionArray(res), nil
}
/**
 * @method
 * @name btcmarkets#fetchMarkets
 * @description retrieves data on all markets for btcmarkets
 * @see https://docs.btcmarkets.net/v3/#tag/Market-Data-APIs/paths/~1v3~1markets/get
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object[]} an array of objects representing market data
 */
func (this *Btcmarkets) FetchMarkets(params ...interface{}) ([]MarketInterface, error) {
    res := <- this.Core.FetchMarkets(params...)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewMarketInterfaceArray(res), nil
}
/**
 * @method
 * @name btcmarkets#fetchTime
 * @description fetches the current integer timestamp in milliseconds from the exchange server
 * @see https://docs.btcmarkets.net/v3/#tag/Misc-APIs/paths/~1v3~1time/get
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {int} the current integer timestamp in milliseconds from the exchange server
 */
func (this *Btcmarkets) FetchTime(params ...interface{}) ( int64, error) {
    res := <- this.Core.FetchTime(params...)
    if IsError(res) {
        return -1, CreateReturnError(res)
    }
    return (res).(int64), nil
}
/**
 * @method
 * @name btcmarkets#fetchBalance
 * @description query for balance and get the amount of funds available for trading or funds locked in orders
 * @see https://docs.btcmarkets.net/v3/#tag/Account-APIs/paths/~1v3~1accounts~1me~1balances/get
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} a [balance structure]{@link https://docs.ccxt.com/#/?id=balance-structure}
 */
func (this *Btcmarkets) FetchBalance(params ...interface{}) (Balances, error) {
    res := <- this.Core.FetchBalance(params...)
    if IsError(res) {
        return Balances{}, CreateReturnError(res)
    }
    return NewBalances(res), nil
}
/**
 * @method
 * @name btcmarkets#fetchOHLCV
 * @description fetches historical candlestick data containing the open, high, low, and close price, and the volume of a market
 * @see https://docs.btcmarkets.net/v3/#tag/Market-Data-APIs/paths/~1v3~1markets~1{marketId}~1candles/get
 * @param {string} symbol unified symbol of the market to fetch OHLCV data for
 * @param {string} timeframe the length of time each candle represents
 * @param {int} [since] timestamp in ms of the earliest candle to fetch
 * @param {int} [limit] the maximum amount of candles to fetch
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {int[][]} A list of candles ordered as timestamp, open, high, low, close, volume
 */
func (this *Btcmarkets) FetchOHLCV(symbol string, options ...FetchOHLCVOptions) ([]OHLCV, error) {

    opts := FetchOHLCVOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var timeframe interface{} = nil
    if opts.Timeframe != nil {
        timeframe = *opts.Timeframe
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchOHLCV(symbol, timeframe, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewOHLCVArray(res), nil
}
/**
 * @method
 * @name btcmarkets#fetchOrderBook
 * @description fetches information on open orders with bid (buy) and ask (sell) prices, volumes and other data
 * @see https://docs.btcmarkets.net/v3/#tag/Market-Data-APIs/paths/~1v3~1markets~1{marketId}~1orderbook/get
 * @param {string} symbol unified symbol of the market to fetch the order book for
 * @param {int} [limit] the maximum amount of order book entries to return
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} A dictionary of [order book structures]{@link https://docs.ccxt.com/#/?id=order-book-structure} indexed by market symbols
 */
func (this *Btcmarkets) FetchOrderBook(symbol string, options ...FetchOrderBookOptions) (OrderBook, error) {

    opts := FetchOrderBookOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchOrderBook(symbol, limit, params)
    if IsError(res) {
        return OrderBook{}, CreateReturnError(res)
    }
    return NewOrderBook(res), nil
}
/**
 * @method
 * @name btcmarkets#fetchTicker
 * @description fetches a price ticker, a statistical calculation with the information calculated over the past 24 hours for a specific market
 * @see https://docs.btcmarkets.net/v3/#tag/Market-Data-APIs/paths/~1v3~1markets~1{marketId}~1ticker/get
 * @param {string} symbol unified symbol of the market to fetch the ticker for
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} a [ticker structure]{@link https://docs.ccxt.com/#/?id=ticker-structure}
 */
func (this *Btcmarkets) FetchTicker(symbol string, options ...FetchTickerOptions) (Ticker, error) {

    opts := FetchTickerOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchTicker(symbol, params)
    if IsError(res) {
        return Ticker{}, CreateReturnError(res)
    }
    return NewTicker(res), nil
}
func (this *Btcmarkets) FetchTicker2(symbol string, options ...FetchTicker2Options) (Ticker, error) {

    opts := FetchTicker2OptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchTicker2(symbol, params)
    if IsError(res) {
        return Ticker{}, CreateReturnError(res)
    }
    return NewTicker(res), nil
}
/**
 * @method
 * @name btcmarkets#fetchTrades
 * @description get the list of most recent trades for a particular symbol
 * @see https://docs.btcmarkets.net/v3/#tag/Market-Data-APIs/paths/~1v3~1markets~1{marketId}~1trades/get
 * @param {string} symbol unified symbol of the market to fetch trades for
 * @param {int} [since] timestamp in ms of the earliest trade to fetch
 * @param {int} [limit] the maximum amount of trades to fetch
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {Trade[]} a list of [trade structures]{@link https://docs.ccxt.com/#/?id=public-trades}
 */
func (this *Btcmarkets) FetchTrades(symbol string, options ...FetchTradesOptions) ([]Trade, error) {

    opts := FetchTradesOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchTrades(symbol, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewTradeArray(res), nil
}
/**
 * @method
 * @name btcmarkets#createOrder
 * @description create a trade order
 * @see https://docs.btcmarkets.net/v3/#tag/Order-Placement-APIs/paths/~1v3~1orders/post
 * @param {string} symbol unified symbol of the market to create an order in
 * @param {string} type 'market' or 'limit'
 * @param {string} side 'buy' or 'sell'
 * @param {float} amount how much of currency you want to trade in units of base currency
 * @param {float} [price] the price at which the order is to be fulfilled, in units of the quote currency, ignored in market orders
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @param {float} [params.triggerPrice] the price at which a trigger order is triggered at
 * @returns {object} an [order structure]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *Btcmarkets) CreateOrder(symbol string, typeVar string, side string, amount float64, options ...CreateOrderOptions) (Order, error) {

    opts := CreateOrderOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var price interface{} = nil
    if opts.Price != nil {
        price = *opts.Price
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.CreateOrder(symbol, typeVar, side, amount, price, params)
    if IsError(res) {
        return Order{}, CreateReturnError(res)
    }
    return NewOrder(res), nil
}
/**
 * @method
 * @name btcmarkets#cancelOrders
 * @description cancel multiple orders
 * @see https://docs.btcmarkets.net/v3/#tag/Batch-Order-APIs/paths/~1v3~1batchorders~1{ids}/delete
 * @param {string[]} ids order ids
 * @param {string} symbol not used by btcmarkets cancelOrders ()
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} an list of [order structures]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *Btcmarkets) CancelOrders(ids interface{}, options ...CancelOrdersOptions) ([]Order, error) {

    opts := CancelOrdersOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.CancelOrders(ids, symbol, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewOrderArray(res), nil
}
/**
 * @method
 * @name btcmarkets#cancelOrder
 * @description cancels an open order
 * @see https://docs.btcmarkets.net/v3/#operation/cancelOrder
 * @param {string} id order id
 * @param {string} symbol not used by btcmarket cancelOrder ()
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} An [order structure]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *Btcmarkets) CancelOrder(id string, options ...CancelOrderOptions) (Order, error) {

    opts := CancelOrderOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.CancelOrder(id, symbol, params)
    if IsError(res) {
        return Order{}, CreateReturnError(res)
    }
    return NewOrder(res), nil
}
/**
 * @method
 * @name btcmarkets#fetchOrder
 * @description fetches information on an order made by the user
 * @see https://docs.btcmarkets.net/v3/#operation/getOrderById
 * @param {string} id the order id
 * @param {string} symbol not used by btcmarkets fetchOrder
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} An [order structure]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *Btcmarkets) FetchOrder(id string, options ...FetchOrderOptions) (Order, error) {

    opts := FetchOrderOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchOrder(id, symbol, params)
    if IsError(res) {
        return Order{}, CreateReturnError(res)
    }
    return NewOrder(res), nil
}
/**
 * @method
 * @name btcmarkets#fetchOrders
 * @description fetches information on multiple orders made by the user
 * @see https://docs.btcmarkets.net/v3/#operation/listOrders
 * @param {string} symbol unified market symbol of the market orders were made in
 * @param {int} [since] the earliest time in ms to fetch orders for
 * @param {int} [limit] the maximum number of order structures to retrieve
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {Order[]} a list of [order structures]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *Btcmarkets) FetchOrders(options ...FetchOrdersOptions) ([]Order, error) {

    opts := FetchOrdersOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchOrders(symbol, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewOrderArray(res), nil
}
/**
 * @method
 * @name btcmarkets#fetchOpenOrders
 * @description fetch all unfilled currently open orders
 * @see https://docs.btcmarkets.net/v3/#operation/listOrders
 * @param {string} symbol unified market symbol
 * @param {int} [since] the earliest time in ms to fetch open orders for
 * @param {int} [limit] the maximum number of  open orders structures to retrieve
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {Order[]} a list of [order structures]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *Btcmarkets) FetchOpenOrders(options ...FetchOpenOrdersOptions) ([]Order, error) {

    opts := FetchOpenOrdersOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchOpenOrders(symbol, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewOrderArray(res), nil
}
/**
 * @method
 * @name btcmarkets#fetchClosedOrders
 * @description fetches information on multiple closed orders made by the user
 * @see https://docs.btcmarkets.net/v3/#operation/listOrders
 * @param {string} symbol unified market symbol of the market orders were made in
 * @param {int} [since] the earliest time in ms to fetch orders for
 * @param {int} [limit] the maximum number of order structures to retrieve
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {Order[]} a list of [order structures]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *Btcmarkets) FetchClosedOrders(options ...FetchClosedOrdersOptions) ([]Order, error) {

    opts := FetchClosedOrdersOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchClosedOrders(symbol, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewOrderArray(res), nil
}
/**
 * @method
 * @name btcmarkets#fetchMyTrades
 * @description fetch all trades made by the user
 * @see https://docs.btcmarkets.net/v3/#operation/getTrades
 * @param {string} symbol unified market symbol
 * @param {int} [since] the earliest time in ms to fetch trades for
 * @param {int} [limit] the maximum number of trades structures to retrieve
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {Trade[]} a list of [trade structures]{@link https://docs.ccxt.com/#/?id=trade-structure}
 */
func (this *Btcmarkets) FetchMyTrades(options ...FetchMyTradesOptions) ([]Trade, error) {

    opts := FetchMyTradesOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchMyTrades(symbol, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewTradeArray(res), nil
}
/**
 * @method
 * @name btcmarkets#withdraw
 * @description make a withdrawal
 * @see https://docs.btcmarkets.net/v3/#tag/Fund-Management-APIs/paths/~1v3~1withdrawals/post
 * @param {string} code unified currency code
 * @param {float} amount the amount to withdraw
 * @param {string} address the address to withdraw to
 * @param {string} tag
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} a [transaction structure]{@link https://docs.ccxt.com/#/?id=transaction-structure}
 */
func (this *Btcmarkets) Withdraw(code string, amount float64, address string, options ...WithdrawOptions) (Transaction, error) {

    opts := WithdrawOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var tag interface{} = nil
    if opts.Tag != nil {
        tag = *opts.Tag
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.Withdraw(code, amount, address, tag, params)
    if IsError(res) {
        return Transaction{}, CreateReturnError(res)
    }
    return NewTransaction(res), nil
}