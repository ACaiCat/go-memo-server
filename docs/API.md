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
  "msg": "success",
  "status": 200,
  "token": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|登录成功|[handler.userLoginResp](#schemahandler.userloginresp)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|请求参数错误|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|邮箱或密码错误|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|false|none||Msg 消息描述|
|» status|integer|false|none||Status 状态码|
|» token|string|false|none||Token JWT令牌|

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
  "msg": "success",
  "status": 200,
  "token": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|令牌刷新成功|[handler.userRefreshResp](#schemahandler.userrefreshresp)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|请求参数错误或Token格式不正确|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|false|none||Msg 消息描述|
|» status|integer|false|none||Status 状态码|
|» token|string|false|none||Token JWT令牌|

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
  "msg": "success",
  "status": 200,
  "token": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|注册成功|[handler.userRegisterResp](#schemahandler.userregisterresp)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|请求参数错误|Inline|
|409|[Conflict](https://tools.ietf.org/html/rfc7231#section-6.5.8)|邮箱或用户名已被使用|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|false|none||Msg 消息描述|
|» status|integer|false|none||Status 状态码|
|» token|string|false|none||Token JWT令牌|

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
  "memo": {
    "content": "string",
    "created_at": "string",
    "end_time": "string",
    "id": 0,
    "start_time": "string",
    "status": 0,
    "title": "string",
    "user_id": 0
  },
  "msg": "success",
  "status": 200
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|创建成功|[handler.memoCreateReqResp](#schemahandler.memocreatereqresp)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» memo|[model.Memo](#schemamodel.memo)|false|none||Memo 备忘录|
|»» content|string|false|none||Content 内容|
|»» created_at|string|false|none||CreatedAt 创建时间|
|»» end_time|string|false|none||EndTime 结束时间|
|»» id|integer|false|none||ID 备忘录ID|
|»» start_time|string|false|none||StartTime 开始时间|
|»» status|[model.Status](#schemamodel.status)|false|none||Status 备忘录状态|
|»» title|string|false|none||Title 标题|
|»» user_id|integer|false|none||UserID 用户ID|
|» msg|string|false|none||Msg 消息描述|
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
  "msg": "success",
  "status": 200
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|删除成功|[handler.memoDeleteResp](#schemahandler.memodeleteresp)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|false|none||Msg 消息描述|
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
  "msg": "success",
  "status": 200
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|修改成功|[handler.memoMarkResp](#schemahandler.memomarkresp)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|备忘录状态无效|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» msg|string|false|none||Msg 消息描述|
|» status|integer|false|none||Status 状态码|

## GET 查询备忘录

GET /api/users/{user_id}/memos/query

查询满足要求的备忘录

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|path|integer| 是 |用户ID|
|page|query|integer| 是 |Page 页码|
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
  "memos": [
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
  ],
  "msg": "success",
  "pagination": {
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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|查询成功|[handler.memoQueryResp](#schemahandler.memoqueryresp)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|备忘录状态无效|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|服务器内部错误|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» memos|[[model.Memo](#schemamodel.memo)]|false|none||Memos 满足条件的备忘录|
|»» content|string|false|none||Content 内容|
|»» created_at|string|false|none||CreatedAt 创建时间|
|»» end_time|string|false|none||EndTime 结束时间|
|»» id|integer|false|none||ID 备忘录ID|
|»» start_time|string|false|none||StartTime 开始时间|
|»» status|[model.Status](#schemamodel.status)|false|none||Status 备忘录状态|
|»» title|string|false|none||Title 标题|
|»» user_id|integer|false|none||UserID 用户ID|
|» msg|string|false|none||Msg 消息描述|
|» pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
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

# 数据模型

<h2 id="tocS_handler.Pagination">handler.Pagination</h2>

<a id="schemahandler.pagination"></a>
<a id="schema_handler.Pagination"></a>
<a id="tocShandler.pagination"></a>
<a id="tocshandler.pagination"></a>

```json
{
  "page": 1,
  "per_page": 20,
  "total": 100
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
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
  },
  "msg": "success",
  "status": 200
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|memo|[model.Memo](#schemamodel.memo)|false|none||Memo 备忘录|
|msg|string|false|none||Msg 消息描述|
|status|integer|false|none||Status 状态码|

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
{
  "msg": "success",
  "status": 200
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|string|false|none||Msg 消息描述|
|status|integer|false|none||Status 状态码|

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
{
  "msg": "success",
  "status": 200
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|string|false|none||Msg 消息描述|
|status|integer|false|none||Status 状态码|

<h2 id="tocS_handler.memoQueryResp">handler.memoQueryResp</h2>

<a id="schemahandler.memoqueryresp"></a>
<a id="schema_handler.memoQueryResp"></a>
<a id="tocShandler.memoqueryresp"></a>
<a id="tocshandler.memoqueryresp"></a>

```json
{
  "memos": [
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
  ],
  "msg": "success",
  "pagination": {
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
|memos|[[model.Memo](#schemamodel.memo)]|false|none||Memos 满足条件的备忘录|
|msg|string|false|none||Msg 消息描述|
|pagination|[handler.Pagination](#schemahandler.pagination)|false|none||Pagination 分页信息|
|status|integer|false|none||Status 状态码|

<h2 id="tocS_handler.userLoginResp">handler.userLoginResp</h2>

<a id="schemahandler.userloginresp"></a>
<a id="schema_handler.userLoginResp"></a>
<a id="tocShandler.userloginresp"></a>
<a id="tocshandler.userloginresp"></a>

```json
{
  "msg": "success",
  "status": 200,
  "token": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|string|false|none||Msg 消息描述|
|status|integer|false|none||Status 状态码|
|token|string|false|none||Token JWT令牌|

<h2 id="tocS_handler.userRefreshResp">handler.userRefreshResp</h2>

<a id="schemahandler.userrefreshresp"></a>
<a id="schema_handler.userRefreshResp"></a>
<a id="tocShandler.userrefreshresp"></a>
<a id="tocshandler.userrefreshresp"></a>

```json
{
  "msg": "success",
  "status": 200,
  "token": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|string|false|none||Msg 消息描述|
|status|integer|false|none||Status 状态码|
|token|string|false|none||Token JWT令牌|

<h2 id="tocS_handler.userRegisterResp">handler.userRegisterResp</h2>

<a id="schemahandler.userregisterresp"></a>
<a id="schema_handler.userRegisterResp"></a>
<a id="tocShandler.userregisterresp"></a>
<a id="tocshandler.userregisterresp"></a>

```json
{
  "msg": "success",
  "status": 200,
  "token": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|msg|string|false|none||Msg 消息描述|
|status|integer|false|none||Status 状态码|
|token|string|false|none||Token JWT令牌|

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

