<?php
// Get PHP to report all errors
error_reporting(E_ALL);

// Some text to output to the page
echo "Middleware Section"."<br>";

// Arr Instantiate
$arr = array();

function cURL_POST(string $url) {
    // Initalize 
    $ch = curl_init();

    // Curl Setopt info
    curl_setopt($ch, CURLOPT_URL, $url); 
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


    // 
    $response = curl_exec($ch);

    curl_close($ch);
}





// Listen section
function Listen($array) : Array {
    $content = file_get_contents("php://input");
    // echo $content."<br>";
    $pairs = explode("&", $content);
    // echo gettype($content)."<br>";
    // echo count($pairs)."<br>";

    for($i = 0; $i < count($pairs); $i++) {
        // Split pairs by key and value
        $key_value = explode("=",$pairs[$i]);

        // Creating a variable for key and values
        $key = $key_value[0];
        $val = $key_value[1];

        // Append key and values to return array
        $array[] = $key;
        $array[] = $val; 
        // echo "$key\t$val"."<br>";
    }
    return $array;
}




$arr = Listen($arr);
// Listen($arr);
// Printing Output of $arr
echo json_encode($arr);

?>