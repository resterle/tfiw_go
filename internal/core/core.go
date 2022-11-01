package core

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"reflect"
)

func RandomId() string {
	b := make([]byte, 5)
	rand.Read(b)
	return base32.HexEncoding.EncodeToString(b)
}

type Option[T any] struct {
	value T
	err   error
}

type RR[T, U any] struct {
	value T
	err   error
	op    P[T, U]
	next  RR[U, any]
}

func Crate[T, U any](value T, op P[T, U]) RR[T, U] {
	return RR[T, U]{
		value: value,
		op:    op,
	}
}

func (r RR[T, U]) Apr(op P[U, any]) RR[U, any] {
	x := RR[U, any]{
		op: op,
	}
	r.next = x
	return x
}

func Wrap[T any](value T) Option[T] {
	return Option[T]{
		value: value,
		err:   nil,
	}
}

func (o Option[T]) Unwrap() (T, error) {
	return o.value, o.err
}

func Map[T, U any](o Option[T], f func(T) (U, error)) Option[U] {
	if o.err == nil {
		r, err := f(o.value)
		return Option[U]{
			value: r,
			err:   err,
		}
	}
	return Option[U]{
		err: o.err,
	}
}

func MapS[T, U any](o Option[T], f func(T) U) Option[U] {
	if o.err == nil {
		r := f(o.value)
		return Option[U]{
			value: r,
			err:   nil,
		}
	}
	return Option[U]{
		err: o.err,
	}
}

func Apply[T any](x Option[T], f func(T) error) Option[T] {
	if x.err == nil {
		err := f(x.value)
		return Option[T]{
			value: x.value,
			err:   err,
		}
	}
	return x
}

type P[T, U any] interface {
	Apply(T) (U, error)
}

type PS[T, U any] struct {
	f func(T) (U, error)
}

func Map2[T, U any](f func(T) (U, error)) PS[T, U] {
	return PS[T, U]{f: f}
}

func Apply2[T any](f func(T) error) PS[T, T] {
	inf := func(v T) (T, error) {
		err := f(v)
		return v, err
	}
	return PS[T, T]{f: inf}
}

func (p PS[T, U]) Apply(t T) (U, error) {
	return p.f(t)
}

func MMulti[T, U any](o Option[T], p P[T, U]) Option[U] {
	v, _ := p.Apply(o.value)
	return Wrap(v)
}

func Multi[T, U any](o Option[T], processors ...any) Option[U] {

	value, err := o.Unwrap()
	if err != nil {
		return Option[U]{err: errors.New("WTF")}
	}

	var x any = value

	for _, p := range processors {
		t := reflect.ValueOf(p)
		z := []reflect.Value{reflect.ValueOf(x)}
		m := t.MethodByName("Apply")
		if !m.IsValid() {
			return Option[U]{err: errors.New("Apply method not found")}
		}
		res := m.Call(z)
		err1 := res[1]

		if err1.IsNil() {
			x = res[0].Interface()
		} else {
			return Option[U]{err: err1.Interface().(error)}
		}

	}
	vc, ok := x.(U)
	if ok {
		return Option[U]{value: vc, err: nil}
	}
	return Option[U]{err: errors.New("Type mismatch in Multi")}
}
