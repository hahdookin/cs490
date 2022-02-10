<?php
include 'cred.php';

try {
    $conn = new PDO("mysql:host=$servername;dbname=$dbname", $username, $password);
    $conn->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
    $stmt = $conn->prepare("SELECT type from users WHERE username = :username AND password = :password;", array(PDO::ATTR_CURSOR => PDO::CURSOR_FWDONLY));
    $stmt->execute(array("username" => $_POST["username"], "password" => hash("sha256",$_POST["password"])));
    //$stmt = $conn->prepare("SELECT type from users WHERE username = ".$_POST["username"]. " AND password = ".$_POST["password"].";");
    $stmt->execute();
  
    // set the resulting array to associative
    $result = $stmt->setFetchMode(PDO::FETCH_ASSOC);
    $result = $stmt->fetchAll()[0];
//     foreach(new TableRows(new RecursiveArrayIterator($stmt->fetchAll())) as $k=>$v) {
//       echo $v;
//     }
  } catch(PDOException $e) {
    echo "Error: " . $e->getMessage();
  }
  $conn = null;

    echo json_encode($result);
?>