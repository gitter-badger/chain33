package mavl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"

	. "code.aliyun.com/chain33/chain33/common"
	"code.aliyun.com/chain33/chain33/common/db"
	"code.aliyun.com/chain33/chain33/types"
)

const testReadLimit = 1 << 20 // Some reasonable limit for wire.Read*() lmt
const (
	strChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" // 62 characters
)

// Constructs an alphanumeric string of given length.
func RandStr(length int) string {
	chars := []byte{}
MAIN_LOOP:
	for {
		val := rand.Int63()
		for i := 0; i < 10; i++ {
			v := int(val & 0x3f) // rightmost 6 bits
			if v >= 62 {         // only 62 characters in strChars
				val >>= 6
				continue
			} else {
				chars = append(chars, strChars[v])
				if len(chars) == length {
					break MAIN_LOOP
				}
				val >>= 6
			}
		}
	}

	return string(chars)
}

func randstr(length int) string {
	return RandStr(length)
}

func i2b(i int) []byte {

	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, i)
	return b_buf.Bytes()
}

func b2i(bz []byte) int {
	var x int
	b_buf := bytes.NewBuffer(bz)
	binary.Read(b_buf, binary.BigEndian, &x)
	return x
}

// Convenience for a new node
func N(l, r interface{}) *MAVLNode {
	var left, right *MAVLNode
	if _, ok := l.(*MAVLNode); ok {
		left = l.(*MAVLNode)
	} else {
		left = NewMAVLNode(i2b(l.(int)), nil)
	}
	if _, ok := r.(*MAVLNode); ok {
		right = r.(*MAVLNode)
	} else {
		right = NewMAVLNode(i2b(r.(int)), nil)
	}

	n := &MAVLNode{
		key:       right.lmd(nil).key,
		value:     nil,
		leftNode:  left,
		rightNode: right,
	}
	n.calcHeightAndSize(nil)
	return n
}

// Setup a deep node
func T(n *MAVLNode) *MAVLTree {
	d := db.NewDB("test", db.MemDBBackendStr, "")
	t := NewMAVLTree(d)

	n.Hash(t)
	t.root = n
	return t
}

// Convenience for simple printing of keys & tree structure
func P(n *MAVLNode) string {
	if n.height == 0 {
		return fmt.Sprintf("%v", b2i(n.key))
	} else {
		return fmt.Sprintf("(%v %v)", P(n.leftNode), P(n.rightNode))
	}
}

// 测试set和get功能
func TestBasic(t *testing.T) {
	var tree *MAVLTree = NewMAVLTree(nil)
	var up bool
	up = tree.Set([]byte("1"), []byte("one"))
	if up {
		t.Error("Did not expect an update (should have been create)")
	}
	up = tree.Set([]byte("2"), []byte("two"))
	if up {
		t.Error("Did not expect an update (should have been create)")
	}
	up = tree.Set([]byte("2"), []byte("TWO"))
	if !up {
		t.Error("Expected an update")
	}
	up = tree.Set([]byte("5"), []byte("five"))
	if up {
		t.Error("Did not expect an update (should have been create)")
	}
	hash := tree.Hash()

	treelog.Info("TestBasic", "roothash", hash)

	//PrintMAVLNode(tree.root)

	// Test 0x00
	{
		idx, val, exists := tree.Get([]byte{0x00})
		if exists {
			t.Errorf("Expected no value to exist")
		}
		if idx != 0 {
			t.Errorf("Unexpected idx %x", idx)
		}
		if string(val) != "" {
			t.Errorf("Unexpected value %v", string(val))
		}
	}

	// Test "1"
	{
		idx, val, exists := tree.Get([]byte("1"))
		if !exists {
			t.Errorf("Expected value to exist")
		}
		if idx != 0 {
			t.Errorf("Unexpected idx %x", idx)
		}
		if string(val) != "one" {
			t.Errorf("Unexpected value %v", string(val))
		}
	}

	// Test "2"
	{
		idx, val, exists := tree.Get([]byte("2"))
		if !exists {
			t.Errorf("Expected value to exist")
		}
		if idx != 1 {
			t.Errorf("Unexpected idx %x", idx)
		}
		if string(val) != "TWO" {
			t.Errorf("Unexpected value %v", string(val))
		}
	}

	// Test "4"
	{
		idx, val, exists := tree.Get([]byte("4"))
		if exists {
			t.Errorf("Expected no value to exist")
		}
		if idx != 2 {
			t.Errorf("Unexpected idx %x", idx)
		}
		if string(val) != "" {
			t.Errorf("Unexpected value %v", string(val))
		}
	}
}

func TestTreeHeightAndSize(t *testing.T) {
	db := db.NewDB("mavltree", "leveldb", "datastore")

	// Create some random key value pairs
	records := make(map[string]string)

	count := 14
	for i := 0; i < count; i++ {
		records[randstr(20)] = randstr(20)
	}

	// Construct some tree and save it
	t1 := NewMAVLTree(db)

	for key, value := range records {
		t1.Set([]byte(key), []byte(value))
	}

	for key, value := range records {
		index, t2value, _ := t1.Get([]byte(key))
		if string(t2value) != value {
			treelog.Info("TestTreeHeightAndSize", "index", index, "key", []byte(key))
		}
	}
	t1.Hash()
	//PrintMAVLNode(t1.root)
	t1.Save()
	if int32(count) != t1.Size() {
		treelog.Error("TestTreeHeightAndSize Size != count", "treesize", t1.Size(), "count", count)
	}
	//treelog.Info("TestTreeHeightAndSize", "treeheight", t1.Height(), "leafcount", count)
	//treelog.Info("TestTreeHeightAndSize", "treesize", t1.Size())
	db.Close()
}

//测试hash，save,load以及节点value值的更新功能
func TestPersistence(t *testing.T) {
	db := db.NewDB("mavltree", "leveldb", "datastore")

	records := make(map[string]string)

	recordbaks := make(map[string]string)

	for i := 0; i < 10; i++ {
		records[randstr(20)] = randstr(20)
	}

	t1 := NewMAVLTree(db)

	for key, value := range records {
		t1.Set([]byte(key), []byte(value))
		//treelog.Info("TestPersistence tree1 set", "key", key, "value", value)
		recordbaks[key] = randstr(20)
	}

	hash := t1.Hash()
	t1.Save()

	treelog.Info("TestPersistence", "roothash1", hash)

	// Load a tree
	t2 := NewMAVLTree(db)
	t2.Load(hash)

	for key, value := range records {
		_, t2value, _ := t2.Get([]byte(key))
		if string(t2value) != value {
			t.Fatalf("Invalid value. Expected %v, got %v", value, t2value)
		}
	}

	// update 5个key的value在hash2 tree中，验证这个5个key在hash和hash2中的值不一样
	var count int = 0
	for key, value := range recordbaks {
		count++
		if count > 5 {
			break
		}
		t2.Set([]byte(key), []byte(value))
		//treelog.Info("TestPersistence insert new node treee2", "key", string(key), "value", string(value))
	}

	hash2 := t2.Hash()
	t2.Save()
	treelog.Info("TestPersistence", "roothash2", hash2)

	// 重新加载hash

	t11 := NewMAVLTree(db)
	t11.Load(hash)

	treelog.Info("------tree11------TestPersistence---------")
	for key, value := range records {
		_, t2value, _ := t11.Get([]byte(key))
		if string(t2value) != value {
			t.Fatalf("tree11 Invalid value. Expected %v, got %v", value, t2value)
		}
	}
	//重新加载hash2
	t22 := NewMAVLTree(db)
	t22.Load(hash2)
	treelog.Info("------tree22------TestPersistence---------")

	//有5个key对应的value值有变化
	for key, value := range records {
		_, t2value, _ := t22.Get([]byte(key))
		if string(t2value) != value {
			treelog.Info("tree22 value update.", "oldvalue", string(value), "newvalue", string(t2value), "key", string(key))
		}
	}
	count = 0
	for key, value := range recordbaks {
		count++
		if count > 5 {
			break
		}
		_, t2value, _ := t22.Get([]byte(key))
		if string(t2value) != value {
			t.Logf("tree2222 Invalid value. Expected %v, got %v,key %v", string(value), string(t2value), string(key))
		}
	}
	db.Close()
}

//测试key:value对的proof证明功能
func TestIAVLProof(t *testing.T) {

	db := db.NewDB("mavltree", "leveldb", "datastore")

	var tree *MAVLTree = NewMAVLTree(db)

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("TestIAVLProof key:%d!", i)
		value := fmt.Sprintf("TestIAVLProof value:%d!", i)
		tree.Set([]byte(key), []byte(value))
	}

	// Persist the items so far
	hash1 := tree.Save()

	// Add more items so it's not all persisted
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("TestIAVLProof KEY:%d!", i)
		value := fmt.Sprintf("TestIAVLProof VALUE:%d!", i)
		tree.Set([]byte(key), []byte(value))
	}

	rootHashBytes := tree.Hash()
	hashetr := ToHex(rootHashBytes)
	hashbyte := FromHex(hashetr)

	treelog.Info("TestIAVLProof", "rootHashBytes", rootHashBytes)
	treelog.Info("TestIAVLProof", "hashetr", hashetr)
	treelog.Info("TestIAVLProof", "hashbyte", hashbyte)

	var KEY9proofbyte []byte

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("TestIAVLProof KEY:%d!", i)
		value := fmt.Sprintf("TestIAVLProof VALUE:%d!", i)
		keyBytes := []byte(key)
		valueBytes := []byte(value)
		_, KEY9proofbyte, _ = tree.Proof(keyBytes)
		value2, proof := tree.ConstructProof(keyBytes)
		if !bytes.Equal(value2, valueBytes) {
			treelog.Info("TestIAVLProof", "value2", string(value2), "value", string(valueBytes))
		}
		if proof != nil {
			istrue := proof.Verify([]byte(key), []byte(value), rootHashBytes)
			if !istrue {
				treelog.Error("TestIAVLProof Verify fail", "keyBytes", string(keyBytes), "valueBytes", string(valueBytes), "roothash", rootHashBytes)
			}
		}
	}

	treelog.Info("TestIAVLProof test Persistence data----------------")

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("TestIAVLProof key:%d!", i)
		value := fmt.Sprintf("TestIAVLProof value:%d!", i)
		keyBytes := []byte(key)
		valueBytes := []byte(value)

		value2, proofbyte, _ := tree.Proof(keyBytes)
		if !bytes.Equal(value2, valueBytes) {
			treelog.Info("TestIAVLProof", "value2", string(value2), "value", string(valueBytes))
		}
		if proofbyte != nil {

			leafNode := types.LeafNode{Key: keyBytes, Value: valueBytes, Height: 0, Size: 1}
			leafHash := leafNode.Hash()

			proof, err := ReadProof(rootHashBytes, leafHash, proofbyte)
			if err != nil {
				treelog.Info("TestIAVLProof ReadProof err ", "err", err)
			}
			istrue := proof.Verify([]byte(key), []byte(value), rootHashBytes)
			if !istrue {
				treelog.Info("TestIAVLProof Verify fail!", "keyBytes", string(keyBytes), "valueBytes", string(valueBytes), "roothash", rootHashBytes)
			}
		}
	}

	roothash := tree.Save()

	//key：value对的proof，hash1中不存在,roothash中存在
	index := 9
	key := fmt.Sprintf("TestIAVLProof KEY:%d!", index)
	value := fmt.Sprintf("TestIAVLProof VALUE:%d!", index)
	keyBytes := []byte(key)
	valueBytes := []byte(value)

	leafNode := types.LeafNode{Key: keyBytes, Value: valueBytes, Height: 0, Size: 1}
	leafHash := leafNode.Hash()

	// verify proof in tree1
	treelog.Info("TestIAVLProof  Verify key proof in tree1 ", "keyBytes", string(keyBytes), "valueBytes", string(valueBytes), "roothash", hash1)

	proof, err := ReadProof(hash1, leafHash, KEY9proofbyte)
	if err != nil {
		treelog.Info("TestIAVLProof ReadProof err ", "err", err)
	}
	istrue := proof.Verify(keyBytes, valueBytes, hash1)
	if !istrue {
		treelog.Info("TestIAVLProof  key not in tree ", "keyBytes", string(keyBytes), "valueBytes", string(valueBytes), "roothash", hash1)
	}

	// verify proof in tree2
	treelog.Info("TestIAVLProof  Verify key proof in tree2 ", "keyBytes", string(keyBytes), "valueBytes", string(valueBytes), "roothash", roothash)

	proof, err = ReadProof(roothash, leafHash, KEY9proofbyte)
	if err != nil {
		treelog.Info("TestIAVLProof ReadProof err ", "err", err)
	}
	istrue = proof.Verify(keyBytes, valueBytes, roothash)
	if istrue {
		treelog.Info("TestIAVLProof  key in tree2 ", "keyBytes", string(keyBytes), "valueBytes", string(valueBytes), "roothash", roothash)
	}
	db.Close()
}

func TestSetAndGetKVPair(t *testing.T) {
	db := db.NewDB("mavltree", "leveldb", "datastore")

	var storeSet types.StoreSet
	var storeGet types.StoreGet

	total := 10
	storeSet.KV = make([]*types.KeyValue, total)
	storeGet.Keys = make([][]byte, total)

	records := make(map[string]string)

	for i := 0; i < total; i++ {
		records[randstr(20)] = randstr(20)
	}
	i := 0
	for key, value := range records {
		var keyvalue types.KeyValue
		keyvalue.Key = []byte(key)
		keyvalue.Value = []byte(value)
		if i < total {
			storeSet.KV[i] = &keyvalue
			storeGet.Keys[i] = []byte(key)
		}
		i++
	}
	// storeSet hash is nil
	storeSet.StateHash = nil
	newhash := SetKVPair(db, &storeSet)

	storeGet.StateHash = newhash
	values := GetKVPair(db, &storeGet)

	i = 0
	for key, value := range records {
		if i < total {
			if !bytes.Equal([]byte(records[string(storeGet.Keys[i])]), values[i]) {
				treelog.Info("TestSetAndGetKVPair", "recordskey", key, "recordsvalue", value)
				treelog.Info("TestSetAndGetKVPair", "value", value, "getvalue", string(values[i]))
				treelog.Info("TestSetAndGetKVPair", "getkey", string(storeGet.Keys[i]))
			}
		}
		i++
	}

	// 在原来的基础上再次插入10个节点

	var storeSet2 types.StoreSet
	var storeGet2 types.StoreGet

	total = 10
	storeSet2.KV = make([]*types.KeyValue, total)
	storeGet2.Keys = make([][]byte, total)

	records2 := make(map[string]string)

	for i := 0; i < total; i++ {
		records2[randstr(20)] = randstr(20)
	}
	i = 0
	for key, value := range records2 {
		var keyvalue types.KeyValue
		keyvalue.Key = []byte(key)
		keyvalue.Value = []byte(value)
		if i < total {
			storeSet2.KV[i] = &keyvalue
			storeGet2.Keys[i] = []byte(key)
		}
		i++
	}
	// storeSet hash is newhash
	storeSet2.StateHash = newhash
	newhash2 := SetKVPair(db, &storeSet2)

	//获取新插入的10个节点
	storeGet2.StateHash = newhash2
	values2 := GetKVPair(db, &storeGet2)

	i = 0
	for _, value := range records2 {
		if i < total {
			if !bytes.Equal([]byte(records2[string(storeGet2.Keys[i])]), values2[i]) {
				treelog.Info("TestSetAndGetKVPair new node", "value", value, "getvalue", string(values2[i]))
			}
		}
		i++
	}

	//获取旧的10个节点
	storeGet.StateHash = newhash2
	values3 := GetKVPair(db, &storeGet)

	i = 0
	for _, value := range records {
		if i < total {
			if !bytes.Equal([]byte(records[string(storeGet.Keys[i])]), values3[i]) {
				treelog.Info("TestSetAndGetKVPair old node", "value", value, "getvalue", string(values3[i]))
			}
		}
		i++
	}

	db.Close()
}

func TestGetAndVerifyKVPairProof(t *testing.T) {
	db := db.NewDB("mavltree", "leveldb", "datastore")

	var storeSet types.StoreSet
	var storeGet types.StoreGet

	total := 10
	storeSet.KV = make([]*types.KeyValue, total)
	storeGet.Keys = make([][]byte, total)

	records := make(map[string]string)

	for i := 0; i < total; i++ {
		records[randstr(20)] = randstr(20)
	}
	i := 0
	for key, value := range records {
		var keyvalue types.KeyValue
		keyvalue.Key = []byte(key)
		keyvalue.Value = []byte(value)
		if i < total {
			storeSet.KV[i] = &keyvalue
			storeGet.Keys[i] = []byte(key)
		}
		i++
	}
	// storeSet hash is nil
	storeSet.StateHash = nil
	newhash := SetKVPair(db, &storeSet)

	i = 0
	for i = 0; i < total; i++ {
		var keyvalue types.KeyValue

		proof := GetKVPairProof(db, newhash, storeGet.Keys[i])

		keyvalue.Key = storeGet.Keys[i]
		keyvalue.Value = []byte(records[string(storeGet.Keys[i])])
		exit := VerifyKVPairProof(db, newhash, keyvalue, proof)
		if !exit {
			treelog.Info("TestGetAndVerifyKVPairProof  Verify proof fail!", "keyvalue", keyvalue.String(), "newhash", newhash)
		}
	}
	db.Close()
}