package bash

import (
"fmt"
"os/exec"
)

func Execute(cmds ...string) {

	for _, cmd:= range cmds{
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			panic("some error found")
		}
		fmt.Printf("%s", out)
	}
}