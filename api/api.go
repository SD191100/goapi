package api

import (
  "encoding/json";
  "net/http";
)

// coin balance params
type CoinBalanceParams struct {
  Username string;
}

type CoinBalanceResponse struct {
  Code int;

  Balance int64;
}

type Error struct {
  Code int;

  Message string;
}
