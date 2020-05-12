package docker

import "fmt"


var (
	tagStorePluginMgr = &TagStorePluginMgr{tagStores:make(map[string]TagStore)}
)

type TagStorePluginMgr struct {
	tagStores map[string]TagStore
}

func (t *TagStorePluginMgr) registerTagStorePlugin(storePlugin TagStore) (err error) {
	_, ok := t.tagStores[storePlugin.Name()]
	if ok {
		err = fmt.Errorf("duplicate registry plugin")
		return
	}
	t.tagStores[storePlugin.Name()] = storePlugin
	return
}

func (t *TagStorePluginMgr) initTagStore(name string, opts ...Option) (tagStore TagStore, err error) {
	tagStore1, ok := t.tagStores[name]
	if !ok {
		err = fmt.Errorf("plugin %s not exists", name)
		return
	}
	tagStore = tagStore1
	err = tagStore1.Init(opts...)
	return
}

func RegisterTagStorePlugin(storePlugin TagStore) (err error) {
	return tagStorePluginMgr.registerTagStorePlugin(storePlugin)
}

func InitTagStore(name string, opts ...Option) (tagStore TagStore, err error) {
	return tagStorePluginMgr.initTagStore(name, opts...)
}
