package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const (
	commentChar string = "#"
	//StartOfSection writ hosts start
	StartOfSection = "# Mirror repository and IP mapping"
	//EndOfSection writ hosts end
	EndOfSection = "# End of Section"
	eol          = "\n"
)

// WriteHosts set rainbond imagehub and ip to local host file
func WriteHosts(hostspath, ip string) error {
	// open hostfile in operator
	hostFile, err := os.OpenFile(hostspath, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	defer hostFile.Close()
	//  check ip and rainbond hub url is exist
	r := bufio.NewReader(hostFile)
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			break
		}
		if line == StartOfSection {
			return nil
		}
	}
	// add rainbond hub url if not exist
	lines := []string{
		eol + StartOfSection + eol,
		ip + " " + "goodrain.me" + eol,
		ip + " " + "region.goodrain.me" + eol,
		EndOfSection + eol,
	}
	writer := bufio.NewWriter(hostFile)
	for _, line := range lines {
		if _, err := writer.WriteString(line); err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}
