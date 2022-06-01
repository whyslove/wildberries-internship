package printtime

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

//Не совсем понял, зачем нужно выходить в этой программе, если есть ошибки
//Гораздо лучше выйте в основной программе, которая исопльзует этот модуль
//То есть эта функция должна возвращать (time.Time, error)
//А уже вызывающий код решает выходить или нет
//Но задание есть задание
func GetCurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при получении времени")
		os.Exit(1)
	}

	fmt.Println(time)
}
