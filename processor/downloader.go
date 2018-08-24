package processor

import (
	"log"
	"github.com/chapterzero/gomposer/provider"
	"io"
	"math/rand"
	"net/http"
	"fmt"
	"strings"
	"os"
	"time"
	"archive/zip"
	"path/filepath"
	"sync"
)

const tempDirectory = "/tmp"

var i int = 2;

type DownloadResult struct {
	status     int
	filePath   string
	err        error
	dependency Dependency
}

func downloadPackagesParallel(dependencies []Dependency, vendorDirectory string) {
	d := make(chan DownloadResult)
	var wg sync.WaitGroup
	wg.Add(len(dependencies))
	rand.Seed(time.Now().UTC().UnixNano())

	for _, dependency := range dependencies {
		go func(dependency Dependency) {
			filePath, err := downloadPackage(dependency.FqPackageName, dependency.Provider, dependency.Version, vendorDirectory)
			downloadResult := DownloadResult {
				status: 1,
				err: err,
				filePath: filePath,
				dependency: dependency,
			}
			log.Println("Download complete")
			if downloadResult.err != nil {
				close(d)
				log.Fatalln(downloadResult.err)
			}
			log.Println("Unzipping...")
			_, err = unzipPackage(downloadResult.filePath, getPackageVendorDir(
				downloadResult.dependency.FqPackageName,
				vendorDirectory,
			))
			if err != nil {
				close(d)
				log.Fatalln("Error while unzipping package ", err)
			}

			deleteFile(downloadResult.filePath)

			d <- downloadResult
			wg.Done()
		}(dependency)
	}

	go func() {
		wg.Wait()
		close(d)
	}()

	for range d {
	}
}

func downloadPackage(packageName string, provider provider.Provider, version Version, vendorDirectory string) (string, error) {
	downloadUrl := provider.GetDownloadUrl(packageName, version.Value)
	filepath := fmt.Sprintf("%v/%v_%v.zip", tempDirectory, rand.Intn(300), strings.Replace(packageName, "/", "_", -1))
	log.Println("Downloading package ", downloadUrl)

	// Create the file
    out, err := os.Create(filepath)
    if err != nil {
		return filepath, err
    }
    defer out.Close()

    // Get the data
    resp, err := http.Get(downloadUrl)
    if err != nil {
		return filepath, err
    }
    defer resp.Body.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    if err != nil {
		return filepath, err
    }

	return filepath, err
}

func deleteFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		log.Println("Failed to delete temporary file ")
	} else {
		log.Println("Deleted temporary file")
	}
}

func unzipPackage(src string, dest string) ([]string, error) {
	var filenames []string

    r, err := zip.OpenReader(src)
    if err != nil {
        return filenames, err
    }
    defer r.Close()

	folderCount := 0

    for _, f := range r.File {

		// skip github bundled folder
		if folderCount == 0 {
			folderCount++
			continue
		}

		// remove first folder
		fNameSplitted := strings.Split(f.Name, "/")
		fNameWithoutRoot := strings.Join(fNameSplitted[1:], "/")

        rc, err := f.Open()
        if err != nil {
            return filenames, err
        }
        defer rc.Close()

        // Store filename/path for returning and using later on
        fpath := filepath.Join(dest, fNameWithoutRoot)

        // Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
        if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
            return filenames, fmt.Errorf("%s: illegal file path", fpath)
        }

        filenames = append(filenames, fpath)

        if f.FileInfo().IsDir() {

            // Make Folder
            os.MkdirAll(fpath, os.ModePerm)

        } else {

            // Make File
            if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
                return filenames, err
            }

            outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return filenames, err
            }

            _, err = io.Copy(outFile, rc)

            // Close the file without defer to close before next iteration of loop
            outFile.Close()

            if err != nil {
                return filenames, err
            }

        }
    }
    return filenames, nil
}

func getPackageVendorDir(fqPackageName string, rootVendorDirectory string) (string) {
	return fmt.Sprintf("%v/%v", rootVendorDirectory, fqPackageName)
}
