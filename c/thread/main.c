/**
 * 新建线程 & 终端显示(清屏)
 */
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <signal.h>
#include <pthread.h>

pthread_t displayThread;

static unsigned int count = 0;

static void clean_exit(int sig)
{
    // 没有该步则out.log不会有数据,kill没有正常关闭退出
    printf("clean_exit: %d\n", sig);
    exit(sig);
}

static void *displayThreadRoutine(void *args)
{
    for (;;) {
        sleep (1);
        //< 清屏
        printf("\033[2J");
        printf("\033[H");
        printf("The Count: %010d\n", count);
    }
    return ((void *)0);
}

int main(int argc, char *argv[])
{
    printf("main start...\n");

    signal(SIGINT, clean_exit);

    pthread_create(&displayThread, NULL, displayThreadRoutine, NULL);
    sleep(10);

    for (;;) {
        count++;
        sleep(1);
    }

    // never get here - but it's proper form to include it
    exit(EXIT_SUCCESS);
}// main end

