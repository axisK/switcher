package main

import (
  "bytes"
  )

  type IRC string

  // address to proxy to
  func (s IRC) Address() string {
    return string(s)
  }

  // identify header as one of IRC
  func (s IRC) Identify(header []byte) bool {
    // first 3 bytes of an IRC connection is NIC, USE, or PAS in case of ZNC
    if bytes.Compare(header, []byte("NIC")) == 0 || bytes.Compare(header, []byte("USE")) == 0 || bytes.Compare(header, []byte("PAS")) == 0 {
      return true
    }

    return false
  }
