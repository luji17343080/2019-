# 模仿 Github，设计一个博客网站的 API  
## Github API
### Current version  
默认情况下，所有https://api.github.com接收v3 版本的REST API的请求，可以通过`Accept header`明确请求此版本
```
Accept: application/vnd.github.v3+json
```

### Schema
所有API访问都是通过HTTPS进行的，并且可以通过访问https://api.github.com。所有数据都以JSON的形式发送和接收。  
```JSON
curl -i https://api.github.com/users/octocat/orgs
HTTP/1.1 200 OK
Server: nginx
Date: Fri, 12 Oct 2012 23:33:14 GMT
Content-Type: application/json; charset=utf-8
Connection: keep-alive
Status: 200 OK
ETag: "a00049ba79152d03380c34652f2cb612"
X-GitHub-Media-Type: github.v3
X-RateLimit-Limit: 5000
X-RateLimit-Remaining: 4987
X-RateLimit-Reset: 1350085394
Content-Length: 5
Cache-Control: max-age=0, private, must-revalidate
X-Content-Type-Options: nosniff
```
包含空白字段，`null`而不是将其省略。  
所有时间戳以ISO 8601格式返回：
> YYYY-MM-DDTHH:MM:SSZ  
  
  
比如访问我的github：`https://api.github.com/users/luji17343080`会得到如下信息：  
```json
{
  "login": "luji17343080",
  "id": 43690730,
  "node_id": "MDQ6VXNlcjQzNjkwNzMw",
  "avatar_url": "https://avatars2.githubusercontent.com/u/43690730?v=4",
  "gravatar_id": "",
  "url": "https://api.github.com/users/luji17343080",
  "html_url": "https://github.com/luji17343080",
  "followers_url": "https://api.github.com/users/luji17343080/followers",
  "following_url": "https://api.github.com/users/luji17343080/following{/other_user}",
  "gists_url": "https://api.github.com/users/luji17343080/gists{/gist_id}",
  "starred_url": "https://api.github.com/users/luji17343080/starred{/owner}{/repo}",
  "subscriptions_url": "https://api.github.com/users/luji17343080/subscriptions",
  "organizations_url": "https://api.github.com/users/luji17343080/orgs",
  "repos_url": "https://api.github.com/users/luji17343080/repos",
  "events_url": "https://api.github.com/users/luji17343080/events{/privacy}",
  "received_events_url": "https://api.github.com/users/luji17343080/received_events",
  "type": "User",
  "site_admin": false,
  "name": null,
  "company": null,
  "blog": "",
  "location": null,
  "email": null,
  "hireable": null,
  "bio": null,
  "public_repos": 15,
  "public_gists": 0,
  "followers": 0,
  "following": 0,
  "created_at": "2018-09-29T02:25:32Z",
  "updated_at": "2019-09-26T09:00:31Z"
}
```


## Blog API
- 参考Github API的基本信息：让https://api.blog.com接收v3 版本的REST API的请求，通过通过`Accept header`明确请求此版本
    ```
    Accept: application/vnd.blog.v3+json
    ```

- 通过访问https://api.blog.com，数据以JSON格式发送和接受（以user为luji为例）
    ```JSON
    curl -i https://api.blog.com/users/luji
    HTTP/1.1 200 OK
    Server: nginx
    Date: Sat, 23 Nov 2019 14:15:14 GMT
    Content-Type: application/json; charset=utf-8
    Connection: keep-alive
    Status: 200 OK
    ETag: "a00049ba79152d03380c34652f2cb612"
    X-GitHub-Media-Type: blog.v3
    X-RateLimit-Limit: 5000
    X-RateLimit-Remaining: 4987
    X-RateLimit-Reset: 1350085394
    Content-Length: 5
    Cache-Control: max-age=0, private, must-revalidate
    X-Content-Type-Options: nosniff
    ```
- API 功能
    - 通过`GET`请求获取所有博客摘要表示  
        ```
        GET /user/luji/allblogs
        ```
    - 通过`GET`请求获取指定名称博客的信息  
        ```
        GET /user/luji/blogname
        ```

    - 使用`POST`请求新建blog  
        ```
        POST /user/luji/blogname
        ```
    - 使用`PUT`请求修改一篇博客内容  
        ```
        PUT /user/luji/blogname
        ```
    - 使用`Delete`请求删除博客内容  
        ```
        Delete /user/luji/blogname
        ```   
- 认证方式  
有两种方法可以通过Blog API v3进行身份验证。需要身份验证的请求在某些地方将返回`404 Not Found`，而不是 `403 Forbidden`。这是为了防止私有存储库意外泄露给未经授权的用户。  
    - 基本认证
        ```json
        curl -u "username" https://api.blog.com
        ```
    - OAuth2令牌（在标头中发送）  
        ```json
        curl -H "Authorization: token OAUTH-TOKEN" https://api.blog.com
        ```
    - 登录限制失败  
    使用无效的凭据进行身份验证将返回`401 Unauthorized`：
        ```json
        curl -i https://api.blog.com -u foo:bar
        HTTP/1.1 401 Unauthorized
        {
        "message": "Bad credentials",
        "documentation_url": "https://developer.blog.com/v3"
        }
        ```  
          
        在短时间内检测到多个具有无效凭据的请求后，API会临时拒绝该用户的所有身份验证尝试（包括具有有效凭据的请求）`403 Forbidden`：
        ```json
        curl -i https://api.blog.com -u valid_username:valid_password
        HTTP/1.1 403 Forbidden
        {
        "message": "Maximum number of login attempts exceeded. Please try again later.",
        "documentation_url": "https://developer.blog.com/v3"
        }   
        ```  
- 参数  
许多API方法采用可选参数。对于`GET`请求，任何未在路径中指定为段的参数都可以作为HTTP查询字符串参数传递：  
    ```json
    curl -i "https://api.blog.com/luji/blogname?state=closed"
    ```  
    对于`POST`、`PUT`和`DELETE`请求，URL中没有包含的参数，应该用“application/ JSON”的内容类型编码为JSON:  
    ```json
    curl -i -u username -d '{"scopes":["public_blogs"]}' https://api.blog.com/authorizations
    ```
- 根断点  
可以通过`GET`向跟断点发出请求，以获取REST API v3支持的所有端点类别：  
    ```json
    curl https://api.github.com
    ```

- 错误处理
    - 发送无效的JSON将导致`400 Bad Request`响应
        ```json
        HTTP/1.1 400 Bad Request
        Content-Length: 35

        {"message":"Problems parsing JSON"}
        ```
    - 发送错误类型的JSON值将导致`400 Bad Request`响应
        ```json
        HTTP/1.1 400 Bad Request
        Content-Length: 40

        {"message":"Body should be a JSON object"}
        ```
    - 发送无效的字段将导致`422 Unprocessable Entity` 响应
        ```json
        HTTP/1.1 422 Unprocessable Entity
        Content-Length: 149

        {
        "message": "Validation Failed",
        "errors": [
            {
            "resource": "Issue",
            "field": "title",
            "code": "missing_field"
            }
        ]
        }
        ```
    下面是可能的验证错误代码：
  
    |错误名称|	描述 |
    | :-------------- | :------------ |
    |missing|	这意味着资源不存在|
    |missing_field|	这意味着尚未设置资源上的必填字段|
    |invalid|	这意味着字段格式无效。该资源的文档应该能够为您提供更多具体信息|
    |already_exists|	这意味着另一个资源具有与此字段相同的值。在必须具有某些唯一键（例如标签名称）的资源中可能会发生这种情况|  
- HTTP重定向  
`API v3`在适当的地方使用`HTTP重定向`。客户应假定任何请求都可能导致重定向。接收HTTP重定向`不是错误`，客户端应遵循该重定向。重定向响应将具有一个 `Location` 标头字段，其中包含客户端应向其重复请求的资源的`URI`:
    |状态码|	描述|
    | :---------------| :------------ |
    |301|	永久重定向：用于发出请求的`URI`已被`Location`标头字段中指定的URI取代。对此资源的此请求以及所有将来的请求都应定向到新URI。|
    |302， 307|	临时重定向：应按`Location`标头字段中指定的URI逐字重复请求，但客户端应继续将原始URI用于以后的请求|  
- HTTP动词  
`API v3`尽可能在每个动作中使用适当的HTTP动词
    
    |动词|	描述|
    | :---------------| :------------ |
    |HEAD|	可以针对任何资源发出以仅获取HTTP标头信息|
    |GET|	用于检索资源|
    |POST|	用于创建资源|
    |PATCH|	用于通过部分JSON数据更新资源。例如，问题资源具有title和body属性，PATCH请求可以接受一个或多个属性以更新资源。PATCH是一个相对较新且不常见的HTTP动词，因此资源端点也接受POST请求|
    |PUT|	用于替换资源或集合，对于PUT没有body属性的请求，请确保将Content-Length标头设置为零|
    |DELETE|	用于删除资源|  
  
- 分页  
默认情况下，返回多个项目的请求将被分页为`30个项目`。可以使用`?page`参数指定其他页面。对于某些资源，还可以使用`?per_page`参数设置最大100的自定义页面大小。请注意，出于技术原因，并非所有端点都遵守该`?per_page`参数  
    ```json
    curl 'https://api.blog.com/user/repos?page=2&per_page=100'
    ```  
- 链接标题  
该链接头包括分页信息：
    ```json
    Link: <https://api.blog.com/user/luji?page=3&per_page=100>; rel="next",
    <https://api.blog.com/user/zz?page=50&per_page=100>; rel="last"
    ```
    该示例包括`换行符`，以提高可读性。
  
    此`Link`响应标头包含一个或多个`超媒体链接`关系，其中一些可能需要扩展为`URI模板`  

    可能的`rel`值为：

    |名称|	描述|
    | :----- | :----- |
    |next|	结果的下一页的链接关系|
    |last|	结果最后一页的链接关系|
    |first|	结果首页的链接关系|
    |prev|	结果前一页的链接关系|  

- 跨源资源共享  
该API支持来自任何来源的`AJAX`请求的跨来源资源共享（`CORS`）  
这是从浏览器发送的示例请求，点击 http://example.com：  
    ```json
    curl -i https://api.blog.com -H "Origin: http://example.com"
    HTTP/1.1 302 Found
    Access-Control-Allow-Origin: *
    Access-Control-Expose-Headers: ETag, Link, X-GitHub-OTP, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset, X-OAuth-Scopes, X-Accepted-OAuth-Scopes, X-Poll-Interval
    ```  
- JSON-P回调  
可以将`?callback`参数添加到任何`GET`调用，以将结果包装在`JSON函数`中。当浏览器希望通过解决跨域问题将`Blog`内容嵌入网页时，通常使用此方法。响应包括与常规API相同的数据输出，以及相关的HTTP标头信息
    ```json
    curl https://api.github.com?callback=foo
    /**/foo({
        "meta": {
            "status": 200,
            "X-RateLimit-Limit": "5000",
            "X-RateLimit-Remaining": "4966",
            "X-RateLimit-Reset": "1372700873",
            "Link": [ // pagination headers and other links
                ["https://api.github.com?page=2", {"rel": "next"}]
            ]
        },
        "data": {
            // the data
        }
    })
    ```  
    可以编写`JavaScript处理程序`来处理回调:
    ```js
    <html>
    <head>
    <script type="text/javascript">
    function foo(response) {
        var meta = response.meta;
        var data = response.data;
        console.log(meta);
        console.log(data);
    }

    var script = document.createElement('script');
    script.src = 'https://api.blog.com?callback=foo';

    document.getElementsByTagName('head')[0].appendChild(script);
    </script>
    </head>

    <body>
        <p>Open up your browser's console.</p>
    </body>
    </html>
    ```
