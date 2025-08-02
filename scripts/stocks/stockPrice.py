import yfinance as yf
import sys

def getPrice(tickerSymbol):
    data = yf.download(tickerSymbol+".NS", period="5d", interval="1d", progress=False, auto_adjust=False)

    if not data.empty:
        if len(data) > 1:
            previous_day = data.iloc[-2]
            prev_close = previous_day['Close'].get(tickerSymbol+".NS", 0)
            return prev_close
        else:
            return 0
    else:
        return 0

if __name__ == "__main__":
    print(getPrice(sys.argv[1]))
