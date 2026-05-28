package luhn

import "strings"
func Valid(id string) bool {
	id = strings.Replace(id, " ", "",-1)
    if len(id) < 2 {
        return false 
    }

    sum := 0 
    for i:= len(id)-1; i>=0; i-- {
        if int(id[i]) < 48 || int(id[i]) > 57{
            return false 
        }
        a:= int(id[i] - '0')
        if (len(id) - (i + 1)) % 2 != 0{
            sum += a*2
            if a>4 {
                sum -= 9 
            }
        }else {
            sum += a
        }
    }
    return sum % 10 == 0
}
