func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	dp := make([][]int, len(matrix))
	for i:= range dp {
		dp[i] = make([]int,len(matrix[0]))
		for j:= range dp[i] {
			dp[i][j]=-1
		}
	}
	ans := 0
	for i:=0;i<len(matrix);i++{
        for j:=0;j<len(matrix[0]);j++{
			res := f(i,j,matrix,dp)
			if i==len(matrix)-1 && j == len(matrix[0])-1 {
				fmt.Println(res)
			}
            ans = max(ans, res)
        }
    }
    return ans*ans
}

func f(i, j int, matrix[][]byte,dp[][]int) int {
	if j == len(matrix[0]) || i == len(matrix) {
		return 0
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
    right,down,rightdiag,res := 0,0,0,0
	right = f(i,j+1,matrix,dp)
	down = f(i+1,j,matrix,dp)
	rightdiag = f(i+1,j+1,matrix,dp)

	if matrix[i][j]=='1' {
		res = 1+min(down, min(right, rightdiag))
	}
	dp[i][j]= res
	return res
}