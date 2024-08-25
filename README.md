### Documentation for `main.go`

This Go program interacts with the `ston.fi` API v1 to fetch various types of data such as assets, farms, pools, and wallet-related information. Below is the documentation for each function along with examples of how to use them.

#### Constants
```go
const baseUrl = "https://api.ston.fi"
```
- `baseUrl`: The base URL for the `ston.fi` API.

#### Functions

##### `getAsset(assetAddress string) (Asset, error)` 1231
Fetches a single asset by its address.

**Parameters:**
- `assetAddress`: The address of the asset.

**Returns:**
- `Asset`: The asset data.
- `error`: Error if the request fails.

**Example:**
```go
asset, err := getAsset("EQD8T***************")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(asset)
```

##### `getAssets() ([]Asset, error)`
Fetches all assets.

**Returns:**
- `[]Asset`: A slice of all assets.
- `error`: Error if the request fails.

**Example:**
```go
assets, err := getAssets()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(assets)
```

##### `getFarms() ([]Farm, error)`
Fetches all farms.

**Returns:**
- `[]Farm`: A slice of all farms.
- `error`: Error if the request fails.

**Example:**
```go
farms, err := getFarms()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(farms)
```

##### `getFarm(farmAddress string) (Farm, error)`
Fetches a single farm by its address.

**Parameters:**
- `farmAddress`: The address of the farm.

**Returns:**
- `Farm`: The farm data.
- `error`: Error if the request fails.

**Example:**
```go
farm, err := getFarm("EQD8T***************")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(farm)
```

##### `getFarmsByPool(poolAddress string) ([]Farm, error)`
Fetches all farms associated with a specific pool.

**Parameters:**
- `poolAddress`: The address of the pool.

**Returns:**
- `[]Farm`: A slice of farms associated with the pool.
- `error`: Error if the request fails.

**Example:**
```go
farms, err := getFarmsByPool("EQD8T***************")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(farms)
```

##### `getSwapPairs() ([]SwapPairs, error)`
Fetches all swap pairs.

**Returns:**
- `[]SwapPairs`: A slice of swap pairs.
- `error`: Error if the request fails.

**Example:**
```go
swapPairs, err := getSwapPairs()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(swapPairs)
```

##### `getPool(poolAddress string) (Pool, error)`
Fetches a single pool by its address.

**Parameters:**
- `poolAddress`: The address of the pool.

**Returns:**
- `Pool`: The pool data.
- `error`: Error if the request fails.

**Example:**
```go
pool, err := getPool("EQD8T***************")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(pool)
```

##### `getPools() ([]Pool, error)`
Fetches all pools.

**Returns:**
- `[]Pool`: A slice of all pools.
- `error`: Error if the request fails.

**Example:**
```go
pools, err := getPools()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(pools)
```

##### `getWalletAsset(walletAddress string, assetAddress string) (Asset, error)`
Fetches a specific asset for a given wallet.

**Parameters:**
- `walletAddress`: The address of the wallet.
- `assetAddress`: The address of the asset.

**Returns:**
- `Asset`: The asset data.
- `error`: Error if the request fails.

**Example:**
```go
asset, err := getWalletAsset("walletAddress", "assetAddress")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(asset)
```

##### `getWalletAssets(walletAddress string) ([]Asset, error)`
Fetches all assets for a given wallet.

**Parameters:**
- `walletAddress`: The address of the wallet.

**Returns:**
- `[]Asset`: A slice of assets for the wallet.
- `error`: Error if the request fails.

**Example:**
```go
assets, err := getWalletAssets("walletAddress")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(assets)
```

##### `getWalletFarm(walletAddress string, farmAddress string) (Farm, error)`
Fetches a specific farm for a given wallet.

**Parameters:**
- `walletAddress`: The address of the wallet.
- `farmAddress`: The address of the farm.

**Returns:**
- `Farm`: The farm data.
- `error`: Error if the request fails.

**Example:**
```go
farm, err := getWalletFarm("walletAddress", "farmAddress")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(farm)
```

##### `getWalletPool(walletAddress string, poolAddress string) (Pool, error)`
Fetches a specific pool for a given wallet.

**Parameters:**
- `walletAddress`: The address of the wallet.
- `poolAddress`: The address of the pool.

**Returns:**
- `Pool`: The pool data.
- `error`: Error if the request fails.

**Example:**
```go
pool, err := getWalletPool("walletAddress", "poolAddress")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(pool)
```

##### `getWalletPools(walletAddress string) ([]Pool, error)`
Fetches all pools for a given wallet.

**Parameters:**
- `walletAddress`: The address of the wallet.

**Returns:**
- `[]Pool`: A slice of pools for the wallet.
- `error`: Error if the request fails.

**Example:**
```go
pools, err := getWalletPools("walletAddress")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(pools)
```