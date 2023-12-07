#include <iostream>
#include <map>
#include <string>
#include <vector>
#include <regex>

int main() {
	std::string line;

	while (std::getline(std::cin, line)) {
		int i1 = line.find(":");
		int i2 = line.find("|");

		std::regex numRegex("(\\d+)");

		std::sregex_iterator winNumsBegin(line.begin() + i1, line.begin() + i2, numRegex);
		std::sregex_iterator winNumsEnd;

		for (std::sregex_iterator i = winNumsBegin; i != winNumsEnd; ++i) {
			std::smatch m = *i;
			std::string s = m.str();
			std::cout << s << ' ';
		}
		std::cout << '\n';
		
		std::sregex_iterator ourNumsBegin(line.begin() + i2, line.end(), numRegex);
		std::sregex_iterator ourNumsEnd;

		for (std::sregex_iterator i = ourNumsBegin; i != ourNumsEnd; ++i) {
			std::smatch m = *i;
			std::string s = m.str();
			std::cout << s << ' ';
		}
		std::cout << '\n';
	}


	std::map<int, int> a;
}