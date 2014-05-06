
# My Webhook Server

This repository has example code for setting up a webserver that accept webhooks. Webhooks are sent 
from NationBuilder to your webserver via [HTTP POST requests](http://en.wikipedia.org/wiki/POST_(HTTP)).

Each server is setup to accept an HTTP POST request with a body  containing the information for a person. An example request body is in the file ```person_created.json```. That file has JSON of the same format if you had a webhook triggered when a person is created.

Examples are provided for: PHP, Ruby, and Go.

## Starting the webhook servers

### Starting the PHP server

On OSX, copy *.php files to ```/Library/WebServer/Documents/```

Make the directory writeable so that the sqlite database file can be modified.
```
chmod a+w /Library/WebServer/Documents
```

In a browser, go here to list the people ```http://localhost/list_people``` 

### Starting the Ruby server
```
bundle install
ruby my_app.rb
```

In a browser, go here to list the people ```http://localhost:4567/list_people``` 


### Starting the Go server
```
go get code.google.com/p/gosqlite/sqlite
go run my_app.go
```

In a browser, go here to list the people ```http://localhost:4567/list_people``` 

## To test individual HTTP POST use [curl(1)](http://en.wikipedia.org/wiki/CURL)

Your web server is now ready to receive webhooks. You can simulate a webhook send out
by NationBuilder using curl.

For Ruby and Go, use curl with this URL
```
curl -X POST -d @person_created.json http://localhost:4567/update_person
```

For PHP, use curl with this URL:
```
curl -X POST -d @person_created.json http://localhost/update_person
```

## Connect webhook server to NationBuilder

For PHP, use this URL ```http://{your machine's external name}/update_person```

For Ruby or Go, use this URL ```http://{your machine's external name}:4567/update_person```

Read more about NationBuilder webhooks here: http://nationbuilder.com/webhooks_overview

![alt tag](https://raw.githubusercontent.com/3dna/my_webhook_server/master/nationbuilder_webhook_setup.png)

## To benchmark

The script below uses ab - Apache HTTP server benchmarking tool - to test the webhooks.

Sqlite3 does not support concurrent writes. So, there will be no improvement with concurrency testing.

For testing, php server use:
```
./benchmark_php.sh
```

For testing, go or ruby server use:
```
./benchmark.sh
```

These benchmarks were run on a MacBook Pro, Mid 2012, 2.3 GHz Intel Core i7.
Here is a summary of the data from {php, ruby, go}/benchmark_out.txt

| Language  | Requests per second |
| ----------|:-------------------:|
| PHP       | 513.04              |
| Ruby      | 236.55              |
| Go        | 638.45              |


