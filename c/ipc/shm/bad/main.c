/*
 * 完全没有同步机制，出错率很大
 */
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/shm.h>
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

int shm_write(struct shm_test_st *pdata)
{
    int num = sizeof(g_test_data) / sizeof(g_test_data[0]);
    int i;
    for (i = 0; i < num; i++) {
        strcpy(pdata->dat, g_test_data[i]);
        printf("Write - %s\n", g_test_data[i]);
        sleep(1);
    }
    printf("Write Over!\n");
    return 0;
}

int shm_read(struct shm_test_st *pdata)
{
    sleep(1);
    char tmp[SHM_TEST_DATALEN];
    while (1) {
        strcpy(tmp, pdata->dat);
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
        ret = shm_write(p_shared);
    } else { // r
        ret = shm_read(p_shared);
    }
    if (shmctl(shmid, IPC_RMID, NULL) < 0) {
        printf("Delete Error - %s\n", strerror(errno));
    }
    if (ret < 0) {
        exit(1);
    }
}
