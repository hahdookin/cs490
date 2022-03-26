[![Go Report Card - Middleware](https://goreportcard.com/badge/github.com/AOrps/cs490/middleware)](https://goreportcard.com/report/github.com/AOrps/cs490/middleware)
# cs490
Distributed 3-layer architecture designed to be a thorough autograder.

## Demo
- [Click Here for Demo](http://exam-central.surge.sh/)

## :triangular_ruler: Architecture
```txt
.------------.               .------------.               .------------.
| Front-end  |   <---S--->   | Middleware |   <---S--->   |  Back-end  |
'------------'               '------------'               '------------'
    /|   POST (Fetch)                                           /|
     |                                                           |  SQL
     |/  JSON                                                    |/
.------------.                                            .------------.
|  Browser   |                                            |  Database  | 
'------------'                                            '------------'

========================================================================
Legend:
     
<---S--->
^1      ^2

^1 - JSON
^2 - POST
```

## :microscope: Technologies
- Language(s): `html`, `javascript`, `go`, `python`, `css`, `shell (bash)`
- Package(s): `nginx`
- Framework(s): `vue`


<!-- :heavy_check_mark: vs :x: vs :soon: -->
| Release Version | Completion
| :-----:  | :-----:
| Alpha     | :heavy_check_mark:
| Beta      | :heavy_check_mark:
| Release   | :soon:


## :books: Resources
- [Debug your Server](https://ist.njit.edu/debug-your-code)
- [AFS Permissions](https://ist.njit.edu/afs-permissions)
- [Reverse Proxy Golang Web App using Nginx on Ubuntu](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04)
- [Enable CORS on Golang](https://flaviocopes.com/golang-enable-cors/)
- [Enable CORS on Nginx](https://enable-cors.org/server_nginx.html)