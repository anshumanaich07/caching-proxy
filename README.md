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
  cd caching-proxy
  make run ARGS="--port 8080 --origin localhost"
```

### server
```bash
  cd server
  make run
```

## Clear cache
```bash
  cd caching-proxy
  make run ARGS="--clear" 
```


* `--port`: Port of the origin server
* `--origin`: URL of the main server, to which the requests will be forwarded to
* `--clear`: Clears the cache


## TODO
- [X] Add mechanism to clear cache
- [X] Add flags on startup for both the servers
- [X] Change the execution method of both the servers
- [X] Write a caching server 
- [X] Refactor the caching server where it checks for cache miss
- [ ] Add tests
