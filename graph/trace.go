package graph

import "fmt"

func (g *GraphNode) Trace(src, dst string) ([][]*GraphNode, error) {
	table := *g.Table
	source, ok := table[src]
	if !ok {
		return nil, fmt.Errorf("could not find node for src '%s'", src)
	}

	destination, ok := table[dst]
	if !ok {
		return nil, fmt.Errorf("could not find node for dst '%s'", dst)
	}

	fmt.Printf("finding path from '%s' -> '%s'\n", source.Value, destination.Value)
	return dfs(source, destination)
}

func dfs(src, dst *GraphNode) ([][]*GraphNode, error) {
	res := make([][]*GraphNode, 0)
	vis := make(visited, 0)

	stack := NewStack()
	for _, s := range src.Children {
		stack.Push([]*GraphNode{src, s})
	}

	for stack.Size() > 0 {
		l := stack.Pop()
		if l == nil {
			return nil, fmt.Errorf("stack empty")
		}
		end := l[len(l)-1]

		if vis.contains(end) {
			continue
		}

		// check if the destination has been found
		if end == dst {
			res = append(res, l)
			continue
		}
		vis = append(vis, end)

		// add the current node's children to the stack
		for _, s := range end.Children {
			nextList := make([]*GraphNode, len(l)+1)
			copy(nextList, l)
			nextList[len(l)] = s
			stack.Push(nextList)
		}
	}

	return res, nil
}

type visited []*GraphNode

func (v visited) contains(n *GraphNode) bool {
	for _, _n := range v {
		if _n == n {
			return true
		}
	}
	return false
}
