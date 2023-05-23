package ch2

import "time"

type Customer struct {
    generationTimestamp   time.Time
    modificationTimestamp time.Time
    recordId              string
    // ...
}
