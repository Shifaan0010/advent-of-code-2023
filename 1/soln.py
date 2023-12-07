import sys
import re


def extract_digits(string):
    return [int(ch) for ch in string if ch.isdigit()]


def extract_digit_strings(string):
    digit_strings = ['one', 'two', 'three', 'four',
                     'five', 'six', 'seven', 'eight', 'nine']

    regex = r'(?=(\d|' + '|'.join(digit_strings) + r'))'

    # print(regex)

    return [int(m.group(1)) if m.group(1).isdigit() else digit_strings.index(
        m.group(1))+1 for m in re.finditer(regex, string)]


def cat_first_last(digits):
    if digits == None or len(digits) == 0:
        return 0

    return 10 * digits[0] + digits[-1]


lines = [line.strip() for line in sys.stdin]

print('Part 1')
print(sum(cat_first_last(extract_digits(l)) for l in lines))

print('Part 2')
print(sum(cat_first_last(extract_digit_strings(l)) for l in lines))
