package main

import "fmt"

type stack struct {
	Data []string
}

func main() {
	var text string
	var s stack

	_, err := fmt.Scanf("%s", &text)

	if err != nil {
		fmt.Println(err)
		return
	}

	temp := ""
	for _, char := range text {
		if temp == "" {
			if char == 'b' || char == 'k' || char == 'a' || char == 'n' {
				temp = string(char)
			}
				
		}else{
			temp += string(char)
		}

		switch temp {
			case "bale":
				s.Data = append(s.Data, "(")
				temp = ""
				break
			case "kheir":
				temp = ""
				leng := len(s.Data)
				if s.Data[leng-1] == "(" {
					s.Data = s.Data[:leng-1]
					}else{
				s.Data = append(s.Data, ")")
				}
				break
			case "areh":
				s.Data = append(s.Data, "{")
				temp = ""
				break
			case "na":
				leng := len(s.Data)
				if s.Data[leng-1] == "{" {
					s.Data = s.Data[:leng-1]
					}else{
				s.Data = append(s.Data, "}")
				}
				temp = ""
				break
		}

		if len(temp) >= 5{
			temp = ""
		}
	}

	if len(s.Data) == 0 {
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
	}
}
