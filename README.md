# HelloDocker

This repository demonstrates a dummy web app, running inside docker container. 

# Build 

```$sh
docker-compose up --build
```
# Run 

```$xslt
curl 127.0.0.1:60235/message
```

# Test Docker Network 

```$xslt
docker exec -it webapp sh
ping mock-server
```
