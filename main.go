/**
 * @note
 * main.go
 *
 * @author	songtianming
 * @date 	2020-05-15
 */
package soduku

type Soduku struct {
	Board [9][9]int
}

func findLeft(exist [9]bool) (possibility []int) {
	for i := 0; i < 9; i++ {
		if !exist[i] {
			possibility = append(possibility, i+1)
		}
	}
	return possibility
}

func (s *Soduku) Fork() *Soduku {
	ret := &Soduku{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			ret.Board[i][j] = s.Board[i][j]
		}
	}
	return ret
}

func (s *Soduku) FindPossibilityForOne(x, y int) (possibility []int) {
	var i, j int
	var exist [9]bool
	for i = 0; i < 9; i++ {
		if i == y || s.Board[x][i] == 0 {
			continue
		}
		exist[s.Board[x][i]] = true
	}
	for i = 0; i < 9; i++ {
		if i == x || s.Board[i][y] == 0 {
			continue
		}
		exist[s.Board[i][y]] = true
	}
	for i = (x / 3) * 3; i < (x/3+1)*3; i++ {
		for j = (y / 3) * 3; j < (x/3+1)*3; j++ {
			if (i == x && j == y) || s.Board[i][j] == 0 {
				continue
			}
			exist[s.Board[i][j]] = true
		}
	}
	return findLeft(exist)
}

//刷新所有位置，遇到只有1种情况的空位就填上并重新开始，直到只剩下多种可能性的情况
func (s *Soduku) FindLeastPossibilityXY() (x int, y int, possibility []int, failed bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.Board[i][j] == 0 { //没填
				p := s.FindPossibilityForOne(i, j)
				if len(p) == 0 { //代表这个数据有问题，有一个地方没有可能性
					failed = true
					return
				}
			}
		}
	}

}

func main() {

}
