<?php
// Get PHP to report all errors
error_reporting(E_ALL);

// making curl more modular
function cURL_POST(string $url, string $username, string $password) : string {
    // Initalize curl resource
    $ch = curl_init();

    // Curl Setopt Info
    curl_setopt($ch, CURLOPT_URL, $url); 
    curl_setopt($ch, CURLOPT_POST, true);
    curl_setopt($ch, CURLOPT_POSTFIELDS, 
        http_build_query(
            array(
                "username" => $username,
                "password" => $password,
            )
        )
    );
    curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

    // output response from curl
    $output = curl_exec($ch);

    // close curl resource to free system resources
    curl_close($ch);

    // return response from curl : string
    return $output;
}

// sends $arr from php://input to curl backend
$response = cURL_POST("https://web.njit.edu/~gmo9/back-end/backend.php", $_POST["username"], $_POST["password"]);

// Printing Output of $arr
echo json_encode($response);
?>