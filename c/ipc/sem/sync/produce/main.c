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
    struct sembuf sem[2];

    ipckey = ftok("/tmp/test", ':');

    semid = semget(ipckey, 1, 0600|IPC_CREAT|IPC_EXCL);
    if (semid < 0) {
        printf("Sem Create Error - %s\n", strerror(errno));
        exit(1);
    }

    sem[0].sem_num = 0;
    sem[1].sem_num = 0;
    sem[0].sem_flg = SEM_UNDO;
    sem[1].sem_flg = SEM_UNDO;
    int loop = 10;
    while(loop--) {
        sem[0].sem_op = 0;
        sem[1].sem_op = 1;
        semop(semid, sem, 2);
        printf("I produce\n");
        sleep(3);
    }
}
