/*
 * 生成守护进程以及其他一些配置
 */
#include <sys/types.h>
#include <sys/stat.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <signal.h>

#define RWRWRW (S_IRUSR|S_IWUSR|S_IRGRP|S_IWGRP|S_IROTH|S_IWOTH)

void clean_exit(int sig)
{
    // 没有该步则out.log不会有数据,kill没有正常关闭退出
    printf("clean_exit: %d\n", sig);
    exit(sig);
}

int main(int argc, char *argv[])
{
    pid_t pid, sid;

    printf("MAIN start...\n");

    signal(SIGTERM, clean_exit);

    /* Fork off the parent process */
    pid = fork();
    if (pid < 0) {
        exit(EXIT_FAILURE);
    }
    /* If we got a good PID, then
       we can exit the parent process. */
    if (pid > 0) {
        printf("DAEMON son id : %d\n", pid);//< 退出父进程,生成一个守护进程,pid==0是子进程
        exit(EXIT_SUCCESS);
    }

    /** 从这开始都是守护进程执行操作(子进程) */

    /* Change the file mode mask */
    umask(0);//< 权限最小所有用户

    /* Open any logs here */

    /* Create a new SID for the child process */
    sid = setsid();
    if (sid < 0) {
        /* Log the failure */
        exit(EXIT_FAILURE);
    }

    /* Change the current working directory */
    if ((chdir("/")) < 0) {//< 改变工作目录
        /* Log the failure */
        exit(EXIT_FAILURE);
    }

    /* Close out the standard file descriptors */
    close(STDIN_FILENO);
    //close(STDOUT_FILENO);//< 关闭标准输出
    //close(STDERR_FILENO);//< 关闭标准错误输出

    /* 重定向标准输出 */
    mkdir("tmp/daemon", RWRWRW);//< 如果已存在会返回-1
    if (freopen("tmp/daemon/out.log", "w", stdout) == NULL) {
        printf("stdout freopen fail...\n");
    }
    if (freopen("tmp/daemon/err.log", "w", stderr) == NULL) {
        printf("stderr freopen fail...\n");
    }

    printf("DAEMON start...\n");
    printf("Hello World!\n");

    /* The Big Loop */
    for(;;)
    {
        sleep(1);
    }

    // never get here - but it's proper form to include it
    exit(EXIT_SUCCESS);
}// main end

