<?php

/*
TODO: 
- Change Server to Backend
*/


// Get PHP to report all errors
error_reporting(E_ALL);

// Some text to output to the page
echo "Middleware Section"."<br>";

// Arr Instantiate (will later become dictionary/hashmap)
$arr = array();

// making curl more modular
function cURL_POST(string $url, array $arr) : string {
    // Initalize curl resource
    $ch = curl_init();

    // Curl Setopt Info
    curl_setopt($ch, CURLOPT_URL, $url); 
    curl_setopt($ch, CURLOPT_POST, true);
    curl_setopt($ch, CURLOPT_POSTFIELDS, 
        http_build_query($arr)
    );
    curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

    // output response from curl
    $response = curl_exec($ch);

    // close curl resource to free system resources
    curl_close($ch);

    // return response from curl : string
    return $response;
}


// Listen input in Server Body
function Listen($array) : Array {
    // Listen for POST request info
    $content = file_get_contents("php://input");

    // Split data into key-value pairs
    $pairs = explode("&", $content);

    for($i = 0; $i < count($pairs); $i++) {
        // Split pairs by key and value
        $key_value = explode("=",$pairs[$i]);

        // Creating a variable for key and values
        $key = $key_value[0];
        $val = $key_value[1];

        // Append key and values to return "array" (dictionary)
        $array[$key] = $val;
        // $array[] = $val; 
    }
    return $array;
}


// set $arr so that odd indexes == keys && even indexes == values 
$arr = Listen($arr);

// sends $arr from php://input to curl backend
$res = cURL_POST("https://web.njit.edu/~aeo34/server.php", $arr);

// Printing Output of $arr
echo json_encode($arr)."<br>";

?>