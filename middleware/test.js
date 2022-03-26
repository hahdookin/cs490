// node-fetch@2
const fetch = require('node-fetch');

const autograde = async function(qid, code) {
    try {
        const response = await fetch("http://ec2-3-136-155-192.us-east-2.compute.amazonaws.com/autograde", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "qid": qid,
                "code": code
            })
        })
        const json = await response.json()
        console.log(json)
    } catch (error) {
        console.log(error)
    }
}
trial = autograde("2", "def DoubleIt(x):\n\treturn x.upper()")





































// (async () => {
//     try {
//         const response = await fetch("http://ec2-3-136-155-192.us-east-2.compute.amazonaws.com/autograde", {
//             method: 'POST',
//             headers: {
//                 'Content-Type': 'application/json'
//             },
//             body: JSON.stringify({
//                 "qid": "2",
//                 "code": "def DoubleIt(x):\n\treturn x.upper()"
//             })
//         })
//         const json = await response.json()

//         console.log(json);
//     } catch (error) {
//         console.log(error);
//     }
// })();

// const autograde = async function(qid, code) {
//     try {
//         const response = await fetch("http://ec2-3-136-155-192.us-east-2.compute.amazonaws.com/autograde", {
//             method: 'POST',
//             headers: {
//                 'Content-Type': 'application/json'
//             },
//             body: JSON.stringify({
//                 "qid": qid,
//                 "code": code
//             })
//         })

//         const resp = await response.json()
//         // console.log(resp)
//         return resp;
//     } catch (error) {
//         console.log(error)
//     }
// }

// // console.log(trial)

// trial = autograde("2", "def DoubleIt(x):\n\treturn x.upper()")
// setTimeout(function() {
//     console.log(trial)
// }, 10000)


// const autograde = async function(qid, code) {
//     try {
//         const response = await fetch("http://ec2-3-136-155-192.us-east-2.compute.amazonaws.com/autograde", {
//             method: 'POST',
//             headers: {
//                 'Content-Type': 'application/json'
//             },
//             body: JSON.stringify({
//                 "qid": qid,
//                 "code": code
//             })
//         }).then((res) => res.json())


//         // response.then(function(result) {
//         //     console.log(result.json())
//         // })

//         // const resp = await response.json()
//         // return resp
//         // }).then(() => response.json())
//             // console.log(resp.json())
//             // return resp.json()
//         // })
        
//         // const resp = await response.json()
//         // console.log(resp)
//         // return resp;
//     } catch (error) {
//         console.log(error)
//     }
// }


// trial = autograde("2", "def DoubleIt(x):\n\treturn x.upper()")
// setTimeout(function() {
//     console.log(trial)
// }, 1000)


// const autograde = async function(qid, code) {
//     try {
//         const response = await fetch("http://ec2-3-136-155-192.us-east-2.compute.amazonaws.com/autograde", {
//             method: 'POST',
//             headers: {
//                 'Content-Type': 'application/json'
//             },
//             body: JSON.stringify({
//                 "qid": qid,
//                 "code": code
//             })
//         })

//         return Promise.resolve(response.json())
//         // const json = await response.json()
//         // console.log(json)
//         // return json
//     } catch (error) {
//         console.log(error)
//     }
// }


// trial = autograde("2", "def DoubleIt(x):\n\treturn x.upper()")
// console.log(trial)

// setTimeout(function() {
//     console.log(trial)
// }, 10000)


// const autograde = async function(qid, code) {
//     try {
//         const response = await fetch("http://ec2-3-136-155-192.us-east-2.compute.amazonaws.com/autograde", {
//             method: 'POST',
//             headers: {
//                 'Content-Type': 'application/json'
//             },
//             body: JSON.stringify({
//                 "qid": qid,
//                 "code": code
//             })
//         })
//         const json = await response.json()

//         // await new Promise((resolve, reject) => setTimeout(resolve, 6000))

//         return json

//     } catch (error) {
//         console.log(error)
//     }
// }


// trial = autograde("2", "def DoubleIt(x):\n\treturn x.upper()")
// console.log(trial)



// const autograde = async function(qid, code) {
//     try {
//         const response = await fetch("http://ec2-3-136-155-192.us-east-2.compute.amazonaws.com/autograde", {
//             method: 'POST',
//             headers: {
//                 'Content-Type': 'application/json'
//             },
//             body: JSON.stringify({
//                 "qid": qid,
//                 "code": code
//             })
//         })
//         const json = await response.json()
//         console.log(json)
//     } catch (error) {
//         console.log(error)
//     }
// }
// trial = autograde("2", "def DoubleIt(x):\n\treturn x.upper()")

