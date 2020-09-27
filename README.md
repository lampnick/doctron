## Doctron description
Doctron is a sample,fast,high quality document convert tool.Supply html convert to pdf, html convert to image(like jpeg,png), add  watermarks to pdf, convert pdf to images etc.

## Online experience
open the following website to have a try.
[Doctron Live Demo](http://doctron.lampnick.com)

## Table of Contents

- [Features](#features)
- [Deploy](#Deploy)
- [Quick Start](#quick-start)
  * [Html convert to pdf](#html-convert-to-pdf)
  * [Html convert to image](#html-convert-to-image)
  * [Pdf add watermark](#pdf-add-watermark)
  * [Pdf convert to image](#pdf-convert-to-image)
- [License](#license)

## Features
- Html convert to pdf/image using chrome kernel to guarantee what you see is what you get.
- Easy deployment.(Using docker,kubernetes.)
- Rich transformation parameters.
- Customize page size from html convert to pdf or image.
- Serverless supported.

## Installing
- Using docker
```
#using default config
docker run -p 8080:8080 --rm --name doctron-alpine lampnick/doctron  
#using custom config
docker run -p 8080:8080 --rm --name doctron-alpine \
-v /path/to/doctron/conf/doctron.yaml:/doctron.yaml \
lampnick/doctron  
```
- Using k8s
```
kubectl apply -f https://github.com/lampnick/doctron/manifests/k8s-doctron.yaml  
```
- From source code
```
git clone https://github.com/lampnick/doctron.git
cd doctron
go run main.go 
```

## Quick Start
### Html convert to pdf
- basic
```
http://127.0.0.1:8080/convert/html2pdf?u=doctron&p=lampnick&url=<url>  
```
- custom size
```
http://127.0.0.1:8080/convert/html2pdf?u=doctron&p=lampnick&url=<url>&marginTop=0&marginLeft=0&marginRight=0&marginbottom=0&paperwidth=4.1  
```
### Html convert to image
- basic
```
http://127.0.0.1:8080/convert/html2image?u=doctron&p=lampnick&url=<url>  
```
- custom size
```
http://127.0.0.1:8080/convert/html2image?u=doctron&p=lampnick&url=<url>&customClip=true&clipX=0&clipY=0&clipWidth=400&clipHeight=1500&clipScale=2&format=jpeg&Quality=80  
```
### Pdf add watermark
- coming soon

### Pdf convert to image
- coming soon

## License

Doctron is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/lampnick/doctron/blob/master/LICENSE)