### [English](README.md) | [中文](README_ZH.md)

### 目录

- [Doctron简介](#doctron简介)
- [在线体验](#在线体验)
- [鼓励一波](#鼓励一波)
- [特性](#特性)
- [安装](#安装)
- [快速开始](#快速开始)
  - [Html转pdf](#html转pdf)
        - [基础转换](#基础转换)
        - [自定义大小](#自定义大小)
        - [支持的参数](#支持的参数)
  - [Html转图片](#html转图片)
        - [基础转换](#基础转换-1)
        - [自定义大小](#自定义大小-1)
        - [支持的参数](#支持的参数-1)
  - [Pdf加水印](#pdf加水印)
        - [添加图片水印](#添加图片水印)
        - [支持的参数](#支持的参数-2)
  - [Pdf转image](#pdf转image)
        - [coming soon](#coming-soon)
- [Doctron Client](#doctron-client)
  - [Doctron go client](#doctron-go-client)
        - [doctron-client-go](#doctron-client-go)
  - [Doctron php client](#doctron-php-client)
        - [doctron-client-php](#doctron-client-php)
- [License](#license)
## Doctron简介
Doctron是基于Docker、无状态、简单、快速、高质量的文档转换服务。目前支持将html转为pdf、图片(使用chrome(Chromium)浏览器内核，保证转换质量)。支持PDF添加水印。

## 在线体验
您可以打开下面的链接在线体验转换质量，由于服务器配置较低，以及网络原因，转换可能会慢一点，实际部署到服务器速度会不一样。
[Doctron Live Demo](http://doctron.lampnick.com)

## 鼓励一波
如果您觉得Doctron这个服务还不错，请给个小星星，fork一下哦。您的鼓励是我前进的动力！


## 特性
- 使用chrome内核保证高质量将HTML转为pdf/图片。
- 简易部署(提供docker镜像,Dockerfile以及k8s yaml配置文件)。
- 支持丰富的转换参数。
- 转为pdf和图片支持自定义大小。
- 无状态服务支持。

## 安装
- 使用Docker
```
#使用默认配置
docker run -p 8080:8080 --rm --name doctron-alpine lampnick/doctron  
#使用自定义配置文件
docker run -p 8080:8080 --rm --name doctron-alpine \
-v <本地doctron.yaml配置文件>:/doctron.yaml \
lampnick/doctron  
```
- 使用k8s
```
kubectl apply -f https://raw.githubusercontent.com/lampnick/doctron/master/manifests/k8s-doctron.yaml
```
- 从源码运行
```
git clone https://github.com/lampnick/doctron.git
cd doctron
go run main.go 
```

- 使用`go get`安装
```
go get -v github.com/lampnick/doctron
下载完成之后直接运行
doctron
```

## 快速开始
### Html转pdf
###### 基础转换
```
http://127.0.0.1:8080/convert/html2pdf?u=doctron&p=lampnick&url=<url>  
```
###### 自定义大小
```
http://127.0.0.1:8080/convert/html2pdf?u=doctron&p=lampnick&url=<url>&marginTop=0&marginLeft=0&marginRight=0&marginbottom=0&paperwidth=4.1  
```
###### 支持的参数
- u/username // doctron 用户名
- p/password // doctron 密码
- uploadKey // 上传到OSS的文件名
- url //需要转换的html URL
- landscape // 横向打印格式.默认false.表示纵向
- displayHeaderFooter // 是否显示页头页尾，默认false.
- printBackground // 是否打印背景。默认false.
- scale // 缩放比例. 默认1.
- paperWidth // 纸张宽度，单位英尺。默认8.5英尺.
- paperHeight // 纸张高度，单位英尺。默认11英尺.
- marginTop // 上外边距，单位英尺。默认纸0.4英尺（1厘米）. 
- marginBottom // 下外边距，单位英尺。默认纸0.4英尺（1厘米）. 
- marginLeft // 左外边距，单位英尺。默认纸0.4英尺（1厘米）. =
- marginRight // 右外边距，单位英尺。默认纸0.4英尺（1厘米）. 
- pageRanges // 需要打印的PDF的页数。默认为空字符串，表示所有页面.
- ignoreInvalidPageRanges // 是否静默的忽略掉不可用的但是成功解析的页面。例如'3-2',默认false.
- WaitingTime // 页面加载后等待时长. 默认为0代表不等待. 单位：毫秒

### Html转图片
###### 基础转换
```
http://127.0.0.1:8080/convert/html2image?u=doctron&p=lampnick&url=<url>  
```
###### 自定义大小
```
http://127.0.0.1:8080/convert/html2image?u=doctron&p=lampnick&url=<url>&customClip=true&clipX=0&clipY=0&clipWidth=400&clipHeight=1500&clipScale=2&format=jpeg&Quality=80  
```
###### 支持的参数
- u/username // doctron 用户名
- p/password // doctron 密码
- uploadKey // 上传到OSS的文件名
- url // 需要转换的html URL
- format // 图片压缩格式(defaults to png)，还支持jpeg.
- quality // jpeg图片压缩质量 [0..100] (jpeg only).
- customClip // 只有设置了这个值，下面的裁剪才会生效.否则不生效.
- clipX // 裁剪X轴方向距离.
- clipY // 裁剪Y轴方向距离.
- clipWidth // 裁剪宽度.
- clipHeight // 裁剪高度.
- WaitingTime // 页面加载后等待时长. 默认为0代表不等待. 单位：毫秒

### Pdf加水印
###### 添加图片水印
```
http://127.0.0.1:8080/convert/pdfAddWatermark?u=doctron&p=lampnick&url=<pdf url>&imageUrl=<image url>
```
###### 支持的参数
- u/username // doctron 用户名
- p/password // doctron 密码
- uploadKey // 上传到OSS的文件名
- url // 需要转换的html URL
- imageUrl // 图片水印URL,支持png/jpeg

### Pdf转image
###### coming soon

## Doctron Client
### Doctron go client
###### [doctron-client-go](https://github.com/lampnick/doctron-client-go)

### Doctron php client
###### [doctron-client-php](https://github.com/lampnick/doctron-client-php)

## License

Doctron is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/lampnick/doctron/blob/master/LICENSE)
