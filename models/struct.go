package models

import (
	"time"
)


type NanChangSrc struct {
	Id  int    `xorm:"not null pk autoincr INT(11)"`
	A1  string `xorm:"VARCHAR(9)"`
	A2  string `xorm:"VARCHAR(9)"`
	A3  string `xorm:"VARCHAR(9)"`
	A4  string `xorm:"VARCHAR(12)"`
	A5  string `xorm:"VARCHAR(11)"`
	A6  string `xorm:"VARCHAR(6)"`
	A7  string `xorm:"VARCHAR(3) index"`
	A8  string `xorm:"VARCHAR(4)"`
	A9  string `xorm:"VARCHAR(20)"`
	A10 string `xorm:"VARCHAR(31)"`
	A11 string `xorm:"VARCHAR(3)"`
	A12 string `xorm:"VARCHAR(14) index "`
	A13 string `xorm:"VARCHAR(108)"`
	A14 string `xorm:"VARCHAR(18)"`
	A15 string `xorm:"VARCHAR(4)"`
	A16 string `xorm:"VARCHAR(3)"`
	A17 string `xorm:"VARCHAR(20)"`
	A18 string `xorm:"VARCHAR(3)"`
	A19 string `xorm:"VARCHAR(20)"`
	A20 string `xorm:"DECIMAL(15,5)"`
	A21 string `xorm:"DECIMAL(15,5)"`
}


type GuangZhouSrc struct {
	Id  int       `xorm:"not null pk autoincr INT(11)"`
	A1  string    `xorm:"VARCHAR(99)"`
	A2  string    `xorm:"VARCHAR(99) index" `
	A3  string    `xorm:"VARCHAR(99)"`
	A4  string    `xorm:"VARCHAR(99)"`
	A5  time.Time `xorm:"DATETIME index"`
	A6  string    `xorm:"VARCHAR(99)"`
	A7  string    `xorm:"VARCHAR(99)"`
	A8  string    `xorm:"VARCHAR(99)"`
	A9  string    `xorm:"VARCHAR(99)"`
	A10 string    `xorm:"VARCHAR(99)"`
	A11 string    `xorm:"VARCHAR(99)"`
	A12 string    `xorm:"VARCHAR(99)"`
	A13 string    `xorm:"VARCHAR(99)"`
	A14 string    `xorm:"VARCHAR(99)"`
	A15 string    `xorm:"VARCHAR(99)"`
	A16 string    `xorm:"VARCHAR(99)"`
	A17 string    `xorm:"VARCHAR(99)"`
	A18 string    `xorm:"VARCHAR(99)"`
	A19 string    `xorm:"VARCHAR(99)"`
	A20 string    `xorm:"VARCHAR(99)"`
	A21 string    `xorm:"VARCHAR(99)"`
	A22 string    `xorm:"VARCHAR(99)"`
	A23 string    `xorm:"VARCHAR(99)"`
	A24 string    `xorm:"VARCHAR(99)"`
	A25 string    `xorm:"VARCHAR(99)"`
	A26 string    `xorm:"VARCHAR(99)"`
	A27 string    `xorm:"VARCHAR(99)"`
	A28 string    `xorm:"VARCHAR(99)"`
	A29 string    `xorm:"VARCHAR(99)"`
	A30 string    `xorm:"VARCHAR(99)"`
	A31 string    `xorm:"VARCHAR(99)"`
	Month string    `xorm:"VARCHAR(4)"`
}

type CalBianMa struct {
	Id                     int    `xorm:"not null pk autoincr INT(9)"`
	Bianma                 string `xorm:"index VARCHAR(99)"`
	Gongyingshang          string `xorm:"VARCHAR(99)"`
	Mingcheng              string `xorm:"VARCHAR(300)"`
	M1shuliang             int    `xorm:"INT(10)"`
	M1danjia               string `xorm:"DECIMAL(15,2)"`
	M1caigoujine           string `xorm:"DECIMAL(15,2)"`
	M1danjiajiangjiajine   string `xorm:"DECIMAL(15,2)"`
	M1danjiajiangjiabili   string `xorm:"DECIMAL(5,4)"`
	M1zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M1zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M2shuliang             int    `xorm:"INT(10)"`
	M2danjia               string `xorm:"DECIMAL(15,2)"`
	M2caigoujine           string `xorm:"DECIMAL(15,2)"`
	M2danjiajiangjiajine   string `xorm:"DECIMAL(15,2)"`
	M2danjiajiangjiabili   string `xorm:"DECIMAL(5,4)"`
	M2zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M2zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M3shuliang             int    `xorm:"INT(10)"`
	M3danjia               string `xorm:"DECIMAL(15,2)"`
	M3caigoujine           string `xorm:"DECIMAL(15,2)"`
	M3danjiajiangjiajine   string `xorm:"DECIMAL(15,2)"`
	M3danjiajiangjiabili   string `xorm:"DECIMAL(5,4)"`
	M3zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M3zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M4shuliang             int    `xorm:"INT(10)"`
	M4danjia               string `xorm:"DECIMAL(15,2)"`
	M4caigoujine           string `xorm:"DECIMAL(15,2)"`
	M4danjiajiangjiajine   string `xorm:"DECIMAL(15,2)"`
	M4danjiajiangjiabili   string `xorm:"DECIMAL(5,4)"`
	M4zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M4zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M5shuliang             int    `xorm:"INT(10)"`
	M5danjia               string `xorm:"DECIMAL(15,2)"`
	M5caigoujine           string `xorm:"DECIMAL(15,2)"`
	M5danjiajiangjiajine   string `xorm:"DECIMAL(15,2)"`
	M5danjiajiangjiabili   string `xorm:"DECIMAL(5,4)"`
	M5zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M5zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M6shuliang             int    `xorm:"INT(10)"`
	M6danjia               string `xorm:"DECIMAL(15,2)"`
	M6caigoujine           string `xorm:"DECIMAL(15,2)"`
	M6danjiajiangjiajine   string `xorm:"DECIMAL(15,2)"`
	M6danjiajiangjiabili   string `xorm:"DECIMAL(5,4)"`
	M6zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M6zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M7shuliang             int    `xorm:"INT(10)"`
	M7danjia               string `xorm:"DECIMAL(15,2)"`
	M7caigoujine           string `xorm:"DECIMAL(15,2)"`
	M7danjiajiangjiajine   string `xorm:"DECIMAL(15,2)"`
	M7danjiajiangjiabili   string `xorm:"DECIMAL(5,4)"`
	M7zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M7zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M8shuliang             int    `xorm:"INT(10)"`
	M8danjia               string `xorm:"DECIMAL(15,2)"`
	M8caigoujine           string `xorm:"DECIMAL(15,2)"`
	M8danjiajiangjiajine   string `xorm:"DECIMAL(15,2)"`
	M8danjiajiangjiabili   string `xorm:"DECIMAL(5,4)"`
	M8zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M8zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M9shuliang             int    `xorm:"INT(10)"`
	M9danjia               string `xorm:"DECIMAL(15,2)"`
	M9caigoujine           string `xorm:"DECIMAL(15,2)"`
	M9danjiajiangjiajine   string `xorm:"DECIMAL(15,2)"`
	M9danjiajiangjiabili   string `xorm:"DECIMAL(5,4)"`
	M9zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M9zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M10shuliang            int    `xorm:"INT(10)"`
	M10danjia              string `xorm:"DECIMAL(15,2)"`
	M10caigoujine          string `xorm:"DECIMAL(15,2)"`
	M10danjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M10danjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M10zongjiajiangjiajine string `xorm:"DECIMAL(15,2)"`
	M10zongjiajiangjiabili string `xorm:"DECIMAL(5,4)"`
	M11shuliang            int    `xorm:"INT(10)"`
	M11danjia              string `xorm:"DECIMAL(15,2)"`
	M11caigoujine          string `xorm:"DECIMAL(15,2)"`
	M11danjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M11danjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M11zongjiajiangjiajine string `xorm:"DECIMAL(15,2)"`
	M11zongjiajiangjiabili string `xorm:"DECIMAL(5,4)"`
	M12shuliang            int    `xorm:"INT(10)"`
	M12danjia              string `xorm:"DECIMAL(12,2)"`
	M12caigoujine          string `xorm:"DECIMAL(12,2)"`
	M12danjiajiangjiajine  string `xorm:"DECIMAL(12,2)"`
	M12danjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M12zongjiajiangjiajine string `xorm:"DECIMAL(12,2)"`
	M12zongjiajiangjiabili string `xorm:"DECIMAL(5,4)"`
}

type CalCompany struct {
	Id                     int    `xorm:"not null pk autoincr INT(11)"`
	Gongyingshang          string `xorm:"index VARCHAR(99)"`
	M1caigoujine           string `xorm:"DECIMAL(15,2)"`
	M1zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M1zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M2caigoujine           string `xorm:"DECIMAL(15,2)"`
	M2zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M2zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M3caigoujine           string `xorm:"DECIMAL(15,2)"`
	M3zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M3zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M4caigoujine           string `xorm:"DECIMAL(15,2)"`
	M4zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M4zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M5caigoujine           string `xorm:"DECIMAL(15,2)"`
	M5zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M5zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M6caigoujine           string `xorm:"DECIMAL(15,2)"`
	M6zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M6zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M7caigoujine           string `xorm:"DECIMAL(15,2)"`
	M7zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M7zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M8caigoujine           string `xorm:"DECIMAL(15,2)"`
	M8zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M8zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M9caigoujine           string `xorm:"DECIMAL(15,2)"`
	M9zongjiajiangjiajine  string `xorm:"DECIMAL(15,2)"`
	M9zongjiajiangjiabili  string `xorm:"DECIMAL(5,4)"`
	M10caigoujine          string `xorm:"DECIMAL(15,2)"`
	M10zongjiajiangjiajine string `xorm:"DECIMAL(15,2)"`
	M10zongjiajiangjiabili string `xorm:"DECIMAL(5,4)"`
	M11caigoujine          string `xorm:"DECIMAL(15,2)"`
	M11zongjiajiangjiajine string `xorm:"DECIMAL(15,2)"`
	M11zongjiajiangjiabili string `xorm:"DECIMAL(5,4)"`
	M12caigoujine          string `xorm:"DECIMAL(15,2)"`
	M12zongjiajiangjiajine string `xorm:"DECIMAL(15,2)"`
	M12zongjiajiangjiabili string `xorm:"DECIMAL(5,4)"`
}
