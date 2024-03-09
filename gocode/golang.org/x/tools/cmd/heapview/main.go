// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// heapview is a tool for viewing Go heap dumps.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

var host = flag.String("host", "", "host addr to listen on")
var port = flag.Int("port", 8080, "service port")

var index = `<!DOCTYPE html>
<script src="js/customelements.js"></script>
<script src="js/typescript.js"></script>
<script src="js/moduleloader.js"></script>
<script>
  System.transpiler = 'typescript';
  System.typescriptOptions = {target: ts.ScriptTarget.ES2015};
  System.locate = (load) => load.name + '.ts';
</script>
<script type="module">
  import {main} from './client/main';
  main();
</script>
`

func toolsDir() string {
	return "/usr/share/gocode/src/golang.org/x/tools"
}

var parseFlags = func() {
	flag.Parse()
}

var addHandlers = func() {
	// Directly serve typescript code in client directory for development.
	http.Handle("/client/", http.StripPrefix("/client",
		http.FileServer(http.Dir(filepath.Join(toolsDir(), "cmd/heapview/client")))))

	// Serve typescript.js and moduleloader.js for development.
	http.HandleFunc("/js/typescript.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/usr/lib/nodejs/typescript/typescript.js")
	})
	http.HandleFunc("/js/moduleloader.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/usr/lib/javascript/es-module-loader-0.17/es6-module-loader.js")
	})
	http.HandleFunc("/js/customelements.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/usr/lib/javascript/webcomponentsjs-custom-element-v0/CustomElements.js")
	})

	// Serve index.html using html string above.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, index)
	})
}

var listenAndServe = func() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), nil)
}

func main() {
	parseFlags()
	addHandlers()
	log.Fatal(listenAndServe())
}
