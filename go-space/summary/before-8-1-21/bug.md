1 导入文件失败， 查看 导入包路径是否准确。



2 

```bash
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

panic: template: array.html:18: missing value for range

goroutine 1 [running]:
html/template.Must(...)
        /usr/local/go/src/html/template/template.go:374
github.com/gin-gonic/gin.(*Engine).LoadHTMLGlob(0xc00032e280, 0x97eec0, 0xd)
        /home/dart/development/GoPath/pkg/mod/github.com/gin-gonic/gin@v1.6.3/gin.go:185 +0x365
main.main()
        /home/dart/DoThinking/GoSpace/GoGin_2021_3_11/v4_composite_data_render/main.go:13 +0x49
exit status 2
```



solution:

```html
  {{range $i, $v := .}}	// 加点
        {{$i}} --> {{$v}} <br>
  {{end}}
```



3.

```bash
2021/03/12 11:12:23 [Recovery] 2021/03/12 - 11:12:23 panic recovered:
GET /map_struct HTTP/1.1
Host: localhost:9000
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7
Cache-Control: max-age=0
Connection: keep-alive
Sec-Ch-Ua: "Google Chrome";v="89", "Chromium";v="89", ";Not A Brand";v="99"
Sec-Ch-Ua-Mobile: ?0
Sec-Fetch-Dest: document
Sec-Fetch-Mode: navigate
Sec-Fetch-Site: none
Sec-Fetch-User: ?1
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36


template: map_struct.html:13: function "user" not defined
/usr/local/go/src/html/template/template.go:374 (0x87b9da)
        Must: panic(err)
/home/dart/development/GoPath/pkg/mod/github.com/gin-gonic/gin@v1.6.3/render/html.go:74 (0x87b53c)
        HTMLDebug.loadTemplate: return template.Must(template.New("").Delims(r.Delims.Left, r.Delims.Right).Funcs(r.FuncMap).ParseGlob(r.Glob))
/home/dart/development/GoPath/pkg/mod/github.com/gin-gonic/gin@v1.6.3/render/html.go:61 (0x87e4f7)
        HTMLDebug.Instance: Template: r.loadTemplate(),
/home/dart/development/GoPath/pkg/mod/github.com/gin-gonic/gin@v1.6.3/context.go:850 (0x8853cc)
        (*Context).HTML: instance := c.engine.HTMLRender.Instance(name, obj)
/home/dart/DoThinking/GoSpace/GoGin_2021_3_11/v4_composite_data_render/handlerpages/composite_data.go:72 (0x896ff2)
        MapAndStructHandler: c.HTML(http.StatusOK, "map_struct.html", m1)
/home/dart/development/GoPath/pkg/mod/github.com/gin-gonic/gin@v1.6.3/context.go:161 (0x8951d0)
        (*Context).Next: c.handlers[c.index](c)
/home/dart/development/GoPath/pkg/mod/github.com/gin-gonic/gin@v1.6.3/recovery.go:83 (0x8951b7)
        RecoveryWithWriter.func1: c.Next()
/home/dart/development/GoPath/pkg/mod/github.com/gin-gonic/gin@v1.6.3/context.go:161 (0x894253)
        (*Context).Next: c.handlers[c.index](c)
/home/dart/development/GoPath/pkg/mod/github.com/gin-gonic/gin@v1.6.3/logger.go:241 (0x894212)
        LoggerWithConfig.func1: c.Next()
/home/dart/development/GoPath/pkg/mod/github.com/gin-gonic/gin@v1.6.3/context.go:161 (0x88b38f)
        (*Context).Next: c.handlers[c.index](c)
/home/dart/development/GoPath/pkg/mod/github.com/gin-gonic/gin@v1.6.3/gin.go:409 (0x88b376)
        (*Engine).handleHTTPRequest: c.Next()
/home/dart/development/GoPath/pkg/mod/github.com/gin-gonic/gin@v1.6.3/gin.go:367 (0x88ae2c)
        (*Engine).ServeHTTP: engine.handleHTTPRequest(c)
/usr/local/go/src/net/http/server.go:2887 (0x69e2a2)
        serverHandler.ServeHTTP: handler.ServeHTTP(rw, req)
/usr/local/go/src/net/http/server.go:1952 (0x6996cc)
        (*conn).serve: serverHandler{c.server}.ServeHTTP(w, w.req)
/usr/local/go/src/runtime/asm_amd64.s:1371 (0x46d780)
        goexit: BYTE    $0x90   // NOP
```



solution:

​	plus a dot	`.user`