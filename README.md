### Documentation

This Go program interacts with the `ston.fi` API v1 to fetch various types of data such as assets, farms, pools, and wallet-related information. Below is the documentation for each function along with examples of how to use them.

### Getting Started
To install the library, run:
```
go get github.com/fateevd/ston-fi-api-go
```

#### Constants
```go
const baseUrl = "https://api.ston.fi"
```
- `baseUrl`: The base URL for the `ston.fi` API.

#### Functions

##### `GetAsset(assetAddress string) (Asset, error)` 1231
Fetches a single asset by its address.

**Parameters:**
- `assetAddress`: The address of the asset.

**Returns:**
- `Asset`: The asset data.
- `error`: Error if the request fails.

**Example:**
```go
asset, err := GetAsset("EQD8T***************")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(asset)
```

##### `GetAssets() ([]Asset, error)`
Fetches all assets.

**Returns:**
- `[]Asset`: A slice of all assets.
- `error`: Error if the request fails.

**Example:**
```go
assets, err := GetAssets()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(assets)
```

##### `GetFarms() ([]Farm, error)`
Fetches all farms.

**Returns:**
- `[]Farm`: A slice of all farms.
- `error`: Error if the request fails.

**Example:**
```go
farms, err := GetFarms()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(farms)
```

##### `GetFarm(farmAddress string) (Farm, error)`
Fetches a single farm by its address.

**Parameters:**
- `farmAddress`: The address of the farm.

**Returns:**
- `Farm`: The farm data.
- `error`: Error if the request fails.

**Example:**
```go
farm, err := GetFarm("EQD8T***************")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(farm)
```

##### `GetFarmsByPool(poolAddress string) ([]Farm, error)`
Fetches all farms associated with a specific pool.

**Parameters:**
- `poolAddress`: The address of the pool.

**Returns:**
- `[]Farm`: A slice of farms associated with the pool.
- `error`: Error if the request fails.

**Example:**
```go
farms, err := GetFarmsByPool("EQD8T***************")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(farms)
```

##### `GetSwapPairs() ([]SwapPairs, error)`
Fetches all swap pairs.

**Returns:**
- `[]SwapPairs`: A slice of swap pairs.
- `error`: Error if the request fails.

**Example:**
```go
swapPairs, err := GetSwapPairs()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(swapPairs)
```

##### `GetPool(poolAddress string) (Pool, error)`
Fetches a single pool by its address.

**Parameters:**
- `poolAddress`: The address of the pool.

**Returns:**
- `Pool`: The pool data.
- `error`: Error if the request fails.

**Example:**
```go
pool, err := GetPool("EQD8T***************")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(pool)
```

##### `GetPools() ([]Pool, error)`
Fetches all pools.

**Returns:**
- `[]Pool`: A slice of all pools.
- `error`: Error if the request fails.

**Example:**
```go
pools, err := GetPools()
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(pools)
```

##### `GetWalletAsset(walletAddress string, assetAddress string) (Asset, error)`
Fetches a specific asset for a given wallet.

**Parameters:**
- `walletAddress`: The address of the wallet.
- `assetAddress`: The address of the asset.

**Returns:**
- `Asset`: The asset data.
- `error`: Error if the request fails.

**Example:**
```go
asset, err := GetWalletAsset("walletAddress", "assetAddress")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(asset)
```

##### `GetWalletAssets(walletAddress string) ([]Asset, error)`
Fetches all assets for a given wallet.

**Parameters:**
- `walletAddress`: The address of the wallet.

**Returns:**
- `[]Asset`: A slice of assets for the wallet.
- `error`: Error if the request fails.

**Example:**
```go
assets, err := GetWalletAssets("walletAddress")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(assets)
```

##### `GetWalletFarm(walletAddress string, farmAddress string) (Farm, error)`
Fetches a specific farm for a given wallet.

**Parameters:**
- `walletAddress`: The address of the wallet.
- `farmAddress`: The address of the farm.

**Returns:**
- `Farm`: The farm data.
- `error`: Error if the request fails.

**Example:**
```go
farm, err := GetWalletFarm("walletAddress", "farmAddress")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(farm)
```

##### `GetWalletPool(walletAddress string, poolAddress string) (Pool, error)`
Fetches a specific pool for a given wallet.

**Parameters:**
- `walletAddress`: The address of the wallet.
- `poolAddress`: The address of the pool.

**Returns:**
- `Pool`: The pool data.
- `error`: Error if the request fails.

**Example:**
```go
pool, err := GetWalletPool("walletAddress", "poolAddress")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(pool)
```

##### `GetWalletPools(walletAddress string) ([]Pool, error)`
Fetches all pools for a given wallet.

**Parameters:**
- `walletAddress`: The address of the wallet.

**Returns:**
- `[]Pool`: A slice of pools for the wallet.
- `error`: Error if the request fails.

**Example:**
```go
pools, err := GetWalletPools("walletAddress")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(pools)
```
