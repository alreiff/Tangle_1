package main

import (
  "fmt"
  "log"
  "os"
  "runtime"

  badger "github.com/dgraph-io/badger/v2"
  options "github.com/dgraph-io/badger/v2/options"
)

func main() {
  dirname := "mainnetdb_bob"
  /**
  // Open the Badger database located in the directory.
  // It will be created if it doesn't exist.
  db, err := badger.Open(badger.DefaultOptions(dirname))
  if err != nil {
	  log.Fatal(err)
  }
  **/

  // The following code is from GoShimmer's source code


  /**
  // assure that the directory exists
	err := createDir(dirname)
	if err != nil {
		log.Fatal(fmt.Errorf("could not create DB directory: %w", err))
	}
  **/
	opts := badger.DefaultOptions(dirname)

	opts.Logger = nil
	opts.SyncWrites = false
	opts.TableLoadingMode = options.MemoryMap
	opts.ValueLogLoadingMode = options.MemoryMap
	opts.CompactL0OnClose = false
	opts.KeepL0InMemory = false
	opts.VerifyValueChecksum = false
	opts.ZSTDCompressionLevel = 1
	opts.Compression = options.None
	opts.MaxCacheSize = 50000000
	opts.EventLogging = false
  if runtime.GOOS == "windows" {
		opts = opts.WithTruncate(true)
	}

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(fmt.Errorf("could not open DB: %w", err))
	}

  // End of GoShimmer's source code

  defer db.Close()

  // Read the database
  err = db.View(func(txn *badger.Txn) error {
    opts := badger.DefaultIteratorOptions
    opts.PrefetchSize = 10
    it := txn.NewIterator(opts)
    defer it.Close()
    for it.Rewind(); it.Valid(); it.Next() {
      item := it.Item()
      k := item.Key()
      err := item.Value(func(v []byte) error {
        fmt.Printf("key=%s, value=%s\n", k, v)
        return nil
      })
      if err != nil {
        return err
      }
    }
    return nil
  })
}
