package main

import "testing"

func Test_cashback(t *testing.T) {
	test:= []struct {
		name string
		amount int
		want int
	}{
		{name:"Have cashback",amount:5000,want:250},
		{name:"Nocashback",amount:1000,want:0},
		{name:"Cashback on bound",amount:3000,want:150},
	}

	for _, tests := range test {
		got:=cashback(tests.amount)
		if got!=tests.want {
				t.Error("For cashback test:", tests.name ,"got:",got, "want:", tests.want)
			}
		}
	}
