package bookDomain

import (
	"errors"
	"fmt"
)

type Book struct {
	ISBN      int
	Title     string
	Publisher string
	Price     float64
}

type BookError struct {
	Code int
	Err  error
}

func (bkerr *BookError) Error() string {
	return fmt.Sprintf("code %d: err %v", bkerr.Code, bkerr.Err)
}

func GetAllBook() ([]Book, error) {
	bks, err := getAllBook()
	if err != nil {
		return bks, &BookError{
			Code: 500,
			Err:  errors.New("there is some server side issue. please try after sometime"),
		}
	} else if len(bks) < 1 {
		return bks, &BookError{
			Code: 204,
			Err:  errors.New("no books found"),
		}
	} else {
		return bks, nil
	}

}

func GetBookByISBN(isbn int) (*Book, error) {
	bk, err := getBookByISBN(isbn)
	if err != nil {
		return bk, &BookError{
			Code: 500,
			Err:  errors.New("there is some server side issue. please try after sometime"),
		}
	} else if bk == nil {
		if err != nil {
			return bk, &BookError{
				Code: 204,
				Err:  errors.New("no book found"),
			}
		}
	}
	return bk, nil

}

func InsertBook(bk Book) (string, error) {
	str, err := insertBook(bk)
	if err != nil {
		return str, &BookError{
			Code: 500,
			Err:  errors.New("there is some server side issue. please try after sometime"),
		}

	} else {
		return str, nil
	}
}
