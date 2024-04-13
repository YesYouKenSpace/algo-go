package problems

var multipliers = map[int]int{
	1: 50,
	2: 10,
	3: 1,
}



var memo = map[string]map[int]int {}
func getDescendents(referral map[string][]string, u string) map[int]int {
	// Func getDescendents(referrals map[string][]string, user string) map[int]int {
	// // if memo[user] != nil return memo[user]
	// // if referrals[user] = nil return map[int]int{}
	// // for each child in referrals[user] return sum(getDescendents(child))
	// }
	if val, ok := memo[u]; ok {
		return val
	}

	if children, ok := referral[u]; ok {
		descendents := map[int]int{
			1: len(referral[u]),
		}
		for _, child := range children {
			descendentsOfChild := getDescendents(referral, child)
			descendents[2] += descendentsOfChild[1]
			descendents[3] += descendentsOfChild[2]
		}
		memo[u] = descendents
		return descendents
	}
	return map[int]int{}
}


func computeScore(referral map[string][]string) map[string]int {
	ans := make(map[string]int, len(referral))
	// for each user, print their score
	for user, _:= range referral {
		descendents := getDescendents(referral, user)

		ans[user] = descendents[1]*multipliers[1] + descendents[2]*multipliers[2] + descendents[3]*multipliers[3]
	}
	return ans
}
