package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	if len(nums) < 2 {
		return Stats{}
	}
	diff := nums[1] - nums[0]
	stats := Stats{Count: len(nums) - 1, Sum: nums[len(nums)-1] - nums[0], Min: diff, Max: diff}
	for i := 1; i < len(nums)-1; i++ {
		diff = nums[i+1] - nums[i]
		stats.Min = min(stats.Min, diff)
		stats.Max = max(stats.Max, diff)
	}
	return stats
}
