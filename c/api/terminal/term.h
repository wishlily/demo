#ifndef TERM_H
#define TERM_H

#include <stdint.h>

#define TERM_INPUT_BUFF_SIZE 64

/* color */
#define TERM_BLACK  0
#define TERM_RED    1
#define TERM_GREEN  2
#define TERM_YELLOW 3
#define TERM_BLUE   4
#define TERM_PURPLE 5
#define TERM_CYAN   6
#define TERM_GRAY   7

#define TERM_DARK_GRAY    60
#define TERM_LIGHT_RED    61
#define TERM_LIGHT_GREEN  62
#define TERM_LIGHT_YELLOW 63
#define TERM_LIGHT_BLUE   64
#define TERM_LIGHT_PURPLE 65
#define TERM_LIGHT_CYAN   66
#define TERM_WHITE        67

/* attributes */
#define TERM_BRIGHT    1
#define TERM_DIM       2
#define TERM_UNDERLINE 4
#define TERM_BLINK     5
#define TERM_REVERSE   7
#define TERM_HIDDEN    8

#define term_clear()                     printf("\033[2J");
/* auto clear the line */
#define term_printf(fmt, args...)        printf("\033[K"fmt, ##args)
#define term_printf_n(len, fmt, args...) printf("\033[%dX"fmt, len, ##args)
#define term_goto(row, col)              printf("\033[%d;%df", row, col)
#define term_up(rows)                    printf("\033[%dA", rows)
#define term_down(rows)                  printf("\033[%dB", rows)
#define term_left(cols)                  printf("\033[%dD", cols)
#define term_right(cols)                 printf("\033[%dC", cols)
#define term_color(bg, fg)               printf("\033[%d;%dm", (bg)+40, (fg)+30)
#define term_set(attr)                   printf("\033[%dm", attr)
#define term_reset()                     printf("\033[0m")

typedef struct term_loc {
	uint16_t x; // col
	uint16_t y; // row
} term_loc_t;

typedef struct term_text {
	char *data;
	struct term_text *next;
} term_text_t;

typedef struct term {
	term_loc_t start;
	term_loc_t limit;
	term_text_t *text;
	term_text_t *last_line;
	uint16_t lines;
	uint16_t page;
	char inputs[TERM_INPUT_BUFF_SIZE];
} term_t;

/**
 *  0----------------x
 *  |................
 *  |................
 *  |...start----col  (from start count)
 *  |..|---------|...
 *  |..|---------|...  (--- is term workspace)
 *  |..|---------|...
 *  |..row-------|... (from start count)
 *  y................
 */
term_t *term_init(term_loc_t start, uint16_t row, uint16_t col);
void term_free(term_t *handler);
/* clear old text & add new text */
int term_text(term_t *handler, char *text);
/* clear old text but save page info */
int term_update(term_t *handler, char *text);
/* add text */
int term_text_add(term_t *handler, char *text);
/* display text */
void term_display(term_t *handler);
void term_page_back(term_t *handler);
void term_page_next(term_t *handler);
/* get line from i/o, only enter not change buff data */
int term_getline(term_t *handler, char *buff, uint32_t buff_size);

#endif