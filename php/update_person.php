<?php
function initialize_database() {
  $db = new SQLite3('people.db');

  $sql = <<<SQL
  CREATE TABLE IF NOT EXISTS people
    (id INTEGER PRIMARY KEY ON CONFLICT REPLACE,
     name STRING,
     email STRING,
     updates INTEGER)
SQL;

  $status = $db->exec($sql);

  if (!$status) {
    exit("Unable to initialize database file (people.db)");
  }

  return $db;
}

function get_update_count($database)
  $results = $db->query('SELECT updates FROM people where id = ' . $id);

  if (!$results) {
    exit("Unable get update count for person's record")
  }

  $updates = 1;
  if ($row = $results->fetchArray()) {
      $updates = $row[0] + 1;
  }
  return $updates;
}

// insert or replace the person's data
function insert($id, $name, $email, $updates) {
  $stmt = $db->prepare('INSERT INTO people(id, name, email, updates) VALUES(?, ?, ?, ?)');
  $stmt->bindValue(1, $id,      SQLITE3_INTEGER);
  $stmt->bindValue(2, $name,    SQLITE3_TEXT);
  $stmt->bindValue(3, $email,   SQLITE3_TEXT);
  $stmt->bindValue(4, $updates, SQLITE3_INTEGER);

  $result = $stmt->execute();

  if (!$result) {
    exit("Unable to insert or replace person's data");
  }
}

function read_http_post() {
  $json_data = file_get_contents("php://input");

  $make_associative_array = true;
  $person = json_decode($json_data, $make_associative_array);

  if (!$person) {
    exit("Invalid json sent");
  }

  return $person;
}


$person = read_http_post();

$id    = $person['id'];
$name  = $person['first_name'] . ' ' . $person['last_name'];
$email = $person['email'];

$db = intialize_database();

$updates = get_update_count();
insert($id, $name, $email, $updates);

?>
