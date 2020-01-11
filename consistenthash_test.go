package consisenthash_cache

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPut(t *testing.T)  {
	c :=New();
	n :=new(Node)
	n.Name="n1"
	c.AddNode(n,4)
	c.Put("a","b")
	fmt.Println(c.Get("a"))
}

func TestRehash(t *testing.T)  {
	c :=New();
	n1 :=new(Node)
	n1.Name="n1"
	c.AddNode(n1,4)
	for i:=0;i<1000;i++{
		c.Put(strconv.Itoa(i),i)
	}
	for _,v :=range c.circle.Values(){
		vn,_:= v.(*VirtualNode)
		fmt.Println("Before:"+strconv.Itoa(len(vn.Data.Elements)))
	}
	n2 :=new(Node)
	n2.Name="n2"
	c.AddNode(n2,4)

	for _,v :=range c.circle.Values(){
		vn,_:= v.(*VirtualNode)
		fmt.Println("After:"+strconv.Itoa(len(vn.Data.Elements)))
	}
}
