from collections import defaultdict
import sys

card_counts = defaultdict(int)

for line in sys.stdin:
    card_no = int(line.split(':')[0].split()[-1])
    win_nums, nums = [{int(n) for n in s.strip().split()} for s in line.split(':')[1].split('|')]

    card_counts[card_no] += 1

    match_count = len(win_nums & nums)

    for i in range(match_count):
        card_counts[card_no + i + 1] += card_counts[card_no]
    
print(card_counts)
print('Total cards:', sum(card_counts.values()))
