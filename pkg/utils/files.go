package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

func CreateBackup(backupN string, username string, database string) error {
	dir := "backups"
	//Mkdir makes a file inside an existing folder
	// o755 is the permisiion mode to allow creation
	// ISExist checks wheter the arhument(err in thsis case) says the file/path exists or not
	if err := os.Mkdir(dir, 0755); err != nil && !os.IsExist(err) {
		return err
	}
	filename := backupN + ".dump"
	f, err := os.Create(filepath.Join(dir, filename))
	if err != nil {
		return err
	}
	//to close the file once its done
	defer f.Close()
	bckupCmd := exec.Command(
		"pg_dump",
		"-U", username,
		"-F", "c",
		database,
	)
	bckupCmd.Stdout = f         //writes file returned to f
	bckupCmd.Stderr = os.Stderr //shows error if any
	if err := bckupCmd.Run(); err != nil {
		return err
	}
	return nil
}
