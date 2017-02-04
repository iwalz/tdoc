package image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSVGRewrite(t *testing.T) {
	d := `<?xml version="1.0" encoding="utf-8"?>
<!-- Generator: Adobe Illustrator 19.1.0, SVG Export Plug-In . SVG Version: 6.00 Build 0)  -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="100px"
	 height="100px" viewBox="0 0 100 100" style="enable-background:new 0 0 100 100;" xml:space="preserve">
<style type="text/css">
.st0{fill:#876929;}
.st1{fill:#624A1E;}
.st2{fill:#FAD791;}
.st3{fill:#D9A741;}
.st4{enable-background:new    ;}
.st5{fill:#FFFFFF;}
.st6{clip-path:url(#SVGID_2_);enable-background:new    ;}
.st7{clip-path:url(#SVGID_4_);enable-background:new    ;}
.st8{clip-path:url(#SVGID_6_);enable-background:new    ;}
</style>
<g id="Layer_1">
<polygon class="st1" points="15.9,45.1 32.5,45.9 43.2,45.2 25.5,21.2 	"/>
<polygon class="st2" points="15.9,55.2 32.5,54.5 43.2,55.2 25.5,79.2 	"/>
<polygon class="st3" points="74.5,44.2 84.1,45.2 84.1,25.9 74.5,21.1 	"/>
<polygon class="st0" points="56.8,25.9 74.5,21.1 74.5,44.2 56.8,45.2 	"/>
</g>
<g id="Layer_2">
</g>
</svg>`
	r := NewRewrite()
	r.SetX(20)
	r.SetY(89)
	r.SetWidth(800)
	r.SetHeight(800)
	r.Data(d)
	assert.Equal(t, `<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x='20px' y='89px' width='800px' 	 height='800px' viewBox="0 0 100 100" style="enable-background:new 0 0 100 100;" xml:space="preserve">
<style type="text/css">
.st0{fill:#876929;}
.st1{fill:#624A1E;}
.st2{fill:#FAD791;}
.st3{fill:#D9A741;}
.st4{enable-background:new    ;}
.st5{fill:#FFFFFF;}
.st6{clip-path:url(#SVGID_2_);enable-background:new    ;}
.st7{clip-path:url(#SVGID_4_);enable-background:new    ;}
.st8{clip-path:url(#SVGID_6_);enable-background:new    ;}
</style>
<g id="Layer_1">
<polygon class="st1" points="15.9,45.1 32.5,45.9 43.2,45.2 25.5,21.2 	"/>
<polygon class="st2" points="15.9,55.2 32.5,54.5 43.2,55.2 25.5,79.2 	"/>
<polygon class="st3" points="74.5,44.2 84.1,45.2 84.1,25.9 74.5,21.1 	"/>
<polygon class="st0" points="56.8,25.9 74.5,21.1 74.5,44.2 56.8,45.2 	"/>
</g>
<g id="Layer_2">
</g>
</svg>`, r.Rewrite())

	assert.NotNil(t, r)
	assert.Equal(t, 20, r.X())
	assert.Equal(t, 89, r.Y())
	assert.Equal(t, 800, r.Width())
	assert.Equal(t, 800, r.Height())
}

func TestStyleRewrite(t *testing.T) {
	d := `<?xml version="1.0" encoding="utf-8"?>
<!-- Generator: Adobe Illustrator 19.1.0, SVG Export Plug-In . SVG Version: 6.00 Build 0)  -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="100px"
	 height="100px" viewBox="0 0 100 100" style="enable-background:new 0 0 100 100;" xml:space="preserve">
<style type="text/css">
.st0{fill:#876929;}
.st1{fill:#624A1E;}
.st2{fill:#FAD791;}
.st3{fill:#D9A741;}
.st4{enable-background:new    ;}
.st5{fill:#FFFFFF;}
.st6{clip-path:url(#SVGID_2_);enable-background:new    ;}
.st7{clip-path:url(#SVGID_4_);enable-background:new    ;}
.st8{clip-path:url(#SVGID_6_);enable-background:new    ;}
</style>
<g id="Layer_1">
<polygon class="st1" points="15.9,45.1 32.5,45.9 43.2,45.2 25.5,21.2 	"/>
<polygon class="st2" points="15.9,55.2 32.5,54.5 43.2,55.2 25.5,79.2 	"/>
<polygon class="st3" points="74.5,44.2 84.1,45.2 84.1,25.9 74.5,21.1 	"/>
<polygon class="st0" points="56.8,25.9 74.5,21.1 74.5,44.2 56.8,45.2 	"/>
</g>
<g id="Layer_2">
</g>
</svg>`
	r := NewRewrite()
	r.SetX(20)
	r.SetY(20)
	r.SetWidth(200)
	r.SetHeight(200)
	r.SetName("test")
	r.Data(d)
	assert.Equal(t, `<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x='20px' y='20px' width='200px' 	 height='200px' viewBox="0 0 100 100" style="enable-background:new 0 0 100 100;" xml:space="preserve">
<style type="text/css">
.test_st0{fill:#876929;}
.test_st1{fill:#624A1E;}
.test_st2{fill:#FAD791;}
.test_st3{fill:#D9A741;}
.st4{enable-background:new    ;}
.st5{fill:#FFFFFF;}
.st6{clip-path:url(#SVGID_2_);enable-background:new    ;}
.st7{clip-path:url(#SVGID_4_);enable-background:new    ;}
.st8{clip-path:url(#SVGID_6_);enable-background:new    ;}
</style>
<g id="Layer_1">
<polygon class="test_st1" points="15.9,45.1 32.5,45.9 43.2,45.2 25.5,21.2 	"/>
<polygon class="test_st2" points="15.9,55.2 32.5,54.5 43.2,55.2 25.5,79.2 	"/>
<polygon class="test_st3" points="74.5,44.2 84.1,45.2 84.1,25.9 74.5,21.1 	"/>
<polygon class="test_st0" points="56.8,25.9 74.5,21.1 74.5,44.2 56.8,45.2 	"/>
</g>
<g id="Layer_2">
</g>
</svg>`, r.Rewrite())

	assert.NotNil(t, r)
	assert.Equal(t, 20, r.X())
	assert.Equal(t, 20, r.Y())
	assert.Equal(t, 200, r.Width())
	assert.Equal(t, 200, r.Height())
	assert.Equal(t, "test", r.Name())
}

func TestAppendixRewrite(t *testing.T) {
	d := `<?xml version="1.0" encoding="utf-8"?>
<!-- Generator: Adobe Illustrator 19.1.0, SVG Export Plug-In . SVG Version: 6.00 Build 0)  -->
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 100 100" style="enable-background:new 0 0 100 100;" xml:space="preserve">
<style type="text/css">
.st0{fill:#876929;}
.st1{fill:#624A1E;}
.st2{fill:#FAD791;}
.st3{fill:#D9A741;}
.st4{enable-background:new    ;}
.st5{fill:#FFFFFF;}
.st6{clip-path:url(#SVGID_2_);enable-background:new    ;}
.st7{clip-path:url(#SVGID_4_);enable-background:new    ;}
.st8{clip-path:url(#SVGID_6_);enable-background:new    ;}
</style>
<g id="Layer_1">
<polygon class="st1" points="15.9,45.1 32.5,45.9 43.2,45.2 25.5,21.2 	"/>
<polygon class="st2" points="15.9,55.2 32.5,54.5 43.2,55.2 25.5,79.2 	"/>
<polygon class="st3" points="74.5,44.2 84.1,45.2 84.1,25.9 74.5,21.1 	"/>
<polygon class="st0" points="56.8,25.9 74.5,21.1 74.5,44.2 56.8,45.2 	"/>
</g>
<g id="Layer_2">
</g>
</svg>`
	r := NewRewrite()
	r.SetX(20)
	r.SetY(20)
	r.SetWidth(200)
	r.SetHeight(200)
	r.SetName("test")
	r.Data(d)
	assert.Equal(t, `<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 100 100" style="enable-background:new 0 0 100 100;" xml:space="preserve" width='200px'  height='200px'  x='20px'  y='20px' >
<style type="text/css">
.test_st0{fill:#876929;}
.test_st1{fill:#624A1E;}
.test_st2{fill:#FAD791;}
.test_st3{fill:#D9A741;}
.st4{enable-background:new    ;}
.st5{fill:#FFFFFF;}
.st6{clip-path:url(#SVGID_2_);enable-background:new    ;}
.st7{clip-path:url(#SVGID_4_);enable-background:new    ;}
.st8{clip-path:url(#SVGID_6_);enable-background:new    ;}
</style>
<g id="Layer_1">
<polygon class="test_st1" points="15.9,45.1 32.5,45.9 43.2,45.2 25.5,21.2 	"/>
<polygon class="test_st2" points="15.9,55.2 32.5,54.5 43.2,55.2 25.5,79.2 	"/>
<polygon class="test_st3" points="74.5,44.2 84.1,45.2 84.1,25.9 74.5,21.1 	"/>
<polygon class="test_st0" points="56.8,25.9 74.5,21.1 74.5,44.2 56.8,45.2 	"/>
</g>
<g id="Layer_2">
</g>
</svg>`, r.Rewrite())

	assert.NotNil(t, r)
	assert.Equal(t, 20, r.X())
	assert.Equal(t, 20, r.Y())
	assert.Equal(t, 200, r.Width())
	assert.Equal(t, 200, r.Height())
	assert.Equal(t, "test", r.Name())

	r.Data("")
	r.Rewrite()
}
