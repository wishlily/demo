vpath %.h include #vpath 指定搜索路径
vpath %.c add
vpath %.c unit_test
objects = test_sub.o add.o test_add.o
test: test_main.c $(objects)
	gcc -I include $^ -o test -lcheck
all: $(objects)
$(objects): %.o: %.c
	gcc -c -I include -I sub $< -o $@
.PHONY: clean
clean:
	rm *.o test