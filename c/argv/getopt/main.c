#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h> /* getopt 时必须包含 unistd 头文件 */

int main (int argc, char **argv)
{
	int index;
	int c;
	opterr = 0;
	/* getopt 用来解析短选项，第一二个参数基本都是相同的 */
	/* 第三个选项传入能接受的参数,':'代表需要值 '::'代表值为可选*/
	/* 如果传入了值，其值保存在 optarg 当中，当其值为必须是，若没有传入值则会报错缺少值 */
	/* getopt 返回匹配的选项 &nbsp; 例如 a/b/c */
	/* 未匹配值时返回？ 解析完时，会返回 -1 跳出 */
	while ((c = getopt (argc, argv, "abc:")) != -1)
		switch (c) {
		case 'a':
			printf("检测到-a选项\n");
			break;
		case 'b':
			printf("检测到-b选项\n");
			break;
		case 'c':
			printf("检测到-c选项，其参数为：%s\n", optarg);
			break;
		/* 用来检测其他意外的值 */
		case '?':
			 /* optopt 为在没有匹配任何选项时，保存的选项值 */
			 /* 例如 -n，没有与上面匹配，则此时 optopt=c */
			if ( optopt == 'c')
				fprintf (stderr, "选项 -%c 需要一个参数.\n", optopt);
			else if ( isprint (optopt ))
				fprintf (stderr, "`-%c'未定义\n", optopt);
			else
				fprintf (stderr, " `\\x%x' 未定义\n", optopt);
			break;
		default:
			abort(); //跳出while循环
		}
	/**
	 * 最后检测命令行给的选项是否已经解析完
	 * optind 是当前解析到选项的索引
	 * 判断依据是如果解析的个数小于 argc，那么就是没有解析完
	 */
	for (index = optind; index < argc; index++)
		printf ("未定义的参数： %s\n", argv[index]);
	return 0;
}