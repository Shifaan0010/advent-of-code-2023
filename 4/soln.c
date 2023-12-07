#include <stdio.h>

typedef struct card {
	int num;
	int *win_nums;
	int *nums;
} card_t;

card_t parse_card(FILE *file) {
	card_t card;

	char ch;
	while ((ch = fgetc(file)) != '\n') {
		if (ch == ' ') {

		}
	}

	return card;
}

int main() {
	
}