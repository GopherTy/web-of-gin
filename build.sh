#!/bin/bash 

# root directory $BASH_SOURCE 变量在脚本文件中可以显示脚本的路径，但是在 shell 命令行中什么都不会输出。
DirRoot=`cd $(dirname $BASH_SOURCE) && pwd`
Target=server
TestItems=()

function manual(){
    echo "l/linux           build for linux(编译到linux)"
    echo "w/windows         build for windows(编译到windows)"
    echo "d/darwin build for macos(编译到macos)"
    echo "t/test            feature test(功能测试)"
}

case $1 in 
    l|linux)
        export GOOS=linux CGO_ENABLED=0 
        cd $DirRoot && go build -ldflags "-s -w" -o "$DirRoot/bin/$Target"
    ;; 
    w|windows)
	export  GOOS=windows CGO_ENABLED=0 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOARCH=amd64
        cd $DirRoot && go build -ldflags "-s -w" -o "$DirRoot/bin/$Target".exe
    ;;
    d|darwin)
       export GOOS=darwin CGO_ENABLED=0 
    	cd $DirRoot && go build -ldflags "-s -w" -o "$DirRoot/bin/$Target"
    ;;
    t|test)
        for i in ${!TestItems[@]}
        do 
            cd "$DirRoot/${TestItems[i]}" && go test
        done
    ;;
    *)
        manual
    ;;
esac
