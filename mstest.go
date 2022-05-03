package main

import (
	"errors"
	"fmt"
)

// social: make([]byte, 9)
// zip: make([]byte, 10)
type SSNode struct {
	social [9]byte
	zip    [10]byte
	next   *SSNode
}
type SSLinkedList interface {
	Find(socialSec string) (*SSNode, error)
	Add(socialSec string, zipcode string) error
	Remove(socialSec string) error
}

func NewSocialSecList(nNodes int) SSLinkedList {
	var head = &SSNode{social: [9]byte{}, zip: [10]byte{}}
	var curr = head
	for i := 1; i < nNodes; i++ {
		curr.next = &SSNode{social: [9]byte{}, zip: [10]byte{}}
		curr = curr.next
	}
	return head
}

func (n *SSNode) Find(socialSec string) (*SSNode, error) {
	if !valid(socialSec) {
		return nil, errors.New("invalid key")
	}
	var curr = n
	for curr != nil {
		if match(curr.social, socialSec) {
			return curr, nil
		}
		curr = curr.next
	}
	return nil, nil
}

func match(social [9]byte, socialStr string) bool {
	return true
}

func valid(socialSec string) bool {
	return true
}

func (n *SSNode) Add(socialSec string, zipcode string) error {
	var curr = n
	var copied bool
	for curr != nil {
		if cap(curr.social) > 0 {
			copy(curr.social, []byte(socialSec))
			copy(curr.zip, []byte(zipcode))
			copied = true
			break
		}
		curr = curr.next
	}
	if !copied {
		return errors.New("out of mem")
	}
	return nil
}

func (n *SSNode) Remove(socialSec string) error {
	var curr = n
	var found bool
	for curr != nil {
		if match(curr.social, socialSec) {
			curr.social = curr.social[:]
			curr.zip = curr.zip[:]
			found = true
			break
		}
		curr = curr.next
	}
	if !found {
		return fmt.Errorf("%s social not found", socialSec)
	}
	return nil
}
