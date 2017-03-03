#include "unit_test.h"
#include <stdlib.h>
#include <stdio.h>

/**
 * yum install check
 * yum install check-devel
 */

int main(void) {
    int n;

    {
        printf (">> Test Sub\n");
        SRunner *sr = srunner_create(make_sub_suite()); // 将Suite加入到SRunner
        srunner_run_all(sr, CK_NORMAL);
        n += srunner_ntests_failed(sr);         // 运行所有测试用例
        srunner_free(sr);
    }

    {
        printf (">> Test Add\n");
        SRunner *sr = srunner_create(make_add_suite()); // 将Suite加入到SRunner
        srunner_run_all(sr, CK_NORMAL);
        n += srunner_ntests_failed(sr);         // 运行所有测试用例
        srunner_free(sr);
    }

    return (n == 0) ? EXIT_SUCCESS : EXIT_FAILURE;
}