#include <stdio.h>
#include <stdlib.h>
/* 使用getopt_long时不需要unistd*/
#include <getopt.h>
/* 用来作为一个标志的值. */
static int verbose_flag;

int main (int argc, char **argv)
{
	int c;
	while (1) {
		static struct option long_options[] = {
			/* 如果指定了 --verbose 那么 verbose_flag=1 */
			{"verbose", no_argument,       &verbose_flag, 1},
			/* 如果指定了--brief 那么 verbose_flag=0 */
			{"brief",   no_argument,       &verbose_flag, 0},

			/* --add = -a 不带参数*/
			{"add",     no_argument,       0, 'a'},
			{"append",  no_argument,       0, 'b'},

			/* --delete=XX = -dXX 带参数*/
			{"delete",  required_argument, 0, 'd'},
			{"create",  required_argument, 0, 'c'},
			{"file",    required_argument, 0, 'f'},

			/* 很明显 5438 不可能用一个 ANSCI 表示出来，那么这个可以直接用 case 就可以*/
			{"a_very_long_arg",    no_argument, 0, 5438},

			/* 最后为空值，来标识结束 */
			{0, 0, 0, 0}
		};

		/* 因为没有了optind，那么就要用一个值来存储当前索引 */
		int option_index = 0;

		/* 前三个参数与 getopt 相同，第四个为 option，第五个为存索引 int 的指针*/
		c = getopt_long (argc, argv, "abc:d:f:", long_options, &option_index);

		/* 当解析完返回 -1 跳出循环 */
		if (c == -1)
			break;

		switch (c) {
		case 0:
			/* 当 flag 不为 0 时，遇到那些选项 getopt_long 返回 0*/
			if (long_options[option_index].flag != 0)
				break;
			printf ("选项 %s", long_options[option_index].name);
			if (optarg)
				printf (" 值 %s", optarg);

			printf ("\n");
			break;
		case 'a':
			puts ("选项 -a");
			break;
		case 'b':
			puts ("选项 -b");
			break;
		//长选项
		case 5438:
			puts ("选项 --a_very_long_arg");
			break;
		case 'c':
			printf ("选项 -c 值为 `%s'\n", optarg);
			break;
		case 'd':
			printf ("选项 -d 值为 `%s'\n", optarg);
			break;
		case 'f':
			printf ("选项 -f 值为 `%s'\n", optarg);
			break;
		case '?':
			/* getopt_long 会自动输出错误，不像 getopt 需要自己检测 */
			break;
		default:
			abort ();
		}
	}
	/* 检测是否设置了标志位 */
	if (verbose_flag)
		puts ("verbose 标志已被设置");

	/* 输入未解析完的选项. */
	if (optind < argc) {
		printf ("无效选项： ");
		while (optind < argc)
			printf ("%s ", argv[optind++]);
		putchar ('\n');
	}
	exit (0);
}