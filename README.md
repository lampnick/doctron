### [English](README.md) | [中文](README_ZH.md)

### Table of Contents

- [Doctron description](#doctron-description)
- [Online experience](#online-experience)
- [Encourage](#encourage)
- [Features](#features)
- [Installing](#installing)
- [Quick Start](#quick-start)
  - [Html convert to pdf](#html-convert-to-pdf)
        - [basic](#basic)
        - [custom size](#custom-size)
        - [support params](#support-params)
  - [Html convert to image](#html-convert-to-image)
        - [basic](#basic-1)
        - [custom size](#custom-size-1)
        - [support params](#support-params-1)
  - [Pdf add watermark](#pdf-add-watermark)
        - [add image watermark](#add-image-watermark)
        - [support params](#support-params-2)
  - [Pdf convert to image](#pdf-convert-to-image)
        - [coming soon](#coming-soon)
- [Doctron Client](#doctron-client)
  - [Doctron go client](#doctron-go-client)
        - [doctron-client-go](#doctron-client-go)
  - [Doctron php client](#doctron-php-client)
        - [doctron-client-php](#doctron-client-php)
- [License](#license)
## Doctron description
Doctron is a Docker-powered,serverless,sample,fast,high quality document convert tool.Supply html convert to pdf(html2pdf), html convert to image(html2image like jpeg,png),which using chrome(Chromium) kernel, add watermarks to pdf, convert pdf to images etc.

## Online experience
open the following website to have a try. Conversion may be slower due to low server configuration.
[Doctron Live Demo](http://doctron.lampnick.com)

## Encourage
If you feel doctron is not bad, give me a star and fork.Star and fork is my greatest encouragement！

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
kubectl apply -f https://raw.githubusercontent.com/lampnick/doctron/master/manifests/k8s-doctron.yaml
```
- From source code
```
git clone https://github.com/lampnick/doctron.git
cd doctron
go run main.go 
```

- install doctron using `go get` 
```
go get -v github.com/lampnick/doctron
When finished then directly run
doctron
```

## Quick Start
### Html convert to pdf
###### basic
```
http://127.0.0.1:8080/convert/html2pdf?u=doctron&p=lampnick&url=<url>  
```
###### custom size
```
http://127.0.0.1:8080/convert/html2pdf?u=doctron&p=lampnick&url=<url>&marginTop=0&marginLeft=0&marginRight=0&marginbottom=0&paperwidth=4.1  
```
###### support params
- u/username // doctron username
- p/password // doctron password
- uploadKey // upload to oss key
- url // need convert html url
- landscape // Paper orientation. core.Defaults to false.
- displayHeaderFooter // Display header and footer. core.Defaults to false.
- printBackground // Print background graphics. core.Defaults to false.
- scale // Scale of the webpage rendering. core.Defaults to 1.
- paperWidth // Paper width in inches. core.Defaults to 8.5 inches.
- paperHeight // Paper height in inches. core.Defaults to 11 inches.
- marginTop // Top margin in inches. core.Defaults to 1cm (~0.4 inches).
- marginBottom // Bottom margin in inches. core.Defaults to 1cm (~0.4 inches).
- marginLeft // Left margin in inches. core.Defaults to 1cm (~0.4 inches).
- marginRight // Right margin in inches. core.Defaults to 1cm (~0.4 inches).
- pageRanges // Paper ranges to print, e.g., '1-5, 8, 11-13'. core.Defaults to the empty string, which means print all pages.
- ignoreInvalidPageRanges // Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. core.Defaults to false.
- WaitingTime // Waiting time after the page loaded. Default 0 means not wait. unit:Millisecond

### Html convert to image
###### basic
```
http://127.0.0.1:8080/convert/html2image?u=doctron&p=lampnick&url=<url>  
```
###### custom size
```
http://127.0.0.1:8080/convert/html2image?u=doctron&p=lampnick&url=<url>&customClip=true&clipX=0&clipY=0&clipWidth=400&clipHeight=1500&clipScale=2&format=jpeg&Quality=80  
```
###### support params
- u/username // doctron username
- p/password // doctron password
- uploadKey // upload to oss key
- url // need convert html url
- format // Image compression format (defaults to png).
- quality // Compression quality from range [0..100] (jpeg only).
- customClip //if set this value, the below clip will work,otherwise not work!
- clipX // Capture the screenshot of a given region only.X offset in device independent pixels (dip).
- clipY // Capture the screenshot of a given region only.Y offset in device independent pixels (dip).
- clipWidth // Capture the screenshot of a given region only.Rectangle width in device independent pixels (dip).
- clipHeight // Capture the screenshot of a given region only.Rectangle height in device independent pixels (dip).
WaitingTime // Waiting time after the page loaded. Default 0 means not wait. unit:Millisecond

### Pdf add watermark
###### add image watermark
```
http://127.0.0.1:8080/convert/pdfAddWatermark?u=doctron&p=lampnick&url=<pdf url>&imageUrl=<image url>
```
###### support params
- u/username // doctron username
- p/password // doctron password
- uploadKey // upload to oss key
- url // need convert html url
- imageUrl // watermark image url,support png/jpeg

### Pdf convert to image
###### coming soon

## Doctron Client
### Doctron go client
###### [doctron-client-go](https://github.com/lampnick/doctron-client-go)

### Doctron php client
###### [doctron-client-php](https://github.com/lampnick/doctron-client-php)

## License

Doctron is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/lampnick/doctron/blob/master/LICENSE)
