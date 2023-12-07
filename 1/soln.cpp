#include <iostream>
#include <string>
#include <vector>
#include <array>
#include <regex>

std::vector<int> extractDigits(const std::string &str) {
	std::vector<int> digits;

	for (char ch : str) {
		if (ch >= '0' && ch <= '9') {
			digits.push_back(ch - '0');
		}
	}

	return digits;
}

std::vector<int> extractDigitStrings(const std::string &str) {
	std::array<std::string, 10> digitStrings = { "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" };

	std::regex digitRegex("(?=(\\d|one|two|three|four|five|six|seven|eight|nine)).");

	std::vector<int> digits;

	std::sregex_iterator next(str.begin(), str.end(), digitRegex);
	std::sregex_iterator end;

	while (next != end) {
		std::smatch match = *next;
		// std::cout << match.str(1) << "\t" << "\t" << match.position() << "\t" << "\n";

		std::string matchStr = match.str(1);

		if ('0' <= matchStr[0] && matchStr[0] <= '9') {
			digits.push_back(matchStr[0] - '0');
		} else {
			const std::string *s = std::find(digitStrings.begin(), digitStrings.end(), matchStr);
			int n = std::distance(digitStrings.cbegin(), s);

			digits.push_back(n);
		}

		// std::cout << digits.back() << '\n';

		next++;
	}

	// std::cout << '\n';

	return digits;
}

int main() {
	std::string line;
	std::vector<std::string> lines;

	while (std::getline(std::cin, line)) {
		lines.push_back(line);
	}

	std::cout << "Part 1\n";
	std::cout << "\n";

	int sum1 = 0;
	for (const std::string &line : lines) {
		std::vector<int> digits = extractDigits(line);

		if (digits.size() > 0) {
			int value = 10 * digits.front() + digits.back();

			sum1 += value;
		}
	}

	std::cout << sum1 << std::endl;

	std::cout << "\n";
	std::cout << "Part 2\n";
	std::cout << "\n";

	int sum2 = 0;
	for (const std::string &line : lines) {
		std::vector<int> digits = extractDigitStrings(line);

		if (digits.size() > 0) {
			int value = 10 * digits[0] + digits.back();

			sum2 += value;
		}
	}

	std::cout << sum2 << std::endl;
}
