# gee_study
从0搭建web框架gee
ref：https://geektutu.com/post/gee.html

1. http 基础
2. 封装 context : 对web框架而言，无非就是根据 request 来构造响应response。但是这两个对象提供的接口都太细粒度了，比如我们构造一个完整的响应其实还需要考虑消息头Header 和 消息体Body的。因此需要进行一些必要的封装
。初次之外，框架还需要支持一些额外的功能，比如解析动态路由，比如 /hello/:name ，再比如框架要支持中间件，中间件的信息存放在哪里呢？ context 随着每个请求的出现而产生，请求的结束而销毁

3. 使用前缀树实现动态路由。具备两个功能：
   1. 参数匹配：例如 /p/:lang/doc，可以匹配 /p/c/doc 和 /p/go/doc。
   2. 通配符*：例如 /static/*filepath，可以匹配/static/fav.ico，也可以匹配/static/js/jQuery.js
