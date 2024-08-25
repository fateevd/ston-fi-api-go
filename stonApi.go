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

func GetAsset(assetAddress string) (Asset, error) {
	response, err := request("/v1/assets/" + assetAddress)
	if err != nil {
		return Asset{}, err
	}
	return getValues[Asset](response), nil
}

func GeAssets() ([]Asset, error) {
	response, err := request("/v1/assets")
	if err != nil {
		return []Asset{}, err
	}
	return getValues[[]Asset](response), nil
}

func GeFarms() ([]Farm, error) {
	response, err := request("/v1/farms")
	if err != nil {
		return []Farm{}, err
	}
	return getValues[[]Farm](response), nil
}

func GeFarm(farmAddress string) (Farm, error) {
	response, err := request("/v1/farms/" + farmAddress)
	if err != nil {
		return Farm{}, err
	}
	return getValues[Farm](response), nil
}

func GeFarmsByPool(poolAddress string) ([]Farm, error) {
	response, err := request("/v1/farms_by_pool/" + poolAddress)
	if err != nil {
		return []Farm{}, err
	}
	return getValues[[]Farm](response), nil
}

func GeSwapPairs() ([]SwapPairs, error) {
	response, err := request("/v1/markets")
	if err != nil {
		return []SwapPairs{}, err
	}
	return getValues[[]SwapPairs](response), nil

}

func GePool(poolAddress string) (Pool, error) {
	response, err := request("/v1/pools/" + poolAddress)
	if err != nil {
		return Pool{}, err
	}
	return getValues[Pool](response), nil

}

func GePools() ([]Pool, error) {
	response, err := request("/v1/pools")
	if err != nil {
		return []Pool{}, err
	}
	return getValues[[]Pool](response), nil
}

func GeWalletAsset(walletAddress string, assetAddress string) (Asset, error) {
	response, err := request("/v1/wallets/" + walletAddress + "/assets/" + assetAddress)
	if err != nil {
		return Asset{}, err
	}
	return getValues[Asset](response), nil
}

func GeWalletAssets(walletAddress string) ([]Asset, error) {
	response, err := request("/v1/wallets/" + walletAddress + "/assets")
	if err != nil {
		return []Asset{}, err
	}
	return getValues[[]Asset](response), nil
}

func GeWalletFarm(walletAddress string, farmAddress string) (Farm, error) {
	response, err := request("/v1/wallets/" + walletAddress + "/farms/" + farmAddress)
	if err != nil {
		return Farm{}, err
	}
	return getValues[Farm](response), nil
}

func GeWalletPool(walletAddress string, poolAddress string) (Pool, error) {
	response, err := request("/v1/wallets/" + walletAddress + "/pools/" + poolAddress)
	if err != nil {
		return Pool{}, err
	}
	return getValues[Pool](response), nil
}

func GeWalletPools(walletAddress string) ([]Pool, error) {
	response, err := request("/v1/wallets/" + walletAddress + "/pools")
	if err != nil {
		return []Pool{}, err
	}
	return getValues[[]Pool](response), nil
}
