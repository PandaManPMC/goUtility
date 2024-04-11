package util

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestRsaSignSha256(t *testing.T) {
	//pubKeyB64 := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF3Z3RLYk1SSVhmRmVzK1lPeDMwbApnUTQ0czJZMktIMDlyNzdFYlBmSlJZOW9Ka2lrR0t3TVVBQnVRVHhXa3M3WjkvZFlIVVZRNkx2UThZQS9kbE51Ck4vU2VEb1pEMFBGYkdGSmJaTlo5SHVpSG93Q1c3MzI3Nk5pQmtVdnhEcFNMRU9HMWxUa0RiT01uUi9TS0s2NjAKcjl3Yy9STmg0bWxYMUgzZjhMZzJtUGJYa2xrQzBLemsxQUFIVWJmNVpwRG1Pd0QyR2dOcnMyL2lTTUM5MmJOSApYVXp2T0d2VjBCb01sd2V5VWM1Z093OHlyUkRtOStzMFVjU3B1RDArOEtqMnRneWppNUdodHZqb2hrTlQ0VUVTCndudnBuZFZZZ2ZTOHNWVkNsWWZLZFV5RjNhVHUxTEQ4dEtHb0swbnhjQkdQaTZRalFadFQvcWFkV2xHTExpaFYKUVFJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg=="
	//priKeyB64 := "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFdlFJQkFEQU5CZ2txaGtpRzl3MEJBUUVGQUFTQ0JLY3dnZ1NqQWdFQUFvSUJBUURDQzBwc3hFaGQ4VjZ6CjVnN0hmU1dCRGppelpqWW9mVDJ2dnNSczk4bEZqMmdtU0tRWXJBeFFBRzVCUEZhU3p0bjM5MWdkUlZEb3U5RHgKZ0Q5MlUyNDM5SjRPaGtQUThWc1lVbHRrMW4wZTZJZWpBSmJ2ZmJ2bzJJR1JTL0VPbElzUTRiV1ZPUU5zNHlkSAo5SW9ycnJTdjNCejlFMkhpYVZmVWZkL3d1RGFZOXRlU1dRTFFyT1RVQUFkUnQvbG1rT1k3QVBZYUEydXpiK0pJCndMM1pzMGRkVE84NGE5WFFHZ3lYQjdKUnptQTdEekt0RU9iMzZ6UlJ4S200UFQ3d3FQYTJES09Ma2FHMitPaUcKUTFQaFFSTENlK21kMVZpQjlMeXhWVUtWaDhwMVRJWGRwTzdVc1B5MG9hZ3JTZkZ3RVkrTHBDTkJtMVArcHAxYQpVWXN1S0ZWQkFnTUJBQUVDZ2dFQkFJajJVemZtYTNYemtuYkVZWllwSFRtMGdnME9qaGVTSHVKWGNtbS9sQTlICkk2b3lCN0ZxYnQ1aEQzRjRWMXNVS2dHK1VqR0c1WThBVW9ERGx3ZTc1OFlUSVNUN1hBNjA3U21EcUFMSzZsSFEKcXp4QWhFalNwTG03WitqWWczTlpJYmR1dVM1MHFaaEgxVWdTc1J3WUdtMHVuajk3V05Ib3JSZk5LUzNOdUt2SwpCSXRibmZDNXA5ZktyaWVhUFM3dDRmRjhBbFNjSkJlOXFzNUxGMno5NVk5bEovUDRrWWxnVUZHa0gvQ2hJK0VDCkFqbmVwQ1lqbiszQUg2UjI2b3lMd1FSVWZVbURtOUMxME5vVlNCcXRQNDZwSnJaKzE5UDh2cEpnTUFrc01BNUwKMWYvU1R6NFJ1cFNQT1I4VTZKSEtTUzl5ekMrblkxcW83M0VKUmVGdXJXRUNnWUVBNTdaTDUrbXRFREhnM0p1cwpYQXp3YkJBbktoNVNZUkdETVZQZGpHa3liTC81cXpsdEJESS9IMEpuYytld1ZCWnhtT1R4TU53ZEczZ0l2SXhOCmxjY0E0OTVaWk42SEJSS2lHNCs2bkFob2U0d09KcGJmbllWV3psWW8ycGl5R0pUd1JiVDFJVU9YeGRqUEhkV2MKVFVUNEFUd05WYmpaNmlxSlBySS9rVDV2UlVVQ2dZRUExbUk0aldGbFBGaktmb1RtVWkydCtjelRTaVhlMXFmNgo5SS8wcVJsWUFqaVdrMjFyNW5LVUtGNlM3azhFSElqUmJVOU1BYjdWQm9aWVJUVEJLd2ZjK01FR0dzclZIMHFDCnBSSy9IanJGZVcyY2xaMmRIWWpxamVETy80R0NpQTYzR1JuQSthWmpmWXo0K0xJUTBOQngvSllkNGhNWEZvUk0KbVEybTBZaHp1YzBDZ1lBME96YVBGM0NvaHVYT05OVThocm9uVWRqU09MV1BKZmh4eFJyYXpOZk9CZFJNMFl3TwpkeGtkZmNWK0xncmtXWTdQelVQRkpNajI2UzdtK2FWL2pyVlhxRVowWTJrQ0xyb1dCbWNsUnd2dVZaclcvZ2w0Ckk1ZDJ1WnRKODBPcUlPQ3NoZWIvMFpIRHltU2RzQW9rck5oT3h1K21sQjJqR0dXSm1YcVV6Z01kRlFLQmdIUlIKR21uckhDaGY3STQwd1ZwNUdsbmNmZzlPK05ieWtVQzhFbnpsR1ZFckx0ZVNtT2FSNkR6M0F6VjFmYitWcER0dwp1TWFCcWNjK2dRb2JrMnNyZXdNa2g1RmwxN2lBanQzTmpCQjB4c3daWXNueW1GcDcvUGM1c0ZZRkNMT1ZlRmFRCkdKbmJZME90aHpBNFBOTnZKVWxza1k0bDJYTUlHUjg1dnZjVTErVkpBb0dBZTNuSGtHUG5mZVhaWExPQUFzSW8KQUZPOW9nMnMrMVZjeEJBbEV2NGpucTlmdVRvVEs5Mjk5Ti8vYTArNFNHaCtBR2NrTjB2MEhPUjVCZFVock53QworUm0wWFFsNDduNWZFaXZqODdyMXhlVlhlOE05Z1lhZ2xqMVFKWlJmdXg4SExwZnZ3OVorNExpcHBSUzdrSVZ3CkYzbGJ4WEJma0pPUEJhN3poNExkRGFNPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ=="

	pubKeyB64 := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwgtKbMRIXfFes+YOx30l" +
		"gQ44s2Y2KH09r77EbPfJRY9oJkikGKwMUABuQTxWks7Z9/dYHUVQ6LvQ8YA/dlNu" +
		"N/SeDoZD0PFbGFJbZNZ9HuiHowCW73276NiBkUvxDpSLEOG1lTkDbOMnR/SKK660" +
		"r9wc/RNh4mlX1H3f8Lg2mPbXklkC0Kzk1AAHUbf5ZpDmOwD2GgNrs2/iSMC92bNH" +
		"XUzvOGvV0BoMlweyUc5gOw8yrRDm9+s0UcSpuD0+8Kj2tgyji5GhtvjohkNT4UES" +
		"wnvpndVYgfS8sVVClYfKdUyF3aTu1LD8tKGoK0nxcBGPi6QjQZtT/qadWlGLLihV" +
		"QQIDAQAB"

	priKeyB64 := "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDCC0psxEhd8V6z" +
		"5g7HfSWBDjizZjYofT2vvsRs98lFj2gmSKQYrAxQAG5BPFaSztn391gdRVDou9Dx" +
		"gD92U2439J4OhkPQ8VsYUltk1n0e6IejAJbvfbvo2IGRS/EOlIsQ4bWVOQNs4ydH" +
		"9IorrrSv3Bz9E2HiaVfUfd/wuDaY9teSWQLQrOTUAAdRt/lmkOY7APYaA2uzb+JI" +
		"wL3Zs0ddTO84a9XQGgyXB7JRzmA7DzKtEOb36zRRxKm4PT7wqPa2DKOLkaG2+OiG" +
		"Q1PhQRLCe+md1ViB9LyxVUKVh8p1TIXdpO7UsPy0oagrSfFwEY+LpCNBm1P+pp1a" +
		"UYsuKFVBAgMBAAECggEBAIj2Uzfma3XzknbEYZYpHTm0gg0OjheSHuJXcmm/lA9H" +
		"I6oyB7Fqbt5hD3F4V1sUKgG+UjGG5Y8AUoDDlwe758YTIST7XA607SmDqALK6lHQ" +
		"qzxAhEjSpLm7Z+jYg3NZIbduuS50qZhH1UgSsRwYGm0unj97WNHorRfNKS3NuKvK" +
		"BItbnfC5p9fKrieaPS7t4fF8AlScJBe9qs5LF2z95Y9lJ/P4kYlgUFGkH/ChI+EC" +
		"AjnepCYjn+3AH6R26oyLwQRUfUmDm9C10NoVSBqtP46pJrZ+19P8vpJgMAksMA5L" +
		"1f/STz4RupSPOR8U6JHKSS9yzC+nY1qo73EJReFurWECgYEA57ZL5+mtEDHg3Jus" +
		"XAzwbBAnKh5SYRGDMVPdjGkybL/5qzltBDI/H0Jnc+ewVBZxmOTxMNwdG3gIvIxN" +
		"lccA495ZZN6HBRKiG4+6nAhoe4wOJpbfnYVWzlYo2piyGJTwRbT1IUOXxdjPHdWc" +
		"TUT4ATwNVbjZ6iqJPrI/kT5vRUUCgYEA1mI4jWFlPFjKfoTmUi2t+czTSiXe1qf6" +
		"9I/0qRlYAjiWk21r5nKUKF6S7k8EHIjRbU9MAb7VBoZYRTTBKwfc+MEGGsrVH0qC" +
		"pRK/HjrFeW2clZ2dHYjqjeDO/4GCiA63GRnA+aZjfYz4+LIQ0NBx/JYd4hMXFoRM" +
		"mQ2m0Yhzuc0CgYA0OzaPF3CohuXONNU8hronUdjSOLWPJfhxxRrazNfOBdRM0YwO" +
		"dxkdfcV+LgrkWY7PzUPFJMj26S7m+aV/jrVXqEZ0Y2kCLroWBmclRwvuVZrW/gl4" +
		"I5d2uZtJ80OqIOCsheb/0ZHDymSdsAokrNhOxu+mlB2jGGWJmXqUzgMdFQKBgHRR" +
		"GmnrHChf7I40wVp5Glncfg9O+NbykUC8EnzlGVErLteSmOaR6Dz3AzV1fb+VpDtw" +
		"uMaBqcc+gQobk2srewMkh5Fl17iAjt3NjBB0xswZYsnymFp7/Pc5sFYFCLOVeFaQ" +
		"GJnbY0OthzA4PNNvJUlskY4l2XMIGR85vvcU1+VJAoGAe3nHkGPnfeXZXLOAAsIo" +
		"AFO9og2s+1VcxBAlEv4jnq9fuToTK9299N//a0+4SGh+AGckN0v0HOR5BdUhrNwC" +
		"+Rm0XQl47n5fEivj87r1xeVXe8M9gYaglj1QJZRfux8HLpfvw9Z+4LippRS7kIVw" +
		"F3lbxXBfkJOPBa7zh4LdDaM="

	dkDataCBC := "HykjzxldgwfSIYc6nyvlUxkqmCR8TEzeyYVxQmAXvk6ZzHtu38ReQml6F9oMI2cc"
	pkb, err := base64.StdEncoding.DecodeString(priKeyB64)
	if nil != err {
		panic(err)
	}
	pbb, err := base64.StdEncoding.DecodeString(pubKeyB64)
	if nil != err {
		panic(err)
	}

	rsa := NewRSA(RSAPemPKCS8)
	sign, err := rsa.RsaSignWithSha256([]byte(dkDataCBC), pkb)

	if nil != err {
		println(err)
		panic(err)
	}
	fmt.Println(fmt.Sprintf("签名=%s", base64.StdEncoding.EncodeToString(sign)))
	fmt.Println(fmt.Sprintf("签名=%s", base64.URLEncoding.EncodeToString(sign)))

	isOk, err := rsa.RsaVerySignWithSha256([]byte(dkDataCBC), sign, pbb)

	if nil != err {
		panic(err)
	}
	println(isOk)

	dkSign := "AigKIsG1TtsC/FIjOonLnuCSiaQOyYhE/bFbt+azCuQPcRaa91Yh9tHmjvIndpLJOU/NpF7U8uRqTyiwCHjFseMpbzS7iZuPG3c4liF4ZJkiRbW8UAaNYDA7lhy/+J2efWqVyXz/wKaTyeHTxA6dETvKG5YJ805mSmqy4Rmae1ySrc7uS8B6nsvUc3kco6hMPbFeb3L5bGp2yk0gVYsAGDvjEqeFVlVyublfKa7+RK8BCgxD2M5tHUChGSk60yyOfPFDPWeogeAWs6Ha1U3p2FjuQBPHltzOIBGu3hZbao1G6VudzLvcwXZlQbc+FDOnBcc7FER/ntMfQh5y9VecCQ=="
	fmt.Println(fmt.Sprintf("密钥比较:%v", dkSign == base64.StdEncoding.EncodeToString(sign)))
}

func TestRsaSignMd5(t *testing.T) {
	//pubKeyB64 := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF3Z3RLYk1SSVhmRmVzK1lPeDMwbApnUTQ0czJZMktIMDlyNzdFYlBmSlJZOW9Ka2lrR0t3TVVBQnVRVHhXa3M3WjkvZFlIVVZRNkx2UThZQS9kbE51Ck4vU2VEb1pEMFBGYkdGSmJaTlo5SHVpSG93Q1c3MzI3Nk5pQmtVdnhEcFNMRU9HMWxUa0RiT01uUi9TS0s2NjAKcjl3Yy9STmg0bWxYMUgzZjhMZzJtUGJYa2xrQzBLemsxQUFIVWJmNVpwRG1Pd0QyR2dOcnMyL2lTTUM5MmJOSApYVXp2T0d2VjBCb01sd2V5VWM1Z093OHlyUkRtOStzMFVjU3B1RDArOEtqMnRneWppNUdodHZqb2hrTlQ0VUVTCndudnBuZFZZZ2ZTOHNWVkNsWWZLZFV5RjNhVHUxTEQ4dEtHb0swbnhjQkdQaTZRalFadFQvcWFkV2xHTExpaFYKUVFJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg=="
	//priKeyB64 := "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFdlFJQkFEQU5CZ2txaGtpRzl3MEJBUUVGQUFTQ0JLY3dnZ1NqQWdFQUFvSUJBUURDQzBwc3hFaGQ4VjZ6CjVnN0hmU1dCRGppelpqWW9mVDJ2dnNSczk4bEZqMmdtU0tRWXJBeFFBRzVCUEZhU3p0bjM5MWdkUlZEb3U5RHgKZ0Q5MlUyNDM5SjRPaGtQUThWc1lVbHRrMW4wZTZJZWpBSmJ2ZmJ2bzJJR1JTL0VPbElzUTRiV1ZPUU5zNHlkSAo5SW9ycnJTdjNCejlFMkhpYVZmVWZkL3d1RGFZOXRlU1dRTFFyT1RVQUFkUnQvbG1rT1k3QVBZYUEydXpiK0pJCndMM1pzMGRkVE84NGE5WFFHZ3lYQjdKUnptQTdEekt0RU9iMzZ6UlJ4S200UFQ3d3FQYTJES09Ma2FHMitPaUcKUTFQaFFSTENlK21kMVZpQjlMeXhWVUtWaDhwMVRJWGRwTzdVc1B5MG9hZ3JTZkZ3RVkrTHBDTkJtMVArcHAxYQpVWXN1S0ZWQkFnTUJBQUVDZ2dFQkFJajJVemZtYTNYemtuYkVZWllwSFRtMGdnME9qaGVTSHVKWGNtbS9sQTlICkk2b3lCN0ZxYnQ1aEQzRjRWMXNVS2dHK1VqR0c1WThBVW9ERGx3ZTc1OFlUSVNUN1hBNjA3U21EcUFMSzZsSFEKcXp4QWhFalNwTG03WitqWWczTlpJYmR1dVM1MHFaaEgxVWdTc1J3WUdtMHVuajk3V05Ib3JSZk5LUzNOdUt2SwpCSXRibmZDNXA5ZktyaWVhUFM3dDRmRjhBbFNjSkJlOXFzNUxGMno5NVk5bEovUDRrWWxnVUZHa0gvQ2hJK0VDCkFqbmVwQ1lqbiszQUg2UjI2b3lMd1FSVWZVbURtOUMxME5vVlNCcXRQNDZwSnJaKzE5UDh2cEpnTUFrc01BNUwKMWYvU1R6NFJ1cFNQT1I4VTZKSEtTUzl5ekMrblkxcW83M0VKUmVGdXJXRUNnWUVBNTdaTDUrbXRFREhnM0p1cwpYQXp3YkJBbktoNVNZUkdETVZQZGpHa3liTC81cXpsdEJESS9IMEpuYytld1ZCWnhtT1R4TU53ZEczZ0l2SXhOCmxjY0E0OTVaWk42SEJSS2lHNCs2bkFob2U0d09KcGJmbllWV3psWW8ycGl5R0pUd1JiVDFJVU9YeGRqUEhkV2MKVFVUNEFUd05WYmpaNmlxSlBySS9rVDV2UlVVQ2dZRUExbUk0aldGbFBGaktmb1RtVWkydCtjelRTaVhlMXFmNgo5SS8wcVJsWUFqaVdrMjFyNW5LVUtGNlM3azhFSElqUmJVOU1BYjdWQm9aWVJUVEJLd2ZjK01FR0dzclZIMHFDCnBSSy9IanJGZVcyY2xaMmRIWWpxamVETy80R0NpQTYzR1JuQSthWmpmWXo0K0xJUTBOQngvSllkNGhNWEZvUk0KbVEybTBZaHp1YzBDZ1lBME96YVBGM0NvaHVYT05OVThocm9uVWRqU09MV1BKZmh4eFJyYXpOZk9CZFJNMFl3TwpkeGtkZmNWK0xncmtXWTdQelVQRkpNajI2UzdtK2FWL2pyVlhxRVowWTJrQ0xyb1dCbWNsUnd2dVZaclcvZ2w0Ckk1ZDJ1WnRKODBPcUlPQ3NoZWIvMFpIRHltU2RzQW9rck5oT3h1K21sQjJqR0dXSm1YcVV6Z01kRlFLQmdIUlIKR21uckhDaGY3STQwd1ZwNUdsbmNmZzlPK05ieWtVQzhFbnpsR1ZFckx0ZVNtT2FSNkR6M0F6VjFmYitWcER0dwp1TWFCcWNjK2dRb2JrMnNyZXdNa2g1RmwxN2lBanQzTmpCQjB4c3daWXNueW1GcDcvUGM1c0ZZRkNMT1ZlRmFRCkdKbmJZME90aHpBNFBOTnZKVWxza1k0bDJYTUlHUjg1dnZjVTErVkpBb0dBZTNuSGtHUG5mZVhaWExPQUFzSW8KQUZPOW9nMnMrMVZjeEJBbEV2NGpucTlmdVRvVEs5Mjk5Ti8vYTArNFNHaCtBR2NrTjB2MEhPUjVCZFVock53QworUm0wWFFsNDduNWZFaXZqODdyMXhlVlhlOE05Z1lhZ2xqMVFKWlJmdXg4SExwZnZ3OVorNExpcHBSUzdrSVZ3CkYzbGJ4WEJma0pPUEJhN3poNExkRGFNPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ=="

	pubKeyB64 := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwgtKbMRIXfFes+YOx30l" +
		"gQ44s2Y2KH09r77EbPfJRY9oJkikGKwMUABuQTxWks7Z9/dYHUVQ6LvQ8YA/dlNu" +
		"N/SeDoZD0PFbGFJbZNZ9HuiHowCW73276NiBkUvxDpSLEOG1lTkDbOMnR/SKK660" +
		"r9wc/RNh4mlX1H3f8Lg2mPbXklkC0Kzk1AAHUbf5ZpDmOwD2GgNrs2/iSMC92bNH" +
		"XUzvOGvV0BoMlweyUc5gOw8yrRDm9+s0UcSpuD0+8Kj2tgyji5GhtvjohkNT4UES" +
		"wnvpndVYgfS8sVVClYfKdUyF3aTu1LD8tKGoK0nxcBGPi6QjQZtT/qadWlGLLihV" +
		"QQIDAQAB"

	priKeyB64 := "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDCC0psxEhd8V6z" +
		"5g7HfSWBDjizZjYofT2vvsRs98lFj2gmSKQYrAxQAG5BPFaSztn391gdRVDou9Dx" +
		"gD92U2439J4OhkPQ8VsYUltk1n0e6IejAJbvfbvo2IGRS/EOlIsQ4bWVOQNs4ydH" +
		"9IorrrSv3Bz9E2HiaVfUfd/wuDaY9teSWQLQrOTUAAdRt/lmkOY7APYaA2uzb+JI" +
		"wL3Zs0ddTO84a9XQGgyXB7JRzmA7DzKtEOb36zRRxKm4PT7wqPa2DKOLkaG2+OiG" +
		"Q1PhQRLCe+md1ViB9LyxVUKVh8p1TIXdpO7UsPy0oagrSfFwEY+LpCNBm1P+pp1a" +
		"UYsuKFVBAgMBAAECggEBAIj2Uzfma3XzknbEYZYpHTm0gg0OjheSHuJXcmm/lA9H" +
		"I6oyB7Fqbt5hD3F4V1sUKgG+UjGG5Y8AUoDDlwe758YTIST7XA607SmDqALK6lHQ" +
		"qzxAhEjSpLm7Z+jYg3NZIbduuS50qZhH1UgSsRwYGm0unj97WNHorRfNKS3NuKvK" +
		"BItbnfC5p9fKrieaPS7t4fF8AlScJBe9qs5LF2z95Y9lJ/P4kYlgUFGkH/ChI+EC" +
		"AjnepCYjn+3AH6R26oyLwQRUfUmDm9C10NoVSBqtP46pJrZ+19P8vpJgMAksMA5L" +
		"1f/STz4RupSPOR8U6JHKSS9yzC+nY1qo73EJReFurWECgYEA57ZL5+mtEDHg3Jus" +
		"XAzwbBAnKh5SYRGDMVPdjGkybL/5qzltBDI/H0Jnc+ewVBZxmOTxMNwdG3gIvIxN" +
		"lccA495ZZN6HBRKiG4+6nAhoe4wOJpbfnYVWzlYo2piyGJTwRbT1IUOXxdjPHdWc" +
		"TUT4ATwNVbjZ6iqJPrI/kT5vRUUCgYEA1mI4jWFlPFjKfoTmUi2t+czTSiXe1qf6" +
		"9I/0qRlYAjiWk21r5nKUKF6S7k8EHIjRbU9MAb7VBoZYRTTBKwfc+MEGGsrVH0qC" +
		"pRK/HjrFeW2clZ2dHYjqjeDO/4GCiA63GRnA+aZjfYz4+LIQ0NBx/JYd4hMXFoRM" +
		"mQ2m0Yhzuc0CgYA0OzaPF3CohuXONNU8hronUdjSOLWPJfhxxRrazNfOBdRM0YwO" +
		"dxkdfcV+LgrkWY7PzUPFJMj26S7m+aV/jrVXqEZ0Y2kCLroWBmclRwvuVZrW/gl4" +
		"I5d2uZtJ80OqIOCsheb/0ZHDymSdsAokrNhOxu+mlB2jGGWJmXqUzgMdFQKBgHRR" +
		"GmnrHChf7I40wVp5Glncfg9O+NbykUC8EnzlGVErLteSmOaR6Dz3AzV1fb+VpDtw" +
		"uMaBqcc+gQobk2srewMkh5Fl17iAjt3NjBB0xswZYsnymFp7/Pc5sFYFCLOVeFaQ" +
		"GJnbY0OthzA4PNNvJUlskY4l2XMIGR85vvcU1+VJAoGAe3nHkGPnfeXZXLOAAsIo" +
		"AFO9og2s+1VcxBAlEv4jnq9fuToTK9299N//a0+4SGh+AGckN0v0HOR5BdUhrNwC" +
		"+Rm0XQl47n5fEivj87r1xeVXe8M9gYaglj1QJZRfux8HLpfvw9Z+4LippRS7kIVw" +
		"F3lbxXBfkJOPBa7zh4LdDaM="

	dkDataCBC := "HykjzxldgwfSIYc6nyvlUxkqmCR8TEzeyYVxQmAXvk6ZzHtu38ReQml6F9oMI2cc"
	pkb, err := base64.StdEncoding.DecodeString(priKeyB64)
	if nil != err {
		panic(err)
	}
	pbb, err := base64.StdEncoding.DecodeString(pubKeyB64)
	if nil != err {
		panic(err)
	}

	rsa := NewRSA(RSAPemPKCS8)
	dataSha256 := GetInstanceByMessageDigest().Sha256Buf(dkDataCBC)

	//sign, err := rsa.RsaSignWithMD5([]byte(dkDataCBC), pkb)
	sign, err := rsa.RsaSignWithMD5(dataSha256, pkb)

	if nil != err {
		println(err)
		panic(err)
	}
	fmt.Println(fmt.Sprintf("签名=%s", base64.StdEncoding.EncodeToString(sign)))
	fmt.Println(fmt.Sprintf("签名=%s", base64.URLEncoding.EncodeToString(sign)))

	//isOk, err := rsa.RsaVerySignWithMD5([]byte(dkDataCBC), sign, pbb)
	isOk, err := rsa.RsaVerySignWithMD5(dataSha256, sign, pbb)

	if nil != err {
		panic(err)
	}
	println(isOk)

	dkSign := "AigKIsG1TtsC/FIjOonLnuCSiaQOyYhE/bFbt+azCuQPcRaa91Yh9tHmjvIndpLJOU/NpF7U8uRqTyiwCHjFseMpbzS7iZuPG3c4liF4ZJkiRbW8UAaNYDA7lhy/+J2efWqVyXz/wKaTyeHTxA6dETvKG5YJ805mSmqy4Rmae1ySrc7uS8B6nsvUc3kco6hMPbFeb3L5bGp2yk0gVYsAGDvjEqeFVlVyublfKa7+RK8BCgxD2M5tHUChGSk60yyOfPFDPWeogeAWs6Ha1U3p2FjuQBPHltzOIBGu3hZbao1G6VudzLvcwXZlQbc+FDOnBcc7FER/ntMfQh5y9VecCQ=="
	dkSign256 := "etN2LIpbyy9soUh7ge042kWLZz6nNmv95pH3eRoLluJzAxM0zUTRCLOSTz206L9NHoyhP0psFV0DwNunjpqe3eeDsz/yx4c1h4dWt2urYoKyGZbXjJbJxWrf8rCWz6+r82KqrHorSG1c2XWnQY+CotYBhcnhz98uC+qSiy/KS46zMyCIYgr0vfR1xvW4UgIN8jbKTghDuBemxBQRQvvgY2poOZLo7AHB4S4LrtkKglgrlmJpl48XGeQL/uNxvLnH/CtzJx4DAwz+Sx7FAtp/2M1772C6mSITBnfMmjE2/RkNiCyXOfQDuBfYYBEK3bJRQsSUyg0wA63Bpa3LEgrN5g=="

	fmt.Println(fmt.Sprintf("密钥比较:%v", dkSign == base64.StdEncoding.EncodeToString(sign)))
	fmt.Println(fmt.Sprintf("sha256密钥比较:%v", dkSign256 == base64.StdEncoding.EncodeToString(sign)))

}

func TestRSA1024(t *testing.T) {
	rsa := NewRSA(RSAPemPKCS8)
	//rsa 密钥文件产生
	fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
	//prvKey, pubKey, err1 := rsa.GenRsaKey1024()
	prvKey, pubKey, err1 := rsa.GenRsaKey2048()
	if nil != err1 {
		t.Fatal(err1)
	}
	println("私钥------------------------------")
	fmt.Println(string(prvKey))
	println("公钥------------------------------")
	fmt.Println(string(pubKey))
	println("------------------------------")

	fmt.Println(base64.URLEncoding.EncodeToString(prvKey))
	fmt.Println(base64.URLEncoding.EncodeToString(pubKey))

	buf, _ := base64.URLEncoding.DecodeString(base64.URLEncoding.EncodeToString(prvKey))
	fmt.Println(string(buf))

	fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	var data = "卧了个槽，这么神奇的吗？？！！！  ԅ(¯﹃¯ԅ) ！！！！！！）"
	fmt.Println("对消息进行签名操作...")
	signData, err2 := rsa.RsaSignWithSha256([]byte(data), prvKey)
	if nil != err2 {
		t.Fatal(err2)
	}

	t.Log(string(signData))
	fmt.Println(fmt.Sprintf("base64.URLEncoding=%s", base64.URLEncoding.EncodeToString(signData)))
	t.Log(fmt.Sprintf("消息的签名信息： %s", hex.EncodeToString(signData)))
	fmt.Println("\n对签名信息进行验证...")

	signData2, _ := hex.DecodeString(hex.EncodeToString(signData))

	isOk, err3 := rsa.RsaVerySignWithSha256([]byte(data), signData2, pubKey)
	if nil != err3 {
		t.Fatal(err3)
	}
	if isOk {
		fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
	}

	fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
	ciphertext, err4 := rsa.RsaEncrypt([]byte(data), pubKey)
	if nil != err4 {
		t.Fatal(err4)
	}
	fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
	sourceData := rsa.RsaDecrypt(ciphertext, prvKey)
	fmt.Println("私钥解密后的数据：", string(sourceData))
}

//func TestRSA10242(t *testing.T) {
//	rsa := GetInstanceByRSA1024()
//
//	priKeyStr := "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCHLvQfTGLj3W7sWiAHDagSbKYiCdfIxWXXC2IVp6R25NSr1cDj8tIssWfYQ+ruA/5DaJXTfwgOA0PPkY0Uob2/PRuchI0Xr4XwqQwAkLcVcjfNZmMVSrC05mD76yL+1BuDnJ31kcivQMwoXSQqzYmZl9lIgDDh5K2PFdJUz1UZUmWNpC2ZHuP3h0vwSpUhUFhJJCuA1hsYblcfIRy6oSpDeeOjgtO2sfSnN8AbYjM9n1Q218qidXv+QUvu+6xuo0nk2+FqbAnjWPvz7cgskLAE62tH1sbEZlBfxe2/8ROxnaIg+7dgKmfqYeQez1HnBGeD9NndtqRWEfosDX+DZnopAgMBAAECggEAY/CnlFgBqBp1vhCnKu/CuNRQQkvqlsixELmeqwnEQg3M2LjvoNZM4bPKZQ1ZKtwS5zzzv2djyhBJ2rPtjDpDMJX5ys4IDWG7cP9ZGzXh1N4bOSQfzoboeuTzAGuG9MRVDwkDkqBTsJUEGjc53NcVilLD1aDIAsjwMx9b301kyZGV1Blf9P7C/JaCxCocHX+nhEgZq7kD9PuvpiKwJ6MDF29BtSE1/B2+8BG3DTGww5MHonv/0bMIt7N/un3FSTQpkQ2qFRB+4qBj1qZgUNt9yLoe+uJtHA/o3ddalXwr495gotRAJ7HFO3hxODjRk7FnJWN+DJF9p7TUM5VeS2nTsQKBgQDYBM4IEoKS2SF1jdHFaYxQV71lG8ERdQlI+eKcm07HNCUK1GtF1zu4NHU7lPHmNPGCuzK/8j+4LkczpfKu0w9o3ky6pI6jPiopXeSlJVUrXWI8mnA0HXBb6sNeUoUlu+ewoawNJGvmnKNbzXWE7tbTWoeAsJLUq6P+F97DfwXEzQKBgQCgNBXHGrbBrOFiUdAnkmEsUk5sNrrjmIBtPKo6YvJmTT0TKPO++6mOakVR7dTNW4HV9J85CdaHjE1j04SZM7FTgu5DLfCc2vFnDV84B5kcLHKTMbfLM3AmshlN5mSYAoRXuvWYEX8FnzBS2pmEibZKLkN9GDKheYVwgLODoXdqzQKBgQCy94gad/tl3i4yTkS04TU2evqWgd/6rpP6ucxdIu6pazIlPseBHUiE3DEkI8olh0dvn9fz3qeb1/t1ds8QuBvULhgzqZHi/OXBT+DWUY+2Va/Ftc2v35PvExi5VHSrRno1hDwex0X90Vgl/pqWf6nLgP0ySRfcyjcblHsiTGJjIQKBgGz3AFkMsoHJNQPK4eoIhk+/K9gu4a8say3htWdBJd6vans9v4yHYCyd28h+G+AR/Z2pZSNGrcREid78X5RUtKg4xhariJ0nzkpprfpOMLYZBVVY28o6km2/dbamnoVGMP37DFEClYMdY6D3TrP3dyW9kenkK4vpO/npkDBYAwGBAoGAQWOPDJMiNchkUw6FqONKHBMJUkWWmOdHuwF2b7mcti4VleYP3yKujTZQecisozDWC3vNEBKJoTQw7/wzIO/imV73Hd0SsQhST56JfsK4Q3GAkeFdc+4Xt2I7blCBrNrlWeAibbF1ULtfijrxT9N5PZzYtME280lEkEPWTdmln9k="
//	pubKeyStr := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAhy70H0xi491u7FogBw2oEmymIgnXyMVl1wtiFaekduTUq9XA4/LSLLFn2EPq7gP+Q2iV038IDgNDz5GNFKG9vz0bnISNF6+F8KkMAJC3FXI3zWZjFUqwtOZg++si/tQbg5yd9ZHIr0DMKF0kKs2JmZfZSIAw4eStjxXSVM9VGVJljaQtmR7j94dL8EqVIVBYSSQrgNYbGG5XHyEcuqEqQ3njo4LTtrH0pzfAG2IzPZ9UNtfKonV7/kFL7vusbqNJ5NvhamwJ41j78+3ILJCwBOtrR9bGxGZQX8Xtv/ETsZ2iIPu3YCpn6mHkHs9R5wRng/TZ3bakVhH6LA1/g2Z6KQIDAQAB"
//
//	priKey, err11 := base64.StdEncoding.DecodeString(priKeyStr)
//	if nil != err11 {
//		t.Fatal(err11)
//	}
//	//fmt.Println(string(priKey))
//	pubKey, err11 := base64.StdEncoding.DecodeString(pubKeyStr)
//	if nil != err11 {
//		t.Fatal(err11)
//	}
//	//fmt.Println(string(pubKey))
//	t.Log("-------------------------")
//
//	//priKey := "-----BEGIN RSA PRIVATE KEY----- MIICXAIBAAKBgQDK4M7qf8rcH5VbQimcecdV6ZzkPaZsNWzyIfaLqmCDfS7kan205w2tklNVdO7MWZJpS97Rb67cS2pwyMvhbXYTjGEXyuhpPrXGgw8Q1hEoGqHQu3ykML8TXPMAe9HTFnUIuKIYMxb/SIfWZqyB/qywtGlWt18fYts1VMQLX2yHEwIDAQABAoGAeX9Ia4c8pbcEazKkWOFVT04od0e0cvlL1XYhgGL4icZeXsynm78Dof8PiQ4ONLMvy390YVjRD3zasdCOyOIU426sotwotN2wegKBGUwwCNU4+lP0VamOad2f51npwzQaLahOduoJOaKPAb+YKxHwNo2Px+pljsUqEYVRCaDl8IECQQDWxDMYqJqhRSFhijme4tbhITyHGgeaNw+u6SVicMfpLfjdf7kxR/dUhNVxRdfJdNpv7QsMzJ55bWPw3QrBbHBfAkEA8dRMsBYo82ZDUwS2JobIlJk2FxF1rLvcJwjiVyd0Yz/0LxENVy0Ji81VZYaHs9cJu4VxlnJnlhcnY3fdHxlVzQJALKO/UyLIcTjjRVjrvSC9NTIpWJOKfP1w3xRK1vlGNCuADNodbibdO84YZ2DzB0aomJcWsuRdFDQuj8QCFk4p1QJBAMgZdustKu1b7NFA0MfINyheLhegZtJrD5ttCnw7NV76iD55yaQcrA119fdv/dGdWXxEytxGBdh3iCwR/nHBMPECQEmieRk0RFoIDaI6tOEfwEX6jCmtOQGL5NhfCcGFtsujm6hmtLZHkZGg65E+eV1QK+8sI6YmDxcCefPOXzCIr6c= -----END RSA PRIVATE KEY-----"
//	//pubKey := "-----BEGIN PUBLIC KEY----- MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDK4M7qf8rcH5VbQimcecdV6ZzkPaZsNWzyIfaLqmCDfS7kan205w2tklNVdO7MWZJpS97Rb67cS2pwyMvhbXYTjGEXyuhpPrXGgw8Q1hEoGqHQu3ykML8TXPMAe9HTFnUIuKIYMxb/SIfWZqyB/qywtGlWt18fYts1VMQLX2yHEwIDAQAB -----END PUBLIC KEY-----"
//
//	//	priKey := `
//	//-----BEGIN RSA PRIVATE KEY-----
//	//MIICXAIBAAKBgQD3UqKEyJR8tfzRUSDozI2p11RYQQQs7caWyIAtsSNgu750Q4sv
//	//r5AlpdVF4zuZayL/hELWQ+EjT/kk8PnvMj4Doi3Heaw8+O4lP4OM0l6AsNfTQEML
//	//Ei3T5EA/28FHCIcHn5Nc/4jB5RQNjq6pNUsmZevenbVA0TC7M2QPbVv9LQIDAQAB
//	//AoGAN9RVj3ff3Q8P1QhlT2ftiqtrBMkYcjPyolL8bFQSUmHPKluc7dTJy1XWAQK8
//	//j3NZ4SgwFkIYbmo9KZOkN9S1npbAzYhD4U4yXAi27zlpLPtbLndoOA4HK3vrk9nF
//	//AiDTmwTVrnJMvX5u7wR+0JqtAw0junZVlRXxDAhT1+uvUZkCQQD9KddoJyud8oFX
//	//oIT/uT2p50X39oV+vKrwf3MSpIprM/477NWhcr7VPHIOS1gNB6Cb4x+DmyayT/Iq
//	//DlFVvDzfAkEA+hgKWj4gpq2GS/4qLue8+owEU25X/fsVctosCjI7AhMeAiVY43J0
//	//hfQ8/sgUDBlW5FK/Iq8jBfUit95Ny6L7cwJATcmSd105yLFfzrXyx8R6Tv9R/2vO
//	///u8nsvfmOr82DNSP9IfD6HSicFC/VucNqgtC7UMvRrfgfv+TkBqQIUDSjwJBAIWF
//	//Qjmtw6bZK8L0njbOmDE3gbO9TJMXcvsPicWjzacs56+DmvJLj/RYUhxAW5ueB6r8
//	//lnkBAfTTTEbYE7atfAsCQHdVtMzrmK/+fdR3uD3t15eKDvTi8BG+e95bBOrsGzlW
//	//fTpZMcUalnNI5LDKgyyLq/TklVLKo9tOFlkRzRr5HoU=
//	//-----END RSA PRIVATE KEY-----
//	//`
//	//	pubKey := `
//	//-----BEGIN PUBLIC KEY-----
//	//MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQD3UqKEyJR8tfzRUSDozI2p11RY
//	//QQQs7caWyIAtsSNgu750Q4svr5AlpdVF4zuZayL/hELWQ+EjT/kk8PnvMj4Doi3H
//	//eaw8+O4lP4OM0l6AsNfTQEMLEi3T5EA/28FHCIcHn5Nc/4jB5RQNjq6pNUsmZeve
//	//nbVA0TC7M2QPbVv9LQIDAQAB
//	//-----END PUBLIC KEY-----
//	//`
//
//	fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
//	var data = "卧了个槽，这么神奇的吗？？！！！  ԅ(¯﹃¯ԅ) ！！！！！！）"
//	fmt.Println("对消息进行签名操作...")
//	signData, err2 := rsa.RsaSignWithSha256([]byte(data), priKey)
//	if nil != err2 {
//		t.Fatal(err2)
//	}
//
//	t.Log(string(signData))
//	t.Log(fmt.Sprintf("消息的签名信息： %s", hex.EncodeToString(signData)))
//	fmt.Println("\n对签名信息进行验证...")
//
//	signData2, _ := hex.DecodeString(hex.EncodeToString(signData))
//
//	isOk, err3 := rsa.RsaVerySignWithSha256([]byte(data), signData2, pubKey)
//	if nil != err3 {
//		t.Fatal(err3)
//	}
//	if isOk {
//		fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
//	}
//
//	fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
//	ciphertext, err4 := rsa.RsaEncrypt([]byte(data), pubKey)
//	if nil != err4 {
//		t.Fatal(err4)
//	}
//	fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
//	sourceData := rsa.RsaDecrypt(ciphertext, priKey)
//	fmt.Println("私钥解密后的数据：", string(sourceData))
//
//}

func TestGenerateKey(t *testing.T) {
	//privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	buf, _ := base64.URLEncoding.DecodeString("LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEZVRNU0FsbnFDMEdxa3pQZ3BTbm1iM25aMwpRakVNY0VhSUFZdVpGT29ydkk3Ujk1b1UrV1NvbGtvTmd0OHd2TmJ2eWZoc2gwejhkRTJpZkkvYkQzUm84eVZJCnprUjB0b2diT1BJMysvaUtiNEExWUZmRmpWaURrWTczQS91UnNvb1pGRmxPTTVYbE83bG5sM0N5b0o0a3kyWUUKbTh4RVlxc3pOWTVvSHdZdEhRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQ==")
	t.Log(string(buf))

}

func TestDKSign(t *testing.T) {

	pubKeyB64 := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDK1dZQ1pSdkg0SVVQdnFCYkgrWWVib2tIcwowM3lpQVZLK0U1U2R2Q0lndkQ1MHI1Tlhra0kzdG9LU2ZDVTFTamdkOWZDU01aRUFJdVJrUkNXdVJUT3Y4VnN5CnNxd1lSZUM4ZFdEMUFtOERhN0tVcUJuRFZ4end4d2FsTlUwWGd6MEVVeWZTb251aXIrS2J5NjFyVlJEUUNoa24KK3RORVo5S08ySW5wVjlNK3N3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	priKeyB64 := "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDZGdJQkFEQU5CZ2txaGtpRzl3MEJBUUVGQUFTQ0FtQXdnZ0pjQWdFQUFvR0JBTDVaZ0psRzhmZ2hRKytvCkZzZjVoNXVpUWV6VGZLSUJVcjRUbEoyOElpQzhQblN2azFlU1FqZTJncEo4SlRWS09CMzE4Skl4a1FBaTVHUkUKSmE1Rk02L3hXekt5ckJoRjRMeDFZUFVDYndOcnNwU29HY05YSFBESEJxVTFUUmVEUFFSVEo5S2llNkt2NHB2TApyV3RWRU5BS0dTZjYwMFJuMG83WWllbFgwejZ6QWdNQkFBRUNnWUJKenVLeWpIUGV4dWRVMGxTakRmcXJPbXRkCnJWT3liZGpyb3lRSlZaM1dHNmdNRHRpUEtFTk0zeFFhUU5FY3JMNjl2MU9kSEdNaExtWnBDcE9oMDJ1S3Jkdk8KbzQ1ZFkveHlhNWpmSUk5MkNKSDR2TVI2L3VWOXRIQkdmOEh5aTN4ckIvUjJOczV1YkNRMThSekR1MmRxVFQrdQp0UVNDNEJCMjQ0NktHRmFCSVFKQkFPM2pXdlcvTFIxcWJFZFczMFNwbEk4cEFjWitER1hJNlJDVWJJOFFmcWlhCkRtQ3hZdW1VU3J1YjkzTGl0UjZ3bUoyc0tVSXFLd1dEUTIzKzUyaW1taThDUVFETTE1U3hKaUVYTW00OWV0aUUKY2NDcGNTNTUvc091dHRSbGd4TmdFN2E5T3lJUlFQdHZoSy9RVGpKV0d1Ump6RVhpRWVMbys3VDE5ZTNFcTZ2QQpxcmE5QWtBS1ZKczRuTXE2d2twZGRycFBZd1hlaWF0WUVWVThmbE1RczBGYm5SM0MrSjJ6T1VEUVgrNDI0M2tGCmRpN2pYRXZrWFB1VnNmc1lUREQ3Yjl4Z3dRdGxBa0E0MVN0TkJ4NHhPRzI3b3dURm9tWG8zUjBlL2Q4Kzd6ejQKdVNnOEJOd3pubDl5V0F3cXdhNmg1Y0F6Z1p5U1Q4K256SHlmVlk2OG16SVAyZTE2TkNNbEFrRUF2UlVUc0tadApBbVY4aWFUM1V0OGJvYXQ1U3ZDMTFBS2JMWHpqOTl5QzViWVp2UGZMb2pvYmt0dFA4bVVsOWtsMVg2NE9ncFp1Ck1LRFJ1cXJOc3NOUkx3PT0KLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K"

	println("----------------------- aes begin ---------------------")
	data := `{"account":"kkkqqq3","nickname":"123"}`
	println(fmt.Sprintf("数据=%s", data))
	secretKey := "D2434fgdgfdgfg12"
	println(fmt.Sprintf("对称密钥=%s", secretKey))

	iv := []byte(secretKey)
	cbc, err := NewAESCryptoCBC([]byte(secretKey))
	if nil != err {
		t.Fatal(err)
	}
	dkDataCBC := "HykjzxldgwfSIYc6nyvlUxkqmCR8TEzeyYVxQmAXvk6ZzHtu38ReQml6F9oMI2cc"
	dataCBC := cbc.EncryptWithIVToBase64URLEncode([]byte(data), iv)
	println(fmt.Sprintf("数据对称加密后=%s，与文档例子比较=%v", dataCBC, dataCBC == dkDataCBC))

	println("----------------------- aes end ---------------------")
	println("----------------------- rsa sign begin ---------------------")

	dataSha := messageDigestInstance.Sha256(dataCBC)
	dkDataSha := "0371f0109423430660fcc13498618983ff238c277f25690144b50568fd3194a8"
	fmt.Println(fmt.Sprintf("加密后的数据sha256后=%s，与文档例子比较=%v", dataSha, dataSha == dkDataSha))

	fmt.Println(fmt.Sprintf("base64后的私钥=%s", priKeyB64))
	fmt.Println(fmt.Sprintf("base64后的公钥=%s", pubKeyB64))

	pkb, err := base64.StdEncoding.DecodeString(priKeyB64)
	if nil != err {
		t.Fatal(err)
	}
	rsa := NewRSA(RSAPemPKCS8)
	si, err := rsa.RsaSignWithSha256([]byte(dataSha), pkb)
	if nil != err {
		t.Fatal(err)
	}
	sis := base64.StdEncoding.EncodeToString(si)
	fmt.Println(fmt.Sprintf("签名base64后=%s", sis))

	pbb, err := base64.StdEncoding.DecodeString(pubKeyB64)
	if nil != err {
		t.Fatal(err)
	}

	sisB, _ := base64.StdEncoding.DecodeString(sis)
	isOk, err := rsa.RsaVerySignWithSha256([]byte(dataSha), sisB, pbb)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("验签结果=%v", isOk))
	println("----------------------- rsa sign end ---------------------")

}

func TestVSign(t *testing.T) {
	dataStr := "uMVBQOuQNOKU4uDK8%2B4Cdp6In%2FrMsFObmv7kNLMqDNc6HNEbmbTtWSNK1IQlYvow"
	signStr := "YQ8WNkRyJWTzFYdU4AgYWwKC1K5xXOInmFIk1KpzDlWse9iVFkh4aIA41kWAnRv2YTmyEGZCmqnw\naX3GCcmEvINRq2u4f8C4fpiWLNK3r1MKzmiaFNQ+vgyZcZXrQ988BwSUYKedn47hC3bgsmeMT+k2\n/7jcY1bLjpY9orRvc2H3IPRDyho0dULvFcmlygckzxz3ZdPiabn6J6Me0v7tDp4BXUdRWwlf4vnA\nGWvzkHdc5UbqPfoOyr+/ftOZyo04suJm1irHXrh66t7lAEKFXIB13NrxAh4MEzR420NjqKb0IJ0c\n6Ll+zBObR5mGULZ7PElmbe85Oah9UZEZUNpxbw==\n"

	//u, err0 := url.QueryUnescape(signStr)
	//if nil != err0 {
	//	t.Fatal(err0)
	//}
	//t.Log(u)

	signBuf, err1 := base64.StdEncoding.DecodeString(signStr)
	if nil != err1 {
		t.Fatal(err1)
	}

	t.Log(dataStr)
	t.Log(string(signBuf))

}
