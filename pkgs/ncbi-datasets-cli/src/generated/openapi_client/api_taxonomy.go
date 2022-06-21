/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/v1/how-tos/genomes/large-download/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// TaxonomyApiService TaxonomyApi service
type TaxonomyApiService service

type ApiTaxNameQueryRequest struct {
	ctx _context.Context
	ApiService *TaxonomyApiService
	taxonQuery string	
	taxRankFilter *V1OrganismQueryRequestTaxRankFilter	
    Headers map[string]string
}

// Set the scope of searched tax ranks
func (r *ApiTaxNameQueryRequest) TaxRankFilter(taxRankFilter V1OrganismQueryRequestTaxRankFilter) *ApiTaxNameQueryRequest {
	r.taxRankFilter = &taxRankFilter
	return r
}

func (r ApiTaxNameQueryRequest) Execute() (V1SciNameAndIds, *_nethttp.Response, error) {
	return r.ApiService.TaxNameQueryExecute(r)
}

/*
TaxNameQuery Get a list of taxonomy names and IDs given a partial taxonomic name

This endpoint retrieves a list of taxonomy names and IDs given a partial taxonomic name of any rank.
 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taxonQuery NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
 @return ApiTaxNameQueryRequest
*/
func (a *TaxonomyApiService) TaxNameQuery(ctx _context.Context, taxonQuery string) ApiTaxNameQueryRequest {
	return ApiTaxNameQueryRequest{
		ApiService: a,
		ctx: ctx,
		taxonQuery: taxonQuery,
	}
}

// Execute executes the request
//  @return V1SciNameAndIds
func (a *TaxonomyApiService) TaxNameQueryExecute(r ApiTaxNameQueryRequest) (V1SciNameAndIds, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1SciNameAndIds
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxonomyApiService.TaxNameQuery")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/taxonomy/taxon_suggest/{taxon_query}"
	localVarPath = strings.Replace(localVarPath, "{"+"taxon_query"+"}", _neturl.PathEscape(parameterToString(r.taxonQuery, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.taxRankFilter != nil {
		localVarQueryParams.Add("tax_rank_filter", parameterToString(*r.taxRankFilter, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// override localVarHeaderParams with the headers passed into the function
	if len(r.Headers) > 0 {
		for k, v := range r.Headers { 
			localVarHeaderParams[k] = v
		}
	}

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuthHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	if localVarHTTPResponse.Header.Get("Content-Type") != "application/json" {
		return localVarReturnValue, localVarHTTPResponse, nil
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTaxonomyFilteredSubtreeRequest struct {
	ctx _context.Context
	ApiService *TaxonomyApiService
	taxons []string	
	specifiedLimit *bool	
	rankLimits *[]V1OrganismRankType	
    Headers map[string]string
}

// Limit to specified species
func (r *ApiTaxonomyFilteredSubtreeRequest) SpecifiedLimit(specifiedLimit bool) *ApiTaxonomyFilteredSubtreeRequest {
	r.specifiedLimit = &specifiedLimit
	return r
}
// Limit to the provided ranks.  If empty, accept any rank.
func (r *ApiTaxonomyFilteredSubtreeRequest) RankLimits(rankLimits []V1OrganismRankType) *ApiTaxonomyFilteredSubtreeRequest {
	r.rankLimits = &rankLimits
	return r
}

func (r ApiTaxonomyFilteredSubtreeRequest) Execute() (V1TaxonomyFilteredSubtreeResponse, *_nethttp.Response, error) {
	return r.ApiService.TaxonomyFilteredSubtreeExecute(r)
}

/*
TaxonomyFilteredSubtree Use taxonomic identifiers to get a filtered taxonomic subtree

Using NCBI Taxonomy IDs or names (common or scientific) at any rank, get a filtered taxonomic subtree that includes the full parent lineage and all immediate children from the selected taxonomic ranks in JSON format.
 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taxons
 @return ApiTaxonomyFilteredSubtreeRequest
*/
func (a *TaxonomyApiService) TaxonomyFilteredSubtree(ctx _context.Context, taxons []string) ApiTaxonomyFilteredSubtreeRequest {
	return ApiTaxonomyFilteredSubtreeRequest{
		ApiService: a,
		ctx: ctx,
		taxons: taxons,
	}
}

// Execute executes the request
//  @return V1TaxonomyFilteredSubtreeResponse
func (a *TaxonomyApiService) TaxonomyFilteredSubtreeExecute(r ApiTaxonomyFilteredSubtreeRequest) (V1TaxonomyFilteredSubtreeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1TaxonomyFilteredSubtreeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxonomyApiService.TaxonomyFilteredSubtree")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/taxonomy/taxon/{taxons}/filtered_subtree"
	localVarPath = strings.Replace(localVarPath, "{"+"taxons"+"}", _neturl.PathEscape(parameterToString(r.taxons, "csv")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.specifiedLimit != nil {
		localVarQueryParams.Add("specified_limit", parameterToString(*r.specifiedLimit, ""))
	}
	if r.rankLimits != nil {
		t := *r.rankLimits
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("rank_limits", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("rank_limits", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// override localVarHeaderParams with the headers passed into the function
	if len(r.Headers) > 0 {
		for k, v := range r.Headers { 
			localVarHeaderParams[k] = v
		}
	}

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuthHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	if localVarHTTPResponse.Header.Get("Content-Type") != "application/json" {
		return localVarReturnValue, localVarHTTPResponse, nil
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTaxonomyFilteredSubtreePostRequest struct {
	ctx _context.Context
	ApiService *TaxonomyApiService
	v1TaxonomyFilteredSubtreeRequest *V1TaxonomyFilteredSubtreeRequest	
    Headers map[string]string
}

func (r *ApiTaxonomyFilteredSubtreePostRequest) V1TaxonomyFilteredSubtreeRequest(v1TaxonomyFilteredSubtreeRequest V1TaxonomyFilteredSubtreeRequest) *ApiTaxonomyFilteredSubtreePostRequest {
	r.v1TaxonomyFilteredSubtreeRequest = &v1TaxonomyFilteredSubtreeRequest
	return r
}

func (r ApiTaxonomyFilteredSubtreePostRequest) Execute() (V1TaxonomyFilteredSubtreeResponse, *_nethttp.Response, error) {
	return r.ApiService.TaxonomyFilteredSubtreePostExecute(r)
}

/*
TaxonomyFilteredSubtreePost Use taxonomic identifiers to get a filtered taxonomic subtree by post

Using NCBI Taxonomy IDs or names (common or scientific) at any rank, get a filtered taxonomic subtree that includes the full parent lineage and all immediate children from the selected taxonomic ranks in JSON format.
 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTaxonomyFilteredSubtreePostRequest
*/
func (a *TaxonomyApiService) TaxonomyFilteredSubtreePost(ctx _context.Context, v1TaxonomyFilteredSubtreeRequest *V1TaxonomyFilteredSubtreeRequest) ApiTaxonomyFilteredSubtreePostRequest {
	return ApiTaxonomyFilteredSubtreePostRequest{
		ApiService: a,
		ctx: ctx,
		v1TaxonomyFilteredSubtreeRequest: v1TaxonomyFilteredSubtreeRequest,
	}
}

// Execute executes the request
//  @return V1TaxonomyFilteredSubtreeResponse
func (a *TaxonomyApiService) TaxonomyFilteredSubtreePostExecute(r ApiTaxonomyFilteredSubtreePostRequest) (V1TaxonomyFilteredSubtreeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1TaxonomyFilteredSubtreeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxonomyApiService.TaxonomyFilteredSubtreePost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/taxonomy/filtered_subtree"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.v1TaxonomyFilteredSubtreeRequest == nil {
		return localVarReturnValue, nil, reportError("v1TaxonomyFilteredSubtreeRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// override localVarHeaderParams with the headers passed into the function
	if len(r.Headers) > 0 {
		for k, v := range r.Headers { 
			localVarHeaderParams[k] = v
		}
	}

	// body params
	localVarPostBody = r.v1TaxonomyFilteredSubtreeRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuthHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	if localVarHTTPResponse.Header.Get("Content-Type") != "application/json" {
		return localVarReturnValue, localVarHTTPResponse, nil
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTaxonomyMetadataRequest struct {
	ctx _context.Context
	ApiService *TaxonomyApiService
	taxons []string	
	returnedContent *V1TaxonomyMetadataRequestContentType	
    Headers map[string]string
}

// Return either tax-ids alone, or entire taxononmy-metadata records
func (r *ApiTaxonomyMetadataRequest) ReturnedContent(returnedContent V1TaxonomyMetadataRequestContentType) *ApiTaxonomyMetadataRequest {
	r.returnedContent = &returnedContent
	return r
}

func (r ApiTaxonomyMetadataRequest) Execute() (V1TaxonomyMetadataResponse, *_nethttp.Response, error) {
	return r.ApiService.TaxonomyMetadataExecute(r)
}

/*
TaxonomyMetadata Use taxonomic identifiers to get taxonomic metadata

Using NCBI Taxonomy IDs or names (common or scientific) at any rank, get metadata about a taxonomic node including taxonomic identifiers, lineage information, child nodes, and gene and genome counts in JSON format.
 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taxons
 @return ApiTaxonomyMetadataRequest
*/
func (a *TaxonomyApiService) TaxonomyMetadata(ctx _context.Context, taxons []string) ApiTaxonomyMetadataRequest {
	return ApiTaxonomyMetadataRequest{
		ApiService: a,
		ctx: ctx,
		taxons: taxons,
	}
}

// Execute executes the request
//  @return V1TaxonomyMetadataResponse
func (a *TaxonomyApiService) TaxonomyMetadataExecute(r ApiTaxonomyMetadataRequest) (V1TaxonomyMetadataResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1TaxonomyMetadataResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxonomyApiService.TaxonomyMetadata")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/taxonomy/taxon/{taxons}"
	localVarPath = strings.Replace(localVarPath, "{"+"taxons"+"}", _neturl.PathEscape(parameterToString(r.taxons, "csv")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.returnedContent != nil {
		localVarQueryParams.Add("returned_content", parameterToString(*r.returnedContent, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// override localVarHeaderParams with the headers passed into the function
	if len(r.Headers) > 0 {
		for k, v := range r.Headers { 
			localVarHeaderParams[k] = v
		}
	}

	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuthHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	if localVarHTTPResponse.Header.Get("Content-Type") != "application/json" {
		return localVarReturnValue, localVarHTTPResponse, nil
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTaxonomyMetadataPostRequest struct {
	ctx _context.Context
	ApiService *TaxonomyApiService
	v1TaxonomyMetadataRequest *V1TaxonomyMetadataRequest	
    Headers map[string]string
}

func (r *ApiTaxonomyMetadataPostRequest) V1TaxonomyMetadataRequest(v1TaxonomyMetadataRequest V1TaxonomyMetadataRequest) *ApiTaxonomyMetadataPostRequest {
	r.v1TaxonomyMetadataRequest = &v1TaxonomyMetadataRequest
	return r
}

func (r ApiTaxonomyMetadataPostRequest) Execute() (V1TaxonomyMetadataResponse, *_nethttp.Response, error) {
	return r.ApiService.TaxonomyMetadataPostExecute(r)
}

/*
TaxonomyMetadataPost Use taxonomic identifiers to get taxonomic metadata by post

Using NCBI Taxonomy IDs or names (common or scientific) at any rank, get metadata about a taxonomic node including taxonomic identifiers, lineage information, child nodes, and gene and genome counts in JSON format.
 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTaxonomyMetadataPostRequest
*/
func (a *TaxonomyApiService) TaxonomyMetadataPost(ctx _context.Context, v1TaxonomyMetadataRequest *V1TaxonomyMetadataRequest) ApiTaxonomyMetadataPostRequest {
	return ApiTaxonomyMetadataPostRequest{
		ApiService: a,
		ctx: ctx,
		v1TaxonomyMetadataRequest: v1TaxonomyMetadataRequest,
	}
}

// Execute executes the request
//  @return V1TaxonomyMetadataResponse
func (a *TaxonomyApiService) TaxonomyMetadataPostExecute(r ApiTaxonomyMetadataPostRequest) (V1TaxonomyMetadataResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V1TaxonomyMetadataResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxonomyApiService.TaxonomyMetadataPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/taxonomy"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.v1TaxonomyMetadataRequest == nil {
		return localVarReturnValue, nil, reportError("v1TaxonomyMetadataRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// override localVarHeaderParams with the headers passed into the function
	if len(r.Headers) > 0 {
		for k, v := range r.Headers { 
			localVarHeaderParams[k] = v
		}
	}

	// body params
	localVarPostBody = r.v1TaxonomyMetadataRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuthHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	if localVarHTTPResponse.Header.Get("Content-Type") != "application/json" {
		return localVarReturnValue, localVarHTTPResponse, nil
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v RpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
