# gitlab-api
Simple application to interact with gitlab api.
This app work with 2 kind of cache : 1) simple string-> string map and 2) Redis,
You can chose which one to use.

## Dependencies

name     | repo
------------- | -------------
  go          | https://golang.org/doc/install 
   gin        | https://github.com/gin-gonic/gin
  go redis    | https://github.com/go-redis/redis 
  docker      | https://www.docker.com/
  
You must have [go](https://golang.org/doc/install) and [gin](https://github.com/gin-gonic/gin),
First you need to install go and then you can run `go get https://github.com/gin-gonic/gin.git` to have gin.

If you want to use redis as cache you need [Docker](https://www.docker.com/) installed on your machine 
also you need to run `go get https://github.com/go-redis/redis` too.

 
## Usage
If you want to use redis as cache:

[Run `docker run --name redis-usdb -p "yourPort":6379 -d redis` to connect redis to port "yourPort".
Set cache_port in config.json to yourPort.]

Build and run **main.go** file(`go run main.go`) to start the app.

(default port is 8080 but if this port of your system is busy you can change it in config.json(listener_port))

## Endpoints
baseUrl: localhost:listener_port 

  `POST baseUrl/token ==> to set your access token and userId body :{ userId: string, gitlabAccessToken: string }` 
  
  `GET baseUrl/projects/:id/repository_tree ==> get repository tree of project with id mentioned in request body and access token for user with userId set in request header with key as <userId>`
  
  `GET baseUrl/ping ==> simple get endpoint used for cloud to hit it and understand our web server is still running`
