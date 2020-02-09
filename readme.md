# beego2ts

## get
```shell script
go get github.com/z-sage/beego2ts
```

## use

```go
GenerateDocs(curpath string) swagger.Swagger
BuildApi(path string)
BuildSwagger(path string)
```

BuildApi生成ts文件，如需js需自行编译
```shell script
tsc api.ts
```

约定：

* get 和 delete 方法只能使用 query 和 path 方式传值
* path方式最多只能传一个值，名称固定为id，path 和 query 方式可以同时使用
* post 和 put 方法只能使用 formdata 和 body 方式传值