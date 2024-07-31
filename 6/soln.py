from math import ceil, floor, sqrt


def ways_to_win(t, d):
    x_min = floor((t - sqrt(t ** 2 - 4 * d)) / 2) + 1
    x_max = ceil((t + sqrt(t ** 2 - 4 * d)) / 2) - 1

    print(x_min, x_max)

    return x_max - x_min + 1


times = input().split()[1:]
distances = input().split()[1:]

prod = 1

for t, d in zip(times, distances):
    w = ways_to_win(int(t), int(d))

    print(w)

    prod *= w

print('Part 1')
print('Product:', prod)

print('Part 2')
print('# Ways:', ways_to_win(int(''.join(times)), int(''.join(distances))))
