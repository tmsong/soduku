/**
 * @note
 * main.go
 *
 * @author	songtianming
 * @date 	2020-05-15
 */
package main

import "fmt"

type Soduku struct {
	Board [9][9]int
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
		exist[s.Board[x][i]-1] = true
	}
	for i = 0; i < 9; i++ {
		if i == x || s.Board[i][y] == 0 {
			continue
		}
		exist[s.Board[i][y]-1] = true
	}
	for i = (x / 3) * 3; i < (x/3+1)*3; i++ {
		for j = (y / 3) * 3; j < (y/3+1)*3; j++ {
			if (i == x && j == y) || s.Board[i][j] == 0 {
				continue
			}
			exist[s.Board[i][j]-1] = true
		}
	}
	for i := 0; i < 9; i++ {
		if !exist[i] {
			possibility = append(possibility, i+1)
		}
	}
	return possibility
}

//刷新所有位置，遇到只有1种情况的空位就填上并重新开始，直到只剩下多种可能性的情况
func (s *Soduku) FindLeastPossibilityXY() (x int, y int, possibility []int, failed bool, result *Soduku) {
	minPossibility := 10
	var filled bool //标志着本次寻找是否找到了一个位置的数
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.Board[i][j] == 0 { //没填
				p := s.FindPossibilityForOne(i, j)
				if len(p) == 0 { //代表这个数据有问题，有一个格子完全没有可能性
					failed = true
					return
				} else if len(p) == 1 {
					s.Board[i][j] = p[0]
					filled = true
				} else if len(p) < minPossibility { //找到了具有更少可能性的格子
					minPossibility = len(p)
					possibility = p
					x = i
					y = j
				}
			}
		}
	}
	if filled {
		return s.FindLeastPossibilityXY() //如果发生了填空行为，那么重新寻找一遍
	} else if minPossibility == 10 {
		result = s
		return
	}
	return
}

func (s *Soduku) Find() (ret *Soduku) {
	x, y, possibility, failed, result := s.FindLeastPossibilityXY()
	if result != nil {
		return result
	}
	if failed {
		return nil
	}
	fmt.Println(x, y, "possibility length:", len(possibility))
	for _, possibleValue := range possibility {
		sodukuCopy := s.Fork()
		sodukuCopy.Board[x][y] = possibleValue
		if result := sodukuCopy.Find(); result != nil {
			return result
		}
	}
	return nil
}

func (s *Soduku) Print() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(s.Board[i][j], " ")
		}
		fmt.Print("\n")
	}
}

func main() {
	s := &Soduku{}
	s.Board = [9][9]int{
		{0, 4, 6, 9, 0, 3, 0, 0, 0},
		{0, 0, 3, 0, 5, 0, 0, 6, 0},
		{9, 0, 0, 0, 0, 2, 0, 0, 3},
		{0, 0, 5, 0, 0, 6, 0, 0, 0},
		{8, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 7, 8, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 5, 0},
		{0, 8, 1, 3, 0, 0, 0, 0, 7},
		{0, 0, 0, 8, 0, 0, 1, 0, 4},
	}
	res := s.Find()
	res.Print()
}
