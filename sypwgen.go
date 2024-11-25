package main

// ver 0.1

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	// Init defaults
	pwlen := 10
	pwnum := 1
	p_uppercase := false
	p_colored := false

	// Init args
	arg_pwlen := flag.Int("l", 10, "Password length")
	arg_pwnum := flag.Int("n", 1, "Count of password")
	arg_cap := flag.Bool("u", false, "Upper case")
	arg_col := flag.Bool("c", false, "Colored syllable")

	// Parse args
	flag.Parse()

	if *arg_pwlen <= 0 {
		pwlen = 10
	} else {
		pwlen = *arg_pwlen
		pwlen += 1
	}

	if *arg_pwnum <= 0 {
		pwnum = 1
	} else {
		pwnum = *arg_pwnum
	}

	if *arg_cap {
		p_uppercase = *arg_cap
	} else {
		p_uppercase = false
	}

	if *arg_col {
		p_colored = *arg_col
	} else {
		p_colored = false
	}

	// Init colors vars
	slog_colors := []string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m", "\033[37m"}
	color_reset := "\033[0m"

	// Syllable list
	slogs := []string{"aq", "aw", "ae", "ar", "at", "ay", "au", "ai", "ao", "ap",
		"as", "ad", "af", "ag", "ah", "aj", "ak", "al", "az", "ax",
		"ac", "av", "ab", "an", "am", "eq", "ew", "er", "et", "ey",
		"eu", "ei", "eo", "ep", "ea", "es", "ed", "ef", "eg", "eh",
		"ej", "ek", "el", "ez", "ex", "ec", "ev", "eb", "en", "em",
		"iq", "iw", "ie", "ir", "it", "iy", "iu", "io", "ip", "ia",
		"is", "id", "if", "ig", "ih", "ij", "ik", "il", "iz", "ix",
		"ic", "iv", "ib", "in", "im", "oq", "ow", "oe", "or", "ot",
		"oy", "ou", "oi", "op", "oa", "os", "od", "of", "og", "oh",
		"oj", "ok", "ol", "oz", "ox", "oc", "ov", "ob", "on", "om",
		"uq", "uw", "ue", "ur", "ut", "uy", "ui", "uo", "up", "ua",
		"us", "ud", "uf", "ug", "uh", "uj", "uk", "ul", "uz", "ux",
		"uc", "uv", "ub", "un", "um", "yq", "yw", "ye", "yr", "yt",
		"yu", "yi", "yo", "yp", "ya", "ys", "yd", "yf", "yg", "yh",
		"yj", "yk", "yl", "yz", "yx", "yc", "yv", "yb", "yn", "ym",
		"qe", "qy", "qu", "qi", "qo", "qa", "we", "wy", "wu", "wi",
		"wo", "wa", "re", "ry", "ru", "ri", "ro", "ra", "te", "ty",
		"tu", "ti", "to", "ta", "pe", "py", "pu", "pi", "po", "pa",
		"se", "sy", "su", "si", "so", "sa", "de", "dy", "du", "di",
		"do", "da", "fe", "fy", "fu", "fi", "fo", "fa", "ge", "gy",
		"gu", "gi", "go", "ga", "he", "hy", "hu", "hi", "ho", "ha",
		"je", "jy", "ju", "ji", "jo", "ja", "ke", "ky", "ku", "ki",
		"ko", "ka", "le", "ly", "lu", "li", "lo", "la", "ze", "zy",
		"zu", "zi", "zo", "za", "xe", "xy", "xu", "xi", "xo", "xa",
		"ce", "cy", "cu", "ci", "co", "ca", "ve", "vy", "vu", "vi",
		"vo", "va", "be", "by", "bu", "bi", "bo", "ba", "ne", "ny",
		"nu", "ni", "no", "na", "me", "my", "mu", "mi", "mo", "ma",
		"azh", "ach", "ash", "ezh", "ech", "esh", "izh", "ich", "ish", "ozh",
		"och", "osh", "uzh", "uch", "ush", "yzh", "ych", "ysh", "zha", "zhe",
		"zhi", "zho", "zhu", "zhy", "cha", "che", "chi", "cho", "chu", "chy",
		"sha", "she", "shi", "sho", "shu", "shy"}

	m := 0
	p_col := ""
	p_tmp_col := ""

	for m < pwnum {
		n := 1
		if p_uppercase {
			for n < pwlen {
				s := slogs[rand.Intn(len(slogs))]
				new_s := ""
				for i := 0; i < len(s); i++ {
					if rand.Intn(42)%2 == 0 {
						new_s += string(s[i])
					} else {
						new_s += strings.ToUpper(string(s[i]))
					}
				}
				if p_colored {
					for {
						if p_tmp_col != p_col {
							break
						}
						p_col = slog_colors[rand.Intn(len(slog_colors))]
					}
					fmt.Print(p_col, new_s, string(color_reset))
					p_tmp_col = p_col
				} else {
					fmt.Print(new_s)
				}
				n += 1
			}
		} else {
			for n < pwlen {
				if p_colored {
					for {
						if p_tmp_col != p_col {
							break
						}
						p_col = slog_colors[rand.Intn(len(slog_colors))]
					}
					fmt.Print(p_col, slogs[rand.Intn(len(slogs))], string(color_reset))
					p_tmp_col = p_col
				} else {
					fmt.Print(slogs[rand.Intn(len(slogs))])
				}
				n += 1
			}
		}
		m += 1
		fmt.Println()
	}
}
