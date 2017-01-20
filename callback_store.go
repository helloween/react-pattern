package main

import (
  "math/rand"
)

type CallbackStore struct {
  store map[int]func(int)
}

func (clbStore CallbackStore) Init() { // TODO: return error
  clbStore.store = make(map[int]func(int))
}

func (clbStore CallbackStore) PushItem(clb func(int)) CallbackHandle {
  var handler CallbackHandle = rand.Int()

  clbStore.store[handler.(int)] = clb

  return handler
}

func (clbStore CallbackStore) DeleteItem(handler CallbackHandle) CallbackHandle {
  delete(clbStore.store, handler.(int))

  return handler
}