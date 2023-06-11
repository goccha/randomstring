package randomstring

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"
)

func CharSet(token string, length int) Generator {
	return func(buf *strings.Builder) (*strings.Builder, error) {
		for i := 0; i < length; i++ {
			v, err := rand.Int(rand.Reader, big.NewInt(int64(len(token))))
			if err != nil {
				return buf, err
			}
			buf.WriteByte(token[v.Int64()])
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

func Numbers(length int) Generator {
	return CharSet("0123456789", length)
}

func Uppers(length int) Generator {
	return CharSet("ABCDEFGHIJKLMNOPQRSTUVWXYZ", length)
}

func Lowers(length int) Generator {
	return CharSet("abcdefghijklmnopqrstuvwxyz", length)
}

func AlphaNumeric(length int) Generator {
	return CharSet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", length)
}

func All(length int) Generator {
	return CharSet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_#$=?@[]!%&'()~|^\\;:,./`{+*}>", length)
}

func Base64(length int) Generator {
	return CharSet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/", length)
}

func Base64Url(length int) Generator {
	return CharSet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_", length)
}

func LowersAlphaNumeric(length int) Generator {
	return CharSet("abcdefghijklmnopqrstuvwxyz0123456789", length)
}

func UppersAlphaNumeric(length int) Generator {
	return CharSet("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", length)
}

func SymbolAll(length int) Generator {
	return CharSet("-_#$=?@[]!%&'()~|^\\;:,./`{+*}>", length)
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
