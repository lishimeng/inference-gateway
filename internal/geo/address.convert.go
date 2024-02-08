package geo

// ConvertAddress 坐标转换地址(GCG2000或WGS84)
func ConvertAddress(lon, lat float64) (address string, err error) {

	resp, err := client.Geo2Address(lon, lat)
	if err != nil {
		return
	}
	address = resp.Result.FormattedAddress
	return
}
