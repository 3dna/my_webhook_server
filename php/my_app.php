<?php
$json_data = file_get_contents("php://input");

$person = json_decode($json_data, true);

// var_dump($person);

// echo "==============================="

// echo $person;

echo $person['email'];

// echo "hello";

?>
