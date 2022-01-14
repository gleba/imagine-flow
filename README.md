# imagine-flow

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gleba/imagine-flow)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/gpanteleev/imagine-flow)
![Docker Image Version (tag latest semver)](https://img.shields.io/docker/v/gpanteleev/imagine-flow/latest)

Image processing http-server with cache in memory 
WIP

Ready for platform
- linux/arm64
- linux/amd64

### docker-compose example
```
version: "3.2"
services:
  imagine:
    container_name: imagine
    image: gpanteleev/imagine-flow:latest
    volumes:
      - /root/web-imagine:/imagine
    expose:
      - "5555"
    ports:
      - 5555:5555
    environment:
      PUBLIC_IMAGES_FOLDER: /imagine
      ALIVE_TIME: 24h      
      PORT: 5555
      URL_PREFIX: /i
```
these env-vars are default

### url patterns
- base resize
  -  URL_PREFIX/re?size=MAX_PIXELSIZE&f=FILENAME&as='FORMAT' (default webp, allow jpg)
  - `http://imagine:5555/i/re?size=100&f=myfolder/mypic.jpg`
- original
  - URL_PREFIX/FILENAME
