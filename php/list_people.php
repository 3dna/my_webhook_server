<?php
$db = new SQLite3('people.db');

echo "<table>";
echo "<tr>";
echo   "<th>ID</th>";
echo   "<th>Name</th>";
echo   "<th>Email</th>";
echo   "<th>Updates</th>";
echo "</tr>";

$results = $db->query('SELECT * FROM people');
while ($row = $results->fetchArray()) {
   echo "<tr>";
     echo "<td>".$row[0]."</td>";
     echo "<td>".$row[1]."</td>";
     echo "<td>".$row[2]."</td>";
     echo "<td>".$row[3]."</td>";
    // var_dump($row);
   echo "</tr>";
}
echo "</table>";

?>