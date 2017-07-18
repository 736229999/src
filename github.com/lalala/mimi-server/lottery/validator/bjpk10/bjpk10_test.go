package bjpk10

import "fmt"

// func Example_Time1() {
// 	fmt.Println(Time1("621992", 358))
// 	//Output:621098 1496365380 1496365620 1496365740
// }

func Example_MakeSaleIssueList() {
	fmt.Println(MakeSaleIssueList("620560", 1))
	//Output: [{620561 1496106420 1496106720 1496106720}]
	fmt.Println(MakeSaleIssueList("620560", DAY_MAX_NO))
	//Output: [{620561 1496106420 1496106720 1496106720}]
}
