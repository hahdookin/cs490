<?php

$ch = curl_init();

// TODO: URL should point to middle man php server
curl_setopt($ch, CURLOPT_URL, "https://web.njit.edu/~cap82/login2.php"); 
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, 
    http_build_query(
        array(
            "username" => $_POST["username"],
            "password" => $_POST["password"],
        )
    )
);
curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

$response = curl_exec($ch);

curl_close($ch);

if (empty($response)) {
	echo "No response :(" . "<br>";
} else {
    // Process the response
    $json = json_decode($response, true);
    switch ($json["type"]) {
        case "student":
            echo "Welcome, student"."<br>";
            break;
        case "teacher":
            echo "Welcome, teacher"."<br>";
            break;
        default:
            echo "Database error"."<br>";
            break;
    }
}

?>
