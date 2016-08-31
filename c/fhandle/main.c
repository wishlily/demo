/*
 * 写入文件保存数据
 *
 * r 以只读方式打开文件，该文件必须存在。
 * r+ 以可读写方式打开文件，该文件必须存在。
 * rb+ 读写打开一个二进制文件，允许读写数据，文件必须存在。
 * w 打开只写文件，若文件存在则文件长度清为0，即该文件内容会消失。若文件不存在则建立该文件。
 * w+ 打开可读写文件，若文件存在则文件长度清为零，即该文件内容会消失。若文件不存在则建立该文件。
 * a 以附加的方式打开只写文件。若文件不存在，则会建立该文件，如果文件存在，写入的数据会被加到文件尾，即文件原先的内容会被保留。（EOF符保留）
 * a+ 以附加方式打开可读写的文件。若文件不存在，则会建立该文件，如果文件存在，写入的数据会被加到文件尾后，即文件原先的内容会被保留。 （原来的EOF符不保留）
 * wb 只写打开或新建一个二进制文件；只允许写数据。
 * wb+ 读写打开或建立一个二进制文件，允许读和写。
 * ab+ 读写打开一个二进制文件，允许读或在文件末追加数据。
 * wx 创建文本文件,只允许写入数据.[C11]
 * wbx 创建一个二进制文件,只允许写入数据.[C11]
 * w+x 创建一个文本文件,允许读写.[C11]
 * wb+x 创建一个二进制文件,允许读写.[C11]
 * w+bx 和"wb+x"相同[C11]
 * 以x结尾的模式为独占模式,文件已存在或者无法创建(一般是路径不正确)都会导致fopen失败.文件以操作系统支持的独占模式打开.[C11]
 *
 * 解析数据
 * 常用函数
 * sscanf
 * strchr 查找某个字符
 * strstr 查找某字符串
 */
#include <stdio.h>
#include <stdlib.h>
#include <signal.h>
#include "parse.h"

FILE *fd = NULL;

static void clean_exit(int sig)
{
    // 没有该步则out.log不会有数据,kill没有正常关闭退出
    printf("clean_exit: %d\n", sig);
    exit(sig);
}

/** 从屏幕获得数据保存到文件,死循环 */
static int get_stdin_to_file(FILE *fd)
{
    fd = fopen("./text.log", "w");
    if (fd == NULL) {
        printf("fopen fail...\n");
    }

    for (;;) {
        char buff[256];
        if (fgets(buff, sizeof(buff), stdin) != NULL) {
            fprintf(fd, "%s\n", buff);
        }
    }
    fclose(fd);
    return 0;
}

int main(int argc, char *argv[])
{
    printf("main start...\n");

    if (argc != 2) {
        printf("usage: %s <config file name>\n", argv[0]);
        exit (EXIT_FAILURE);
    }

    signal(SIGINT, clean_exit);

    printf("parse: %s\n", parse(argv[1]));

    get_stdin_to_file(fd);

    // never get here - but it's proper form to include it
    exit(EXIT_SUCCESS);
}// main end

