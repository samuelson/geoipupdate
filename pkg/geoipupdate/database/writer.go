package database

// ZeroMD5 is the default value provided as an MD5 hash for a non-existent
// database.
const ZeroMD5 = "00000000000000000000000000000000"

// Writer provides an interface for writing a database to a target location.
type Writer interface {
	Write(*ReadResult) error
	GetHash(editionID string) (string, error)
}
