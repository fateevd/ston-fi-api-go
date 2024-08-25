package stonApi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseUrl = "https://api.ston.fi"

func request(path string) ([]byte, error) {
	fmt.Println("Requesting", baseUrl+path)
	resp, err := http.Get(baseUrl + path)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getValues[T any](data []byte) T {
	var dataMap map[string]interface{}
	var result T
	err := json.Unmarshal(data, &dataMap)
	if err != nil {
		return result
	}

	for _, value := range dataMap {
		jsonData, err := json.Marshal(value)
		if err != nil {
			return result
		}
		err = json.Unmarshal(jsonData, &result)
		if err != nil {
			return result
		}
		break
	}
	return result
}

func getAsset(assetAddress string) (Asset, error) {
	response, err := request("/v1/assets/" + assetAddress)
	if err != nil {
		return Asset{}, err
	}
	return getValues[Asset](response), nil
}

func getAssets() ([]Asset, error) {
	response, err := request("/v1/assets")
	if err != nil {
		return []Asset{}, err
	}
	return getValues[[]Asset](response), nil
}

func getFarms() ([]Farm, error) {
	response, err := request("/v1/farms")
	if err != nil {
		return []Farm{}, err
	}
	return getValues[[]Farm](response), nil
}

func getFarm(farmAddress string) (Farm, error) {
	response, err := request("/v1/farms/" + farmAddress)
	if err != nil {
		return Farm{}, err
	}
	return getValues[Farm](response), nil
}

func getFarmsByPool(poolAddress string) ([]Farm, error) {
	response, err := request("/v1/farms_by_pool/" + poolAddress)
	if err != nil {
		return []Farm{}, err
	}
	return getValues[[]Farm](response), nil
}

func getSwapPairs() ([]SwapPairs, error) {
	response, err := request("/v1/markets")
	if err != nil {
		return []SwapPairs{}, err
	}
	return getValues[[]SwapPairs](response), nil

}

func getPool(poolAddress string) (Pool, error) {
	response, err := request("/v1/pools/" + poolAddress)
	if err != nil {
		return Pool{}, err
	}
	return getValues[Pool](response), nil

}

func getPools() ([]Pool, error) {
	response, err := request("/v1/pools")
	if err != nil {
		return []Pool{}, err
	}
	return getValues[[]Pool](response), nil
}

func getWalletAsset(walletAddress string, assetAddress string) (Asset, error) {
	response, err := request("/v1/wallets/" + walletAddress + "/assets/" + assetAddress)
	if err != nil {
		return Asset{}, err
	}
	return getValues[Asset](response), nil
}

func getWalletAssets(walletAddress string) ([]Asset, error) {
	response, err := request("/v1/wallets/" + walletAddress + "/assets")
	if err != nil {
		return []Asset{}, err
	}
	return getValues[[]Asset](response), nil
}

func getWalletFarm(walletAddress string, farmAddress string) (Farm, error) {
	response, err := request("/v1/wallets/" + walletAddress + "/farms/" + farmAddress)
	if err != nil {
		return Farm{}, err
	}
	return getValues[Farm](response), nil
}

func getWalletPool(walletAddress string, poolAddress string) (Pool, error) {
	response, err := request("/v1/wallets/" + walletAddress + "/pools/" + poolAddress)
	if err != nil {
		return Pool{}, err
	}
	return getValues[Pool](response), nil
}

func getWalletPools(walletAddress string) ([]Pool, error) {
	response, err := request("/v1/wallets/" + walletAddress + "/pools")
	if err != nil {
		return []Pool{}, err
	}
	return getValues[[]Pool](response), nil
}
