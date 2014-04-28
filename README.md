# Code for making your own server for processing NationBuilder webhooks

This a collection of example code for NationBuilder webhook servers. Each server is setup to accept
an HTTP POST request with a body containing the information for a person. An example request body
is in the file person_created.json. That file has JSON of the same format if you had a webhook triggered
when a person is created.

Examples are provided for: PHP, Ruby, and Go.
The URL is localhost:4567 for each of the servers.

## Starting the webhook servers

### PHP
Copy the file file my_app.php in where your Apache webserver files live

The default port is not easy to change, use :4000(?)

### Ruby
```
bundle install
ruby app.rb
```

### Go
`go run my_app.go`

## To benchmark
`./benchmark.sh`

## To test individual HTTP POST use curl(1)

```
curl -X POST -d @person_created.json http://localhost:4567/

curl -X POST -d @person_created.json http://tranquil-reef-3487.herokuapp.com/
```
