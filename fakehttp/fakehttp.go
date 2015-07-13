package fakehttp

func Get(resource string, lastTimestamp int) []byte {
	var resp string
	if resource == "cat" {
		if lastTimestamp == 0 {
			resp = `{"items":[{"name":"bobby"},{"name":"gray"}],"has_more":true,"last_timestamp":1}`
		} else {
			resp = `{"items":[{"name":"alive"},{"name":"dead"}],"has_more":false,"last_timestamp":2}`
		}
	} else {
		if lastTimestamp == 0 {
			resp = `{"items":[{"flavor":"peach"},{"flavor":"prune"}],"has_more":true,"last_timestamp":1}`
		} else {
			resp = `{"items":[{"flavor":"apple"},{"flavor":"meat"}],"has_more":false,"last_timestamp":2}`
		}
	}
	return []byte(resp)
}
