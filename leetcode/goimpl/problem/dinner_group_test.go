package problem

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 题目一：公司有n个组,>=5，每组人数相同，>=2人，需要进行随机的组队吃饭。
// 要求：
// 1. 两两一队或三人一队，不能落单
// 2. 两人队、三人队各自的队伍数均不得少于2
// 3. 一个人只出现一次
// 4. 队伍中所有人不能来自相同组
// 5. 随机组队，重复执行程序得到的结果不一样，总队伍数也不能一样
// 6. 注释注释注释
// 注：要同时满足条件1-6,
// 举例：
// GroupList = [ # 小组列表
// ['小名', '小红', '小马', '小丽', '小强'],
// ['大壮', '大力', '大1', '大2', '大3'],
// ['阿花', '阿朵', '阿蓝', '阿紫', '阿红'],
// ['A', 'B', 'C', 'D', 'E'],
// ['一', '二', '三', '四', '五'],
// ['建国', '建军', '建民', '建超', '建跃'],
// ['爱民', '爱军', '爱国', '爱辉', '爱月']
// ]
// 输入：GroupList
// 示例输出: [小强 大3] [阿红 E] [五 建跃] [爱月 小名] [大壮 阿花] [A 一] [建国 爱民] [小红 大力] [阿朵 B] [二 建军] [爱军 小马] [大1 阿蓝] [C 三] [建民 爱国 小丽] [大2 阿紫 D] [四 建超 爱辉]

func dinnerGroups(groups [][]string) (teams [][]string, g2, g3 int) {
	var (
		m                = len(groups)
		n                = len(groups[0])
		total            = m * n
		r                = rand.New(rand.NewSource(time.Now().UnixNano()))
		g2Count, g3Count = computeGroupCount(total, r)

		split = map[int]bool{} // 已分配的人，key为坐标"row_col"
	)
	g2, g3 = g2Count, g3Count

	var (
		p1Pos, p2Pos, p3Pos int
	)
	for g2Count > 0 {
		p1Pos, p2Pos = getNextG2(total, n, split, r)
		g := []string{groups[p1Pos/n][p1Pos%n], groups[p2Pos/n][p2Pos%n]}
		teams = append(teams, g)
		g2Count--
	}
	for g3Count > 0 {
		p1Pos, p2Pos, p3Pos = getNextG3(total, n, split, r)
		g := []string{groups[p1Pos/n][p1Pos%n], groups[p2Pos/n][p2Pos%n], groups[p3Pos/n][p3Pos%n]}
		teams = append(teams, g)
		g3Count--
	}
	return
}

func getNextG2(total int, n int, split map[int]bool, r *rand.Rand) (p1Pos, p2Pos int) {
	p1Pos = r.Intn(total)
	var exists bool
	for {
		exists = split[p1Pos]
		if !exists {
			break
		}
		p1Pos = r.Intn(total)
	}

	p2Pos = r.Intn(total)
	for {
		exists = split[p2Pos]
		if !exists &&
			(p1Pos != p2Pos) &&
			(p1Pos/n != p2Pos/n) {
			break
		}
		p2Pos = r.Intn(total)
	}
	split[p1Pos] = true
	split[p2Pos] = true
	return
}

func getNextG3(total int, n int, split map[int]bool, r *rand.Rand) (p1Pos, p2Pos, p3Pos int) {
	p1Pos, p2Pos = getNextG2(total, n, split, r)
	p3Pos = r.Intn(total)
	var exists bool
	for {
		exists = split[p3Pos]
		if !exists &&
			(p3Pos != p1Pos && p3Pos != p2Pos) &&
			(p3Pos/n != p1Pos/n && p3Pos/n != p2Pos/n) {
			break
		}
		p3Pos = r.Intn(total)
	}
	split[p3Pos] = true
	return
}

func computeGroupCount(total int, r *rand.Rand) (g2Count, g3Count int) {
	g2Count, g3Count = 2, 2
	var (
		left                 = total - 10
		leftG2, leftG3       = 0, 0
		leftG2Max, leftG3Max = 0, 0
	)
	if left%2 == 0 {
		leftG2Max = left / 2
		if leftG2Max <= 2 {
			leftG2 = leftG2Max
			left -= leftG2 * 2
		} else {
			leftG2 = r.Intn(leftG2Max + 1)
			left -= leftG2 * 2
			if left >= 3 {
				leftG3 += left / 3
				left -= leftG3 * 3
			}
		}
	} else {
		leftG3Max = left / 3
		if leftG3Max <= 1 {
			leftG3 = leftG3Max
			left -= leftG3 * 3
		} else {
			leftG3 = r.Intn(leftG3Max + 1)
			left -= leftG3 * 3
			if left >= 2 {
				leftG2 += left / 2
				left -= leftG2 * 2
			}
		}
	}

	if left == 2 {
		leftG2 += 1
	} else if left == 1 {
		leftG2 -= 1
		if leftG2 == -1 {
			leftG2 = 2
			leftG3 -= 1
		} else {
			leftG3 += 1
		}
	}
	g2Count += leftG2
	g3Count += leftG3
	return
}

func Test_dinnerGroups(t *testing.T) {
	groups := [][]string{
		{"小名", "小红", "小马", "小丽", "小强"},
		{"大壮", "大力", "大1", "大2", "大3"},
		{"阿花", "阿朵", "阿蓝", "阿紫", "阿红"},
		{"A", "B", "C", "D", "E"},
		{"一", "二", "三", "四", "五"},
		{"建国", "建军", "建民", "建超", "建跃"},
		{"爱民", "爱军", "爱国", "爱辉", "爱月"},
	}

	r := make(chan struct {
		gs [][]string
		g2 int
		g3 int
	})

	ticker := time.NewTicker(1 * time.Millisecond)
loop:
	for {
		select {
		case <-ticker.C:
			go func() {
				gs, g2, g3 := dinnerGroups(groups)
				r <- struct {
					gs [][]string
					g2 int
					g3 int
				}{
					gs, g2, g3,
				}
			}()
		case d := <-r:
			fmt.Println("g2Count:", d.g2, " g3Count:", d.g3)
			fmt.Println(d.gs)
			break loop
		}
	}
}
