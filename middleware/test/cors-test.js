const http = require('http');

// Data to be sent to point
const data = JSON.stringify({
    qid: '2',
    code: 'def DoubleIt(x):\n\treturn x.upper()'
});

// Header Options
const options = {
    hostname: 'localhost',
    port: 8087,
    path: '/autograde',
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Content-Length': data.length
    }
};

// The Request
const req = http.request(options, (res) => {
    let data = '';

    console.log('Status Code:', res.statusCode);
    
    res.on('data', (d) => {
        data += d;
    });

    res.on('end', () => {
        console.log('Body: ', JSON.parse(data));
    });

}).on("error", (err) => {
    console.log("Error: ", err.message);
});

req.write(data);
req.end();