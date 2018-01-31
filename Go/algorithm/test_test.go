
// 转义sql，将单独一个'替换为''，其他情况不替换
func SqlEscape(sql string) string {
	const apostropheConst = "'"
	apostrophe := rune(apostropheConst[0])
	size := len(sql)
	runes := make([]rune, 0, size+10)

	pre := false
	cur := false

	for i, r := range sql {
		cur = r == apostrophe
		next := false
		if i < size-1 {
			next = rune(sql[i+1]) == apostrophe
		}
		if !next && cur && !pre {
			runes = append(runes, apostrophe)
		}
		runes = append(runes, r)
		pre = cur
	}

	return string(runes)
}

func SqlEscape1(sql string) string {
	const apostropheConst = "'"
	count := strings.Count(sql, apostropheConst)
	if count <= 0 {
		return sql
	}
	apostrophe := rune(apostropheConst[0])
	size := len(sql)
	runes := make([]rune, 0, size+count*2)
	for i := size - 1; i >= 0; i-- {
		r := rune(sql[i])
		runes = append(runes, r)
		if i == size-1 {
			if r == apostrophe {
				count--
			}
			if size > 1 && rune(sql[size-2]) != apostrophe {
				runes = append(runes, r)
			}
			continue
		}
		if count <= 0 {
			runes = append(runes, []rune(sql[:i])...)
			break
		}
		if r == apostrophe && i > 0 && rune(sql[i+1]) != apostrophe && rune(sql[i-1]) != apostrophe {
			runes = append(runes, r)
			count--
		} else if r == apostrophe && i == 0 && rune(sql[i+1]) != apostrophe {
			runes = append(runes, r)
			count--
		}
	}
	convert(runes)
	return string(runes)
}

func SqlEscape3(sql string) string {
	const apostropheConst = "'"
	apostrophe := rune(apostropheConst[0])
	ret := bytes.NewBuffer(nil)
	size := len(sql)
	for i, r := range sql {
		if i == 0 {
			ret.WriteRune(r)
			if size > 1 && r == apostrophe && rune(sql[1]) != apostrophe {
				ret.WriteRune(apostrophe)
			}
			continue
		}
		if r == apostrophe && i < size-1 && rune(sql[i+1]) != apostrophe && rune(sql[i-1]) != apostrophe {
			ret.WriteRune(apostrophe)
			ret.WriteRune(apostrophe)
		} else if r == apostrophe && i == size-1 && rune(sql[i-1]) != apostrophe {
			ret.WriteRune(apostrophe)
			ret.WriteRune(apostrophe)
		} else {
			ret.WriteRune(r)
		}
	}
	return ret.String()
}

func convert(runes []rune) {
	size := len(runes)
	helfSize := size / 2
	for i := 0; i < helfSize; i++ {
		runes[i], runes[size-1-i] = runes[size-1-i], runes[i]
	}
}

