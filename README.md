# 一个博客相关的项目后端

## 使用技术：完全是Go语言，主要使用Gin框架和Mysql数据库。

## 中间件用到了其它技术，并且有QQ官方登录接口。

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

#### ⭐12、QQ官方登录



## 定义的Gorm模型

![image-20240423100331497](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20240423100331497.png)

## 对应的Mysql数据库表结构

![image-20240423100605452](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20240423100605452.png)



## 已实现接口

![image-20240423101004251](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20240423101004251.png)

### 对于QQ登录：

#### 1、使用很久之前自己用Notion搭的Blog，关联的域名到*腾讯应用开放平台*申请QQ互联，跟这个blog项目不是一个！因为要申请，所以拿之前Notion来申请，申请好了这边这个项目测试看看。要有appid，key，各种之类的。下面是我成功的整个流程

![QQ截图20240424201646](C:\Users\Lenovo\Desktop\QQ截图20240424201646.png)

#### 2、开启SwitchHosts

![QQ截图20240424201958](C:\Users\Lenovo\Desktop\QQ截图20240424201958.png)

#### 3、进入相应回调地址，得到熟悉的画面，回调地址也没必要藏着掖着：https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=102107572&redirect_uri=https://www.blog.dzcs.online/login/callback/qq

![image-20240424202148160](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20240424202148160.png)

#### 4、编写后端代码。

![image-20240424202355636](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20240424202355636.png)

#### 5、第3步登陆成功后，会有一个code，使用apifox请求携带此code测试，效果如下：

##### 其中第一大段JSON是qq用户的信息，包括qq昵称，qq上你设置的性别，qq头像地址，open_id，其中open_id是每个qq用户唯一的。

##### 第二大段是我给的登录成功后的一个jwt的token字符串，表明测试成功了。

![image-20240424202458095](C:\Users\Lenovo\AppData\Roaming\Typora\typora-user-images\image-20240424202458095.png)

#### 6、看一眼数据库，最后一个用户即刚刚用qq注册登录的。使用唯一open_id做主键即可。

![QQ截图20240424202752](C:\Users\Lenovo\Desktop\QQ截图20240424202752.png)