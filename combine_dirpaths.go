package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const INDENT_SPACES = 2;

type Node struct {
  ParentNode  *Node;
  Name        string;
  Level       int;
  ChildNodes  []*Node;
};

func printNode(node *Node, spaces int) {
  _spaces := strings.Repeat(" ", spaces);
  for _, val := range node.ChildNodes {
    fmt.Printf("%s %s\n", _spaces, val.Name);
    printNode(val, spaces + 1);
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

// Helper function which backtracks a given node, constructing it's full path to the
// parent nodes.
func backtrackNodeName(node *Node) string {
  if node.ParentNode != nil {
    return backtrackNodeName(node.ParentNode) + node.Name;
  }
  return node.Name;
}

func flattenMapToPaths (node *Node) []string {
  paths := []string{};

  // Interative DFS path construction.
  stack := []*Node{node};
  discovered := []*Node{};
  for len(stack) != 0 {
    // Pop off the stack.
    parentNode := stack[len(stack) - 1];
    stack = stack[:len(stack) - 1];

    // Check if node was already traversed.
    if !eltInArray(parentNode, stack) {
      discovered = append(discovered, parentNode);

      for _, childNode := range parentNode.ChildNodes {
        // fmt.Println(parentNode.Name, childNode.Name, len(childNode.ChildNodes));
        // Construct path of leaf nodes.
        if len(childNode.ChildNodes) == 0 {
          path := backtrackNodeName(childNode);
          paths = append(paths, path);
          // fmt.Println(" - ", path);
        }
        stack = append(stack, childNode)
      }
    }
  }

  return paths;
}

func findNodeWithLevel(node *Node, level int) *Node {
  n := node;
  for n != nil {
    if n.Level == level {
      return n;
    }
    n = n.ParentNode;
  }
  return n;
}

func _processRecursivePaths(scanner *bufio.Scanner, parentNode *Node) {
  prefixSpacesRe := regexp.MustCompile("^\\s+");

	for scanner.Scan() {
		line := scanner.Text();
    prefixSpacesMatch := prefixSpacesRe.FindAllString(line, -1)
    currentLevel := 0
    if len(prefixSpacesMatch) > 0 {
      currentLevel = len(prefixSpacesMatch[0]) / INDENT_SPACES;
    }

    // Keep track of children
    childNode := &Node{};
    childNode.Name = strings.Trim(line, " ");
    childNode.Level = currentLevel;

    // Recursively include sub-directories.
		if strings.HasSuffix(line, "/") {
      // Go up one level
      actualParent := parentNode;
      if currentLevel != 0 && currentLevel == parentNode.Level {
        actualParent = findNodeWithLevel(parentNode, currentLevel - 1);
      }

      // We are at the root level.
      if currentLevel == 0{
        actualParent = findNodeWithLevel(parentNode, -1);
      }

      // Apply child to the correct parent.
      childNode.ParentNode = actualParent;
      actualParent.ChildNodes = append(actualParent.ChildNodes, childNode);

      // New recursive directory node.
			_processRecursivePaths(scanner, childNode);
		} else {
      parentNode.ChildNodes = append(parentNode.ChildNodes, childNode);
      childNode.ParentNode = parentNode;
    }
	}
}

func processDirectory(scanner *bufio.Scanner) []string {
  pathMp := &Node{
    Level: -1,
  };
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

  paths := processDirectory(scanner);
  for _, path := range paths {
    fmt.Println(path);
  }
}

