package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct{
	head *bucketNode
}

type bucketNode struct {
	key string
	next *bucketNode
}

func (h *HashTable) Insert(key string){
	index :=hash(key)
	h.array[index].insert(key)
}

// func (h *HashTable) Search(key string) bool {
// 	index :=hash(key)
// }

// func (h *HashTable) Delete(key string) {
// 	index :=hash(key)
// }

func (b *bucket) insert(k string){
	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

func (b *bucket) search(k string) bool{
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}


func hash(key string) int{
	sum := 0
	for _,v := range key {
		sum+=int(v)
	}
	return sum & ArraySize
}

func Init() *HashTable{
	result := &HashTable{}
	for i := range result.array{
		result.array[i] = &bucket{}
	}
	return result
}

func main(){
	testHashTable := Init()
	fmt.Println(testHashTable)

	testBucket := &bucket{}
	testBucket.insert("RANDY")
	fmt.Println(testBucket.search("RANDY"))
	fmt.Println(testBucket.search("ERIC"))
}

