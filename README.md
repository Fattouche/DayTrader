# Day_Trader

Day trader is a stock exchange platform built for performance and reliability.

## Design

This project is implemented in 2 different ways:

1. Django(Python)
- Rest communication
- Nginx LB
- Redis job queue
- Django_rq workers
- Memcached for caching
- Mysql DB


2. Golang
- GRPC communication
- Nginx LB
- Memcached for caching
- Mysql Db

## Development

To start using either app cd into the respective directory and run `docker-compose up`

## Generator

To run the generator, cd into the respective directory and then `cd workload_generator` and run `docker-compose up`

## Quote server

To run the quote server, `cd quote_server` and run `docker-compose up`

## Testing

### Golang

To test the golang app, `cd golang/day_trader/test_infrastructure` then run `docker-compose up`

__Note - Since this uses a real db and cache, you must docker-compose down every time you want to retest__
