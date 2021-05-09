package main

import (
	"testing"
)

func TestMyBigNumber_sum(t *testing.T) {
	m := MyBigNumber{}
	//
	err := m.sum("12", "12")
	if err != nil {
		t.Errorf("m.sum(\"12\", \"12\") return err: %s; want nil", err.Error())
	}

	err = m.sum("12", "1269")
	if err != nil {
		t.Errorf("m.sum(\"12\", \"1269\") return err: %s; want nil", err.Error())
	}

	err = m.sum("12", "98")
	if err != nil {
		t.Errorf("m.sum(\"12\", \"98\") return err: %s; want nil", err.Error())
	}

	err = m.sum("12", "88")
	if err != nil {
		t.Errorf("m.sum(\"12\", \"88\") return err: %s; want nil", err.Error())
	}

	err = m.sum("123", "88")
	if err != nil {
		t.Errorf("m.sum(\"123\", \"88\") return err: %s; want nil", err.Error())
	}

	err = m.sum("abc", "98")
	if err != nil && err.Error() != "chuỗi thứ nhất không phải kiểu số" {
		t.Errorf("Output expect chuỗi thứ nhất không phải kiểu số instead of %v", err.Error())
	}

	err = m.sum("23", "acb")
	if err != nil && err.Error() != "chuỗi thứ hai không phải kiểu số" {
		t.Errorf("Output expect chuỗi thứ nhất không phải kiểu số instead of %v", err.Error())
	}
}
