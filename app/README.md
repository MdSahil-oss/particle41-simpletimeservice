# Build & Run container

If you want to build the container image of this microservice execute following commands

```bash
$ docker build -t mdsahiloss/simple-time-service:1.0 .
```

Once the image has built then run following command to run an instance of container image

```bash
$ docker run -p 8080:8080 mdsahiloss/simple-time-service:1.0
```


