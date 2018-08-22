#include "chart.h"
#include <curses.h>
#include <errno.h>
#include <fcntl.h>
#include <float.h>
#include <getopt.h>
#include <limits.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>
#include <unistd.h>

static size_t dsize, dlen, dptr = -1;
static struct ohlc {
	double open;
	double high;
	double low;
	double close;
} *data;

/* grow the data buffer to fill the screen */
static void resize(void)
{
	int width = getmaxx(stdscr);
	size_t newsize = dsize ? dsize : 1;
	while (newsize < (size_t)width)  /* round up to the next binary number */
		newsize <<= 1;
	if (newsize != dsize) {
		struct ohlc *newdata = malloc(newsize * sizeof *newdata);
		size_t len = dptr + 1;
		memcpy(newdata, data, len * sizeof *newdata);
		memcpy(newdata + newsize - dsize + len, data + len,
				(dlen - len) * sizeof *newdata);
		free(data);
		dsize = newsize;
		data = newdata;
	}
}

/* add a new point to the data set */
static void addpoint(double y)
{
	dptr = (dptr + 1) % dsize;
	data[dptr] = (struct ohlc){
		.open = y,
		.high = y,
		.low = y,
		.close = y,
	};
	if (dlen < dsize)
		++dlen;
}

/* update the most recent point already in the data set */
static void updpoint(double y)
{
	if (y > data[dptr].high)
		data[dptr].high = y;
	else if (y < data[dptr].low)
		data[dptr].low = y;
	data[dptr].close = y;
}

#define DRAW_FLAG_RANGE 0x1  /* the '-r <min>,<max>' argument was specified */
#define DRAW_FLAG_TMO_EN 0x2  /* is timeout() enabled? */
#define DRAW_FLAG_HAS_COLOR 0x4  /* does the terminal support color? */
struct draw {
	unsigned int flags;
	const char *title;
	int (*drawpoint)(const struct draw *, WINDOW *, int, int,
			const struct ohlc *);
	int maxy, maxx;
	size_t begin, end;
	double dmin, dmax, drange;
	int margin;
};

/* calculate the screen-y value from a data-point-y value */
static int screeny(const struct draw *drw, int screen_range, double y)
{
	return (drw->dmax - y) / drw->drange * (screen_range - 1) + 1;
}

/* draw a data point using the "dot" style and return its screen y-value */
static int drawpoint_dot(const struct draw *drw, WINDOW *win, int screen_range,
		int x, const struct ohlc *pt)
{
	int y = screeny(drw, screen_range, pt->close);
	mvwaddch(win, y, x, ACS_BULLET | A_BOLD);
	return y;
}

/* draw a data point using the "plus" style and return its screen y-value */
static int drawpoint_plus(const struct draw *drw, WINDOW *win, int screen_range,
		int x, const struct ohlc *pt)
{
	int y = screeny(drw, screen_range, pt->close);
	mvwaddch(win, y, x, '+' | A_BOLD);
	return y;
}

#define OHLC_CH_RIGHT 0x1
#define OHLC_CH_UP 0x2
#define OHLC_CH_DOWN 0x4
#define OHLC_CH_LEFT 0x8
static chtype ohlc_ch[0x10];

/* draw a data point using the "ohlc" style and return its screen y-value */
static int drawpoint_ohlc(const struct draw *drw, WINDOW *win, int screen_range,
		int x, const struct ohlc *pt)
{
	attr_t attrs_orig;
	short colors_orig;
	wattr_get(win, &attrs_orig, &colors_orig, 0);
	if (drw->flags & DRAW_FLAG_HAS_COLOR) {
		if (pt->open < pt->close)
			wattron(win, COLOR_PAIR(COLOR_GREEN));
		else if (pt->open > pt->close)
			wattron(win, COLOR_PAIR(COLOR_RED));
	} else if (pt->open > pt->close)
		wattron(win, A_BOLD);
	int open = screeny(drw, screen_range, pt->open);
	int high = screeny(drw, screen_range, pt->high);
	int low = screeny(drw, screen_range, pt->low);
	int close = screeny(drw, screen_range, pt->close);
	mvwvline(win, high, x, ACS_VLINE, low - high + 1);
	unsigned int ch = OHLC_CH_LEFT;
	if (pt->low < pt->open)
		ch |= OHLC_CH_DOWN;
	if (pt->high > pt->open)
		ch |= OHLC_CH_UP;
	if (open == close) {
		ch |= OHLC_CH_RIGHT;
	} else {
		mvwaddch(win, open, x, ohlc_ch[ch]);
		ch = OHLC_CH_RIGHT;
		if (pt->low < pt->close)
			ch |= OHLC_CH_DOWN;
		if (pt->high > pt->close)
			ch |= OHLC_CH_UP;
	}
	mvwaddch(win, close, x, ohlc_ch[ch]);
	wattr_set(win, attrs_orig, colors_orig, 0);
	return close;
}

/* draw it! */
static void drawchart(struct draw *drw)
{
	/* find some metrics to make everything fit in the terminal nicely */
	getmaxyx(stdscr, drw->maxy, drw->maxx);
	drw->end = (dptr + 1) % dsize;

	/* find the domain */
	size_t maxx = drw->maxx - 2;  /* -2 to account for the border */
	drw->begin = (drw->end - (dlen < maxx ? dlen : maxx)) % dsize;

	/* find the range */
	if (!(drw->flags & DRAW_FLAG_RANGE)) {
		drw->dmin = DBL_MAX;
		drw->dmax = -DBL_MAX;
		for (size_t i = drw->begin; i != drw->end; i = (i + 1) % dsize) {
			if (data[i].low < drw->dmin)
				drw->dmin = data[i].low;
			if (data[i].high > drw->dmax)
				drw->dmax = data[i].high;
		}
	}
	drw->drange = drw->dmax - drw->dmin;

	/* find the maximum width of the margin for the scale on the right */
	drw->margin = snprintf(0, 0, "%lg", drw->dmax);
	int n = snprintf(0, 0, "%lg", drw->dmin);
	if (n > drw->margin)
		drw->margin = n;
	double dlast = dlen ? data[dptr].close : 0.0;
	n = snprintf(0, 0, "%lg", dlast);
	if (n > drw->margin)
		drw->margin = n;

	/* adjust the domain with the newly-found chart width (still
	 * accounting for the border) */
	maxx = drw->maxx - drw->margin - 2;
	drw->begin = (drw->end - (dlen < maxx ? dlen : maxx)) % dsize;

	/* draw the chart */
	WINDOW *win = newwin(drw->maxy, drw->maxx - drw->margin, 0, 0);
	int y = 0, x = 1, maxy = drw->maxy - 2;  /* -2 for the border */
	for (size_t i = drw->begin; i != drw->end; i = (i + 1) % dsize)
		y = drw->drawpoint(drw, win, maxy, x++, data + i);
	x = getmaxx(win) - 1;  /* fix x to the right side of the chart */
	box(win, 0, 0);
	if (drw->title) {
		wattron(win, A_BOLD | A_REVERSE);
		mvwprintw(win, 0, 4, " %s ", drw->title);
		wattroff(win, A_BOLD | A_REVERSE);
	}
	mvwaddch(win, 1, x, ACS_RTEE);
	mvwaddch(win, maxy, x, ACS_RTEE);
	mvwaddch(win, y, x, ACS_RTEE);
	wrefresh(win);
	delwin(win);

	/* draw the scale */
	win = newwin(drw->maxy, drw->margin, 0, drw->maxx - drw->margin);
	mvwprintw(win, 1, 0, "%lg", drw->dmax);
	mvwprintw(win, maxy, 0, "%lg", drw->dmin);
	mvwprintw(win, y, 0, "%lg", dlast);
	wclrtoeol(win);
	wrefresh(win);
	delwin(win);
}

void execute(float points[])
{

	// for(int x=0; x<var2; x++)
	//  {
	// 		 printf("Value of var_arr[%d] is: %d \n", x, *var1);
	// 		 /*increment pointer for next element fetch*/
	// 		 var1++;
	//  }


	// struct draw drw = {
	// 	.drawpoint = drawpoint_dot,
	// };
	//
  // drw.drawpoint = drawpoint_ohlc;
  // drw.title = "Title";

  //
	// /* if a filename was specified, open it and dup it to STDIN_FILENO */
	// int fd_stdin = STDIN_FILENO;
	// if (optind < argc && strcmp(argv[optind], "-")) {
	// 	const char *fn = argv[optind];
	// 	if (!drw.title)
	// 		drw.title = fn;
	// 	int fd = open(fn, O_RDONLY);
	// 	if (fd < 0) {
	// 		fprintf(stderr, "%s: %s: %s\n", argv[0], fn,
	// 				strerror(errno));
	// 		return 1;
	// 	}
	// 	fd_stdin = dup(STDIN_FILENO);
	// 	dup2(fd, STDIN_FILENO);
	// 	close(fd);
	// }

	/* initialize curses */
// 	if (!initscr())
// 		return;
// 	if (has_colors()) {
// 		start_color();
// #ifdef NCURSES_VERSION
// 		use_default_colors();
// 		init_pair(COLOR_RED, COLOR_RED, -1);
// 		init_pair(COLOR_GREEN, COLOR_GREEN, -1);
// #else
// 		init_pair(COLOR_RED, COLOR_RED, COLOR_BLACK);
// 		init_pair(COLOR_GREEN, COLOR_GREEN, COLOR_BLACK);
// #endif
// 		drw.flags |= DRAW_FLAG_HAS_COLOR;
// 	}
// 	curs_set(0);
// 	noecho();
// 	resize();
//
// 	/* bits: [left][down][up][right] */
// 	ohlc_ch[0x0] = '0';
// 	ohlc_ch[0x1] = '1';
// 	ohlc_ch[0x2] = '2';
// 	ohlc_ch[0x3] = ACS_LLCORNER;
// 	ohlc_ch[0x4] = '4';
// 	ohlc_ch[0x5] = ACS_ULCORNER;
// 	ohlc_ch[0x6] = ACS_VLINE;
// 	ohlc_ch[0x7] = ACS_LTEE;
// 	ohlc_ch[0x8] = '8';
// 	ohlc_ch[0x9] = ACS_HLINE;
// 	ohlc_ch[0xa] = ACS_LRCORNER;
// 	ohlc_ch[0xb] = ACS_BTEE;
// 	ohlc_ch[0xc] = ACS_URCORNER;
// 	ohlc_ch[0xd] = ACS_TTEE;
// 	ohlc_ch[0xe] = ACS_RTEE;
// 	ohlc_ch[0xf] = ACS_PLUS;
//
// 	/* input loop */
// 	char buf[BUFSIZ], *ptr = buf;
// 	size_t n = sizeof buf;
// 	while (1) {
// 		for (int status; status = getnstr(ptr, n), status != ERR;) {
// 			if (status == KEY_RESIZE) {
// 				endwin();
// 				resize();
// 				drw.flags |= DRAW_FLAG_TMO_EN;
// 				break;  /* redraw immediately */
// 			}
// 			if (*buf != ',') {
// 				double y = strtod(buf, &ptr);
// 				if (ptr == buf)
// 					continue;
// 				addpoint(y);
// 			} else {
// 				ptr = buf;
// 			}
// 			while (*ptr == ',') {
// 				char *end;
// 				double y = strtod(++ptr, &end);
// 				if (end == ptr)
// 					break;
// 				updpoint(y);
// 				ptr = end;
// 			}
// 			ptr = buf;
// 			n = sizeof buf;
// 			/* aggregate input lines that are within 20ms of each
// 			 * other to avoid flooding the terminal with updates
// 			 * that are coming too rapidly to be seen */
// 			timeout(20);
// 			drw.flags |= DRAW_FLAG_TMO_EN;
// 		}
// 		if (!(drw.flags & DRAW_FLAG_TMO_EN))
// 			/* ERR must have been the result of EOF, not timeout */
// 			break;
// 		drawchart(&drw);
// 		refresh();
// 		size_t len = strlen(ptr);
// 		ptr += len;
// 		n -= len;
// 		timeout(-1);
// 		drw.flags &= ~DRAW_FLAG_TMO_EN;
// 	}
//
// 	/* if the input was from a file (i.e., this wasn't a live stream), wait
// 	 * for a keypress before exiting so that the user has a chance to see
// 	 * the nice chart I made for him or her */
// 	if (fd_stdin != STDIN_FILENO) {
// 		dup2(fd_stdin, STDIN_FILENO);
// 		char *title = malloc(snprintf(0, 0, "%s (closed)", drw.title));
// 		sprintf(title, "%s (closed)", drw.title);
// 		drw.title = title;
// 		while (1) {
// 			drawchart(&drw);
// 			refresh();
// 			if (getch() != KEY_RESIZE)
// 				break;
// 			endwin();
// 		}
// 	}

	/* shutdown */
	endwin();
}
