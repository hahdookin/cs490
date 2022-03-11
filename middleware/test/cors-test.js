// const fetch = require('node_modules/node-fetch');

// const middle = 'http://ec2-3-136-155-192.us-east-2.compute.amazonaws.com';

// const postToAutograder = async function(qid, code) {
//     const res = await fetch(`${middle}/autograde`, {
//         method: 'POST',
//         headers: {
//             'Content-Type': 'application/json'
//         },
//         body: serialize({
//             qid: qid,
//             code: code
//         })
//     });
//     const resp = await res.json();
//     return resp;
// };

// postToAutograder("2","def DoubleIt(x):\n\treturn x.upper()")
// .then(data => {console.log(data)})
// .catch(err => {console.log(err)})

// ---

const http = require('http');

const data = JSON.stringify({
    qid: '2',
    code: 'def DoubleIt(x):\n\treturn x.upper()'
});

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


const req = http.request(options, (res) => {
    let data = '';

    console.log('Status Code:', res.statusCode);
    
    res.on('data', (d) => {
        data += d;
        // process.stdout.write(d)
    });

    res.on('end', () => {
        console.log('Body: ', JSON.parse(data));
    });

    // console.log(JSON.parse(data))

}).on("error", (err) => {
    console.log("Error: ", err.message);
});

req.write(data);
req.end();

// console.log(req)

// console.log(val)