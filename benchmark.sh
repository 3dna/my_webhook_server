#!/bin/bash
ab -n 100 -c 1 -p person_created.json -T application/json http://localhost:4567/update_person
