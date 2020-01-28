# welcome to my golang project to crawl movies 👋

<p align="center">
  <a href="https://github.com/dytt-bryce/server">
    <img src="https://github.com/dytt-bryce/server/workflows/pack%20server/badge.svg" alt='pack server'>
  </a>
  <img alt="version" src="https://img.shields.io/badge/version-1.0.0-green.svg?cacheSeconds=2592000">
</p>

* [live api](http://guygubaby.top:5000/api/v1/movies)
* [电影淘淘](https://www.dytt.com/)

## how to run

clone the project and run the follow command 

``` bash
go run main.go
```

then the server will listen port `5000` 

## api

### movies `电影` 

* url `/api/v1/movies` 
* params 
  + page (default is `1` )
* response

``` json
[
  {
    "Name": "釜山行 ",
    "URL": "http://m.dytt.com/xiazai/id959.html",
    "Cover": "http://m.dytt.com/upload/vod/2019-08-30/15671778360.png",
    "Score": "",
    "Resolution": "HD"
  },
  {
    "Name": "土耳其战狼电影版",
    "URL": "http://m.dytt.com/xiazai/id67785.html",
    "Cover": "http://m.dytt.com/upload/vod/2020-01-18/157928775811.jpg",
    "Score": "",
    "Resolution": "HD"
  }
]
```

### soap `电视剧` 

* url `/api/v1/soap` 
* params 
  + page (default is `1` )
* response
  + same as top

### zongyi `综艺` 

* url `/api/v1/zongyi` 
* params 
  + page (default is `1` )
* response
  + same as top

### cartoon `动漫` 

* url `/api/v1/cartoon` 
* params 
  + page (default is `1` )
* response
  + same as top

