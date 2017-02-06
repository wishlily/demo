#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <termios.h>
#include <unistd.h>

#include "term.h"

term_t *term_init(term_loc_t start, uint16_t row, uint16_t col)
{
	term_t *handler = (term_t *)calloc(1, sizeof(term_t));
	if (handler == NULL) {
		return NULL;
	}
	handler->start.x = start.x;
	handler->start.y = start.y;
	handler->limit.x = col;
	handler->limit.y = row - 2;
	return handler;
}

static void term_text_free(term_t *handler)
{
	term_t *fp = handler;
	term_text_t *ptr = fp->text;
	while (ptr) {
		if (ptr->data) {
			free(ptr->data);
		}
		term_text_t *ptr_next = ptr->next;
		free(ptr);
		ptr = ptr_next;
	}
	fp->text = NULL;
	fp->last_line = NULL;
	fp->lines = 0;
	fp->page = 0;
}

static term_text_t *term_text_new_line(char *text, uint16_t len)
{
	term_text_t *ptr = (term_text_t *)calloc(1, sizeof(term_text_t));
	if (ptr == NULL) {
		return NULL;
	}
	// NOTE: len + 1 because '\0' need 1 size space
	ptr->data = (char *)calloc(len+1, sizeof(char));
	if (ptr->data == NULL) {
		free(ptr);
		return NULL;
	}
	memcpy(ptr->data, text, len);
	return ptr;
}

static int term_text_add_line(term_t *handler, char *text, uint16_t len)
{
	term_t *fp = handler;
	term_text_t *line = term_text_new_line(text, len);
	if (line == NULL) {
		return -1;
	}
	if (fp->last_line) {
		fp->last_line->next = line;
	} else {
		fp->text = line;
	}
	fp->last_line = line;
	fp->lines++;
	return 0;
}

static int term_text_add_lines(term_t *handler, char *text)
{
	if (strlen(text) <= 0) {
		return 0;
	}
	term_t *fp = handler;
	int limit = fp->limit.x;
	char *line = text;
	while (line < text + strlen(text)) {
		int len = strlen(line);
		if (len > limit) {
			len = limit;
		}
		if (term_text_add_line(fp, line, len) < 0) {
			return -1;
		}
		line += len;
	}
	return 0;
}

int term_text_add(term_t *handler, char *text)
{
	int len = strlen(text);
	char *line = strdup(text);
	if (line == NULL) {
		return -1;
	}
	char *line_head = line;
	while (line < line_head + len) {
		char *ptr = strchr(line, '\n');
		if (ptr) { // have CR
			*ptr = '\0';
			if (term_text_add_lines(handler, line) < 0) {
				free(line_head);
				return -1;
			}
			*ptr = '\n';
			line = ptr + 1;
			continue;
		}
		// the end
		if (term_text_add_lines(handler, line) < 0) {
			free(line_head);
			return -1;
		}
		line += strlen(line);
	}
	free(line_head);
	return 0;
}

int term_text(term_t *handler, char *text)
{
	term_text_free(handler);
	return term_text_add(handler, text);
}

int term_update(term_t *handler, char *text)
{
	uint16_t page = handler->page;
	term_text_free(handler);
	handler->page = page;
	return term_text_add(handler, text);
}

#define HELP_PAGE_BACK "back"
#define HELP_PAGE_NEXT "next"

static void term_display_help(term_t *handler)
{
	term_t *fp = handler;
	term_loc_t help_loc = {fp->start.x, fp->start.y + fp->limit.y};
	term_loc_t cmd_loc = {fp->start.x, fp->start.y + fp->limit.y + 1};

	if (fp->lines > fp->limit.y) { // more page
		term_goto(help_loc.y, help_loc.x);
		term_color(TERM_GRAY, TERM_BLACK);
		term_printf_n(strlen(HELP_PAGE_BACK), "%s", HELP_PAGE_BACK);
		term_goto(help_loc.y, help_loc.x + fp->limit.x - (int)strlen(HELP_PAGE_NEXT));
		term_printf_n(strlen(HELP_PAGE_NEXT), "%s", HELP_PAGE_NEXT);
		term_reset();
	}
	// clear cmd & set cursor
	term_goto(cmd_loc.y, cmd_loc.x);
	term_printf_n(fp->limit.x, "$ ");
	if (strlen(fp->inputs)) {
		term_printf_n(fp->limit.x-2, "%s", fp->inputs);
	}
}

static void term_display_page(term_t *handler, uint16_t page)
{
	term_t *fp = handler;
	int i;
	term_text_t *data = fp->text;
	// page
	int lines = page * fp->limit.y;
	for (i = 0; i < lines; i++) {
		if (data && data->data) {
			data = data->next;
		} else {
			break;
		}
	}
	// print
	for (i = 0; i < fp->limit.y; i++) {
		term_goto(i + fp->start.y, fp->start.x);
		if (data && data->data) {
			term_printf_n(fp->limit.x, "%s", data->data);
			data = data->next;
		} else {
			term_printf_n(fp->limit.x, ""); // null line
		}
	}
	term_display_help(fp);
	fp->page = page;
}

void term_display(term_t *handler)
{
	term_display_page(handler, handler->page);
	fflush(stdout);
}

void term_page_back(term_t *handler)
{
	term_t *fp = handler;
	if (fp->page > 0) {
		fp->page--;
	}
}

void term_page_next(term_t *handler)
{
	term_t *fp = handler;
	uint16_t page_number = fp->lines / fp->limit.y;
	if (fp->lines % fp->limit.y) {
		page_number++;
	}
	if (fp->page < page_number - 1) {
		fp->page++; // from zero
	}
}

/* reads from keypress, doesn't echo */
static int getch(void)
{
	struct termios oldattr, newattr;
	int ch;
	tcgetattr( STDIN_FILENO, &oldattr );
	newattr = oldattr;
	newattr.c_lflag &= ~( ICANON | ECHO );
	tcsetattr( STDIN_FILENO, TCSANOW, &newattr );
	ch = getchar();
	tcsetattr( STDIN_FILENO, TCSANOW, &oldattr );
	return ch;
}

int term_getline(term_t *handler, char *buff, uint32_t buff_size)
{
	term_t *fp = handler;
	int len = 0;
	while (1) {
		int c = getch();
		switch (c) {
		case '\b':
		case 127:
			if (len > 0) {
				len--;
				fp->inputs[len] = '\0';
			}
			break;
		case '\n':
			/* only enter not change */
			if (len < TERM_INPUT_BUFF_SIZE) {
				fp->inputs[len] = '\0';
			}
			fp->inputs[TERM_INPUT_BUFF_SIZE-1] = '\0';
			if (len) {
				strncpy(buff, fp->inputs, buff_size);
				buff[buff_size-1] = '\0';
				memset(fp->inputs, 0, TERM_INPUT_BUFF_SIZE);
			}
			return len;
		default:
			if (len < TERM_INPUT_BUFF_SIZE) {
				fp->inputs[len++] = c;
			}
		}
	}
	return len;
}

void term_free(term_t *handler)
{
	term_text_free(handler);
	free(handler);
}
