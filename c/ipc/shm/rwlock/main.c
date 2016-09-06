/*
 * 读写锁，适合写少读多
 * 0: write
 * 1: read
 */
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/shm.h>
#include <sys/sem.h>
#include <errno.h>
#include <string.h>
#include <unistd.h>

#define SHM_TEST_DATALEN 256

struct shm_test_st {
    char dat[SHM_TEST_DATALEN];
};

const char g_test_data[][SHM_TEST_DATALEN] = {
    {"Nothing gold can stay"},
    {"Nature's first green is gold"},
    {"Her hardest hue to hold"},
    {"Her early leaf's a flower"},
    {"But only so an hour"},
    {"Then leaf subsides leaf"},
    {"So Eden sank to grief"},
    {"So down gose down to day"},
    {"Ending"}
};

int shm_write(struct shm_test_st *pdata, int semid)
{
    struct sembuf sem[3];
    sem[0].sem_num = 0; // w
    sem[1].sem_num = 1; // r
    sem[2].sem_num = 0; // w
    sem[0].sem_flg = SEM_UNDO;
    sem[1].sem_flg = SEM_UNDO;
    sem[2].sem_flg = SEM_UNDO;
    int num = sizeof(g_test_data) / sizeof(g_test_data[0]);
    int i;
    for (i = 0; i < num; i++) {
        sem[0].sem_op = 0; // wait write free
        sem[1].sem_op = 0; // wait read free
        sem[2].sem_op = 1; // write +1 -> going
        semop(semid, sem, 3);
        printf("Write - %s\n", g_test_data[i]);
        strcpy(pdata->dat, g_test_data[i]);
        sem[0].sem_op = -1;
        semop(semid, sem, 1); // write -1 -> over
        sleep(1);
    }
    printf("Write Over!\n");
    return 0;
}

int shm_read(struct shm_test_st *pdata, int semid)
{
    struct sembuf sem[2];
    sem[0].sem_num = 0; // w
    sem[1].sem_num = 1; // r
    sem[0].sem_flg = SEM_UNDO;
    sem[1].sem_flg = SEM_UNDO;
    sleep(1);
    char tmp[SHM_TEST_DATALEN];
    while (1) {
        sem[0].sem_op = 0; // write is free
        sem[1].sem_op = 1; // read +1 -> going
        semop(semid, sem, 2);
        strcpy(tmp, pdata->dat);
        sem[0].sem_num = 1; // r
        sem[0].sem_op = -1; // read -1 -> over
        semop(semid, sem, 1);
        int num = sizeof(g_test_data) / sizeof(g_test_data[0]);
        int ret = -1;
        int i;
        for (i = 0; i < num; i++) {
            if (strcmp(tmp, g_test_data[i]) == 0) {
                ret = 0;
                if (i == (num - 1)) {
                    printf("Read Over\n");
                    return 0;
                }
            }
        }
        if (ret < 0) {
            printf("Read Error - %s\n", tmp);
        }
        sleep(0);
    }
    return -1;
}

int main (int argc, char **argv)
{
    if (argc < 2) {
        printf("daemon w/r\n");
        exit(1);
    }
    key_t ipckey = ftok("/tmp/test", ':');
    int shmid = shmget(ipckey, sizeof(struct shm_test_st), 0600 | IPC_CREAT);
    if (shmid < 0) {
        printf("Error - %s\n", strerror(errno));
        exit(1);
    }
    int semid = semget(ipckey, 2, 0666|IPC_CREAT);
    if (semid < 0) {
        printf("Error - %s\n", strerror(errno));
        exit(1);
    }
    void *addr = NULL;
    if (argv[1][0] == 'w') {
        addr = shmat(shmid, NULL, 0);
    } else { // r
        addr = shmat(shmid, NULL, SHM_RDONLY);
    }
    if (addr == (void *)-1) {
        printf("Addr Error - %s\n", strerror(errno));
        exit(1);
    }
    struct shm_test_st *p_shared = (struct shm_test_st *)addr;
    int ret;
    if (argv[1][0] == 'w') {
        ret = shm_write(p_shared, semid);
    } else { // r
        ret = shm_read(p_shared, semid);
    }
    if (shmctl(shmid, IPC_RMID, NULL) < 0) {
        printf("Delete Error - %s\n", strerror(errno));
    }
    if (ret < 0) {
        exit(1);
    }
}
