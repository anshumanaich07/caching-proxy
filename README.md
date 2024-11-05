## Caching Proxy working mechanism
* Request from CLI is sent to the Caching server  
* The Caching server checks if it already has the response for it, meaning if it already had received the similar request 
* If the request and its response exists, it returns
* Else, it forwards it to the actual server.

## Technologies 
* Golang for backend
* Redis for caching
* MongoDB for database
 

## Run 
### caching server
```bash
  cd caching-server
  caching-proxy --port <number> --origin <url>
```

### server
```bash
  cd server
  make run
```

## Clear cache
```bash
  caching-proxy --clear-cache
```

* `--port`: Port of the caching server
* `--origin`: URL of the main server, to which the requests will be forwarded to
* `--clear-cache`: Clears the cache


## TODO
- [ ] Change the execution method of both the servers
- [X] Write a caching server 
- [ ] Refactor the caching server where it checks for cache miss
