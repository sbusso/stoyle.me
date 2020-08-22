### Official repository for [stoyle.me URL Shortener](https://stoyle.me)

#### Why?

I build this for learning purpose. Enjoy the application for free. No trackers.

#### Features

- Fast
- Custom URL starting as short as 1 char
- 24 hours default URL expiry

#### Tech Stack

- Golang
- Redis

### API Documentation

> API endpoint: https://api.stoyle.me/v1/

#### API Payload

- "url" - Original URL
- "short" - Custom short URL(Optional)
- "expiry" - Time to expire: int(hours)

#### API Response

- "url" - Original URL
- "short" - Custom short URL
- "expiry" - Time to expire: int(hours)
- "rate_limit" - # of API calls remaining: int
- "rate_limit_reset" - Time to rate limit reset: int(minutes)

> API is rate limited to 10 calls every 30mins. Have fun.
