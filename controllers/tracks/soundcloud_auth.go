package tracks

import (
	"encoding/json"
	"fmt"
	"github.com/MKA-Nigeria/mkanmedia-go/models"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"io/ioutil"
	. "net/http"
	"net/url"
	"strings"
)

type AuthImpl struct {
	//http client
	Client *Client
	Cache  *redis.Client
}

func (auth AuthImpl) GetToken() (*models.TokenResponse, *error) {

	var tokenModel models.TokenResponse
	result, _ := auth.Cache.Get(SoundCloudTokenKey).Result()

	tokenModel, err := tokenModel.TokenResponseFromJson(result)

	if *err == nil {
		fmt.Println("reusing token ", tokenModel)
		return &tokenModel, nil
	}

	fmt.Println("fetching token")
	token, err := auth.authenticate()
	if err != nil {
		return nil, err
	}

	data, _ := token.ToJson()
	//store token in redis and expire it in one hour
	err_ := auth.Cache.Set(SoundCloudTokenKey, data, TokenCacheTimeOut).Err()
	if err != nil {
		return nil, &err_
	}

	return token, &err_
}

func makeSoundCloudParams() url.Values {
	param := url.Values{}
	env := viper.GetString("env")

	param.Add("client_id", viper.GetString(env+".sources.soundcloud.sound_cloud_client_id"))
	param.Add("grant_type", "client_credentials")
	param.Add("client_secret", viper.GetString(env+".sources.soundcloud.sound_cloud_client_secret"))
	return param
}

func (auth AuthImpl) authenticate() (*models.TokenResponse, *error) {
	param := makeSoundCloudParams()
	newRequest, err := NewRequest(MethodPost, authEndPoint, strings.NewReader(param.Encode()))
	if err != nil {
		return nil, &err
	}

	newRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	newRequest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")

	resp, err := auth.Client.Do(newRequest)
	if err != nil {
		return nil, &err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= StatusBadRequest {
		return nil, &err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &err
	}

	token := models.TokenResponse{}

	if err := json.Unmarshal(body, &token); err != nil {
		return nil, &err
	}

	return &token, nil
}
