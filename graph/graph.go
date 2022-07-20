package graph

import (
	"fmt"
	"strings"
)

type Dependency struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (d Dependency) String() string {
	if d.Version != "na" {
		return fmt.Sprintf("%s@%s", d.Name, d.Version)
	} else {
		return d.Name
	}
}

type GraphNode struct {
	Value    *Dependency            `json:"value"`
	Children map[string]*GraphNode  `json:"children"`
	Parents  map[string]*GraphNode  `json:"parents"`
	Table    *map[string]*GraphNode `json:"table"`
}

func New(input string) (*GraphNode, error) {
	table := make(map[string]*GraphNode)

	lines := strings.Split(input, "\n")

	firstLine := lines[0]
	// build root
	left, right, err := splitDependencies(firstLine)
	if err != nil {
		return nil, err
	}
	parent, err := constructDependency(left)
	if err != nil {
		return nil, err
	}
	child, err := constructDependency(right)
	if err != nil {
		return nil, err
	}

	var root *GraphNode
	root = &GraphNode{
		Value: parent,
		Children: map[string]*GraphNode{
			right: {
				Value: child,
				Parents: map[string]*GraphNode{
					left: root,
				},
				Children: make(map[string]*GraphNode),
				Table:    &table,
			}},
		Parents: make(map[string]*GraphNode),
		Table:   &table,
	}
	table[parent.String()] = root

	// iterate through the rest of the lines
	for _, line := range lines[1:] {
		// get dependency and dependent from line
		dependencyStr, dependentStr, err := splitDependencies(line)
		if err != nil {
			return nil, err
		}

		var ok bool
		var parentNode, childNode *GraphNode

		// find the graph node for parent if it exists
		// if not, create it
		if parentNode, ok = table[dependencyStr]; !ok {
			parent, err = constructDependency(dependencyStr)
			if err != nil {
				return nil, err
			}
			parentNode = &GraphNode{
				Value:    parent,
				Children: make(map[string]*GraphNode),
				Parents:  make(map[string]*GraphNode),
				Table:    &table,
			}
			table[dependencyStr] = parentNode
		}

		// find the graph node for the child if it exists
		// if not create it
		if childNode, ok = table[dependentStr]; !ok {
			child, err = constructDependency(dependentStr)
			if err != nil {
				return nil, err
			}
			childNode = &GraphNode{
				Value:    child,
				Children: make(map[string]*GraphNode),
				Parents:  make(map[string]*GraphNode),
				Table:    &table,
			}
			table[dependentStr] = childNode
		}

		// update parent's children and children's parents
		parentNode.Children[dependentStr] = childNode
		childNode.Parents[dependencyStr] = parentNode
	}

	return root, nil
}

func splitDependencies(line string) (string, string, error) {
	if dd := strings.Split(line, " "); len(dd) == 2 {
		return dd[0], dd[1], nil
	}
	return "", "", fmt.Errorf("could not get dependencies from line '%s'", line)
}

func constructDependency(dep string) (*Dependency, error) {
	if nv := strings.Split(dep, "@"); len(nv) == 2 {
		return &Dependency{
			Name:    nv[0],
			Version: nv[1],
		}, nil
	} else if len(nv) == 1 {
		return &Dependency{
			Name:    nv[0],
			Version: "na",
		}, nil
	}
	return nil, fmt.Errorf("could not parse name and version from dependency '%s'", dep)
}
