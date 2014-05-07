<?php

// read the HTTP POST json data and parse it

$json_data = file_get_contents("php://input");

$make_associative_array = true;
$person = json_decode($json_data, $make_associative_array);

$id    = $person['id'];
$email = $person['email'];
$name  = $person['first_name'] . ' ' . $person['last_name'];

// create sql table

$db = new SQLite3('people.db');

$sql = <<<SQL
CREATE TABLE IF NOT EXISTS people
  (id INTEGER PRIMARY KEY ON CONFLICT REPLACE,
   name STRING,
   email STRING,
   updates INTEGER)
SQL;
$stmt = $db->exec($sql);

// get the update count

$results = $db->query('SELECT updates FROM people where id = ' . $id);

$updates = 1;
if ($row = $results->fetchArray()) {
    $updates = $row[0] + 1;
}

// insert or replace the person's data

$stmt = $db->prepare('INSERT INTO people(id, name, email, updates) VALUES(?, ?, ?, ?)');
$stmt->bindValue(1, $id,      SQLITE3_INTEGER);
$stmt->bindValue(2, $name,    SQLITE3_TEXT);
$stmt->bindValue(3, $email,   SQLITE3_TEXT);
$stmt->bindValue(4, $updates, SQLITE3_INTEGER);

$result = $stmt->execute();

?>
