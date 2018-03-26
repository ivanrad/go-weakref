package weakref

import (
	"log"
	"runtime"
	"testing"
)

type Point3D struct {
	x, y, z int
}

func makeWeakReference() *WeakRef {
	point := Point3D{1, 2, 3}
	weakRef := NewWeakRef(&point)
	return weakRef
}
func TestWeakRef(t *testing.T) {
	weakRef := makeWeakReference()
	if weakRef == nil {
		t.Fatal("unexpected: weakRef is nil")
	}
	if weakRef.GetTarget() == nil {
		t.Fatal("unexpected: weakRef.Get() is nil")
	}
	t.Log(weakRef.GetTarget().(*Point3D))
	if weakRef.IsAlive() == false {
		t.Fatal("unexpected: weakRef.IsAlive() is false")
	}

	// ouch
	for i := 1; i < 10; i++ {
		runtime.Gosched()
		runtime.GC()
	}
	if weakRef == nil {
		log.Fatal("unexpected: weakRef is nil after GC")
	}
	if weakRef.GetTarget() != nil {
		log.Fatal("unexpected: weakRef.Get() is not nil after GC")
	}
	if weakRef.IsAlive() == true {
		log.Fatal("unexpected: weakRef.IsAlive() is true")
	}
}
