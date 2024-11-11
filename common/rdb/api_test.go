package rdb

import (
	"context"
	"testing"
)

func TestClient_NewSingle_Set(t *testing.T) {
	redisHost := "127.0.0.1:6379"
	client := NewSingle(redisHost, "")
	clientWrapper := NewClientWrapper(client)
	defer clientWrapper.Close()

	err := clientWrapper.Set(context.TODO(), "key1", "test1", 0)
	if err != nil {
		t.Error(err)
	}

	val, err := clientWrapper.Get(context.TODO(), "key1")
	if err != nil {
		t.Error(err)
	}

	if val != "test1" {
		t.Error("not test1")
	}
}

func TestClient_NewSentinel_Set(t *testing.T) {
	redisHost := []string{"127.0.0.1:63791", "127.0.0.1:63792", "127.0.0.1:63793"}
	client := NewSentinel("", redisHost, "")
	clientWrapper := NewClientWrapper(client)
	defer clientWrapper.Close()

	err := clientWrapper.Set(context.TODO(), "key1", "test1", 0)
	if err != nil {
		t.Error(err)
	}

	val, err := clientWrapper.Get(context.TODO(), "key1")
	if err != nil {
		t.Error(err)
	}

	if val != "test1" {
		t.Error("not test1")
	}
}

func TestClient_NewCluster_Set(t *testing.T) {
	redisHost := []string{"127.0.0.1:16379", "127.0.0.1:26379", "127.0.0.1:36379"}
	client := NewCluster(redisHost, "")
	clientWrapper := NewClientWrapper(client)
	defer clientWrapper.Close()

	err := clientWrapper.Set(context.TODO(), "key1", "test1", 0)
	if err != nil {
		t.Error(err)
	}

	val, err := clientWrapper.Get(context.TODO(), "key1")
	if err != nil {
		t.Error(err)
	}

	if val != "test1" {
		t.Error("not test1")
	}
}
