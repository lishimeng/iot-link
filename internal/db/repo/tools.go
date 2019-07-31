package repo

import "strings"

func checkErr(err error) error {

	if err != nil {
		if !strings.Contains(err.Error(), "LastInsertId") {
			return err
		}
	}
	return nil
}

type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
}

func GetLimit(pageNo int, pageSize int) (limit int, start int) {
	limit = pageSize
	start = (pageNo - 1) * pageSize
	return limit, start
}
