# axios的使用

## 一、安装

```
npm install axios --save
```

## 二、使用

1.get请求

```
// 为给定 ID 的 user 创建请求
axios.get('/user?ID=12345')
  .then(function (response) {   // 成功的回调
    console.log(response);
  })
  .catch(function (error) {  // 失败的回调
    console.log(error);
  });
```

2.get请求加参数

```
// 上面的请求也可以这样做
axios.get('/user', {    // 参数是在配置中加，配置中还可以指定header等信息
    params: {  
      ID: 12345
    }
  })
  .then(function (response) {
    console.log(response);
  })
  .catch(function (error) {
    console.log(error);
  });
```

3.post请求

```
axios.post('/user', {
    firstName: 'Fred',
    lastName: 'Flintstone'
  })
  .then(function (response) {
    console.log(response);
  })
  .catch(function (error) {
    console.log(error);
  });
```

4.使用统一配置

```
const instance = axios.create({
  baseURL: 'https://some-domain.com/api/',
  timeout: 1000,
  headers: {'X-Custom-Header': 'foobar'}
});


instance.get('/user', {    // 参数是在配置中加，配置中还可以指定header等信息
    params: {  
      ID: 12345
    }
  })
  .then(function (response) {
    console.log(response);
  })
  .catch(function (error) {
    console.log(error);
  });
```

