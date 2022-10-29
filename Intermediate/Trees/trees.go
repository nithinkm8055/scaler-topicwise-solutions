package Trees

import (
	"fmt"

	"github.com/nithinkm8055/scaler-topicwise-solutions/Intermediate/Trees/Trees1/Assignment"
)

func Trees() {

	{
		C := Assignment.TreeNode_new(3)
		B := Assignment.TreeNode_new(2)
		B.Left = C

		D := Assignment.TreeNode_new(6)

		A := Assignment.TreeNode_new(1)
		A.Right = B
		A.Left = D

		fmt.Println(Assignment.InorderTraversal(A))
	}

}
