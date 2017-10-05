package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	INCR  string = "ﾝｺﾞ"
	DECR  string = "ﾝｺﾞwww"
	NEXT  string = "ﾝｺﾞｺﾞ"
	PREV  string = "ﾝｺﾞｺﾞwww"
	READ  string = "ﾝｺﾞ??"
	WRITE string = "ﾝｺﾞ!!"
	OPEN  string = "ﾝｺﾞｺﾞ??"
	CLOSE string = "ﾝｺﾞｺﾞ!!"
)

var MAX_BUFFER_SIZE = 3000

func main() {
	// 引数のチェック
	if len(os.Args) < 2 {
		fmt.Println("invalid args.")
		return
	}

	// 引数で指定したファイルをオープン
	fp, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// ファイル読み込み
	scanner := bufio.NewScanner(fp)
	src := ""
	for scanner.Scan() {
		src += scanner.Text()
	}

	// ファイルの内容を"ﾝ"で分割
	cmds := strings.Split(src, "ﾝ")
	cmds = cmds[1:]

	// 分割時に消えた"ﾝ"を復元
	for index, cmd := range cmds {
		cmds[index] = "ﾝ" + cmd
	}

	// 色々宣言
	b := make([]byte, MAX_BUFFER_SIZE, MAX_BUFFER_SIZE)
	index := 0
	buf := make([]byte, 1)
	cmdIndex := 0

	// 本処理
	for {
		cmd := cmds[cmdIndex]
		fmt.Println(cmd)
		switch cmd {
		case INCR:
			index++
		case DECR:
			index--
		case NEXT:
			b[index]++
		case PREV:
			b[index]--
		case WRITE:
			buf[0] = b[index]
			os.Stdout.Write(buf)
		case READ:
			os.Stdin.Read(buf)
			b[index] = buf[0]
		case OPEN:
			if b[index] == 0 {
				n := 0
				for {
					cmdIndex++
					cmd = cmds[cmdIndex]
					if cmd == OPEN {
						n++
					} else if cmd == CLOSE {
						n--
						if n < 0 {
							break
						}
					}
				}
			}
		case CLOSE:
			if b[index] != 0 {
				n := 0
				for {
					cmdIndex--
					cmd = cmds[cmdIndex]
					if cmd == CLOSE {
						n++
					} else if cmd == OPEN {
						n--
						if n < 0 {
							break
						}
					}
				}
			}
		}
		cmdIndex++
		if cmdIndex >= len(cmds) {
			break
		}
	}
}
