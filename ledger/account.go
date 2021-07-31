package ledger

import (
	"fmt"
	"regexp"
	"strings"
)

// Regular expression string that matches valid account name components.
// Categories are:
//   Lu: Uppercase letters.
//   L: All letters.
//   Nd: Decimal numbers.
const (
	kAccountCompTypeRe = `[\p{Lu}][\p{L}\p{Nd}-]*`
	kAccountCompNameRe = `[\p{Lu}\p{Nd}][\p{L}\p{Nd}-]*`
	kSep               = ":"
)

var kAccountRE *regexp.Regexp

func init() {
	kAccountRE = regexp.MustCompile(fmt.Sprintf("^(?:%s)(?:%s%s)+$", kAccountCompTypeRe, kSep, kAccountCompNameRe))
}

func IsAccountValid(account string) bool {
	return kAccountRE.MatchString(account)
}

func JoinAccount(sl []string) string {
	return strings.Join(sl, kSep)
}

func SplitAccount(account string) []string {
	return strings.Split(account, kSep)
}

func ParentAccount(account string) string {
	if len(account) == 0 {
		return ""
	}
	components := SplitAccount(account)

	return strings.Join(components[:len(components)-1], kSep)
}

func LeafAccount(account string) string {
	if len(account) == 0 {
		return ""
	}
	startpos := strings.LastIndex(account, kSep)
	return account[startpos+1:]
}

func AccountSansRoot(account string) string {
	if len(account) == 0 {
		return ""
	}
	components := strings.Split(account, kSep)

	return strings.Join(components[1:], kSep)
}

func AccountRoot(account string, numComponents int) string {
	if len(account) == 0 {
		return ""
	}
	kLen := len(kSep)
	index := 0
	count := 0
	for index < len(account)-kLen {
		if account[index:index+kLen] == kSep {
			count += 1
		}
		if count == numComponents {
			return account[0:index]
		}
		index += 1
	}
	return account
}

func HasAccountComponent(account, component string) bool {
	regexpc, _ := regexp.Compile(fmt.Sprintf("(^|:)%s(:|$)", component))
	return regexpc.Match([]byte(account))
}

func CommonPrefix(ss []string) string {
	// Special cases first
	switch len(ss) {
	case 0:
		return ""
	case 1:
		return ss[0]
	}
	// LCP of min and max (lexicographically)
	// is the LCP of the whole set.
	min, max := ss[0], ss[0]
	for _, s := range ss[1:] {
		switch {
		case s < min:
			min = s
		case s > max:
			max = s
		}
	}
	minParts := SplitAccount(min)
	maxParts := SplitAccount(max)
	var i = 0
	for i = 0; i < len(minParts) && i < len(maxParts); i++ {
		if minParts[i] != maxParts[i] {
			return JoinAccount(minParts[:i])
		}
	}
	return JoinAccount(minParts)
}
