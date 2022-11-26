# proto_proj
公共协议工程文档

-- common协议的文件夹 通过 common开头
-- 其他的服务以服务为后缀就好


编译脚本：
（1）先手动编译出依赖的proto的pb文件
（2）在手动编译上层的proto的pb文件 


# protoc基本使用
- 搜索proto文件的路径参数

-IPATH, --proto_path=PATH:它表示的是我们要在哪个路径下搜索.proto文件，这个参数既可以用-I指定，也可以使用--proto_path=指定。
如果不指定该参数，则默认在当前路径下进行搜索；另外，该参数也可以指定多次，这也意味着我们可以指定多个路径进行搜索。
注意：是proto文件的搜索路径，而不是其他的。

- proto文件位置参数 

具体的，真的要编译的proto文件，通过这个参数指定。
运行原理是：先通过PATH，找到一个父路径，然后，根据本参数，来找到具体的，需要编译的文件。 

- 语言插件参数

由于protoc支持各种不同的语言，比如cpp或者golang。
这个参数，就是用来指定使用的生成proto哪种语言桩代码的插件。
比如：如果是生成golang代码，则使用的是 --go_out=对应的是protoc-gen-go。

每种语言的插件，都需要单独自己安装的。

参考： https://juejin.cn/post/6949927882126966820 
---
# 最新版
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# 指定版本
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.3.0
---

## package & go_package
package参数针对的是protobuf，是proto文件的命名空间，它的作用是为了避免我们定义的接口，或者message出现冲突。应该是针对proto文件而言的，而不是桩代码。




## 脚本使用
### 如何只编译proto
### 如何编译普通proto&interface
### 如何编译http的proto&http_interface
这个的关键是要解决proto引用外部包的问题：
http://t.zoukankan.com/yisany-p-14875488.html  
总体是，将对应的库拷贝到gopath对应的目录下，然后引用就好了。
---
git clone https://github.com/googleapis/googleapis.git
---
