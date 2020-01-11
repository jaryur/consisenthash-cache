package consisenthash_cache

import (
	"github.com/emirpasic/gods/maps/treemap"
	"hash/fnv"
	"math"
	"math/rand"
	"strconv"
	"sync"
)


type ConsistentHashCircle struct {
	circle *treemap.Map
	mutex sync.Mutex
}

type Node struct {
	Name string
}

type VirtualNode struct {
	node *Node
	hash string
	Data *CacheNode
}
type CacheNode struct {
	Elements map[string]interface{}
}

func New() *ConsistentHashCircle{
	c:= new(ConsistentHashCircle)
	c.circle=treemap.NewWithIntComparator();
	c.circle.Put(math.MaxInt32,newVNode())
	c.circle.Put(math.MinInt32,newVNode())
	return c
}

func newVNode() *VirtualNode{
	vnode:=new(VirtualNode)
	vnode.node=nil
	vnode.Data=new(CacheNode)
	vnode.Data.Elements=make(map[string]interface{})
	return vnode;
}

func  (c *ConsistentHashCircle) AddNode(node *Node,expandFactor int){
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for i:=0;i <expandFactor;i++{
		newnode:=newVNode();
		hash :=c.hash(node.Name+"_"+strconv.Itoa(rand.Int()))
		key,value :=c.circle.Floor(hash-1)
		vn,_ :=value.(*VirtualNode)
		c.circle.Put(hash, newnode)
		nextNode :=value.(*VirtualNode);
		if(key != nil && value != nil){
			for fk,fv :=range nextNode.Data.Elements{
				if(c.hash(fk) <= hash){
					newnode.Data.Elements[fk]=fv
					delete(vn.Data.Elements,fk)
				}
			}
		}
	}
}

func (c *ConsistentHashCircle) Put(k string,v interface{}){
	_,value :=c.circle.Floor(c.hash(k))
	vn,_ :=value.(*VirtualNode)
	vn.Data.Elements[k]=v
}

func (c *ConsistentHashCircle) Get(k string)(v interface{}){
	_,value :=c.circle.Floor(c.hash(k))
	vn,_ :=value.(*VirtualNode)
	return vn.Data.Elements[k]
}

func (c *ConsistentHashCircle) RouteNode(k interface{})(v *VirtualNode){
	_,value :=c.circle.Floor(k)
	vn,_ :=value.(*VirtualNode)
	return vn
}

func (c *ConsistentHashCircle)  NodeSize() int{
	return c.circle.Size()
}

func (c *ConsistentHashCircle) hash(k string) int  {
	f:=fnv.New32a();
	f.Write([]byte(k))
	return int(f.Sum32())
}

