<h1 align="center">
    Govo
</h1>

<div align="center">
  基于gin + Vue + Redis + Docker + MySQL的入门 GoWeb 应用
</div>
 <br/>
 <br/>
<p align="center">
  <a href="https://github.com/Nitrosaccharose/go-web/blob/main/LICENSE"><img src="https://img.shields.io/github/license/Nitrosaccharose/go-web" alt="License"></a>
  <br/>
  本教程使用AGPL-3.0开源协议
<p/>

## 🔍 项目预览
<div align=center>
    <img  src="/img/g1.gif" height="70%" width="70%"/>
</div>
<div align=center>
    <img  src="/img/g2.gif" height="60%" width="60%"/>
</div>
<div align=center>
    <img  src="/img/img1.png" height="45%" width="45%"/>
    <img  src="/img/img2.png" height="45%" width="45%"/>
    <img  src="/img/img3.png" height="45%" width="45%"/>
    <img  src="/img/img4.png" height="45%" width="45%"/>
</div>


## 🧭 项目简介
本项目是Go语言实现GoWeb应用的一个练手项目，项目采用前后端分离的形式来实现。项目实现了User类的CRUD，可与MySQL进行数据通信。实现了KVTObject类，可与Redis进行数据通信。


## 🧰 项目选型
                

| 技术          | 	简介                                                                                                             |
|-------------|-----------------------------------------------------------------------------------------------------------------|
| Vue	        | Vue是一套用于构建用户界面的渐进式框架，具有轻量、高效、易用等特点。Vue的核心库只关注视图层，非常容易上手，可以与其它第三方库或既有项目整合。                                       |
| Vite	       | Vite是一款基于浏览器原生 ES 模块构建的前端构建工具，开发模式下利用浏览器原生 ES 模块直接导入文件的方式来提高构建速度和开发效率。                                          |
| Vue-router	 | Vue-router是Vue.js官方的路由管理器，用于创建单页面应用。它通过URL映射到组件，实现了前端路由和组件的耦合。                                                  |
| ElementPlus | 	ElementPlus是一款基于Vue 3.0的UI组件库，包含丰富的组件和强大的功能。它具有简单易用、高效稳定等特点，可以快速搭建优秀的用户界面。                                     |
| Axios	      | Axios是一个基于Promise的HTTP客户端，可以用于浏览器和Node.js环境中，支持请求拦截、响应拦截、请求取消等特性，提供了一种优雅、简洁的方式来处理HTTP请求。                        |
| Go	         | Go是一种高效、可靠的编程语言，具有良好的并发编程和内存管理机制，被广泛应用于网络编程、系统编程、云计算、人工智能等领域。                                                   |
| gin	        | Gin是一个使用Go语言编写的高性能Web框架，具有轻量级、高性能、易用等特点，被广泛应用于RESTful API的开发。                                                   |
| gorm	       | GORM是一个使用Go语言编写的ORM框架，支持多种数据库，包括MySQL、PostgreSQL、SQLite等，具有简单易用、功能丰富等特点。                                        |
| go-redis	   | go-redis是Go语言的Redis客户端库，提供了完整的Redis命令封装和一些额外的功能，具有高性能、易用等特点。                                                    |
| Docker	     | Docker是一款开源的容器化平台，可以将应用程序及其依赖项打包到一个可移植的容器中，从而实现应用程序的快速部署、跨平台移植等优势。Docker还支持镜像管理、容器网络、存储卷等功能，使得应用程序在不同环境中运行更加方便。 |
| Redis	      | Redis是一个高性能的键值存储系统，支持多种数据结构，包括字符串、哈希、列表、集合、有序集合等。Redis被广泛应用于缓                                                   |
| MySQL       | MySQL是一款开源的关系型数据库管理系统，具有成本低廉、易于使用、稳定性高等特点，被广泛应用于Web应用程序、数据仓库、企业应用等领域。                                           |

## 🍔 食用方法
使用以下命令运行后端：
```
go run main.go
```
<br/>
使用以下命令运行前端：

```
cd go-web-front-project
npm run dev
```
<br/>

使用以下命令构建Redis集群

```
docker compose up -d
docker run -it --rm --network=cluster_default redis redis-cli -h cluster-redis-node-0-1 -p 7000 -c
```

## 🧧 赞助
如果喜欢的话，请作者喝杯咖啡吧！（赞助一点点也行喔！非常感谢QWQ）
<div align=center>
  <img  src="/img/2dcode.png" height="45%" width="45%"/>
</div>

## 🤙 联系我
📥 Email: 931911761@qq.com
