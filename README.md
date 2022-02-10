# cs490


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
- Languages: `php`, `css`, `html`, `javascript`
- Packages: `curl`, `php-curl`


<!-- :heavy_check_mark: vs :x: vs :soon: -->
| Release Version | Completion
| :-----:  | :-----:
| Alpha     | :heavy_check_mark:
| Beta      | :soon:
| Release   | :x:


## :books: Resources
- [Debug your Server](https://ist.njit.edu/debug-your-code)
- [AFS Permissions](https://ist.njit.edu/afs-permissions)