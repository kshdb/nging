# Nging V3

![Nging's logo](https://github.com/admpub/nging/blob/master/public/assets/backend/images/nging-gear.png?raw=true)

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/admpub/nging)

> 注意：这是Nging V3源代码，旧版V2.x、V1.x已经转移到 [v2分支](https://github.com/admpub/nging/tree/v2) [v1分支](https://github.com/admpub/nging/tree/v1)

    注意：目前只支持安装到MySQL的Nging无缝升级，暂不支持SQLite安装方式的升级（推荐安装到MySQL，版本需不低于MySQL v5.7.18+）。

    升级步骤：
    0. 备份数据库和旧版可执行文件；
    1. 停止旧版本程序的运行；
    2. 将新版本所有文件复制到旧版文件目录里进行覆盖；
    3. 启动新版本程序；
    4. 登录后台检查各项功能是否正常；
    5. 升级完毕


Nging是一个网站服务程序，可以代替Nginx或Apache来搭建Web开发测试环境，并附带了实用的周边工具，例如：计划任务、MySQL管理、Redis管理、FTP管理、SSH管理、服务器管理等。

本软件项目不仅仅实现了一些网站服务工具，本身还是一个具有很好扩展性的通用网站后台管理系统，通过本项目，您可以很轻松的构建一个全新的网站项目，省去从头构建项目的麻烦，减少重复性劳动。

当您基于本项目来构建新软件的时候，您可以根据需要来决定是否使用本系统的网站服务工具，这取决于您是否在`main.go`中导入包：
```go
import (
	_ "github.com/admpub/nging/application/initialize/manager"
)
```

## 可执行文件下载

* [最新版下载地址](http://dl.webx.top/nging/latest/)

* [最新版备用地址](http://dl2.webx.top/nging/latest/)

## 版本说明

### V3 说明

1. 简化：移除上传文件后需要移动文件的功能，简化上传处理。
2. 修复：caddy服务申请HTTPS失败
3. 修复：MySQL管理工具管理数据库账号的bug，以及修复对MySQL8的支持
4. 改进：客户端加密方式改用国密SM2
5. 改进：安装程序增加数据编码选项，可以自由选择安装为utf8还是utf8mb4，以便于更好的兼容低于MySQL5.7的数据库
6. 新增：基于云存储增加云备份功能
7. 增加：arm-5、arm-6、arm-7、arm64(即arm-8) 二进制文件安装包，以便于支持树莓派设备

### V2 说明

相较于Nging V1.x进行了相当大的改进，包括增加了几个重量级的新功能、提高了稳定性、降低了CPU占用、修复了界面细节上的bug、作为通用后台基础框架这个定位来进行的项目结构优化

## 安装

1. 安装Nging

    1). 自动安装方式:

    ```sh
    sudo sh -c "$(wget https://raw.githubusercontent.com/admpub/nging/master/nging-installer.sh -O -)"
    ```

    或

    ```sh
    sudo wget https://raw.githubusercontent.com/admpub/nging/master/nging-installer.sh -O ./nging-installer.sh && sudo chmod +x ./  nging-installer.sh && sudo ./nging-installer.sh
    ```

    nging-installer.sh 脚本支持的命令如下

    命令 | 说明
    :--- | :---
    `./nging-installer.sh` 或 `./nging-installer.sh install` | 安装(自动下载nging并启动为系统服务)
    `./nging-installer.sh upgrade` 或 `./nging-installer.sh up` | 升级
    `./nging-installer.sh uninstall` 或 `./nging-installer.sh un` | 卸载

    2). 手动安装方式:  
    下载相应平台的安装包，解压缩到当前目录，进入目录执行名为“nging”的可执行程序(在Linux系统，执行之前请赋予nging可执行权限)。 例如在Linux64位系统，分别执行以下命令：

    ```sh
    cd ./nging_linux_amd64
    chmod +x ./nging
    ./nging
    ```

2. 配置Nging:  
    打开浏览器，访问网址 <http://localhost:9999/setup> ，
    在页面中配置数据库和管理员账号信息进行安装。

安装成功后，使用管理员账号登录。

## 开机自动运行

1. 首先，安装为服务，执行命令 `./nging service install`
2. 启动服务，执行命令 `./nging service start`

与服务相关的命令：

命令 | 说明
:--- | :---
`./nging service install` | 安装服务
`./nging service start` | 启动服务
`./nging service stop` | 停止服务
`./nging service restart` | 重启服务
`./nging service uninstall` | 卸载服务


## 0、基本功能

### 一. 系统设置

目前，在“系统设置”中提供了以下设置组：

1. `系统`: 目前提供了API密钥和调试模式的开关
2. `SMTP`: 配置用于发送邮件的SMTP服务器的相关设置
3. `日志`: 在这里可以设置如何保存和输出本系统所生成的日志


### 二. 用户管理

提供对后台用户的修改、添加和删除等操作。可以为后台用户分配相应的角色，从而对其进行权限控制

### 三. 角色管理

一个角色就是一个权限集合。在本系统中我们将权限分为了操作权限和指令权限。  

1. 操作权限：基于网址操作来进行权限判断
2. 指令权限：用于指定该角色可以执行的指令(这里的指令是指系统命令)

### 四. 邀请码

邀请码用来邀请其他人来注册成为本系统的后台用户。您可以指定受邀人注册后所拥有的角色，从而限制其操作权限。

### 五. 验证码

在这里可以查看短信验证码和邮箱验证码的发送记录以及使用时间

### 六. 指令集

为了对系统命令的执行进行权限控制，可以定义某个系统命令为指令

### 七. 修改个人资料

可以在此修改自己的头像、邮箱、手机号、密码等信息

### 八. 使用两步验证

如果想要提高账号的安全性，请启用两步验证。本系统已经实现了对两步验证的完整支持。


## Ⅰ、特色功能

### 一. Web服务

Web服务采用caddy作为内核。  
caddy 是类似于nginx或apache的网站服务软件。caddy的配置文件比nginx更加简洁易用。
Nging采用了图形化界面来配置caddy，这使得对caddy的配置变得更加的容易。

设置网站服务的一般流程为：添加分组 - 添加网站 - 启动服务

#### 1. 创建分组

首先，我们来添加一个分组。点击左边竖向导航栏上的`网站管理`展开子菜单，点击`添加分组`打开创建界面，在表单录入框`组名称`中输入名称后提交

#### 2. 添加网站

然后，我们来添加一个网站。点击左边竖向导航栏上的`网站管理`展开子菜单，点击`添加网站`打开创建界面，这里向我们展示了创建网站服务所要配置的一些录入框，我们现在分别来说明各个录入框的作用:

- `分组`：如果网站特别多，需要按组进行归类，您可以在这里选择相应分组，这里我们选择第一步添加的分组  
- `监听地址`：可以是网址、域名或IP。如果不提供端口则默认为2005，如果不提供协议(比如https://)则默认为http，如果不提供IP和域名则默认为0.0.0.0，如果只想允许本地访问请设置为localhost或127.0.0.1。  
  域名支持通配符“\*”(例如\*.admpub.com)和环境变量(环境变量用花括号括起来，例如localhost:{$PORT})。  
  例子: https://admpub.com:443 admpub.com:80 :8080 127.0.0.1:9999  
  监听地址如有多个，用空格隔开。  

- `网站位置`：网站文件夹在服务器上的绝对路径  
- `网站名称`：网站的名字
- `日志文件`：输入保存访问日志文件的路径
- `默认首页`：在访问某个目录的时候，默认的首页文件。如有多个，用空格隔开。如不填写，默认为：index.html index.htm index.txt default.html default.htm default.txt  
- `响应Header`：可以对网址中的某个起始路径的访问设置响应头参数，支持添加、删除和修改Header  
    > 删除Header：在Header名称前使用减号(-)并将其值留空
- `HTTPS`：支持手动和自动设置SSL证书。  
  在手动模式下，需要分别填写证书和私钥文件位置；在自动模式下必须填写电子邮箱地址以便自动获取Let's encrypt 证书  
- `GZIP`：压缩数据，用于减少页面或静态文件的网络传输尺寸，可以通过指定扩展名的方式来对特定文件进行压缩  
- `FastCGI`：FastCGI代理用于将请求转发到FastCGI服务。尽管此指令最常见的用途是为PHP站点提供服务，但默认情况下它是通用的。支持设置多个不同的请求路径。  
- `Proxy`：反向代理功能用于将接收到的请求转发到指定的后台。通过设置多个后台地址我们能实现服务器的负载均衡。  
- `文件服务`：用于提供文件的直接访问  
- `IP过滤`：通过设置过滤规则来禁止某些IP或国家来访问我们的网站  
- `网址重写`：将请求的网址重写为其它值。常用于资源重定向、伪静态等场景  
- `备注`：记录当前网站的一些附加说明  
- `网站状态`：如果是`启用`状态则表示该网站能够上线  
- `是否重启`：修改配置后，必须重启Web服务后才会生效  

#### 3. 启动服务

Web服务是默认启动的。  

可以在`网站列表`页面点击`重新应用配置`来重新生成Caddy配置文件和重启服务。  

当然，你也可以在左边的导航菜单中点击`服务器`链接然后再点击展开的子菜单`服务管理`，在打开的页面中点击Web服务的控制按钮来进行单纯的重启和关闭

### 二. 服务器

1. 服务器信息

   展示服务器的操作系统、CPU、磁盘、内存、网络等信息。

   在这个页面中，使用了进度条和百分比来直观的展示数据和资源的占用量。

2. 网络端口

   列出占用网络端口的进程，并且支持关闭进程

3. 执行命令
  
   执行简单的控制台命令。

   用聊天对话框的方式来与机器进行简单的交互。

4. 服务管理

   控制Web和FTP服务的启动和关闭

5. 进程值守

   进程保活，即当执行的程序被异常关闭后自动重新启动此程序

   1. 添加配置

       在左边的导航菜单中点击`服务器`链接然后再点击展开的子菜单`进程值守`打开列表页，点击列表右上角的`添加配置`按钮，我们来添加一个配置，在配置表单中我们可以录入以下信息：

      - `状态`：指定是否启用当前配置。如果启用则会自动开始值守  
      - `名称`：描述一下要值守的这个进程的名字  
      - `命令`：进程的启动命令  
      - `环境变量`：如果启动的程序需要一些特定的环境变量，您可以在这里进行设置。输入的格式为“`环境变量的名称`=`环境变量的值`”，比如“`NAME=Nging`”。如有多个，一行一个  
      - `命令参数`：其实在`命令`输入框中已经可以带参数了，但是如果参数比较多，则可以在这里添加，格式为`-c=value`，如有多个则一行一个  
      - `工作目录`：有的时候我们需要切换到某个目录下之后再执行程序，对于这种情况在这里输入这个目录的路径即可  
      - `信息日志文件`：值守的信息日志保存位置。如果不填，则输出到控制台  
      - `错误日志文件`：错误日志保存位置。如果不填，则输出到控制台  
      - `重试次数`：连续重试的次数。如果连续启动程序都失败了并且达到重试次数的限制，则退出值守  
      - `启动延时`：程序异常退出后，等待多长时间后再启动。格式为数字和单位字母的组合，有效的单位有："ns", "us" (or "µs"), "ms", "s", "m", "h" 分别表示 "纳秒", "微秒", "毫秒", "秒", "分钟", "小时"  
      - `心跳检测间隔`：也就是多久检测一次程序运行状态。默认1m，即1分钟。格式同`启动延时`  
      - `说明`：该配置的附加说明  
      - `调试模式`：如果开启，则会在控制台输出调试信息  

   2. 启动和关闭值守

       在列表页点击“启用状态”列中的复选框可以控制值守程序的开和关

### 三. FTP账号

1. 添加账号分组

    对于有大量FTP账号的情况，可以通过建立分组来进行归类

2. 添加FTP账号

    要使用Nging提供的FTP功能，必须首先要创建一个FTP账号。  

    FTP客户端使用这里创建的账号来登录。

3. 开关FTP服务

    在左边的导航菜单中点击`服务器`链接然后再点击展开的子菜单`服务管理`，在打开的页面中点击FTP服务的控制按钮来进行重启和关闭

    为了避免与原有FTP服务发生冲突，本系统的FTP服务默认端口为`233`，可以通过修改配置文件config/config.yaml中ftp节点内的port字段值来指定其它端口。

### 四. 数据采集

数据采集模块提供了强大的采集功能，它包含了以下特色：

- 灵活便捷的配置表单
- 支持无限级页面
- 支持多种页面格式(html,json,text)
- 支持多类型采集规则(regexp,regexp2,goquery)
- 丰富的过滤器和验证器
- 支持多种浏览器引擎(chromedp，webdriver等)
- 三种去重验证机制
- 支持代理
- 支持自动导出采集数据到不同的数据源(WebAPI,DSN)
- 支持定时采集

1. 添加采集规则

    在左边的导航菜单中点击`数据采集`链接然后再点击展开的子菜单`新建规则`，我们来添加一个新规则，在配置表单中我们可以录入以下信息：

    - `规则名称`：首先您得为规则取一个名字  
    - `分组`：为当前规则选择一个分组  
    - `说明`：规则附加说明  
    - `判断重复`：选择一个判断重复的方式，这样可以避免重复采集相同的数据
    - `入口页面网址`：在这里，您可以指定多个入口页面网址，每一个网址单独放在一行，支持使用Go语言的模板语法循环输出网址，此外，每一个网址中都支持使用`数字范围`标签，在数字范围标签中，连续的范围我们使用`-`来指定，不连续的范围使用`,`来罗列，同时支持指定步进值，且步进值与范围值之间用`:`分隔，默认步进值为1，例如：`{(`1-9,11,13-19:2`)}`。  

       ```plain
      例如：http://www.admpub.com/{(1-2)}.html 会生成网址：  
      http://www.admpub.com/1.html  
      http://www.admpub.com/2.html  
       ```

    - `最大超时`：请求页面时的最大等待时长(秒)  
    - `间歇时间`：每个页面的随机等待的秒数范围  
    - `代理地址`：例如：`http://admpub:123456@123.123.123.123:8080`。支持格式`protocol://user:password@ip:port`  
        * `protocol` - 支持http、https、socks5
        * `user` - 用户名(选填)
        * `password` - 密码(选填)
        * `ip` - IP地址
        * `port` - 端口
    - `页面格式`：目标页面的内容格式，支持的有HTML/JSON/Text  
    - `浏览引擎`：目前支持standard/chromedp/webdriver这三种，其中standard为普通方式，速度最快，默认为standard。chromedp和webdriver为调用chrome浏览器，速度相对较慢  
    - `页面类型`：可以选择“列表页”和“内容页”选项，请根据所采集页面的实际类型来选择  
    - `页面字符集`：比如`gbk`或`utf-8`。此项为选填，不填写的情况下会根据网页内容自动判断  
    - `区域规则`：指定所要采集页面上的某个区域的匹配规则。这里支持两种类型的规则：  

      1. 选择器规则：即类似于jQuery匹配页面元素的规则，比如：`div.container > h1`
      2. 正则匹配规则：需要在规则前添加前缀`regexp:`或`regexp2:`。其中`regexp:`为Go语言原生正则表达式规则；`regexp2:`为兼容 Perl5 和 .NET 的正则表达式规则（相关文档：<https://github.com/admpub/regexp2>）  

    - `规则`：此处指定以下规则  

      1. 数据采集规则：所要采集的数据匹配规则，如果指定了`区域规则`则在该区域内进行查找
      2. 数据保存变量名：指定数据保存到哪个变量里
      3. 数据过滤器和验证器：多个过滤器或验证器之间用管道符“|”分隔，以下划线“_”开头的为验证器（不符合验证器规则的数据会被跳过）

    然后可以点击“添加下一级页面”链接来对添加第二级页面的规则。页面的层级数量没有限制。

2. 测试采集规则

点击“测试”按钮可以进行采集测试，为了快速获取测试结果，每一级页面只采集一条

3. 执行采集

点击“采集”按钮进行手动采集

4. 将采集加入计划任务

也可以将采集任务加入到计划任务中，进行自动定时采集

5. 添加导出规则

    定义采集数据的导出方式。如果希望在采集的时候自动将采集到的数据导入其它系统，请在这里添加导出规则。
    目前支持两种导出方式：

    1、导入到数据库；
    2、提交到API接口（将JSON数据POST提交到API接口）

### 五. 计划任务

crontab的完美替代，并采用了图形化配置界面，支持记录日志、发送错误报告邮件等

### 六. 离线下载

这里的离线下载支持并行下载。  
对于需要很长时间才能下载完成的大文件，我们只需要建立一个离线下载任务并且启动，然后就可以关闭浏览器去做其它的事情了，Nging会自动在后台帮您下载。

### 七. 云服务

1. 云存储账号 可以管理所有与Amazon S3 API兼容的对象存储文件，比如亚马逊AWS、阿里云OSS、腾讯云COS、网易云NOS、百度云BOS、华为云OBS、七牛云Kodo等
2. 文件备份 可以通过配置将指定目录下的文件备份到云存储

### 八. 数据库管理

1. 管理MySQL
2. 管理Redis
3. MySQL表结构比较/同步

### 九. FRP内网穿透

当你想要将局域网内的电脑暴露到外网，以便于外网的用户也能访问到您的网站服务时，这时候就需要用到内网穿透。

内网穿透支持服务端模式和客户端模式，全图形化的配置界面让配置变得非常容易。

要成功使用内网穿透功能，必须在局域网电脑上启动客户端模式，在提供外网的服务器或VPS上启动服务端模式。

### 十. SSH管理

SSH管理功能实现了SSH的Web客户端管理功能，您可以在Nging网页上进行SSH命令行交互操作，也可以通过Nging以SFTP方式来进行文件的上传、下载、删除和编辑


## Ⅱ、先睹为快

### 运行

[![安装](https://gitee.com/admpub/nging/raw/master/preview/preview_cli.png?raw=true)](https://gitee.com/admpub/nging/raw/master/preview/preview_cli.png)

### 安装：

[![安装](https://gitee.com/admpub/nging/raw/master/preview/preview_install.png?raw=true)](https://gitee.com/admpub/nging/raw/master/preview/preview_install.png)

### 登录：

[![登录](https://gitee.com/admpub/nging/raw/master/preview/preview_login.png?raw=true)](https://gitee.com/admpub/nging/raw/master/preview/preview_login.png)

### 系统信息：

[![系统信息](https://gitee.com/admpub/nging/raw/master/preview/preview_sysinfo.png?raw=true)](https://gitee.com/admpub/nging/raw/master/preview/preview_sysinfo.png)

### 实时状态：

[![实时状态](https://user-images.githubusercontent.com/512718/59155431-376ebe00-8abc-11e9-8d29-cee91978e574.png)](https://user-images.githubusercontent.com/512718/59155431-376ebe00-8abc-11e9-8d29-cee91978e574.png)


### 在线编辑文件：

[![在线编辑文件](https://gitee.com/admpub/nging/raw/master/preview/preview_editfile.png?raw=true)](https://gitee.com/admpub/nging/raw/master/preview/preview_editfile.png)

### 添加计划任务：

[![添加计划任务](https://gitee.com/admpub/nging/raw/master/preview/preview_task.png?raw=true)](https://gitee.com/admpub/nging/raw/master/preview/preview_task.png)

### MySQL数据库管理：

[![MySQL数据库管理](https://gitee.com/admpub/nging/raw/master/preview/preview_listtable.png?raw=true)](https://gitee.com/admpub/nging/raw/master/preview/preview_listtable.png)

## Ⅲ、开发环境下的启动方式

- 第一步： 安装GO环境(必须1.12.1版以上)，配置GOPATH、GOROOT环境变量，并将`%GOROOT%/bin`和`%GOPATH%/bin`加入到PATH环境变量中
- 第二步： 执行命令`go get github.com/admpub/nging`
- 第三步： 进入`%GOPATH%/src/github.com/admpub/nging/`目录中启动`run_first_time.bat`(linux系统启动`run_first_time.sh`)
- 第四步： 打开浏览器，访问网址`http://localhost:8080/setup`，在页面中配置数据库账号和管理员账号信息进行安装
- 第五步： 安装成功后会自动跳转到登录页面，使用安装时设置的管理员账号进行登录


请注意，本系统的源代码基于AGPL协议发布，不管您使用本系统的完整代码还是部分代码，都请遵循AGPL协议。  
> 如果需要更宽松的商业授权协议，请联系我购买授权。
