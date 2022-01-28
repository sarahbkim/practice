package main

import (
	"testing"
)

func TestTrie_Search(t *testing.T) {
	type args struct {
		commands []string
		inputs   []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "basic",
			args: args{
				commands: commands,
				inputs:   inputs,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := NewTrie()
			dict := map[string]struct{}{}
			for i := 0; i < len(tt.args.commands); i++ {
				switch tt.args.commands[i] {
				case "insert":
					trie.Insert(tt.args.inputs[i])
					dict[inputs[i]] = struct{}{}
				case "search":
					found := trie.Search(tt.args.inputs[i])
					_, ok := dict[tt.args.inputs[i]]
					if !(found == ok) {
						t.Fatalf("%d: expected %s search to be %t", i, tt.args.inputs[i], ok)
					}
				default:
					t.Fatal("unknown command")
				}
			}
		})
	}
}
