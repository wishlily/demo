/*
 * 消息队列，服务端
 */
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <errno.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/stat.h>
#include <sys/msg.h>
#include <signal.h>

#define MSG_FILE "server.c"
#define BUFFER 255
#define PERM S_IRUSR|S_IWUSR

int msgid;

static void clean_exit(int sig)
{
    // 从内核中移除消息
    if (msgctl(msgid, IPC_RMID, NULL) == -1) {
        fprintf(stderr,"Remove Message Error：%s\a\n",strerror(errno));
        exit(1);
    }
    exit(sig);
}

struct msgtype {
    long mtype;
    char buffer[BUFFER+1];
};

int main()
{
    signal(SIGINT, clean_exit);

    struct msgtype msg;
    key_t key;
    // 获得消息键
    if((key=ftok(MSG_FILE,'a'))==-1) {
        fprintf(stderr,"Creat Key Error：%s\a\n",strerror(errno));
        exit(1);
    }
    // 创建消息
    if((msgid=msgget(key,PERM|IPC_CREAT|IPC_EXCL))==-1) {
        fprintf(stderr,"Creat Message Error：%s\a\n",strerror(errno));
        exit(1);
    }
    while(1) {
        msgrcv(msgid,&msg,sizeof(struct msgtype),1,0);
        fprintf(stderr,"Server Receive：%s\n",msg.buffer);
        msg.mtype=2;
        msgsnd(msgid,&msg,sizeof(struct msgtype),0);
    }
    exit(0);
}
