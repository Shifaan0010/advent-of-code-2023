from collections import defaultdict
import sys
import re


def issymbol(ch):
    return not ch.isdigit() and not ch == '.'


def neighbors(lines, i, m):
    for y in range(i-1, i+2):
        for x in range(m.span()[0]-1, m.span()[1]+1):
            if y >= 0 and y < len(lines) and x >= 0 and x < len(lines[0]):
                # print((x, y), lines[y][x])
                yield lines[y][x], (y, x)


def part_numbers(lines):
    for y, line in enumerate(lines):
        for m in re.finditer(r'(\d+)', line):
            if any(issymbol(ch) for ch, _ in neighbors(lines, y, m)):
                yield m, y


def gear_ratios(lines):
    gears = defaultdict(list)

    # for y, line in enumerate(lines):
    #     for x, ch in enumerate(line):
    #         if ch == '*':
    #             gears[(y,x)] = []

    for m, y in part_numbers(lines):
        for ch, pos in neighbors(lines, y, m):
            if ch == '*':
                gears[pos].append(int(m.groups()[0]))

    return [parts[0] * parts[1] for parts in gears.values() if len(parts) == 2]


lines = [l.strip() for l in sys.stdin]

print('Part 1')
print(sum([int(m.groups()[0]) for m, _ in part_numbers(lines)]))

print('Part 2')
print(sum(gear_ratios(lines)))
