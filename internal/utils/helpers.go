package util

//Dependencies
import "fmt"

//Converts a uint64 to a human readable string representing a data size
func ByteCountIEC(b int64) string {
    const unit = 1024

	if b < unit {
        return fmt.Sprintf("%dB", b)
    }

    div, exp := int64(unit), 0
    for n := b / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }

    return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
