package registry

type tagsResponse struct {
	Tags []string `json:"tags"`
}

// TODO: https://github.com/docker/distribution/blob/master/docs/spec/api.md#pagination-1
// , page, max-tags int
// Tags returns the tags for a specific repository.
func (r *Registry) Tags(repository string) ([]string, error) {
	url := r.url("/v2/%s/tags/list", repository)
	r.Logf("registry.tags url=%s repository=%s", url, repository)

	var response tagsResponse
	if _, err := r.getJSON(url, &response, false); err != nil {
		return nil, err
	}

	return response.Tags, nil
}
