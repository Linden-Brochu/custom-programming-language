package src

import "os"

func exist(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}
