
# A server for receiving NationBuilder webhooks

In this repository is code for setting up your webserver that accept webhooks. Webhooks are sent 
from NationBuilder to your webserver via [HTTP POST requests](http://en.wikipedia.org/wiki/POST_(HTTP)).

Each server is setup to accept an HTTP POST request with a body  containing the information for a person. An example request body is in the file person_created.json.  That file has JSON of the same format if you had a webhook triggered when a person is created.

Examples are provided for: PHP, Ruby, and Go.
The URL is localhost:4567 for each of the servers.

## Starting the webhook servers

### Starting the PHP server
Copy the file file my_app.php in where your Apache webserver files live

The default port is not easy to change, use :4000(?)

### Starting the Ruby server
```
bundle install
ruby my_app.rb
```

### Starting the Go server
```
go run my_app.go
```

## To test individual HTTP POST use [curl(1)](http://en.wikipedia.org/wiki/CURL)

```
curl -X POST -d @person_created.json http://localhost:4567/

curl -X POST -d @person_created.json http://tranquil-reef-3487.herokuapp.com/
```

## To benchmark
```
./benchmark.sh
```
