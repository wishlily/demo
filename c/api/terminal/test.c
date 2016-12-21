#include <stdio.h>
#include <stdlib.h>
#include <mcheck.h>
#include <stdbool.h>

#include "term.h"

#ifdef DEBUG
#define dbg_printf(fmt, args...)  fprintf(stderr, fmt, ##args)
#else
#define dbg_printf(fmt, args...)
#endif
#define info_printf(fmt, args...) fprintf(stderr, "[INFO]: "fmt, ##args)
#define err_printf(fmt, args...)  fprintf(stderr, "[ERROR]%s<%d>: "fmt, __FILE__, __LINE__, ##args)

#define TEST_ARRAY_NUM 13
char *test_array[TEST_ARRAY_NUM] = {
	{"Devouring Time blunt thou the lion's paws,"},
	{"And make the earth devour her own sweet brood,"},
	{"Pluck the keen teeth from the fierce tiger's jaws,\n"},
	{"And burn the long-lived phoenix, in her blood,\nMake glad and sorry seasons as thou fleet'st,"},
	{"And do whate'er thou wilt swift-footed Time"},
	{"If you shed tears when you miss the sun, you also miss the stars."},
	{"To the wide world and all her fading sweets:"},
	{"But I forbid thee one most heinous crime,"},
	{"O carve not with thy hours my love's fair brow,"},
	{"Him in thy course untainted do allow,"},
	{"For beauty's pattern to succeeding men."},
	{"Yet do thy worst old Time: despite thy wrong,"},
	{"My love shall in my verse ever live young."}
};
char *test_array2[TEST_ARRAY_NUM] = {
	{"Devouring Time blunt thou the lion's paws,"},
	{"$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$"},
	{"Pluck the keen teeth from the fierce tiger's jaws,\n"},
	{"And burn the long-lived phoenix, in her blood,\nMake glad and sorry seasons as thou fleet'st,"},
	{"And do whate'er thou wilt swift-footed Time"},
	{"If you shed tears when you miss the sun, you also miss the stars."},
	{"To the wide world and all her fading sweets:"},
	{"But I forbid thee one most heinous crime,"},
	{"O carve not with thy hours my love's fair brow,"},
	{"&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&\n**************************"},
	{"For beauty's pattern to succeeding men."},
	{"Yet do thy worst old Time: despite thy wrong,"},
	{"My love shall in my verse ever live young."}
};

int test_text(term_t *tp)
{
	int ret = term_text(tp, test_array[0]);
	if (ret < 0) {
		return ret;
	}
	dbg_printf("%d, %d\n", tp->start.x, tp->start.y);
	int i;
	term_text_t *lp = tp->text;
	for (i = 0; i < tp->lines; i++) {
		if (lp) {
			dbg_printf("[%d] %s\n", i+1,lp->data);
			lp = lp->next;
		}
	}
	term_display(tp);
	return 0;
}

int test_text_add(term_t *tp)
{
	int ret = term_text(tp, test_array[0]);
	if (ret < 0) {
		return ret;
	}
	int i;
	for (i = 1; i < TEST_ARRAY_NUM; i++) {
		if (term_text_add(tp, test_array[i]) < 0) {
			return -1;
		}
	}
	term_display(tp);
	term_page_next(tp);
	term_display(tp);
	term_page_back(tp);
	term_display(tp);
	return 0;
}

int test_getline(term_t *tp)
{
	char cmd[5];
	while (1) {
		int len = term_getline(tp, cmd, 5);
		info_printf("[%d,%d]%s\n", len, strlen(cmd), cmd);
		if (!strcmp(cmd, "exit")) {
			break;
		}
	}
	return 0;
}

bool running = true;

void *displayThreadRoutine(void *args)
{
	term_t *tp = (term_t *)args;
	while(running) {
		term_display(tp);
		usleep(200);
	}
}

int main()
{
	/* $ export MALLOC_TRACE=malloc.log */
	mtrace();

	term_loc_t start = {12, 4};
	term_t *tp = term_init(start, 10, 20);
	if (tp == NULL) {
		err_printf("term init failed.\n");
		exit(1);
	}

	int test_number = 0;

#if 0
	term_clear();
	term_goto(5, 5);
	term_printf("Hello World!\n");
	term_goto(5, 5);
	term_printf_n(12, "%s", "Hello!");
	test_number++;
#endif

#if 0
	term_clear();
	if (test_text(tp) < 0) {
		goto fail;
	}
	test_number++;
#endif

#if 0
	term_clear();
	if (test_text_add(tp) < 0) {
		goto fail;
	}
	test_number++;
#endif

#if 0
	if (test_getline(tp) < 0) {
		goto fail;
	}
	test_number++;
#endif

#if 1
	term_clear();
	int ret = term_text(tp, test_array[0]);
	if (ret < 0) {
		goto fail;
	}
	int i;
	for (i = 1; i < TEST_ARRAY_NUM; i++) {
		if (term_text_add(tp, test_array[i]) < 0) {
			goto fail;
		}
	}
	pthread_t display_thread;
	pthread_create(&display_thread, NULL, displayThreadRoutine, tp);
	char cmd[32];
	while (1) {
		/* 直接回车不改变 buff 数据，所以回车默认执行上次的命令 */
		int len = term_getline(tp, cmd, 10);
		if (!strcmp(cmd, "exit")) {
			running = false;
			break;
		} else if (!strcmp(cmd, "back") || !strcmp(cmd, "b")) {
			term_page_back(tp);
		} else if (!strcmp(cmd, "next") || !strcmp(cmd, "n")) {
			term_page_next(tp);
		} else if (!strcmp(cmd, "2")) {
			int ret = term_update(tp, test_array2[0]);
			if (ret < 0) {
				goto fail;
			}
			int i;
			for (i = 1; i < TEST_ARRAY_NUM; i++) {
				if (term_text_add(tp, test_array2[i]) < 0) {
					goto fail;
				}
			}
		} else if (!strcmp(cmd, "1")) {
			int ret = term_update(tp, test_array[0]);
			if (ret < 0) {
				goto fail;
			}
			int i;
			for (i = 1; i < TEST_ARRAY_NUM; i++) {
				if (term_text_add(tp, test_array[i]) < 0) {
					goto fail;
				}
			}
		}
	}
	pthread_join(display_thread, NULL);
#endif

	printf("\n");
	info_printf("test successful.\n");
	term_free(tp);
	exit(0);
fail:
	err_printf("\nterm test[%d] failed.\n", test_number);
	term_free(tp);
	exit(1);
}