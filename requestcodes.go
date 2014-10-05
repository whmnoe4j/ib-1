package ib

var REQUEST RequestHeader

type RequestHeader struct {
	CODE    RequestCodes
	VERSION RequestVersions
}

type RequestCodes struct {
	MARKET_DATA      int64
	MARKET_DEPTH     int64
	CONTRACT_DATA    int64
	HISTORICAL_DATA  int64
	ACCOUNT_VALUE    int64
	PLACE_ORDER      int64
	CANCEL_ORDER     int64
	OPEN_ORDERS      int64
	ALL_OPEN_ORDERS  int64
	AUTO_OPEN_ORDERS int64
	REQ_IDS          int64
	EXERCISE_OPTS    int64
	GLOBAL_CANCEL    int64
	REALTIMEBARS int64
}

type RequestVersions struct {
	MARKET_DATA      int64
	MARKET_DEPTH     int64
	CONTRACT_DATA    int64
	HISTORICAL_DATA  int64
	ACCOUNT_VALUE    int64
	PLACE_ORDER      int64
	CANCEL_ORDER     int64
	OPEN_ORDERS      int64
	ALL_OPEN_ORDERS  int64
	AUTO_OPEN_ORDERS int64
	REQ_IDS          int64
	EXERCISE_OPTS    int64
	GLOBAL_CANCEL    int64
	REALTIMEBARS int64
}

func init() {
	REQUEST.CODE.MARKET_DATA = 1
	REQUEST.CODE.MARKET_DEPTH = 10
	REQUEST.CODE.CONTRACT_DATA = 9
	REQUEST.CODE.HISTORICAL_DATA = 20
	REQUEST.CODE.ACCOUNT_VALUE = 6
	REQUEST.CODE.PLACE_ORDER = 3
	REQUEST.CODE.CANCEL_ORDER = 4
	REQUEST.CODE.OPEN_ORDERS = 5
	REQUEST.CODE.ALL_OPEN_ORDERS = 16
	REQUEST.CODE.AUTO_OPEN_ORDERS = 15
	REQUEST.CODE.REQ_IDS = 8
	REQUEST.CODE.EXERCISE_OPTS = 21
	REQUEST.CODE.GLOBAL_CANCEL = 58
	REQUEST.CODE.REALTIMEBARS = 50

	REQUEST.VERSION.MARKET_DATA = 10
	REQUEST.VERSION.MARKET_DEPTH = 4
	REQUEST.VERSION.CONTRACT_DATA = 7
	REQUEST.VERSION.HISTORICAL_DATA = 5
	REQUEST.VERSION.ACCOUNT_VALUE = 2
	//	REQUEST.VERSION.PLACE_ORDER =
	//	REQUEST.VERSION.CANCEL_ORDER
	//	REQUEST.VERSION.OPEN_ORDERS
	//	REQUEST.VERSION.ALL_OPEN_ORDERS
	//	REQUEST.VERSION.AUTO_OPEN_ORDERS
	//	REQUEST.VERSION.REQ_IDS
	//	REQUEST.VERSION.EXERCISE_OPTS
	//	REQUEST.VERSION.GLOBAL_CANCEL
	REQUEST.VERSION.REALTIMEBARS = 2
}
