package closest

import (
	"bytes"
	"fmt"
	"testing"
)

func TestClosestNraw1(t *testing.T) {
	targetData := []byte(
		`>Target1
ATGATC
>Target2
ATGATG
>Target3
ATTAGG
>Target4
ATTATG
>Target5
ATTATT
`)

	queryData := []byte(
		`>Query1
ATGATG
>Query2
ATGATC
>Query3
ATTATT
`)

	target := bytes.NewReader(targetData)

	query := bytes.NewReader(queryData)

	out := new(bytes.Buffer)

	err := ClosestN(2, -1.0, query, target, "raw", out, false, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,closest
Query1,Target2;Target1
Query2,Target1;Target2
Query3,Target5;Target4
` {
		t.Errorf("problem in TestClosestNraw1()")
		fmt.Println(string(out.Bytes()))
	}
}

func TestClosestNraw2(t *testing.T) {

	target := bytes.NewReader(targetData)
	query := bytes.NewReader(queryData)

	out := new(bytes.Buffer)

	err := ClosestN(10, -1.0, query, target, "raw", out, false, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,closest
query1,target9;target7;target8;target6;target10;target1;target4;target3;target5;target2
query2,target10;target4;target6;target3;target5;target1;target9;target7;target2;target8
` {
		t.Errorf("problem in TestClosestNraw2()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(5, -1.0, query, target, "raw", out, false, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,closest
query1,target9;target7;target8;target6;target10
query2,target10;target4;target6;target3;target5
` {
		t.Errorf("problem in TestClosestNraw2()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(0, 0.0022, query, target, "raw", out, false, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,closest
query1,target9;target7;target8;target6
query2,target10;target4;target6;target3;target5;target1;target9;target7;target2;target8
` {
		t.Errorf("problem in TestClosestNraw2()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(5, 0.0022, query, target, "raw", out, false, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,closest
query1,target9;target7;target8;target6
query2,target10;target4;target6;target3;target5
` {
		t.Errorf("problem in TestClosestNraw2()")
	}
}

func TestClosestNraw2Table(t *testing.T) {

	target := bytes.NewReader(targetData)
	query := bytes.NewReader(queryData)

	out := new(bytes.Buffer)

	err := ClosestN(10, -1.0, query, target, "raw", out, true, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,target,distance
query1,target9,0.002111932
query1,target7,0.002161977
query1,target8,0.002193871
query1,target6,0.002197654
query1,target10,0.002205301
query1,target1,0.002356075
query1,target4,0.002377798
query1,target3,0.002407351
query1,target5,0.002460697
query1,target2,0.002491127
query2,target10,0.000000000
query2,target4,0.000038867
query2,target6,0.000038950
query2,target3,0.000077583
query2,target5,0.000117357
query2,target1,0.000117366
query2,target9,0.000156937
query2,target7,0.000158165
query2,target2,0.000312671
query2,target8,0.000428399
` {
		t.Errorf("problem in TestClosestNraw2Table()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(5, -1.0, query, target, "raw", out, true, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,target,distance
query1,target9,0.002111932
query1,target7,0.002161977
query1,target8,0.002193871
query1,target6,0.002197654
query1,target10,0.002205301
query2,target10,0.000000000
query2,target4,0.000038867
query2,target6,0.000038950
query2,target3,0.000077583
query2,target5,0.000117357
` {
		t.Errorf("problem in TestClosestNraw2Table()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(0, 0.0022, query, target, "raw", out, true, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,target,distance
query1,target9,0.002111932
query1,target7,0.002161977
query1,target8,0.002193871
query1,target6,0.002197654
query2,target10,0.000000000
query2,target4,0.000038867
query2,target6,0.000038950
query2,target3,0.000077583
query2,target5,0.000117357
query2,target1,0.000117366
query2,target9,0.000156937
query2,target7,0.000158165
query2,target2,0.000312671
query2,target8,0.000428399
` {
		t.Errorf("problem in TestClosestNraw2Table()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(5, 0.0022, query, target, "raw", out, true, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,target,distance
query1,target9,0.002111932
query1,target7,0.002161977
query1,target8,0.002193871
query1,target6,0.002197654
query2,target10,0.000000000
query2,target4,0.000038867
query2,target6,0.000038950
query2,target3,0.000077583
query2,target5,0.000117357
` {
		t.Errorf("problem in TestClosestNraw2Table()")
	}
}

func TestClosestNsnp(t *testing.T) {

	target := bytes.NewReader(targetData)
	query := bytes.NewReader(queryData)

	out := new(bytes.Buffer)

	err := ClosestN(10, -1.0, query, target, "snp", out, false, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,closest
query1,target10;target9;target7;target8;target6;target1;target4;target3;target5;target2
query2,target10;target6;target4;target3;target1;target5;target9;target7;target2;target8
` {
		t.Errorf("problem in TestClosestNsnp()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(5, -1.0, query, target, "snp", out, false, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,closest
query1,target10;target9;target7;target8;target6
query2,target10;target6;target4;target3;target1
` {
		t.Errorf("problem in TestClosestNsnp()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(0, 12, query, target, "snp", out, false, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,closest
query1,
query2,target10;target6;target4;target3;target1;target5;target9;target7;target2;target8
` {
		t.Errorf("problem in TestClosestNsnp()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(5, 12, query, target, "snp", out, false, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,closest
query1,
query2,target10;target6;target4;target3;target1
` {
		t.Errorf("problem in TestClosestNsnp()")
	}
}

func TestClosestNsnpTable(t *testing.T) {

	target := bytes.NewReader(targetData)
	query := bytes.NewReader(queryData)

	out := new(bytes.Buffer)

	err := ClosestN(10, -1.0, query, target, "snp", out, true, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,target,distance
query1,target10,53
query1,target9,62
query1,target7,63
query1,target8,65
query1,target6,65
query1,target1,69
query1,target4,70
query1,target3,71
query1,target5,72
query1,target2,73
query2,target10,0
query2,target6,1
query2,target4,1
query2,target3,2
query2,target1,3
query2,target5,3
query2,target9,4
query2,target7,4
query2,target2,8
query2,target8,11
` {
		t.Errorf("problem in TestClosestNsnpTable()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(5, -1.0, query, target, "snp", out, true, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,target,distance
query1,target10,53
query1,target9,62
query1,target7,63
query1,target8,65
query1,target6,65
query2,target10,0
query2,target6,1
query2,target4,1
query2,target3,2
query2,target1,3
` {
		t.Errorf("problem in TestClosestNsnpTable()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(0, 12, query, target, "snp", out, true, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,target,distance
query2,target10,0
query2,target6,1
query2,target4,1
query2,target3,2
query2,target1,3
query2,target5,3
query2,target9,4
query2,target7,4
query2,target2,8
query2,target8,11
` {
		t.Errorf("problem in TestClosestNsnpTable()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(5, 12, query, target, "snp", out, true, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,target,distance
query2,target10,0
query2,target6,1
query2,target4,1
query2,target3,2
query2,target1,3
` {
		t.Errorf("problem in TestClosestNsnpTable()")
	}
}

func TestClosestNtn93(t *testing.T) {

	target := bytes.NewReader(targetData)
	query := bytes.NewReader(queryData)

	out := new(bytes.Buffer)

	err := ClosestN(10, -1.0, query, target, "raw", out, false, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,closest
query1,target9;target7;target8;target6;target10;target1;target4;target3;target5;target2
query2,target10;target4;target6;target3;target5;target1;target9;target7;target2;target8
` {
		t.Errorf("problem in TestClosestNtn93()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(5, -1.0, query, target, "raw", out, false, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,closest
query1,target9;target7;target8;target6;target10
query2,target10;target4;target6;target3;target5
` {
		t.Errorf("problem in TestClosestNtn93()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(0, 0.0022, query, target, "raw", out, false, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,closest
query1,target9;target7;target8;target6
query2,target10;target4;target6;target3;target5;target1;target9;target7;target2;target8
` {
		t.Errorf("problem in TestClosestNtn93()")
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(5, 0.0022, query, target, "raw", out, false, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,closest
query1,target9;target7;target8;target6
query2,target10;target4;target6;target3;target5
` {
		t.Errorf("problem in TestClosestNtn93()")
	}
}

func TestClosestNtn93Table(t *testing.T) {

	target := bytes.NewReader(targetData)
	query := bytes.NewReader(queryData)

	out := new(bytes.Buffer)

	err := ClosestN(10, -1.0, query, target, "tn93", out, true, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,target,distance
query1,target9,0.002115549
query1,target7,0.002165809
query1,target8,0.002197920
query1,target6,0.002201555
query1,target10,0.002209135
query1,target1,0.002360527
query1,target4,0.002382286
query1,target3,0.002411993
query1,target5,0.002465526
query1,target2,0.002496279
query2,target10,0.000000000
query2,target4,0.000038870
query2,target6,0.000038951
query2,target3,0.000077595
query2,target5,0.000117374
query2,target1,0.000117377
query2,target9,0.000156954
query2,target7,0.000158190
query2,target2,0.000312766
query2,target8,0.000428590
` {
		t.Errorf("problem in TestClosestNtn93Table()")
		fmt.Println(string(out.Bytes()))
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(5, -1.0, query, target, "tn93", out, true, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,target,distance
query1,target9,0.002115549
query1,target7,0.002165809
query1,target8,0.002197920
query1,target6,0.002201555
query1,target10,0.002209135
query2,target10,0.000000000
query2,target4,0.000038870
query2,target6,0.000038951
query2,target3,0.000077595
query2,target5,0.000117374
` {
		t.Errorf("problem in TestClosestNtn93Table()")
		fmt.Println(string(out.Bytes()))
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(0, 0.0022, query, target, "tn93", out, true, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,target,distance
query1,target9,0.002115549
query1,target7,0.002165809
query1,target8,0.002197920
query2,target10,0.000000000
query2,target4,0.000038870
query2,target6,0.000038951
query2,target3,0.000077595
query2,target5,0.000117374
query2,target1,0.000117377
query2,target9,0.000156954
query2,target7,0.000158190
query2,target2,0.000312766
query2,target8,0.000428590
` {
		t.Errorf("problem in TestClosestNtn93Table()")
		fmt.Println(string(out.Bytes()))
	}

	target = bytes.NewReader(targetData)
	query = bytes.NewReader(queryData)

	out = new(bytes.Buffer)

	err = ClosestN(5, 0.0022, query, target, "tn93", out, true, 2)
	if err != nil {
		t.Error(err)
	}

	if string(out.Bytes()) != `query,target,distance
query1,target9,0.002115549
query1,target7,0.002165809
query1,target8,0.002197920
query2,target10,0.000000000
query2,target4,0.000038870
query2,target6,0.000038951
query2,target3,0.000077595
query2,target5,0.000117374
` {
		t.Errorf("problem in TestClosestNtn93Table()")
		fmt.Println(string(out.Bytes()))
	}
}
