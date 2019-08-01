package log

// A Warp Sort Map For Text Logger
// 这个map存在于每一个entry中，因此不存在并发问题
// 不要在Formatter接口中使用异步读写,否则会导致竞争问题
type KvMap struct {
	ks    []string
	kv    map[string]interface{}
	state int32
}

func newKvMap() *KvMap {
	return &KvMap{ks: make([]string, 0), kv: make(map[string]interface{})}
}

func (kvm *KvMap) Put(k string, v interface{}) {
	if _, ok := kvm.kv[k]; !ok {
		kvm.kv[k] = v
		kvm.ks = append(kvm.ks, k)
	}
}

func (kvm *KvMap) Get(k string) interface{} {
	return kvm.kv[k]
}

func (kvm *KvMap) Range(fn func(k string, v interface{}) bool) {
	for _, k := range kvm.ks {
		if !fn(k, kvm.kv[k]) {
			break
		}
	}
}

func (kvm *KvMap) Keys() []string {
	return kvm.ks
}

func (kvm *KvMap) Map() map[string]interface{} {
	return kvm.kv
}

func (kvm *KvMap) Reset() {
	for _, k := range kvm.ks { // 应该还有更好的回收方式
		delete(kvm.kv, k)
	}
	kvm.ks = kvm.ks[:0]
}
