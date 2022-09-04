# message-board

遵循`RESTful`规范

- [x] 密码加盐
- [x] 预防SQL注入
- [x] JWT鉴权
- [ ] 分页查询

## 接口设计

### 用户相关 `/user`

#### POST `/user/login` 登录

+ 登录账号
+ 仅支持邮箱登录

##### 请求参数

| 名称            | 位置     | 类型     | 必选    | 说明 |
|---------------|--------|--------|-------|--|
| body          | body   | object | true | none |
| » email       | body   | string | true | 邮箱 |
| » password    | body   | string | true | 密码 |

#### POST `/user/register` 注册

+ 注册账号
+ 仅支持邮箱注册

##### 请求参数

| 名称            | 位置     | 类型     | 必选    | 说明 |
|---------------|--------|--------|-------|--|
| body          | body   | object | true | none |
| » name        | body   | string | true | 用户名 |
| » email       | body   | string | true | 邮箱 |
| » verify      | body   | string | true | 邮箱验证码 |
| » password    | body   | string | true | 密码 |

#### POST `/user/changepass` 修改密码 & 忘记密码

+ 更改密码

##### 请求参数

| 名称            | 位置     | 类型     | 必选    | 说明 |
|---------------|--------|--------|-------|--|
| body          | body   | object | true | none |
| » email       | body   | string | true | 邮箱 |
| » password       | body   | string | true | 密码 |
| » verify       | body   | string | true | 验证码 |

#### POST `/user/verify` 发送验证码

+ 请勿频繁请求

##### 请求参数

| 名称            | 位置     | 类型     | 必选    | 说明 |
|---------------|--------|--------|-------|--|
| body          | body   | object | false | none |
| » email       | body   | string | false | 邮箱 |

#### POST `/user/detail` 更改用户信息

+ 更改用户信息
+ 只用填需要修改的部分即可
+ 需要鉴权

##### 请求参数

| 名称            | 位置     | 类型     | 必选    | 说明 |
|---------------|--------|--------|-------|--|
| body          | body   | object | true | none |
| » avatar       | body   | string | false | 头像 |
| » bio       | body   | string | false | 描述 |

#### GET `/user/detail` 获取用户信息

+ 获取用户信息
+ email name uid任选其一即可

##### 请求参数

| 名称            | 位置     | 类型     | 必选    | 说明 |
|---------------|--------|--------|-------|--|
| name       | query   | string | false | 名称 |
| email      | query   | string | false | 邮箱 |
| uid      | query   | number | false | id |

### 留言版功能 `/message`

#### PUT `/message` 留言

+ 创建一条parent为null的留言记录
+ 需要鉴权

##### 请求参数

| 名称        | 位置   | 类型     | 必选   | 说明   |
|-----------|------|--------|------|------|
| body      | body | object | true | none |
| » message | body | string | true | 信息   |

#### PUT `/message/{id}` 回复指定留言

+ 创建一条parent为id的留言记录
+ 需要鉴权

##### 请求参数

| 名称        | 位置   | 类型     | 必选   | 说明      |
|-----------|------|--------|------|---------|
| id        | path | string | true | 回复信息的id |
| body      | body | string | true | none    |
| » message | body | string | true | 信息      |

#### GET `/message` 获取留言详情

+ 获取树顶层的所有节点

##### 请求参数

无

#### GET `/message/{id}` 获取指定条留言的回复

+ 获取id对应节点的所有下一级节点

##### 请求参数

| 名称  | 位置   | 类型     | 必选   | 说明      |
|-----|------|--------|------|---------|
| id  | path | string | true | 回复信息的id |

