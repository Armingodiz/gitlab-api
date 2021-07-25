package cache


type Cache  interface {
	Set(string,string) error
	Get(string) (string,error)
}


/*func GetNewCache(port int)Cache {
	if port==-1{
		return inMemoryCache{
			mcache := make(map[string]string)
		}
	}else{
		return GetRedisCache(port)
	}
}*/


