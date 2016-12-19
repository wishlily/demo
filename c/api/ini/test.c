# include <stdio.h>
# include <string.h>
# include <error.h>
# include <errno.h>
# include <inttypes.h>
# include <arpa/inet.h>
# include <stdlib.h>
# include <mcheck.h>

# include "ini.h"

int main()
{
	/* $ export MALLOC_TRACE=malloc.log */
	mtrace();

	ini_t *conf = ini_load("test.ini");
	if (conf == NULL)
		error(1, errno, "ini_load fail");

	int test_number = 0;
	do {
		uint8_t want[5] = {23,34,5,7,0};
		uint8_t a[10];
		int ret = ini_read_uint8s(conf, "mytest", "array", a, 10, NULL);
		if (ret < 0) {
			printf("ERR test%d: %d\n", test_number, ret);
			break;
		}
		if (memcmp(want, a, sizeof(want)) || ret != sizeof(want)) {
			printf("ERR test%d: [%d]", test_number, ret);
			int i;
			for (i = 0; i < ret; i++) {
				printf("%d,", a[i]);
			}
			printf("\n");
			break;
		}
	} while(0);

	test_number++;
	do {
		uint8_t want[5] = {1,2,3,4,5};
		uint8_t a[5];
		int ret = ini_read_uint8s(conf, "mytest", "array_def", a, 5, want);
		if (ret < 0) {
			printf("ERR test%d: %d\n", test_number, ret);
			break;
		}
		if (memcmp(want, a, sizeof(want)) || ret != sizeof(want)) {
			printf("ERR test%d: [%d]", test_number, ret);
			int i;
			for (i = 0; i < ret; i++) {
				printf("%d,", a[i]);
			}
			printf("\n");
			break;
		}
	} while(0);

	ini_free(conf);
	return 0;
}

