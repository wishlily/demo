#include "check.h"
#include "unit_test.h"
#include "add.h"
START_TEST(test_add) {
    fail_unless(add(6, 2) == 8, "error, 6 + 2 != 8");
}
END_TEST

START_TEST(test_add_double) {
    fail_unless(add_double(6.6, 2.2) == 8.8, "error, 6.6 + 2.2 != 8.8");
}
END_TEST

Suite * make_add_suite(void) {
    Suite *s = suite_create("add");       // 建立Suite
    TCase *tc_add = tcase_create("add");  // 建立测试用例集
    suite_add_tcase(s, tc_add);           // 将测试用例加到Suite中
    tcase_add_test(tc_add, test_add);     // 测试用例加到测试集中
    tcase_add_test(tc_add, test_add_double);
    return s;
}