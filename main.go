/**
 * @note
 * main.go
 *
 * @author	songtianming
 * @date 	2020-05-15
 */
package soduku

type Soduku struct {
	Board [9][9]int64
}

func isP

func (s *Soduku) FindPossibility(x, y int) (ret []int64) {
	var i int
	var exist [9]bool
	for i = 0; i < 9; i++ {
		if i == y || s.Board[x][i]==0{
			continue
		}
		exist[s.Board[x][i]]=true
	}
}

func main() {

}
