# 1  移植后的文件 无法执行

```bash
> vue-project@0.1.0 serve
> vue-cli-service serve

sh: vue-cli-service: command not found
```



## 解决方案：

在命令行中先进入该文件的路径，然后在输入
`npm install`

就会下载node_modules文件
下载完成后，使用 npm run serve 运行项目





# 2 测试自定义shallowRef 出现问题

嵌套对象可以访问





# 3 Suspense(不确定的)

异步访问 细节





# 4 [Failed to resolve loader: less-loader You may need to install it. Failed to resolve loader: less-loader](https://www.cnblogs.com/xqschool/p/14120297.html)



```
npm install less ``-``-``save``-``dev
```