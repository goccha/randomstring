package randomstring

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func CharSet(token string, length int, maxLength ...int) Generator {
	return func(buf *strings.Builder, seed *rand.Rand) (*strings.Builder, error) {
		if len(maxLength) > 0 && maxLength[0] > length {
			minLen := length
			maxLen := maxLength[0]
			length = seed.Intn(maxLen-minLen+1) + minLen
		}
		for i := 0; i < length; i++ {
			v := seed.Intn(len(token))
			buf.WriteByte(token[v])
		}
		return buf, nil
	}
}

func Merge(builder *strings.Builder) Generator {
	return func(buf *strings.Builder, seed *rand.Rand) (*strings.Builder, error) {
		if buf.Len() > 0 {
			capacity := builder.Cap()
			buf.WriteString(builder.String())
			builder.Reset()
			if buf.Cap() > capacity {
				builder.Grow(buf.Cap())
			} else {
				builder.Grow(capacity)
			}
			builder.WriteString(buf.String())
		}
		return builder, nil
	}
}

func Grow(length int) Generator {
	return func(buf *strings.Builder, seed *rand.Rand) (*strings.Builder, error) {
		buf.Grow(length)
		return buf, nil
	}
}

func Fix(token string, options ...string) Generator {
	if len(options) > 0 {
		tokens := make([]string, 0, len(options)+1)
		tokens = append(tokens, token)
		tokens = append(tokens, options...)
		return func(buf *strings.Builder, seed *rand.Rand) (*strings.Builder, error) {
			buf.WriteString(tokens[seed.Intn(len(tokens))])
			return buf, nil
		}
	}
	return func(buf *strings.Builder, seed *rand.Rand) (*strings.Builder, error) {
		buf.WriteString(token)
		return buf, nil
	}
}

func Now(layout string, loc ...*time.Location) Generator {
	return func(buf *strings.Builder, seed *rand.Rand) (*strings.Builder, error) {
		now := time.Now()
		if len(loc) > 0 {
			now = now.In(loc[0])
		}
		buf.WriteString(now.Format(layout))
		return buf, nil
	}
}

func Numbers(length int, maxLength ...int) Generator {
	return CharSet("0123456789", length, maxLength...)
}

func Uppers(length int, maxLength ...int) Generator {
	return CharSet("ABCDEFGHIJKLMNOPQRSTUVWXYZ", length, maxLength...)
}

func Lowers(length int, maxLength ...int) Generator {
	return CharSet("abcdefghijklmnopqrstuvwxyz", length, maxLength...)
}

func AlphaNumeric(length int, maxLength ...int) Generator {
	return CharSet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", length, maxLength...)
}

func All(length int, maxLength ...int) Generator {
	return CharSet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_#$=?@[]!%&'()~|^\\;:,./`{+*}>", length, maxLength...)
}

func Base64(length int, maxLength ...int) Generator {
	return CharSet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/", length, maxLength...)
}

func Base64Url(length int, maxLength ...int) Generator {
	return CharSet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_", length, maxLength...)
}

func Alphabet(length int, maxLength ...int) Generator {
	return CharSet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", length, maxLength...)
}

func LowersAlphaNumeric(length int, maxLength ...int) Generator {
	return CharSet("abcdefghijklmnopqrstuvwxyz0123456789", length, maxLength...)
}

func UppersAlphaNumeric(length int, maxLength ...int) Generator {
	return CharSet("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", length, maxLength...)
}

func SymbolAll(length int, maxLength ...int) Generator {
	return CharSet("-_#$=?@[]!%&'()~|^\\;:,./`{+*}>", length, maxLength...)
}

func Format(format string, a ...any) Generator {
	return func(buf *strings.Builder, seed *rand.Rand) (*strings.Builder, error) {
		buf.WriteString(fmt.Sprintf(format, a...))
		return buf, nil
	}
}

type Generator func(buf *strings.Builder, seed *rand.Rand) (*strings.Builder, error)

var _builder = New()

func Build(gen ...Generator) (string, error) {
	return _builder.Build(gen...)
}

func Gen(gen ...Generator) string {
	return _builder.Sync(gen...)
}

type Builder struct {
	seed *rand.Rand
	mux  sync.Mutex
}

func (b *Builder) Build(gen ...Generator) (string, error) {
	buf := &strings.Builder{}
	var err error
	for _, f := range gen {
		buf, err = f(buf, b.seed)
		if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func (b *Builder) Gen(gen ...Generator) string {
	v, err := b.Build(gen...)
	if err != nil {
		panic(err)
	}
	return v
}

func (b *Builder) Sync(gen ...Generator) string {
	b.mux.Lock()
	defer b.mux.Unlock()
	return b.Gen(gen...)
}

func New() *Builder {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Builder{seed: seed, mux: sync.Mutex{}}
}
