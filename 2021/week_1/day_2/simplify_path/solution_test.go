package simplify_path

import "testing"

func Test_simplifyPath(t *testing.T) {
	t.Log(simplifyPath("/home/") == "/home")
	t.Log(simplifyPath("/../") == "/")
	t.Log(simplifyPath("/home//foo/") == "/home/foo")
	t.Log(simplifyPath("/a/./b/../../c/") == "/c")
	t.Log(simplifyPath("/a//b////c/d//././/..") == "/a/b/c")
	t.Log(simplifyPath("/.../") == "/...")
	t.Log(simplifyPath("///") == "/")
	t.Log(simplifyPath("/..hidden") == "/..hidden")

}
