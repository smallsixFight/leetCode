package simplify_path

import "testing"

func TestSimplifyPath(t *testing.T) {
	t.Log("/111" == SimplifyPath("111"))
	t.Log("/home" == SimplifyPath("/home/"))
	t.Log("/" == SimplifyPath("/../"))
	t.Log("/home/foo" == SimplifyPath("/home//foo/"))
	t.Log("/c" == SimplifyPath("/a/./b/../../c/"))
	t.Log("/c" == SimplifyPath("/a/../../b/../c//.//"))
	t.Log("/a/b/c" == SimplifyPath("/a//b////c/d//././/.."))
	t.Log(SimplifyPath("/a//b////c/d//././/../.."))
	t.Log(SimplifyPath("///"))
	t.Log(SimplifyPath("/..."))
	t.Log(SimplifyPath("/.a"))
	t.Log(SimplifyPath("/..hidden"))
}
