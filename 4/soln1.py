import sys

total_score = 0

for line in sys.stdin:
    win_nums, nums = [{int(n) for n in s.strip().split()} for s in line.split(':')[1].split('|')]

    # print(win_nums)
    # print(nums)

    score = 2 ** (len(win_nums & nums) - 1) if len(win_nums & nums) > 0 else 0

    # print(score)

    total_score += score

print('Total points:', total_score)

