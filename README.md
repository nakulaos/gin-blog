# 一个博客相关的项目后端

## 使用技术：完全是Go语言，主要使用Gin框架和Mysql数据库。

### 中间用到的其它技术

#### 1、Gorm对象关系映射框架

#### 2、logrus日志等级打印

#### 3、JWT鉴权

#### 4、Redis做用户退出时的JWT失效校验，即对过期时间做缓存

#### 5、读取yaml配置

#### 6、本地、七牛云对象存储（COS）图片

#### 7、发送验证码绑定邮箱

#### 8、使用Session存入服务端，校验验证码合法性

#### 9、Swagger文档导出

#### 10、Air项目热加载

#### 11、flag命令行执行参数定义

## 定义的Gorm模型

![image-20240423100331497](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20240423100331497.png)

## 对应的Mysql数据库表结构

![image-20240423100605452](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20240423100605452.png)



## 已实现接口

![image-20240423101004251](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20240423101004251.png)