/**
 * gdb daemon
 * gdb 调试测试使用
 * l -- list
 * b -- break + 行号/函数名
 * info b -- 显示断点信息
 * r -- run运行
 * c -- continue程序继续运行
 * n -- 单步运行 不进入函数
 * step -- 单步运行 进入函数
 * u -- 退出循环体
 * p -- print + 变量名
 * bt -- 查看函数堆栈
 * finish -- 结束
 * q -- 退出
 * disable -- 停止点
 * enable -- 停止点
 * clear -- 清除所有的已定义的停止点
 * delete -- 删除指定的断点
 */
#include <stdio.h>

int func(int n)
{
    int sum=0,i;
    for(i = 0; i < n; i++) {
        sum += i;
    }
    return sum;
}

int main(void)
{
    int i;
    long result = 0;
    for(i = 1; i <= 100; i++) {
        result += i;
    }

    printf("result[1-100] = %ld \n", result);
    printf("result[1-250] = %d \n", func(250));
    return 0;
}

