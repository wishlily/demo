#include <stdio.h>
#include <stdlib.h>
#include "parse.h"

// 存放结果
static char parse_data[64] = {0};

static int get_string(FILE *fc, char **pline)
{
    size_t len = 0;
    int read = 0;
    if (*pline) {
        printf("%s pline is not NULL\n", __func__);
        return -1;
    }
    read = getline(pline, &len, fc);
    //printf("getline %d %d\n", (int)read, (int)len);
    if (read <= 0) {
        return -1;
    }
    return read;
}

static int get_int(FILE *fc)
{
    char *pline = NULL;
    int data = 0;
    if (get_string(fc, &pline) < 0) {
        return 0;
    }

    //printf("%s: 0x%02X\n", pline, pline[0]);
    if (sscanf(pline, "%d", &data) != 1) {
        printf("sscanf fail\n");
        return 0;
    }

    //printf("data: %d\n", data);
    if (pline) {
        free(pline);
    }
    return data;
}

char *parse(char *filename)
{
    FILE *fc;
    int year,month,date,hh,mm,ss;

    printf("Parsing file %s\n", filename);
    fc = fopen(filename, "r");
    if (fc == NULL) {
        printf("Could not open file %s\n", filename);
        return parse_data;
    }

    year = get_int(fc);
    month = get_int(fc);
    date = get_int(fc);
    hh = get_int(fc);
    mm = get_int(fc);
    ss = get_int(fc);

    sprintf(parse_data,
        "%04d-%02d-%02d, %02d:%02d:%02d\n",
        year, month, date,
        hh, mm, ss);

    fclose(fc);
    return parse_data;
}

