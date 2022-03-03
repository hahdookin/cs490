# cs490
Distributed 3-layer architecture designed to be a thorough autograder.

## Demo
- [Click Here for Demo](https://afsaccess4.njit.edu/~cap82/)

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
^2 - POST (Curl)
```

## :microscope: Technologies
- Languages: `php`, `css`, `html`, `javascript`, `go`
- Packages: `curl`, `php-curl`, `nginx`


<!-- :heavy_check_mark: vs :x: vs :soon: -->
| Release Version | Completion
| :-----:  | :-----:
| Alpha     | :heavy_check_mark:
| Beta      | :soon:
| Release   | :x:


## :books: Resources
- [Debug your Server](https://ist.njit.edu/debug-your-code)
- [AFS Permissions](https://ist.njit.edu/afs-permissions)
- [Reverse Proxy Golang Web App using Nginx on Ubuntu](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04)