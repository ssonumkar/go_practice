package main

import (
	"fmt"

	"practice.com/sagar/go_practice/fileio/fileoperations"
)

func main() {

	var ch int
	var fileName string
	fmt.Println("Enter the file name :")
	fmt.Scanf("%s", &fileName)
	for {

		fmt.Println("1.CREATE\n2.READ\n3.WRITE\n4.APPEND")
		fmt.Scanf("%d", &ch)
		switch ch {
		case 1:
			{
				fileoperations.Create(fileName)
			}
		case 2:
			{
				fmt.Println(fileoperations.Read(fileName))
			}
		case 3:
			{
				var data string
				fmt.Scanf("%s", &data)
				fileoperations.Write(fileName, data)
			}

		case 4:
			{
				var data string
				fmt.Scanf("%s", &data)
				fileoperations.Append(fileName, data)
			}
		}
		var op int
		fmt.Println("Press 1 to continue? ")
		fmt.Scanf("%d", &op)
		if op != 1 {
			break
		}
	}

}
