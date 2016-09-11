// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command(
		"python3", "-c",
		"import json, dis, sys; json.dump(dis.opname, fp=sys.stdout)",
	)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	var opnames []string
	err = json.NewDecoder(buf).Decode(&opnames)
	if err != nil {
		log.Fatalf("error decoding opnames: %v\n", err)
	}

	o, err := os.Create("opcodes_gen.go")
	if err != nil {
		log.Fatal(err)
	}
	defer o.Close()

	_, err = o.WriteString(header)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(o, "// Python-3 Opcodes\nconst (\n")
	for i, v := range opnames {
		if v[0] == '<' {
			v = fmt.Sprintf("%03d", i)
		}
		fmt.Fprintf(o, "\tOp_%s Opcode = %d\n", v, i)
	}
	fmt.Fprint(o, ")\n\n")

	fmt.Fprint(o, "func (op Opcode) String() string {\n\tswitch op {\n")
	for i, v := range opnames {
		fmt.Fprintf(o, "\t\tcase %d: return %q\n", i, v)
	}
	fmt.Fprint(o, "\tdefault:panic(fmt.Errorf(\"invalid opcode value %%d\",byte(op)))\n}\nreturn \"\"}\n\n")

	err = o.Close()
	if err != nil {
		log.Fatal(err)
	}

}

const header = `// Copyright 2016 The pygo Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pygo

`
