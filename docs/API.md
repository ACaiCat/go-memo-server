---
title: Memo
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.30"

---

# Memo

A demo of memo api

Base URLs:

Web: <a href="https://github.com/ACaiCat">ACaiCat</a> 
License: <a href="https://mit-license.org/">The MIT License (MIT)</a>

# Authentication

* API Key (ApiKeyAuth)
    - Parameter Name: **Authorization**, in: header. Description for what is this security definition being used

# 认证

## POST 用户登录

POST /api/auth/login

处理用户登录并返回JWT令牌

> Body 请求参数

```yaml
email: ""
password: ""

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» email|body|string| 否 |Email 用户邮箱|
|» password|body|string| 否 |Password 用户密码|

> 返回示例

> 200 Response

```json
{
  "data": {
    "token": "string",
    "user_id": 0
  },
  "msg": "success",
  "pagination": {
    "max_page": 5,
    "page": 1,
    "per_page": 20,
    "total": 100
  },
  "status": 200
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|登录成功|[handler.BaseResp-handler_userLoginResp](#schemahandler.baseresp-handler_userloginresp)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|请求参数错误|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|邮箱或密码错误|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[handler.userLoginResp](#schemahandler.userloginresp)|false|none||Data 数据|
|»» token|string|false|none||Token JWT令牌|
|»» user_id|integer|false|none||UserID 用户ID|
|» msg|string|false|none||Msg 消息描述|
|» pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
|»» max_page|integer|false|none||MaxPage 最大页数|
|»» page|integer|false|none||Page 当前页码|
|»» per_page|integer|false|none||PerPage 每页数量|
|»» total|integer|false|none||Total 总记录数|
|» status|integer|false|none||Status 状态码|

## POST 刷新用户登录令牌

POST /api/auth/refresh

刷新用户的JWT令牌

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |Authorization 认证字符串|

> 返回示例

> 200 Response

```json
{
  "data": {
    "token": "string"
  },
  "msg": "success",
  "pagination": {
    "max_page": 5,
    "page": 1,
    "per_page": 20,
    "total": 100
  },
  "status": 200
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|令牌刷新成功|[handler.BaseResp-handler_userRefreshResp](#schemahandler.baseresp-handler_userrefreshresp)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|请求参数错误或Token格式不正确|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[handler.userRefreshResp](#schemahandler.userrefreshresp)|false|none||Data 数据|
|»» token|string|false|none||Token JWT令牌|
|» msg|string|false|none||Msg 消息描述|
|» pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
|»» max_page|integer|false|none||MaxPage 最大页数|
|»» page|integer|false|none||Page 当前页码|
|»» per_page|integer|false|none||PerPage 每页数量|
|»» total|integer|false|none||Total 总记录数|
|» status|integer|false|none||Status 状态码|

## POST 用户注册

POST /api/auth/register

处理用户注册并返回JWT令牌

> Body 请求参数

```yaml
email: ""
name: ""
password: ""

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» email|body|string| 是 |Email 用户邮箱|
|» name|body|string| 是 |Name 用户名|
|» password|body|string| 是 |Password 用户密码|

> 返回示例

> 200 Response

```json
{
  "data": {
    "token": "string",
    "user_id": 0
  },
  "msg": "success",
  "pagination": {
    "max_page": 5,
    "page": 1,
    "per_page": 20,
    "total": 100
  },
  "status": 200
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|注册成功|[handler.BaseResp-handler_userRegisterResp](#schemahandler.baseresp-handler_userregisterresp)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|请求参数错误|Inline|
|409|[Conflict](https://tools.ietf.org/html/rfc7231#section-6.5.8)|邮箱或用户名已被使用|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[handler.userRegisterResp](#schemahandler.userregisterresp)|false|none||Data 数据|
|»» token|string|false|none||Token JWT令牌|
|»» user_id|integer|false|none||UserID 用户ID|
|» msg|string|false|none||Msg 消息描述|
|» pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
|»» max_page|integer|false|none||MaxPage 最大页数|
|»» page|integer|false|none||Page 当前页码|
|»» per_page|integer|false|none||PerPage 每页数量|
|»» total|integer|false|none||Total 总记录数|
|» status|integer|false|none||Status 状态码|

# 备忘录

## POST 创建备忘录

POST /api/users/{user_id}/memos/create

创建一条新的备忘录

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|path|integer| 是 |用户ID|
|content|header|string| 是 |Content 内容|
|end_timestamp|header|integer| 是 |EndTimestamp 结束时间戳|
|start_timestamp|header|integer| 是 |StartTimestamp 开始时间戳|
|title|header|string| 是 |Title 标题|

> 返回示例

> 200 Response

```json
{
  "data": {
    "memo": {
      "content": null,
      "created_at": null,
      "end_time": null,
      "id": null,
      "start_time": null,
      "status": null,
      "title": null,
      "user_id": null
    }
  },
  "msg": "success",
  "pagination": {
    "max_page": 5,
    "page": 1,
    "per_page": 20,
    "total": 100
  },
  "status": 200
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|创建成功|[handler.BaseResp-handler_memoCreateReqResp](#schemahandler.baseresp-handler_memocreatereqresp)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[handler.memoCreateReqResp](#schemahandler.memocreatereqresp)|false|none||Data 数据|
|»» memo|[model.Memo](#schemamodel.memo)|false|none||Memo 备忘录|
|»»» content|string|false|none||Content 内容|
|»»» created_at|string|false|none||CreatedAt 创建时间|
|»»» end_time|string|false|none||EndTime 结束时间|
|»»» id|integer|false|none||ID 备忘录ID|
|»»» start_time|string|false|none||StartTime 开始时间|
|»»» status|[model.Status](#schemamodel.status)|false|none||Status 备忘录状态|
|»»» title|string|false|none||Title 标题|
|»»» user_id|integer|false|none||UserID 用户ID|
|» msg|string|false|none||Msg 消息描述|
|» pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
|»» max_page|integer|false|none||MaxPage 最大页数|
|»» page|integer|false|none||Page 当前页码|
|»» per_page|integer|false|none||PerPage 每页数量|
|»» total|integer|false|none||Total 总记录数|
|» status|integer|false|none||Status 状态码|

#### 枚举值

|属性|值|
|---|---|
|status|0|
|status|1|
|status|2|

## POST 删除备忘录

POST /api/users/{user_id}/memos/delete

删除备忘录

> Body 请求参数

```json
{
  "memo_ids": [
    0
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|path|integer| 是 |用户ID|
|body|body|[handler.memoDeleteReq](#schemahandler.memodeletereq)| 是 |none|

> 返回示例

> 200 Response

```json
{
  "data": {},
  "msg": "success",
  "pagination": {
    "max_page": 5,
    "page": 1,
    "per_page": 20,
    "total": 100
  },
  "status": 200
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|删除成功|[handler.BaseResp-handler_memoDeleteResp](#schemahandler.baseresp-handler_memodeleteresp)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[handler.memoDeleteResp](#schemahandler.memodeleteresp)|false|none||Data 数据|
|» msg|string|false|none||Msg 消息描述|
|» pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
|»» max_page|integer|false|none||MaxPage 最大页数|
|»» page|integer|false|none||Page 当前页码|
|»» per_page|integer|false|none||PerPage 每页数量|
|»» total|integer|false|none||Total 总记录数|
|» status|integer|false|none||Status 状态码|

## POST 标记备忘录

POST /api/users/{user_id}/memos/mark

标记备忘录的状态

> Body 请求参数

```json
{
  "memo_ids": [
    0
  ],
  "status": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|path|integer| 是 |用户ID|
|body|body|[handler.memoMarkReq](#schemahandler.memomarkreq)| 是 |none|

> 返回示例

> 200 Response

```json
{
  "data": {},
  "msg": "success",
  "pagination": {
    "max_page": 5,
    "page": 1,
    "per_page": 20,
    "total": 100
  },
  "status": 200
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|修改成功|[handler.BaseResp-handler_memoMarkResp](#schemahandler.baseresp-handler_memomarkresp)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|备忘录状态无效|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[handler.memoMarkResp](#schemahandler.memomarkresp)|false|none||Data 数据|
|» msg|string|false|none||Msg 消息描述|
|» pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
|»» max_page|integer|false|none||MaxPage 最大页数|
|»» page|integer|false|none||Page 当前页码|
|»» per_page|integer|false|none||PerPage 每页数量|
|»» total|integer|false|none||Total 总记录数|
|» status|integer|false|none||Status 状态码|

## GET 查询备忘录

GET /api/users/{user_id}/memos/query

查询满足要求的备忘录

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|path|integer| 是 |用户ID|
|page|query|integer| 否 |Page 页码|
|per_page|query|integer| 否 |PerPage 每页条目数|
|query|query|string| 否 |Query 关键词|
|status|query|integer| 是 |Status 状态|

#### 枚举值

|属性|值|
|---|---|
|status|0|
|status|1|
|status|2|

> 返回示例

> 200 Response

```json
{
  "data": {},
  "msg": "success",
  "pagination": {
    "max_page": 5,
    "page": 1,
    "per_page": 20,
    "total": 100
  },
  "status": 200
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|查询成功|[handler.BaseResp-handler_memoMarkResp](#schemahandler.baseresp-handler_memomarkresp)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|备忘录状态无效|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[handler.memoMarkResp](#schemahandler.memomarkresp)|false|none||Data 数据|
|» msg|string|false|none||Msg 消息描述|
|» pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
|»» max_page|integer|false|none||MaxPage 最大页数|
|»» page|integer|false|none||Page 当前页码|
|»» per_page|integer|false|none||PerPage 每页数量|
|»» total|integer|false|none||Total 总记录数|
|» status|integer|false|none||Status 状态码|

# 数据模型

<h2 id="tocS_handler.BaseResp-handler_memoCreateReqResp">handler.BaseResp-handler_memoCreateReqResp</h2>

<a id="schemahandler.baseresp-handler_memocreatereqresp"></a>
<a id="schema_handler.BaseResp-handler_memoCreateReqResp"></a>
<a id="tocShandler.baseresp-handler_memocreatereqresp"></a>
<a id="tocshandler.baseresp-handler_memocreatereqresp"></a>

```json
{
  "data": {
    "memo": {
      "content": null,
      "created_at": null,
      "end_time": null,
      "id": null,
      "start_time": null,
      "status": null,
      "title": null,
      "user_id": null
    }
  },
  "msg": "success",
  "pagination": {
    "max_page": 5,
    "page": 1,
    "per_page": 20,
    "total": 100
  },
  "status": 200
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|data|[handler.memoCreateReqResp](#schemahandler.memocreatereqresp)|false|none||Data 数据|
|msg|string|false|none||Msg 消息描述|
|pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
|status|integer|false|none||Status 状态码|

<h2 id="tocS_handler.BaseResp-handler_memoDeleteResp">handler.BaseResp-handler_memoDeleteResp</h2>

<a id="schemahandler.baseresp-handler_memodeleteresp"></a>
<a id="schema_handler.BaseResp-handler_memoDeleteResp"></a>
<a id="tocShandler.baseresp-handler_memodeleteresp"></a>
<a id="tocshandler.baseresp-handler_memodeleteresp"></a>

```json
{
  "data": {},
  "msg": "success",
  "pagination": {
    "max_page": 5,
    "page": 1,
    "per_page": 20,
    "total": 100
  },
  "status": 200
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|data|[handler.memoDeleteResp](#schemahandler.memodeleteresp)|false|none||Data 数据|
|msg|string|false|none||Msg 消息描述|
|pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
|status|integer|false|none||Status 状态码|

<h2 id="tocS_handler.BaseResp-handler_memoMarkResp">handler.BaseResp-handler_memoMarkResp</h2>

<a id="schemahandler.baseresp-handler_memomarkresp"></a>
<a id="schema_handler.BaseResp-handler_memoMarkResp"></a>
<a id="tocShandler.baseresp-handler_memomarkresp"></a>
<a id="tocshandler.baseresp-handler_memomarkresp"></a>

```json
{
  "data": {},
  "msg": "success",
  "pagination": {
    "max_page": 5,
    "page": 1,
    "per_page": 20,
    "total": 100
  },
  "status": 200
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|data|[handler.memoMarkResp](#schemahandler.memomarkresp)|false|none||Data 数据|
|msg|string|false|none||Msg 消息描述|
|pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
|status|integer|false|none||Status 状态码|

<h2 id="tocS_handler.BaseResp-handler_userLoginResp">handler.BaseResp-handler_userLoginResp</h2>

<a id="schemahandler.baseresp-handler_userloginresp"></a>
<a id="schema_handler.BaseResp-handler_userLoginResp"></a>
<a id="tocShandler.baseresp-handler_userloginresp"></a>
<a id="tocshandler.baseresp-handler_userloginresp"></a>

```json
{
  "data": {
    "token": "string",
    "user_id": 0
  },
  "msg": "success",
  "pagination": {
    "max_page": 5,
    "page": 1,
    "per_page": 20,
    "total": 100
  },
  "status": 200
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|data|[handler.userLoginResp](#schemahandler.userloginresp)|false|none||Data 数据|
|msg|string|false|none||Msg 消息描述|
|pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
|status|integer|false|none||Status 状态码|

<h2 id="tocS_handler.BaseResp-handler_userRefreshResp">handler.BaseResp-handler_userRefreshResp</h2>

<a id="schemahandler.baseresp-handler_userrefreshresp"></a>
<a id="schema_handler.BaseResp-handler_userRefreshResp"></a>
<a id="tocShandler.baseresp-handler_userrefreshresp"></a>
<a id="tocshandler.baseresp-handler_userrefreshresp"></a>

```json
{
  "data": {
    "token": "string"
  },
  "msg": "success",
  "pagination": {
    "max_page": 5,
    "page": 1,
    "per_page": 20,
    "total": 100
  },
  "status": 200
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|data|[handler.userRefreshResp](#schemahandler.userrefreshresp)|false|none||Data 数据|
|msg|string|false|none||Msg 消息描述|
|pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
|status|integer|false|none||Status 状态码|

<h2 id="tocS_handler.BaseResp-handler_userRegisterResp">handler.BaseResp-handler_userRegisterResp</h2>

<a id="schemahandler.baseresp-handler_userregisterresp"></a>
<a id="schema_handler.BaseResp-handler_userRegisterResp"></a>
<a id="tocShandler.baseresp-handler_userregisterresp"></a>
<a id="tocshandler.baseresp-handler_userregisterresp"></a>

```json
{
  "data": {
    "token": "string",
    "user_id": 0
  },
  "msg": "success",
  "pagination": {
    "max_page": 5,
    "page": 1,
    "per_page": 20,
    "total": 100
  },
  "status": 200
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|data|[handler.userRegisterResp](#schemahandler.userregisterresp)|false|none||Data 数据|
|msg|string|false|none||Msg 消息描述|
|pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
|status|integer|false|none||Status 状态码|

<h2 id="tocS_handler.Pagination">handler.Pagination</h2>

<a id="schemahandler.pagination"></a>
<a id="schema_handler.Pagination"></a>
<a id="tocShandler.pagination"></a>
<a id="tocshandler.pagination"></a>

```json
{
  "max_page": 5,
  "page": 1,
  "per_page": 20,
  "total": 100
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|max_page|integer|false|none||MaxPage 最大页数|
|page|integer|false|none||Page 当前页码|
|per_page|integer|false|none||PerPage 每页数量|
|total|integer|false|none||Total 总记录数|

<h2 id="tocS_handler.memoCreateReqResp">handler.memoCreateReqResp</h2>

<a id="schemahandler.memocreatereqresp"></a>
<a id="schema_handler.memoCreateReqResp"></a>
<a id="tocShandler.memocreatereqresp"></a>
<a id="tocshandler.memocreatereqresp"></a>

```json
{
  "memo": {
    "content": "string",
    "created_at": "string",
    "end_time": "string",
    "id": 0,
    "start_time": "string",
    "status": 0,
    "title": "string",
    "user_id": 0
  }
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|memo|[model.Memo](#schemamodel.memo)|false|none||Memo 备忘录|

<h2 id="tocS_handler.memoDeleteReq">handler.memoDeleteReq</h2>

<a id="schemahandler.memodeletereq"></a>
<a id="schema_handler.memoDeleteReq"></a>
<a id="tocShandler.memodeletereq"></a>
<a id="tocshandler.memodeletereq"></a>

```json
{
  "memo_ids": [
    0
  ]
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|memo_ids|[integer]|true|none||MemoIDs 备忘录ID|

<h2 id="tocS_handler.memoDeleteResp">handler.memoDeleteResp</h2>

<a id="schemahandler.memodeleteresp"></a>
<a id="schema_handler.memoDeleteResp"></a>
<a id="tocShandler.memodeleteresp"></a>
<a id="tocshandler.memodeleteresp"></a>

```json
{}

```

### 属性

*None*

<h2 id="tocS_handler.memoMarkReq">handler.memoMarkReq</h2>

<a id="schemahandler.memomarkreq"></a>
<a id="schema_handler.memoMarkReq"></a>
<a id="tocShandler.memomarkreq"></a>
<a id="tocshandler.memomarkreq"></a>

```json
{
  "memo_ids": [
    0
  ],
  "status": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|memo_ids|[integer]|true|none||MemoIDs 备忘录ID|
|status|[model.Status](#schemamodel.status)|true|none||Status 状态|

<h2 id="tocS_handler.memoMarkResp">handler.memoMarkResp</h2>

<a id="schemahandler.memomarkresp"></a>
<a id="schema_handler.memoMarkResp"></a>
<a id="tocShandler.memomarkresp"></a>
<a id="tocshandler.memomarkresp"></a>

```json
{}

```

### 属性

*None*

<h2 id="tocS_handler.userLoginResp">handler.userLoginResp</h2>

<a id="schemahandler.userloginresp"></a>
<a id="schema_handler.userLoginResp"></a>
<a id="tocShandler.userloginresp"></a>
<a id="tocshandler.userloginresp"></a>

```json
{
  "token": "string",
  "user_id": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|token|string|false|none||Token JWT令牌|
|user_id|integer|false|none||UserID 用户ID|

<h2 id="tocS_handler.userRefreshResp">handler.userRefreshResp</h2>

<a id="schemahandler.userrefreshresp"></a>
<a id="schema_handler.userRefreshResp"></a>
<a id="tocShandler.userrefreshresp"></a>
<a id="tocshandler.userrefreshresp"></a>

```json
{
  "token": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|token|string|false|none||Token JWT令牌|

<h2 id="tocS_handler.userRegisterResp">handler.userRegisterResp</h2>

<a id="schemahandler.userregisterresp"></a>
<a id="schema_handler.userRegisterResp"></a>
<a id="tocShandler.userregisterresp"></a>
<a id="tocshandler.userregisterresp"></a>

```json
{
  "token": "string",
  "user_id": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|token|string|false|none||Token JWT令牌|
|user_id|integer|false|none||UserID 用户ID|

<h2 id="tocS_model.Memo">model.Memo</h2>

<a id="schemamodel.memo"></a>
<a id="schema_model.Memo"></a>
<a id="tocSmodel.memo"></a>
<a id="tocsmodel.memo"></a>

```json
{
  "content": "string",
  "created_at": "string",
  "end_time": "string",
  "id": 0,
  "start_time": "string",
  "status": 0,
  "title": "string",
  "user_id": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|content|string|false|none||Content 内容|
|created_at|string|false|none||CreatedAt 创建时间|
|end_time|string|false|none||EndTime 结束时间|
|id|integer|false|none||ID 备忘录ID|
|start_time|string|false|none||StartTime 开始时间|
|status|[model.Status](#schemamodel.status)|false|none||Status 备忘录状态|
|title|string|false|none||Title 标题|
|user_id|integer|false|none||UserID 用户ID|

<h2 id="tocS_model.Status">model.Status</h2>

<a id="schemamodel.status"></a>
<a id="schema_model.Status"></a>
<a id="tocSmodel.status"></a>
<a id="tocsmodel.status"></a>

```json
0

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|*anonymous*|integer|false|none||none|

#### 枚举值

|属性|值|
|---|---|
|*anonymous*|0|
|*anonymous*|1|
|*anonymous*|2|

