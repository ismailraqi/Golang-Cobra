package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/kopy"

	"github.com/spf13/cobra"
)

var comdirCmd = &cobra.Command{
	Use:   "comdir",
	Short: "Compress the directory or a folder.",
	Long:  `comdir this is a long description for this comdir command.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// To make directory path separator a universal, in Linux "/" and in Windows "\" to auto change
		// depends on the user's OS using the filepath.FromSlash organic Go's library.
		src := filepath.FromSlash(args[0])
		dst := filepath.FromSlash(args[1])

		IgnoreFilesOrFolders := []string{".txt", ".jpg", "folder_name"}

		msg := `Start compressing the directory or a folder:`
		fmt.Println(msg, src)
		itrlog.Infow(msg, "src", src, "dst", dst, "log_time", time.Now().Format(itrlog.LogTimeFormat))

		// Compose the zip filename
		fnWOext := kopy.FileNameWOExt(filepath.Base(src)) // Returns a filename without an extension.
		zipDir := fnWOext + kopy.ComFileFormat

		// To make directory path separator a universal, in Linux "/" and in Windows "\" to auto change
		// depends on the user's OS using the filepath.FromSlash organic Go's library.
		zipDest := filepath.FromSlash(path.Join(dst, zipDir))

		// Start compressing the entire directory or a folder using the tar + gzip
		var buf bytes.Buffer
		if err := kopy.CompressDIR(src, &buf, IgnoreFilesOrFolders); err != nil {
			fmt.Println(err)
			itrlog.Errorw("error", "err", err, "log_time", time.Now().Format(itrlog.LogTimeFormat))
			return
		}

		// write the .tar.gzip
		os.MkdirAll(dst, os.ModePerm) // Create the root folder first
		fileToWrite, err := os.OpenFile(zipDest, os.O_CREATE|os.O_RDWR, os.FileMode(600))
		if err != nil {
			itrlog.Errorw("error", "err", err, "log_time", time.Now().Format(itrlog.LogTimeFormat))
			panic(err)
		}
		if _, err := io.Copy(fileToWrite, &buf); err != nil {
			itrlog.Errorw("error", "err", err, "log_time", time.Now().Format(itrlog.LogTimeFormat))
			panic(err)
		}
		defer fileToWrite.Close()

		msg = `Done compressing the directory or a folder:`
		fmt.Println(msg, src)
		itrlog.Infow(msg, "dst", zipDest, "log_time", time.Now().Format(itrlog.LogTimeFormat))
	},
}

func init() {
	rootCmd.AddCommand(comdirCmd)
}
