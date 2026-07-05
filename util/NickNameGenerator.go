package util

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/bits"
	"unicode/utf8"
)

// NickNameGenerator 用于生成不重复昵称。
type NickNameGenerator struct {
	adjectives []string
	nouns      []string
	suffixBase uint32
	minLen     int
	maxLen     int

	// 所有满足长度要求的 (adjIdx, nounIdx)
	pairs []pair

	// 可生成的总数量
	total uint64

	// 用于打乱 id
	seed uint64
	mask uint64
}

type pair struct {
	adj  uint32
	noun uint32
}

// NewNickNameGenerator 创建昵称生成器。
// suffixBase:
//
//	100   -> 00~99
//	1000  -> 000~999
//	10000 -> 0000~9999
//
// english-adjectives.txt english-nouns.txt https://gist.github.com/hugsy/8910dc78d208e40de42deb29e62df913
func NewNickNameGenerator(
	adjectives []string,
	nouns []string,
	minLen int,
	maxLen int,
	suffixBase uint32,
) (*NickNameGenerator, error) {
	if len(adjectives) == 0 {
		return nil, fmt.Errorf("empty adjectives")
	}
	if len(nouns) == 0 {
		return nil, fmt.Errorf("empty nouns")
	}
	if minLen <= 0 || maxLen < minLen {
		return nil, fmt.Errorf("invalid length range")
	}
	if suffixBase == 0 {
		return nil, fmt.Errorf("invalid suffixBase")
	}

	g := &NickNameGenerator{
		adjectives: adjectives,
		nouns:      nouns,
		minLen:     minLen,
		maxLen:     maxLen,
		suffixBase: suffixBase,
	}
	suffixLen := g.digits(suffixBase - 1)

	// 预计算所有满足长度要求的组合
	for ai, a := range adjectives {
		al := utf8.RuneCountInString(a)

		for ni, n := range nouns {
			nl := utf8.RuneCountInString(n)

			l := al + nl + suffixLen
			if l < minLen || l > maxLen {
				continue
			}

			g.pairs = append(g.pairs, pair{
				adj:  uint32(ai),
				noun: uint32(ni),
			})
		}
	}

	if len(g.pairs) == 0 {
		return nil, fmt.Errorf("no valid combinations")
	}

	g.total = uint64(len(g.pairs)) * uint64(suffixBase)

	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return nil, err
	}
	g.seed = binary.LittleEndian.Uint64(b[:])

	g.mask = g.nextMask(g.total)

	return g, nil
}

// Total 返回最多能生成多少个不重复昵称。
func (g *NickNameGenerator) Total() uint64 {
	return g.total
}

// Nickname 根据 id 生成昵称。
// 要求：0 <= id < Total()
func (g *NickNameGenerator) Nickname(id uint64) (string, error) {
	if id >= g.total {
		return "", fmt.Errorf("id out of range")
	}

	v := g.permute(id)

	pairIdx := v / uint64(g.suffixBase)
	num := v % uint64(g.suffixBase)

	p := g.pairs[pairIdx]

	return fmt.Sprintf(
		"%s%s%0*d",
		g.adjectives[p.adj],
		g.nouns[p.noun],
		g.digits(g.suffixBase-1),
		num,
	), nil
}

// --------------------------------
// internal
// --------------------------------

// cycle walking，保证结果落在 [0,total)
func (g *NickNameGenerator) permute(v uint64) uint64 {
	x := v

	for {
		y := g.feistel(x&g.mask, g.seed, g.mask)

		if y < g.total {
			return y
		}

		x = y
	}
}

// 64 位 Feistel
func (g *NickNameGenerator) feistel(x, key, mask uint64) uint64 {
	n := bits.Len64(mask + 1)
	half := n / 2

	leftMask := uint64((1 << half) - 1)

	l := x >> half
	r := x & leftMask

	for i := uint64(0); i < 4; i++ {
		f := g.mix(r ^ key ^ i)
		l, r = r, l^(f&leftMask)
	}

	return ((l << half) | r) & mask
}

func (g *NickNameGenerator) mix(x uint64) uint64 {
	x ^= x >> 33
	x *= 0xff51afd7ed558ccd
	x ^= x >> 33
	x *= 0xc4ceb9fe1a85ec53
	x ^= x >> 33
	return x
}

// 生成 >= n 的全 1 mask
// n=1_000_000 -> 0xFFFFF
func (g *NickNameGenerator) nextMask(n uint64) uint64 {
	if n <= 1 {
		return 1
	}

	b := bits.Len64(n - 1)
	return (uint64(1) << b) - 1
}

func (g *NickNameGenerator) digits(v uint32) int {
	d := 1
	for v >= 10 {
		d++
		v /= 10
	}
	return d
}
