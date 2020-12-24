#!/bin/bash
#echo $@

# import
. ./print.sh

PRINT echo "Test Start ..."
# export DEBUG=true
DEBUG echo "Debug print ..."

function help() {
    echo "Usage: ${0##*/} [-options]"
    echo ""
    echo "where options include:"
    echo "    -a|along            xxx, eg: --along xxx"
}

#-o或--options选项后面接可接受的短选项，如ab:c::，表示可接受的短选项为-a -b -c，其中-a选项不接参数，-b选项后必须接参数，-c选项的参数为可选的
#-l或--long选项后面接可接受的长选项，用逗号分开，冒号的意义同短选项。
#-n选项后接选项解析错误时提示的脚本名字
ARGS=`getopt -o hab:c:: --long along,blong:,clong:: -n 'example.sh' -- "$@"`
if [ $? != 0 ]; then
    echo "Terminating..."
    help && exit 1
fi

#echo $ARGS
#将规范化后的命令行参数分配至位置参数（$1,$2,...)
eval set -- "${ARGS}"

while true
do
    case "$1" in
        -a|--along)
            echo "Option a";
            shift
            ;;
        -b|--blong)
            echo "Option b, argument $2";
            shift 2
            ;;
        -c|--clong)
            case "$2" in
                "")
                    echo "Option c, no argument";
                    shift 2
                    ;;
                *)
                    echo "Option c, argument $2";
                    shift 2;
                    ;;
            esac
            ;;
        --)
            shift
            break
            ;;
        -h)
            help && exit 0
            ;;
        *)
            echo "Internal error!"
            help && exit 1
            ;;
    esac
done

#处理剩余的参数
for arg in $@
do
    echo "processing $arg"
done

# ./main.sh -b 123 -a -c456 file1 file2