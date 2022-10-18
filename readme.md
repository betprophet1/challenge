# Simplebet
### start-up
```sh
  bash start.sh
```
service running at *localhost:8080*
### log test cases
```sh
  docker logs -f simplebet_http_test
```
### Features supported
- Place an wager
- Buy full or part an wagere
- List wager

### Techincal supported
- Lock on DB for buying an wager
- Cache sold counts

### Nice to have features
- [x] Rate limit middleware for each user's buying requests
- [x] Lock on multiple requests from an user (incase user sends multiple requests at same time under valid rate limit)
- Integration tests