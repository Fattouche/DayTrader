# Day_Trader

Day trader is a stock exchange platform built for performance and reliability.

## Development

To start using the django app simply run `docker-compose up`. This will create one replica
of the web app. The nginx loadbalancer will be responsible for loadbalancing requests towards the web application. To test that the
app is working properly simply go to `localhost` which shows the site through the loadbalancer. If you want to see each individual web app, run `docker-compose ps` and copy the port number on the left side of the colon and use that as the hostname `localhost<port>`.

## Generator

To run the generator, `cd workload_generator/golang` and run `docker-compose up`

## Quote server

To run the quote server, `cd quote_server` and run `docker-compose up`
