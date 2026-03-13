package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/OpenListTeam/OpenList/v4/internal/bootstrap"
	"github.com/OpenListTeam/OpenList/v4/wing/filehash"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	hashJson  string
	hashForce bool
)

var HashCmd = &cobra.Command{
	Use:   "hash [path]",
	Short: "Calculate and store file hashes in the metadata database",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.InitConfig()
		if err := filehash.Init(); err != nil {
			logrus.Fatalf("failed to init filehash db: %v", err)
		}
		defer filehash.Close()

		targetPath, err := filepath.Abs(args[0])
		if err != nil {
			logrus.Fatalf("failed to get absolute path: %v", err)
		}

		var results []*filehash.FileMetadata

		err = filepath.Walk(targetPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			// Check cache
			if !hashForce {
				if m, err := filehash.GetMetadataByPath(path, info.Size()); err == nil {
					logrus.Infof("Found in cache: %s", path)
					results = append(results, m)
					return nil
				}
			}

			logrus.Infof("Calculating hash for: %s", path)
			m, err := filehash.CalculateMetadata(path)
			if err != nil {
				logrus.Errorf("failed to calculate metadata for %s: %v", path, err)
				return nil
			}

			if err := filehash.SaveMetadata(m, path); err != nil {
				logrus.Errorf("failed to save metadata for %s: %v", path, err)
			}

			results = append(results, m)
			return nil
		})

		if err != nil {
			logrus.Errorf("error walking path: %v", err)
		}

		if hashJson != "" {
			f, err := os.Create(hashJson)
			if err != nil {
				logrus.Fatalf("failed to create json file: %v", err)
			}
			defer f.Close()
			enc := json.NewEncoder(f)
			enc.SetIndent("", "  ")
			if err := enc.Encode(results); err != nil {
				logrus.Fatalf("failed to encode json: %v", err)
			}
			logrus.Infof("Results exported to %s", hashJson)
		} else {
			for _, m := range results {
				fmt.Printf("--- %s ---\n%s\n", m.Name, m.ToHumanReadable())
			}
		}
	},
}

func init() {
	HashCmd.Flags().StringVarP(&hashJson, "json", "j", "", "export results to a json file")
	HashCmd.Flags().BoolVarP(&hashForce, "force", "f", false, "force recalculate hash")
	RootCmd.AddCommand(HashCmd)
}
