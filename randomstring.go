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

func Now(layout string) Generator {
	return func(buf *strings.Builder) (*strings.Builder, error) {
		buf.WriteString(time.Now().Format(layout))
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
