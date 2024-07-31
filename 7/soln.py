from collections import Counter
import sys


def hand_type(hand: str, joker=False) -> int:
    counts = Counter(hand)

    if joker and 'J' in counts:
        # print('joker')
        # print(counts)
        freq = counts.pop('J')

        if len(counts) == 0:
            counts = Counter({'A': 5})
        else:
            most_common = counts.most_common(1)[0][0]
            counts[most_common] += freq

        # print(counts)
        # print()

    most_common, most_common_freq = counts.most_common(1)[0]

    if most_common_freq == 5:
        return 6
    elif most_common_freq == 4:
        return 5
    elif list(sorted(counts.values())) == [2, 3]:
        return 4
    elif most_common_freq == 3:
        return 3
    elif list(sorted(counts.values())) == [1, 2, 2]:
        return 2
    elif most_common_freq == 2:
        return 1
    elif most_common_freq == 1:
        return 0


def card_val(card: str, joker: bool = False) -> int:
    # print(card)
    if not joker:
        return '23456789TJQKA'.index(card)
    else:
        return 'J23456789TQKA'.index(card)


def hand_to_key(hand: str, joker=False) -> tuple[int, tuple]:
    return (hand_type(hand, joker=joker), tuple(card_val(card, joker=joker) for card in hand))


hands = []

for line in sys.stdin:
    hand, bid = line.split()

    bid = int(bid)

    print(hand, bid)

    hands.append((hand, bid))


print('\nPart 1')

sorted_hands = sorted(hands, key=lambda h: hand_to_key(h[0]))

# print(hands)
# print([hand_to_key(h[0]) for h in hands])

print('Total winnings:', sum((i+1) * bid for i,
      (_, bid) in enumerate(sorted_hands)))

print('\nPart 2')

sorted_hands = sorted(hands, key=lambda h: hand_to_key(h[0], joker=True))

# print(hands)
# print([hand_to_key(h[0]) for h in hands])

print('Total winnings:', sum((i+1) * bid for i,
      (_, bid) in enumerate(sorted_hands)))
