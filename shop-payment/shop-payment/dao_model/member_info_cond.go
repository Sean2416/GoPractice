package dao_model

import "time"

type MemberInfoCond struct {
	ID           *int64     `where:"id ="`
	IDs          []int64    `where:"id in"`
	ClientSns    []int64    `where:"client_sn in"`
	CreatedStart *time.Time `where:"created_at >"`
	CreatedEnd   *time.Time `where:"created_at <"`
}


