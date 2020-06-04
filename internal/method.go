package internal

import (
	"reflect"
	"strconv"
	"strings"
)

func argTypeList(args, types []string) string {
	out := make([]string, 0, len(args))
	for k, v := range args {
		out = append(out, v+" "+types[k])
	}
	return strings.Join(out, ", ")
}

func typeArg(pre string) func(string) string {
	m := map[uint8]int64{}
	return func(s string) string {
		s = strings.ToLower(s)
		arg := s[0]
		for ; !('a' <= arg && arg <= 'z'); s = s[1:] {
			if len(s) == 0 {
				arg = 'a'
				break
			}
			arg = s[0]
		}
		out := pre + string(arg)
		if n, ok := m[arg]; ok {
			out += strconv.FormatInt(n, 10)
		}
		m[arg]++
		return out
	}
}

type Method struct {
	name              string
	inArgs, inTypes   []string
	outArgs, outTypes []string
}

func BuildMethod(rm reflect.Method) (Method, error) {
	t := rm.Type
	m := Method{
		name:     rm.Name,
		outTypes: make([]string, t.NumOut()),
		outArgs:  make([]string, t.NumOut()),
	}
	if l := t.NumIn() - 1; l > 0 {
		m.inTypes, m.inArgs = make([]string, t.NumIn()), make([]string, l)
	}
	inArgs := typeArg("")
	outArgs := typeArg("o")
	for i := 1; i < t.NumIn(); i++ {
		t := t.In(i).String()
		m.inTypes[i-1] = t
		m.inArgs[i-1] = inArgs(t)
	}
	for k := range m.outArgs {
		t := t.Out(k).String()
		m.outTypes[k] = t
		m.outArgs[k] = outArgs(t)
	}
	return m, nil
}

func (m Method) IntTypes() []string {
	return m.inTypes
}

func (m Method) InArgs() []string {
	return m.inArgs
}

func (m Method) InArgsList() string {
	return strings.Join(m.inArgs, ", ")
}

func (m Method) OutTypes() []string {
	return m.outTypes
}

func (m Method) OutArgs() []string {
	return m.outArgs
}

func (m Method) OutArgsList() string {
	return strings.Join(m.outArgs, ", ")
}

func (m Method) Name() string {
	return m.name
}

func (m Method) Definition() string {
	in, out := argTypeList(m.inArgs, m.inTypes), argTypeList(m.outArgs, m.outTypes)
	if out != "" {
		out = " (" + out + ")"
	}
	return m.name + "(" + in + ")" + out
}
