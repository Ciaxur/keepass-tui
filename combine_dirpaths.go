package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
  ParentNode *Node;
  Path map[string]*Node;
};

func printNode(node *Node, spaces int) {
  _spaces := strings.Repeat(" ", spaces);
  for key, val := range node.Path {
    fmt.Printf("%s %s\n", _spaces, key);

    if val != nil {
      printNode(val, spaces + 1);
    }
  }
}

// Helper function which checks if element is in an array.
func eltInArray(elt *Node, arr []*Node) bool {
  for _, _elt := range arr {
    if elt == _elt {
      return true;
    }
  }
  return false;
}

func flattenMapToPaths (node *Node) []string {
  paths := []string{};

  // Interative DFS path construction.
  stack := []*Node{node};
  discovered := []*Node{};
  for len(stack) != 0 {
    // Pop off the stack.
    n := stack[len(stack) - 1];
    stack = stack[:len(stack) - 1];

    if n == nil {
      continue;
    }

    // Check if node was already traversed.
    if !eltInArray(n, stack) {
      discovered = append(discovered, n);

      for key, edge := range n.Path {
        if edge == nil {
          fmt.Println(key);
        }
        stack = append(stack, edge)
      }
    }
  }

  return paths;
}

func _processRecursivePaths(scanner *bufio.Scanner, parentNode *Node) {
	for scanner.Scan() {
		line := scanner.Text();
		if strings.HasSuffix(line, "/") {
      // New recursive directory node.
      node := &Node{};
      node.ParentNode = parentNode;
      parentNode.Path[line] = node;
			_processRecursivePaths(scanner, node);
		} else {
      if parentNode.Path == nil {
        parentNode.Path = make(map[string]*Node);
      }

      parentNode.Path[line] = nil;
		}
	}
}

func processDirectory(scanner *bufio.Scanner) []string {
  pathMp := &Node{};
  _processRecursivePaths(scanner, pathMp);
  // printNode(pathMp, 0);
  return flattenMapToPaths(pathMp);
}


func main() {
  // Grab filepath as a cli argument.
  if len(os.Args) < 2 {
    fmt.Println("FAIL: Expected filepath argument");
    os.Exit(1);
  }
  filepath := os.Args[1];

	file, err := os.Open(filepath);
	if err != nil {
		fmt.Println("Error opening file:", err);
		return;
	}
	defer file.Close();

	scanner := bufio.NewScanner(file);

  // TODO: print em yo
  processDirectory(scanner);
}

