# RMANT
RMANT - Redis Mark AgeNT. It is a simple package to provide marking system with (temporary lived marks) based on REDIS database.
## Setup
```go
// First of all you need to create MarkAgent Service 
ma := rmant.MarkAgent(redisClient, rmant.Conf{
    Prefix: "myCoolRedisPrefix",
    MarkTTL: 5 * time.Hour,
})
// Then you can mark some entity with agent
ma.Mark(ma.Key("pussy"))
```