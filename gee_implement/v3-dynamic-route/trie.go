package main

import "strings"

// -------------------------------------------
// @file          : trie.go
// @author        : binshow
// @time          : 2022/7/25 1:10 PM
// @description   : 前缀树实现
// -------------------------------------------


type node struct {
	pattern  string  	// 待匹配路由 例如 /p/:lang
	part     string		// 路由中的一部分，例如 :lang
	children []*node	// 子节点，例如 [doc, tutorial, intro]
	isWild	 bool		// 是否精确匹配，part 含有 : 或 * 时为true
}


// 在当前节点的子节点中找到 存储 part的节点 ， 注意这里是第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild == true {
			return child
		}
	}
	return nil
}

// 在当前节点的子节点中找到 存储 part的节点 ， 注意这里是所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node , 0)
	for _, child := range n.children {
		if child.part == part || child.isWild == true {
			nodes = append(nodes , child)
		}
	}
	return nodes
}


// 路由的注册和匹配 就分别对应的 trie树的节点的插入和查询

// 节点的插入，递归的查找每一层的节点，如果没有匹配到当前 part 的节点没新建一个
// 需要注意的一点是：/p/:lang/doc只有在第三层节点，即doc节点，pattern才会设置为/p/:lang/doc。p和:lang节点的pattern属性皆为空。
// 因此，当匹配结束时，我们可以使用n.pattern == ""来判断路由规则是否匹配成功。
// 例如，/p/python虽能成功匹配到:lang，但:lang的pattern值为空，因此匹配失败。

func (n *node) insert(pattern string, parts []string, height int) {
	// 递归截止条件：所有的 parts 都插入完了
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:     part,
			isWild:   part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children , child)
	}

	child.insert(pattern , parts , height+1)

}

// 节点的查找: 同样也是递归查询每一层的节点，退出规则是，匹配到了*，匹配失败，或者匹配到了第len(parts)层节点
func (n *node) search(parts []string , height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part , "*") {
		if n.pattern == "" {	// 匹配失败
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts , height+1)
		if result != nil {
			return result
		}
	}
	return nil
}