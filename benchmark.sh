#!/bin/bash
# ab -n 10000 -c 8 -p person_created.json -T application/json http://localhost:4567/person
ab -n 10000 -c 8 -p person_created.json -T application/json http://localhost/my_app2.php
