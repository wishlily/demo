/*
 * 同步，消费者
 */
#include <sys/types.h>
#include <sys/sem.h>
#include <sys/ipc.h>
#include <string.h>
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main (void)
{
    key_t ipckey;
    int semid;
    struct sembuf sem[1];

    ipckey = ftok("/tmp/test", ':');

    semid = semget(ipckey, 1, 0600|IPC_CREAT);
    if (semid < 0) {
        printf("Sem Create Error - %s\n", strerror(errno));
        exit(1);
    }
    // int val = semctl(semid, 0, GETVAL);
    // printf("Sem Val %d\n", val);
    sem[0].sem_num = 0;
    sem[0].sem_flg = SEM_UNDO;
    int loop = 10;
    while(loop--) {
        sem[0].sem_op = -1;
        semop(semid, sem, 1);
        printf("I consume\n");
        sleep(0);
    }
    if (semctl(semid, 0, IPC_RMID) == -1) {
        printf("Sem Delete Error - %s\n", strerror(errno));
        exit(1);
    }
    printf("Sem Delete!\n");
}
