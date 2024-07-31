import heapq
import itertools
import sys
from typing import Iterable


class Mapping:
    def __init__(self, source, dest, length):
        self.source = source
        self.dest = dest
        self.length = length

    def contains(self, val):
        return self.source <= val < self.source + self.length

    def mapped(self, val):
        return val - self.source + self.dest

    def __str__(self):
        return f'({self.source} {self.dest} {self.length})'

    def __repr__(self):
        return f'({self.source} {self.dest} {self.length})'


class ValueMapper:
    def __init__(self, mappings):
        self.mappings = mappings

    def map(self, value):
        for m in self.mappings:
            if m.contains(value):
                return m.mapped(value)
        return value


class Range:
    def __init__(self, start, stop):
        self.start = start
        self.stop = stop

    def __lt__(self, other):
        return self.start < other.start

    def __repr__(self):
        return f'Range({self.start}, {self.stop})'
    
    @staticmethod
    def merge(ranges):
        starts = sorted([r.start for r in ranges], reverse=True)
        stops = sorted([r.stop for r in ranges], reverse=True)

        curr_start = -1
        curr_stop = -1
        count = 0
        while len(starts) > 0 or len(stops) > 0:
            if len(stops) == 0 or (len(starts) > 0 and starts[-1] <= stops[-1]):
                val = starts.pop()
                # print(val)
                if count == 0:
                    curr_start = val
                count += 1
            else:
                curr_stop = stops.pop()
                count -= 1

            # print(curr_start, curr_stop, count)

            if count == 0:
                yield Range(curr_start, curr_stop)
            


class RangeMapper:
    def __init__(self, mappings: list[Mapping]):
        self.mappings_source_asc = sorted(mappings, key=lambda m: m.source)
        # self.mappings_source_stop_asc = sorted(mappings, key=lambda m: m.source+m.length)

    def map(self, ranges: list[Range]):
        heapq.heapify(ranges)

        mapped_ranges = []

        map_idx = 0
        while len(ranges) > 0:
            # print(str(ranges).ljust(80), str(mapped_ranges).ljust(40), map_idx, self.mappings_source_asc[map_idx] if map_idx < len(self.mappings_source_asc) else None)

            r = heapq.heappop(ranges)

            if map_idx == len(self.mappings_source_asc):
                mapped_ranges.append(r)
                continue

            while map_idx < len(self.mappings_source_asc):
                m = self.mappings_source_asc[map_idx]
                m_start = m.source
                m_stop = m.source + m.length

                if r.start < m_start:
                    if r.stop <= m_start:
                        mapped_ranges.append(Range(r.start, r.stop))
                    else:
                        # print(r, m)
                        if r.start != m_start:
                            # print('push', Range(r.start, m_start))
                            heapq.heappush(ranges, Range(r.start, m_start))
                            
                        if m_start != r.stop:
                            # print('push', Range(m_start, r.stop))
                            heapq.heappush(ranges, Range(m_start, r.stop))
                    break
                elif r.start < m_stop:
                    if r.stop <= m_stop:
                        mapped_ranges.append(Range(m.mapped(r.start), m.mapped(r.stop)))
                    else:
                        if r.start != m_stop:
                            heapq.heappush(ranges, Range(r.start, m_stop))
                        
                        if m_stop != r.stop:
                            heapq.heappush(ranges, Range(m_stop, r.stop))
                    break
                else:
                    map_idx += 1
                    if map_idx == len(self.mappings_source_asc):
                        mapped_ranges.append(r)

        # print(str(ranges).ljust(80), mapped_ranges)
        # print()

        return list(Range.merge(mapped_ranges))


# range_mapper = RangeMapper([Mapping(98, 50, 2), Mapping(50, 52, 48)])
# rs = range_mapper.map([Range(50, 55), Range(79, 102)])
# print(rs)
# print(list(Range.merge(rs)))

# sys.exit()

seeds = []
seed_to_soil = []
soil_to_fertilizer = []
fertilizer_to_water = []
water_to_light = []
light_to_temperature = []
temperature_to_humidity = []
humidity_to_location = []

curr_map = []
for line in sys.stdin:
    if line.startswith('seeds:'):
        seeds = [int(n) for n in line.split()[1:]]
    elif line.startswith('seed-to-soil map:'):
        curr_map = seed_to_soil
    elif line.startswith('soil-to-fertilizer map:'):
        curr_map = soil_to_fertilizer
    elif line.startswith('fertilizer-to-water map:'):
        curr_map = fertilizer_to_water
    elif line.startswith('water-to-light map:'):
        curr_map = water_to_light
    elif line.startswith('light-to-temperature map:'):
        curr_map = light_to_temperature
    elif line.startswith('temperature-to-humidity map:'):
        curr_map = temperature_to_humidity
    elif line.startswith('humidity-to-location map:'):
        curr_map = humidity_to_location
    else:
        try:
            dest, source, length = [int(n) for n in line.split()]
            curr_map.append(Mapping(source, dest, length))
        except ValueError as e:
            pass

# print(seeds)
# print(seed_to_soil)
# print(soil_to_fertilizer)
# print(fertilizer_to_water)
# print(water_to_light)
# print(light_to_temperature)
# print(temperature_to_humidity)
# print(humidity_to_location)

print('Part 1')
# print('# seeds:', len(seeds))

val_mappers = [ValueMapper(ms) for ms in (seed_to_soil, soil_to_fertilizer, fertilizer_to_water,
               water_to_light, light_to_temperature, temperature_to_humidity, humidity_to_location)]

locations = []
for seed in seeds:
    val = seed
    for mapper in val_mappers:
        val = mapper.map(val)

    locations.append(val)

print('Min location:', min(locations))

# -----------------------------------------------------------------------------

print('\nPart 2')

seed_ranges = [Range(start, start+length)
               for start, length in zip(seeds[::2], seeds[1::2])]

# # brute force
# new_seeds = list(itertools.chain.from_iterable(seed_ranges))

# # print('# seeds:', len(new_seeds))

# print('Min location:', min(map_seeds_to_locations([seed_to_soil, soil_to_fertilizer, fertilizer_to_water,
#       water_to_light, light_to_temperature, temperature_to_humidity, humidity_to_location], new_seeds)))

range_mappers = [RangeMapper(ms) for ms in (seed_to_soil, soil_to_fertilizer, fertilizer_to_water,
                                            water_to_light, light_to_temperature, temperature_to_humidity, humidity_to_location)]

ranges = seed_ranges

for mapper in range_mappers:
    print(ranges)
    ranges = mapper.map(ranges)

print(ranges)

print('Min location:', min(r.start for r in ranges))
