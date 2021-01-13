package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-yaml/yaml"
)

type Character struct {
	name    string
	skill   int
	stamina int
	luck    int
}

func (c *Character) calcAttack() int {
	attackPoint := c.skill + roleDice(2, 6)
	return attackPoint
}

func (c *Character) calcDamaged() {
	c.stamina -= 2
}

type Monster struct {
	name    string
	skill   int
	stamina int
	attack  int
}

func (m *Monster) calcAttack() int {
	attackPoint := m.skill + roleDice(2, 6)
	return attackPoint
}

func (m *Monster) calcDamaged() {
	m.stamina -= 2
}

func main() {
	data, _ := readMonsterYamlFile()
	numberOfMonsterTypes := len(data)

	var c Character
	c.name = "Tomcha"
	c.skill = roleDice(1, 6) + 6
	c.stamina = roleDice(2, 6) + 12
	c.luck = roleDice(2, 6)

	fmt.Println(c)
	wc := 0
	for c.stamina > 0 {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(numberOfMonsterTypes)

		var m Monster
		m.name = data[n]["name"]
		m.skill, _ = strconv.Atoi(data[n]["skill"])
		m.stamina, _ = strconv.Atoi(data[n]["stamina"])
		m.attack, _ = strconv.Atoi(data[n]["attack"])

		battle(&c, &m)
		wc += 1
	}
	fmt.Printf("%s は %d 回の戦闘に勝利した後、死亡しました。\n", c.name, wc)
}

func readMonsterYamlFile() ([]map[string]string, error) {
	// monster.yamlを []byte として読み込みます。
	buf, err := ioutil.ReadFile("./monster.yaml")
	if err != nil {
		fmt.Println(err)
		return []map[string]string{}, err
	}
	// []byte を []map[string]string に変換します。
	data, err := ReadOnSliceMap(buf)
	if err != nil {
		return []map[string]string{}, err
	}
	return data, nil
}

// yaml形式の[]byteを渡すと[]map[string]stringに変換してくれる関数です。
func ReadOnSliceMap(fileBuffer []byte) ([]map[string]string, error) {
	data := make([]map[string]string, 20)
	// ここで変換を行っています。
	// []byteを渡すとデータ型に合わせて上手い事マッピングしてくれます。
	err := yaml.Unmarshal(fileBuffer, &data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func roleDice(times int, hedron int) int {
	rand.Seed(time.Now().UnixNano())
	total := 0
	for i := 0; i < times; i++ {
		total += rand.Intn(hedron) + 1
	}
	return total
}

func battle(c *Character, m *Monster) {
	var winner, looser string
loop:
	for {
		if c.stamina <= 0 || m.stamina <= 0 {
			break loop
		}
		ca := c.calcAttack()
		ma := m.calcAttack()
		if ca > ma {
			m.calcDamaged()
		} else if ma > ca {
			c.calcDamaged()
		}
	}

	if c.stamina > m.stamina {
		winner = c.name
		looser = m.name
	} else {
		winner = m.name
		looser = c.name
	}
	fmt.Println(winner, "is win")
	fmt.Println(looser, "is loose")
}
