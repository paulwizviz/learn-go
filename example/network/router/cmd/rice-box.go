package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "bundle.js",
		FileModTime: time.Unix(1583335519, 0),

		Content: string("function msg(){\n    alert(\"Hello javascript\")\n}"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1583335552, 0),

		Content: string("<html>\n\n<head>\n    <script type=\"text/javascript\" src=\"message.js\"></script>\n</head>\n\n<body>\n    <p>Welcome to JavaScript</p>\n    <form>\n        <input type=\"button\" value=\"click\" onclick=\"msg()\" />\n    </form>\n</body>\n\n</html>"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1583334696, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "bundle.js"
			file3, // "index.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`../web`, &embedded.EmbeddedBox{
		Name: `../web`,
		Time: time.Unix(1583334696, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"bundle.js":  file2,
			"index.html": file3,
		},
	})
}
