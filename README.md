We would like you to create a crypto notification app. The features it should include:

Create a server which will be able to take in the following rest APIs

> 60$ > 60100 
> Create a notification. Line items may include current price of BTC, (nudging a user to invest > 10ms)
  market trade volume, intra day high price, market cap
> Send a notification to an email
> Wishlisted (BTC)
> List sent notifications (sent, outstanding, failed etc.)
> Delete a notification


> Trigger the notification sync (order, wishlisted, invested)
> Trigger the notification async (high_chance)

Trigger Notification Request:
{
    "type": "high_change",
    "coin": "BTC",
    "channel": ["email", "whatsapp", "sms", "push"],
    "data": {
        "market_trade_volume": 1882.919,
        "intra_day_high_price": 600022.090,
        "market_cap": 2392893293.989,
        "ordered_amt": 1223982
    }
}

id > email

Requirements:
- extensible to support other coins
- extensible to support other channels
- list sent notifications based on status (paginated)
- persist those notifications

Assumptions:
- email functionality is already in place
- auth layer in place