
# 使用方式：
# sh protoc_gen.sh ${输入的协议的目录名称} ${协议类型}
# 如： sh goproto.sh wx_user_login_svr http
# 表示按照wx_user_login_svr下的协议，生成最终的协议桩代码，以http的协议对外提供服务。   


# 1. 参数指定输入协议文件
# 第一个参数为输入的协议文件
target=$1 
targetPBFile=${target}.proto
targetOutDir=$1
# 第二个参数指定编译的协议类型. "http"表示支持http协议的桩代码编译，否则，为普通的grpc-go代码编译. 
type=$2 


echo '开始生成pb协议文件,目标文件:'${targetPBFile}', 输出目录:'${targetOutDir}', 模式:'${type}

# 3. 用于生成pb文件
# --go_out参数：
# plugins=grpc 表示 protoc-gen-go 插件使用grpc的工作方式
# paths=source_relative:${targetOutDir} 表示通过相对目录来指定 生成的pb桩代码的路径
# --go_out=plugins=grpc:. 这是默认path使用import方式，表示在当前目录下，使用go_package的包名的全路径，来创建代码输出目录 (不太实用！！)
# protoc --proto_path=${target} --proto_path=./ --go_out=plugins=grpc,paths=source_relative:${targetOutDir} ${targetPBFile}  ## 正解
# protoc --proto_path=${target} --proto_path=./ --go_out=plugins=grpc:. ${targetPBFile} # 不太适用，不能在当前目录层次创建输出目录

if [ "$type" = "" ]; then # 变量比较方式
    protoc --proto_path=${target} --proto_path=./ --go_out=plugins=grpc,paths=source_relative:${targetOutDir} ${targetPBFile};
else 
    # 参考文章： https://www.cnblogs.com/cxt618/p/15647316.html
    protoc --proto_path=${target} --proto_path=$GOPATH/src/googleapis --proto_path=./ --go_out=plugins=grpc,paths=source_relative:${targetOutDir} \
     --grpc-gateway_out ${targetOutDir} --grpc-gateway_opt paths=source_relative ${targetPBFile};
     # --grpc-gateway_opt paths=source_relative 表示按照相对路径进行查找
     # --grpc-gateway_out ${targetOutDir} 表示输出文件的目录
fi  

echo ${targetPBFile}'生成pb协议文件完成'


