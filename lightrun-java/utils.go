package lightrunJava

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	rt "runtime"
	"strings"
)

var javaVersionCommand = []string{"java", "-version"}
var macOSJavaTestCommand = "/usr/libexec/java_home"

func checkForMacJVM() (ok bool, err error) {
	// test for java differently because of the shim application
	var out bytes.Buffer
	cmd := exec.Command(macOSJavaTestCommand)
	cmd.Stdout = &out
	cmd.Stderr = &out
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("failed check for macOS jvm: %v, out: %v", err, strings.Replace(strings.Replace(out.String(), "\n", " ", -1), `"`, `\"`, -1))
		return false, err
	}
	return true, nil
}

func javaVersionInfo() (version, runtime, vm string, err error) {
	var out bytes.Buffer

	if rt.GOOS == "darwin" {
		_, err = checkForMacJVM()
		if err != nil {
			err = fmt.Errorf("failed to check java version: %v", err)
			return
		}
	}

	cmd := exec.Command(javaVersionCommand[0], javaVersionCommand[1:]...)
	cmd.Stdout = &out
	cmd.Stderr = &out
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("failed to check java version: %v", err)
		return
	}

	version, runtime, vm = parseJavaVersionOutput(out.String())
	return
}

func parseJavaVersionOutput(infoString string) (version, runtime, vm string) {
	infoString = strings.TrimSpace(infoString)

	lines := strings.Split(infoString, "\n")
	if strings.Contains(lines[0], "Picked up _JAVA_OPTIONS") {
		lines = lines[1:]
	}

	if len(lines) < 3 {
		// unexpected output format, don't attempt to parse output for version
		return "", "", ""
	}

	versionString := strings.TrimSpace(lines[0])

	re := regexp.MustCompile(`version "([^"]*)"`)
	if match := re.FindStringSubmatch(lines[0]); len(match) == 2 {
		versionString = match[1]
	}

	return versionString, strings.TrimSpace(lines[1]), strings.TrimSpace(lines[2])
}

func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}