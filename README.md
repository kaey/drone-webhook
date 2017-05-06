# drone-webhook

Example .drone.yml
```
pipeline:
  deploy:
    image: kaey/drone-webhook
    url: https://deploy-handler.corp.org
    method: post
    body: "{}"
```

Building:
```
$ CGO_ENABLED=0 go build -v
$ docker build -t kaey/drone-webhook .
```
