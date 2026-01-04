# 备忘录接口文档

## 接口概览（总计7个）

### 备忘录

| **路径**                                         | **功能**   | **请求方式** | **是否需要鉴权** |
|--------------------------------------------------|------------|--------------|------------------|
| [/api/users/{user_id}/memos/create](#创建备忘录) | 创建备忘录 | POST         | true             |
| [/api/users/{user_id}/memos/delete](#删除备忘录) | 删除备忘录 | POST         | true             |
| [/api/users/{user_id}/memos/mark](#标记备忘录)   | 标记备忘录 | POST         | true             |
| [/api/users/{user_id}/memos/query](#查询备忘录)  | 查询备忘录 | GET          | true             |

### 认证

| **路径**                               | **功能**         | **请求方式** | **是否需要鉴权** |
|----------------------------------------|------------------|--------------|------------------|
| [/api/auth/login](#用户登录)           | 用户登录         | POST         | false            |
| [/api/auth/refresh](#刷新用户登录令牌) | 刷新用户登录令牌 | POST         | false            |
| [/api/auth/register](#用户注册)        | 用户注册         | POST         | false            |

## 接口详情

### 备忘录

### 创建备忘录

[返回概览](#备忘录)

---

POST /api/users/{user_id}/memos/create  
Content-Type: application/json

请求参数：

| **来源** | **参数**        | **描述**                  | **类型** | **约束** | **说明** |
|----------|-----------------|---------------------------|----------|----------|----------|
| header   | content         | Content 内容              | string   | 必填     |          |
| header   | end_timestamp   | EndTimestamp 结束时间戳   | integer  | 必填     |          |
| header   | start_timestamp | StartTimestamp 开始时间戳 | integer  | 必填     |          |
| header   | title           | Title 标题                | string   | 必填     |          |
| path     | user_id         | 用户ID                    | integer  | 必填     |          |

请求示例：

```
Header:
content: content
end_timestamp: 1
start_timestamp: 1
title: title
```

---

Content-Type: application/json

响应参数：

| **参数**   | **描述**            | **类型** | **说明** |
|------------|---------------------|----------|----------|
| data       | Data 数据           | object   |          |
| msg        | Msg 消息描述        | string   |          |
| pagination | Pagination 分页信息 | object   |          |
| status     | Status 状态码       | integer  |          |

响应示例：

```json
{
  "data": null,
  "msg": "success",
  "pagination": null,
  "status": 200
}
```

### 删除备忘录

[返回概览](#备忘录)

---

POST /api/users/{user_id}/memos/delete  
Content-Type: application/json

请求参数：

| **来源** | **参数** | **描述**         | **类型**      | **约束** | **说明** |
|----------|----------|------------------|---------------|----------|----------|
| body     | memo_ids | MemoIDs 备忘录ID | integer array | 必填     |          |
| path     | user_id  | 用户ID           | integer       | 必填     |          |

请求示例：

```json
{
  "memo_ids": [
    1
  ]
}
```

---

Content-Type: application/json

响应参数：

| **参数**   | **描述**            | **类型** | **说明** |
|------------|---------------------|----------|----------|
| data       | Data 数据           | object   |          |
| msg        | Msg 消息描述        | string   |          |
| pagination | Pagination 分页信息 | object   |          |
| status     | Status 状态码       | integer  |          |

响应示例：

```json
{
  "data": null,
  "msg": "success",
  "pagination": null,
  "status": 200
}
```

### 标记备忘录

[返回概览](#备忘录)

---

POST /api/users/{user_id}/memos/mark  
Content-Type: application/json

请求参数：

| **来源** | **参数** | **描述**         | **类型**      | **约束** | **说明** |
|----------|----------|------------------|---------------|----------|----------|
| body     | memo_ids | MemoIDs 备忘录ID | integer array | 必填     |          |
| body     | status   | Status 状态      | object        | 必填     |          |
| path     | user_id  | 用户ID           | integer       | 必填     |          |

请求示例：

```json
{
  "memo_ids": [
    1
  ],
  "status": null
}
```

---

Content-Type: application/json

响应参数：

| **参数**   | **描述**            | **类型** | **说明** |
|------------|---------------------|----------|----------|
| data       | Data 数据           | object   |          |
| msg        | Msg 消息描述        | string   |          |
| pagination | Pagination 分页信息 | object   |          |
| status     | Status 状态码       | integer  |          |

响应示例：

```json
{
  "data": null,
  "msg": "success",
  "pagination": null,
  "status": 200
}
```

### 查询备忘录

[返回概览](#备忘录)

---

GET /api/users/{user_id}/memos/query

请求参数：

| **来源** | **参数** | **描述**           | **类型** | **约束** | **说明** |
|----------|----------|--------------------|----------|----------|----------|
| query    | page     | Page 页码          | integer  | 非必填   |          |
| query    | per_page | PerPage 每页条目数 | integer  | 非必填   |          |
| query    | query    | Query 关键词       | string   | 非必填   |          |
| query    | status   | Status 状态        | integer  | 必填     |          |
| path     | user_id  | 用户ID             | integer  | 必填     |          |

请求示例：

```
Query:
/api/users/{user_id}/memos/query?page=1&per_page=1&query=query&status=1
```

---

Content-Type: application/json

响应参数：

| **参数**   | **描述**            | **类型** | **说明** |
|------------|---------------------|----------|----------|
| data       | Data 数据           | object   |          |
| msg        | Msg 消息描述        | string   |          |
| pagination | Pagination 分页信息 | object   |          |
| status     | Status 状态码       | integer  |          |

响应示例：

```json
{
  "data": null,
  "msg": "success",
  "pagination": null,
  "status": 200
}
```

### 认证

### 用户登录

[返回概览](#认证)

---

POST /api/auth/login  
Content-Type: application/x-www-form-urlencoded

请求参数：

| **来源** | **参数** | **描述**          | **类型** | **约束** | **说明** |
|----------|----------|-------------------|----------|----------|----------|
| formData | email    | Email 用户邮箱    | string   | 非必填   |          |
| formData | password | Password 用户密码 | string   | 非必填   |          |

请求示例：

```
Form Data:
email: email
password: password
```

---

Content-Type: application/json

响应参数：

| **参数**   | **描述**            | **类型** | **说明** |
|------------|---------------------|----------|----------|
| data       | Data 数据           | object   |          |
| msg        | Msg 消息描述        | string   |          |
| pagination | Pagination 分页信息 | object   |          |
| status     | Status 状态码       | integer  |          |

响应示例：

```json
{
  "data": null,
  "msg": "success",
  "pagination": null,
  "status": 200
}
```

### 刷新用户登录令牌

[返回概览](#认证)

---

POST /api/auth/refresh  
Content-Type: application/json

请求参数：

| **来源** | **参数**      | **描述**                 | **类型** | **约束** | **说明** |
|----------|---------------|--------------------------|----------|----------|----------|
| header   | Authorization | Authorization 认证字符串 | string   | 必填     |          |

请求示例：

```
Header:
Authorization: Authorization
```

---

Content-Type: application/json

响应参数：

| **参数**   | **描述**            | **类型** | **说明** |
|------------|---------------------|----------|----------|
| data       | Data 数据           | object   |          |
| msg        | Msg 消息描述        | string   |          |
| pagination | Pagination 分页信息 | object   |          |
| status     | Status 状态码       | integer  |          |

响应示例：

```json
{
  "data": null,
  "msg": "success",
  "pagination": null,
  "status": 200
}
```

### 用户注册

[返回概览](#认证)

---

POST /api/auth/register  
Content-Type: application/x-www-form-urlencoded

请求参数：

| **来源** | **参数** | **描述**          | **类型** | **约束** | **说明** |
|----------|----------|-------------------|----------|----------|----------|
| formData | email    | Email 用户邮箱    | string   | 必填     |          |
| formData | name     | Name 用户名       | string   | 必填     |          |
| formData | password | Password 用户密码 | string   | 必填     |          |

请求示例：

```
Form Data:
email: email
name: name
password: password
```

---

Content-Type: application/json

响应参数：

| **参数**   | **描述**            | **类型** | **说明** |
|------------|---------------------|----------|----------|
| data       | Data 数据           | object   |          |
| msg        | Msg 消息描述        | string   |          |
| pagination | Pagination 分页信息 | object   |          |
| status     | Status 状态码       | integer  |          |

响应示例：

```json
{
  "data": null,
  "msg": "success",
  "pagination": null,
  "status": 200
}
```
