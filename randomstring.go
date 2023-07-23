package randomstring

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func CharSet(token string, length int, maxLength ...int) Generator {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func(buf *strings.Builder) (*strings.Builder, error) {
		if len(maxLength) > 0 && maxLength[0] > length {
			min := length
			max := maxLength[0]
			length = seed.Intn(max-min+1) + min
		}
		for i := 0; i < length; i++ {
			v := seed.Intn(len(token))
			buf.WriteByte(token[v])
		}
		return buf, nil
	}
}

func Builder(builder *strings.Builder) Generator {
	return func(buf *strings.Builder) (*strings.Builder, error) {
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
	return func(buf *strings.Builder) (*strings.Builder, error) {
		buf.Grow(length)
		return buf, nil
	}
}

func Fix(token string) Generator {
	return func(buf *strings.Builder) (*strings.Builder, error) {
		buf.WriteString(token)
		return buf, nil
	}
}

func Now(layout string, loc ...*time.Location) Generator {
	return func(buf *strings.Builder) (*strings.Builder, error) {
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
	return func(buf *strings.Builder) (*strings.Builder, error) {
		buf.WriteString(fmt.Sprintf(format, a...))
		return buf, nil
	}
}

type Generator func(buf *strings.Builder) (*strings.Builder, error)

func Build(gen ...Generator) (string, error) {
	buf := &strings.Builder{}
	var err error
	for _, f := range gen {
		buf, err = f(buf)
		if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func Gen(gen ...Generator) string {
	v, err := Build(gen...)
	if err != nil {
		panic(err)
	}
	return v
}
