package hash

import (
	"os"
)

func HashFunction(plainName, plainPassword string) string{
   key1 := os.Getenv("KEY1")
   key2 := os.Getenv("KEY2")
	plainText := plainName + key1 + plainPassword + key2	
	//plainText := "revpxqnq@?=vnzrevpxvnzo@D!"
	hash := make([]byte, len(plainText))

	for i := 0; i < len(plainText); i++{
		char := plainText[i]
		switch {
			case char >= 'a' && char <= 'z' :
			   hash[i] = 'a' + (char - 'a' + 13)%26
			case char >= 'A' && char <= 'Z':
			    hash[i] = 'A' + (char - 'A' + 13)%26
		    case char >= '0' && char <= '9' :
				hash[i] = '0' + (char - '0' + 13 )
			default:
			hash[i] = char
		}
	}
	
	return string(hash)
}
