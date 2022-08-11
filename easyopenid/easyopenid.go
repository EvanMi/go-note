package easyopenid

import "errors"

func GetUserId(appUuid string, openId string) (string, error) {
	return calc(appUuid, openId, func(i1, i2 int) int { return i1 - i2 })
}

func GetOpenId(appUuid string, userUuid string) (string, error) {
	return calc(appUuid, userUuid, func(i1, i2 int) int { return i1 + i2 })
}

func calc(firstUuid string, secondUuid string, calcFunc func(int, int) int) (string, error) {
	res := [32]rune{}
	firstUuidLen := len(firstUuid)
	secondLen := len(secondUuid)
	if firstUuidLen != 32 || secondLen != 32 {
		return "", errors.New("error length")
	}
	firstUuidRuneArr := []rune(firstUuid)
	secondUuidRuneArr := []rune(secondUuid)
	for i := 0; i < 32; i++ {
		firstUuidInt, firstUuidIntErr := charToInt(firstUuidRuneArr[i])
		secondUuidInt, secondUuidIntErr := charToInt(secondUuidRuneArr[i])
		if firstUuidIntErr != nil || secondUuidIntErr != nil {
			return "", firstUuidIntErr
		}
		resInt := calcFunc(secondUuidInt, firstUuidInt)
		resRune, resErr := intToChar(resInt)
		if resErr != nil {
			return "", resErr
		}
		res[i] = resRune
	}
	return string(res[:]), nil
}

func charToInt(c rune) (int, error) {
	switch c {
	case '0':
		return 0, nil
	case '1':
		return 1, nil
	case '2':
		return 2, nil
	case '3':
		return 3, nil
	case '4':
		return 4, nil
	case '5':
		return 5, nil
	case '6':
		return 6, nil
	case '7':
		return 7, nil
	case '8':
		return 8, nil
	case '9':
		return 9, nil
	case 'a':
		return 10, nil
	case 'b':
		return 11, nil
	case 'c':
		return 12, nil
	case 'd':
		return 13, nil
	case 'e':
		return 14, nil
	case 'f':
		return 15, nil
	case 'g':
		return 16, nil
	case 'h':
		return 17, nil
	case 'i':
		return 18, nil
	case 'j':
		return 19, nil
	case 'k':
		return 20, nil
	case 'l':
		return 21, nil
	case 'm':
		return 22, nil
	case 'n':
		return 23, nil
	case 'o':
		return 24, nil
	case 'p':
		return 25, nil
	case 'q':
		return 26, nil
	case 'r':
		return 27, nil
	case 's':
		return 28, nil
	case 't':
		return 29, nil
	case 'u':
		return 30, nil
	default:
		return -1, errors.New("illegal character found")
	}
}

func intToChar(i int) (rune, error) {
	switch i {
	case 0:
		return '0', nil
	case 1:
		return '1', nil
	case 2:
		return '2', nil
	case 3:
		return '3', nil
	case 4:
		return '4', nil
	case 5:
		return '5', nil
	case 6:
		return '6', nil
	case 7:
		return '7', nil
	case 8:
		return '8', nil
	case 9:
		return '9', nil
	case 10:
		return 'a', nil
	case 11:
		return 'b', nil
	case 12:
		return 'c', nil
	case 13:
		return 'd', nil
	case 14:
		return 'e', nil
	case 15:
		return 'f', nil
	case 16:
		return 'g', nil
	case 17:
		return 'h', nil
	case 18:
		return 'i', nil
	case 19:
		return 'j', nil
	case 20:
		return 'k', nil
	case 21:
		return 'l', nil
	case 22:
		return 'm', nil
	case 23:
		return 'n', nil
	case 24:
		return 'o', nil
	case 25:
		return 'p', nil
	case 26:
		return 'q', nil
	case 27:
		return 'r', nil
	case 28:
		return 's', nil
	case 29:
		return 't', nil
	case 30:
		return 'u', nil
	default:
		return 'x', errors.New("illegal int found")
	}
}
